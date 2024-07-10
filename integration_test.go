package main_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"specmatic-order-bff-grpc-go/internal/com/store/order/bff/config"
	"strconv"
	"testing"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type testEnvironment struct {
	ctx               context.Context
	domainServiceHost testcontainers.Container
	domainServicePort string
	bffServiceHost    testcontainers.Container
	bffServicePort    string
	config            *config.Config
}

func TestIntegration(t *testing.T) {

	// SETUP (domain grpc server and bff grpc server )
	env, err := setup(t)
	if err != nil {
		t.Fatalf("Setup failed: %v", err)
	}

	// RUN (on test container)
	runTests(t, env)

	// TEAR DOWN
	defer teardown(t, env)
}

func setup(t *testing.T) (*testEnvironment, error) {
	env := &testEnvironment{
		ctx: context.Background(),
	}

	var err error

	// Load configuration
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}
	env.config = cfg

	printHeader(t, 1, "Starting Domain Service")
	env.domainServiceHost, env.domainServicePort, err = startDomainService(env)
	if err != nil {
		return nil, fmt.Errorf("could not start domain service container: %w", err)
	}

	printHeader(t, 2, "Starting BFF Service")
	env.bffServiceHost, env.bffServicePort, err = startBFFService(t, env, env.domainServicePort)
	if err != nil {
		return nil, fmt.Errorf("could not start bff service container: %w", err)
	}

	return env, nil
}

func runTests(t *testing.T, env *testEnvironment) {
	printHeader(t, 3, "Starting tests")
	testLogs, err := runTestContainer(env.ctx, env.bffServicePort)
	if err != nil {
		t.Fatalf("Could not run test container: %s", err)
	}

	// Print test outcomes
	t.Log("Test Results:")
	t.Log(testLogs)
}

func teardown(t *testing.T, env *testEnvironment) {
	if env.bffServiceHost != nil {
		if err := env.bffServiceHost.Terminate(env.ctx); err != nil {
			t.Logf("Failed to terminate BFF container: %v", err)
		}
	}
	if env.domainServiceHost != nil {
		if err := env.domainServiceHost.Terminate(env.ctx); err != nil {
			t.Logf("Failed to terminate stub container: %v", err)
		}
	}
}

func startDomainService(env *testEnvironment) (testcontainers.Container, string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current directory: %v", err)
	}

	// Create a nat.Port from the string port
	port, err := nat.NewPort("tcp", env.config.Backend.Port)
	if err != nil {
		return nil, "", fmt.Errorf("invalid port number: %w", err)
	}

	req := testcontainers.ContainerRequest{
		Image:        "znsio/specmatic-grpc-trial",
		ExposedPorts: []string{port.Port() + "/tcp"},
		Cmd:          []string{"stub"},
		Mounts: testcontainers.Mounts(
			testcontainers.BindMount(filepath.Join(pwd, "specmatic.yaml"), "/usr/src/app/specmatic.yaml"),
		),
		WaitingFor: wait.ForLog("Stub server is running"),
	}

	stubContainer, err := testcontainers.GenericContainer(env.ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, "", err
	}

	domainServicePort, err := stubContainer.MappedPort(env.ctx, port)
	if err != nil {
		return nil, "", err
	}

	return stubContainer, domainServicePort.Port(), nil
}

func startBFFService(t *testing.T, env *testEnvironment, domainServicePort string) (testcontainers.Container, string, error) {
	dockerfilePath := "Dockerfile"

	bffImageName := "specmatic-order-bff-grpc-go"
	buildCmd := exec.Command("docker", "build", "-t", bffImageName, "-f", dockerfilePath, ".")
	// enable the following for detailed logs of bff service docerization.
	// buildCmd.Stdout = os.Stdout
	// buildCmd.Stderr = os.Stderr
	if err := buildCmd.Run(); err != nil {
		return nil, "", fmt.Errorf("could not build BFF image: %w", err)
	}

	// Create a nat.Port from the string port
	port, err := nat.NewPort("tcp", env.config.BFFServer.Port)
	if err != nil {
		return nil, "", fmt.Errorf("invalid port number: %w", err)
	}

	req := testcontainers.ContainerRequest{
		Image: bffImageName,
		Env: map[string]string{
			"DOMAIN_SERVER_PORT": domainServicePort,
			"DOMAIN_SERVER_HOST": "host.docker.internal",
		},
		ExposedPorts: []string{port.Port() + "/tcp"},
		WaitingFor:   wait.ForLog("Starting gRPC server"),
	}

	t.Log("Container created")

	bffContainer, err := testcontainers.GenericContainer(env.ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, "", err
	}

	bffPort, err := bffContainer.MappedPort(env.ctx, port)
	if err != nil {
		return nil, "", err
	}

	return bffContainer, bffPort.Port(), nil
}

func runTestContainer(ctx context.Context, bffPort string) (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current directory: %v", err)
	}

	bffPortInt, err := strconv.Atoi(bffPort)
	if err != nil {
		return "", fmt.Errorf("invalid port number: %w", err)
	}

	req := testcontainers.ContainerRequest{
		Image: "znsio/specmatic-grpc-trial",
		Env: map[string]string{
			"SPECMATIC_GENERATIVE_TESTS": "true",
		},
		Cmd: []string{"test", fmt.Sprintf("--port=%d", bffPortInt), "--host=host.docker.internal"},
		Mounts: testcontainers.Mounts(
			testcontainers.BindMount(filepath.Join(pwd, "specmatic.yaml"), "/usr/src/app/specmatic.yaml"),
		),
		WaitingFor: wait.ForLog("Tests completed"),
	}

	testContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return "", err
	}
	defer testContainer.Terminate(ctx)

	// Streaming testing logs to terminal
	logReader, err := testContainer.Logs(ctx)
	if err != nil {
		return "", err
	}
	defer logReader.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, logReader)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func printHeader(t *testing.T, stepNum int, title string) {
	t.Log("")
	t.Logf("======== STEP %d =========", stepNum)
	t.Log(title)
	t.Log("=========================")
	t.Log("")
}
