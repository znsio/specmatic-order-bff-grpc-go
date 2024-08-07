package main_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"specmatic-order-bff-grpc-go/internal/com/store/order/bff/config"
	"strconv"
	"testing"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/network"
	"github.com/testcontainers/testcontainers-go/wait"
)

type testEnvironment struct {
	ctx                      context.Context
	domainServiceContainer   testcontainers.Container
	domainServiceDynamicPort string
	kafkaServiceContainer    testcontainers.Container
	kafkaServiceDynamicPort  string
	bffServiceContainer      testcontainers.Container
	bffServiceDynamicPort    string
	dockerNetwork            *testcontainers.DockerNetwork
	config                   *config.Config
}

func TestContract(t *testing.T) {
	env := setUpEnv(t)

	// setUp (start domain service stub with specmatic-grpc and bff server in container)
	setUp(t, env)

	// RUN (run specmatic-grpc test in container)
	runTests(t, env)

	// TEAR DOWN
	defer tearDown(t, env)
}

func setUpEnv(t *testing.T) *testEnvironment {
	config, err := config.LoadConfig()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	return &testEnvironment{
		ctx:    context.Background(),
		config: config,
	}
}

func setUp(t *testing.T, env *testEnvironment) {
	var err error

	// Create a sub net and store in env.
	newNetwork, err := network.New(env.ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		require.NoError(t, newNetwork.Remove(env.ctx))
	})
	env.dockerNetwork = newNetwork

	printHeader(t, 1, "Starting Domain Service")
	env.domainServiceContainer, env.domainServiceDynamicPort, err = startDomainService(env)
	if err != nil {
		t.Fatalf("could not start domain service container: %v", err)
	}

	printHeader(t, 2, "Starting Kafka Service")
	env.kafkaServiceContainer, env.kafkaServiceDynamicPort, err = startKafkaMock(env)
	if err != nil {
		t.Fatalf("could not start domain service container: %v", err)
	}

	printHeader(t, 3, "Starting BFF Service")
	env.bffServiceContainer, env.bffServiceDynamicPort, err = startBFFService(t, env)
	if err != nil {
		t.Fatalf("could not start bff service container: %v", err)
	}
}

func runTests(t *testing.T, env *testEnvironment) {
	printHeader(t, 4, "Starting tests")
	testLogs, err := runTestContainer(env)
	if err != nil {
		t.Fatalf("Could not run test container: %s", err)
	}

	// Print test outcomes
	t.Log("Test Results:")
	t.Log(testLogs)
}

func tearDown(t *testing.T, env *testEnvironment) {
	if env.bffServiceContainer != nil {
		if err := env.bffServiceContainer.Terminate(env.ctx); err != nil {
			t.Logf("Failed to terminate BFF container: %v", err)
		}
	}
	if env.kafkaServiceContainer != nil {
		if err := env.kafkaServiceContainer.Terminate(env.ctx); err != nil {
			t.Logf("Failed to terminate Kafka container: %v", err)
		}
	}
	if env.domainServiceContainer != nil {
		if err := env.domainServiceContainer.Terminate(env.ctx); err != nil {
			t.Logf("Failed to terminate stub container: %v", err)
		}
	}
}

func startDomainService(env *testEnvironment) (testcontainers.Container, string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current directory: %v", err)
	}

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
		Networks: []string{
			env.dockerNetwork.Name,
		},
		NetworkAliases: map[string][]string{
			env.dockerNetwork.Name: {"order-api-mock"},
		},
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

func startKafkaMock(env *testEnvironment) (testcontainers.Container, string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, "", fmt.Errorf("Error getting current directory: %v", err)

	}

	port, err := nat.NewPort("tcp", env.config.KafkaService.Port)
	if err != nil {
		return nil, "", fmt.Errorf("invalid port number: %w", err)
	}

	req := testcontainers.ContainerRequest{
		Name:         "specmatic-kafka",
		Image:        "znsio/specmatic-kafka-trial:0.22.10",
		ExposedPorts: []string{port.Port() + "/tcp"},
		Networks: []string{
			env.dockerNetwork.Name,
		},
		NetworkAliases: map[string][]string{
			env.dockerNetwork.Name: {"specmatic-kafka"},
		},
		Cmd: []string{"virtualize"},
		Env: map[string]string{
			"KAFKA_EXTERNAL_HOST": env.config.KafkaService.Host,
			"KAFKA_EXTERNAL_PORT": port.Port(),
		},
		Mounts: testcontainers.Mounts(
			testcontainers.BindMount(filepath.Join(pwd, "specmatic.yaml"), "/usr/src/app/specmatic.yaml"),
		),
		WaitingFor: wait.ForLog("Listening on topics: (product-queries)").WithStartupTimeout(2 * time.Minute),
	}

	kafkaC, err := testcontainers.GenericContainer(env.ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		fmt.Printf("Error starting Kafka mock container: %v", err)
	}

	mappedPort, err := kafkaC.MappedPort(env.ctx, port)
	if err != nil {
		fmt.Printf("Error getting mapped port for Kafka mock: %v", err)
	}

	return kafkaC, mappedPort.Port(), nil
}

func startBFFService(t *testing.T, env *testEnvironment) (testcontainers.Container, string, error) {

	port, err := nat.NewPort("tcp", env.config.BFFServer.Port)
	if err != nil {
		return nil, "", fmt.Errorf("invalid port number: %w", err)
	}

	dockerfilePath := "Dockerfile"
	contextPath := "."

	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    contextPath,
			Dockerfile: dockerfilePath,
		},
		Env: map[string]string{
			"DOMAIN_SERVER_PORT": env.config.Backend.Port,
			"DOMAIN_SERVER_HOST": "order-api-mock",
			"KAFKA_PORT":         env.config.KafkaService.Port,
			"KAFKA_HOST":         "specmatic-kafka",
		},
		Networks: []string{
			env.dockerNetwork.Name,
		},
		NetworkAliases: map[string][]string{
			env.dockerNetwork.Name: {"bff-service"},
		},
		ExposedPorts: []string{env.config.BFFServer.Port + "/tcp"},
		WaitingFor:   wait.ForLog("Starting gRPC server"),
	}

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

func runTestContainer(env *testEnvironment) (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current directory: %v", err)
	}

	bffPortInt, err := strconv.Atoi(env.config.BFFServer.Port)
	// bffPortInt, err := strconv.Atoi(env.bffServiceDynamicPort)
	if err != nil {
		return "", fmt.Errorf("invalid port number: %w", err)
	}

	req := testcontainers.ContainerRequest{
		Image: "znsio/specmatic-grpc-trial",
		Env: map[string]string{
			"SPECMATIC_GENERATIVE_TESTS": "true",
		},
		Cmd: []string{"test", fmt.Sprintf("--port=%d", bffPortInt), "--host=bff-service"},
		Mounts: testcontainers.Mounts(
			testcontainers.BindMount(filepath.Join(pwd, "specmatic.yaml"), "/usr/src/app/specmatic.yaml"),
		),
		Networks: []string{
			env.dockerNetwork.Name,
		},
		WaitingFor: wait.ForLog("Passed Tests:"),
	}

	testContainer, err := testcontainers.GenericContainer(env.ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return "", err
	}
	defer testContainer.Terminate(env.ctx)

	// Streaming testing logs to terminal
	logReader, err := testContainer.Logs(env.ctx)
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
