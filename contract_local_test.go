package main_test

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"specmatic-order-bff-grpc-go/internal/com/store/order/bff/config"
	"specmatic-order-bff-grpc-go/internal/com/store/order/bff/tests"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go/network"
)

var bffServiceCmd *exec.Cmd

func TestContractLocal(t *testing.T) {
	env := setUpEnviron(t)

	// setUp (start domain service stub with specmatic-grpc and bff server in container)
	setUpServices(t, env)

	// RUN (run specmatic-grpc test in container)
	runContractTest(t, env)

	// TEAR DOWN
	defer tearDownAll(t, env)
}

func setUpEnviron(t *testing.T) *tests.TestEnvironment {
	config, err := config.LoadConfig()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	return &tests.TestEnvironment{
		Ctx:                  context.Background(),
		Config:               config,
		ExpectedMessageCount: 10,
	}
}

func setUpServices(t *testing.T, env *tests.TestEnvironment) {
	var err error

	// Create a sub net and store in env.
	newNetwork, err := network.New(env.Ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		require.NoError(t, newNetwork.Remove(env.Ctx))
	})
	env.DockerNetwork = newNetwork

	printHeaderr(t, 1, "Starting Domain Service")
	env.DomainServiceContainer, env.DomainServiceDynamicPort, err = tests.StartDomainService(env)
	if err != nil {
		t.Fatalf("could not start domain service container: %v", err)
	}

	printHeaderr(t, 2, "Starting Kafka Service")
	env.KafkaServiceContainer, env.KafkaServiceDynamicPort, err = tests.StartKafkaMock(env)
	if err != nil {
		t.Fatalf("could not start domain service container: %v", err)
	}

	printHeaderr(t, 3, "Starting BFF Service")
	bffServiceCmd, env.BffServiceDynamicPort, err = startBFFService(env, t)
	if err != nil {
		t.Fatalf("could not start bff service container: %v", err)
	}
}

func startBFFService(env *tests.TestEnvironment, t *testing.T) (*exec.Cmd, string, error) {
	cmd := exec.Command("go", "run", "./cmd/main.go")
	cmd.Env = os.Environ()

	// Set environment variables
	cmd.Env = append(os.Environ(),
		"DOMAIN_SERVER_PORT="+env.DomainServiceDynamicPort,
		"DOMAIN_SERVER_HOST=localhost",
		"KAFKA_PORT="+env.KafkaServiceDynamicPort,
		"KAFKA_HOST=localhost",
		// Add any other environment variables your app needs
	)

	// Buffer to store the command output
	var outputBuf bytes.Buffer
	cmd.Stdout = &outputBuf
	cmd.Stderr = &outputBuf

	// Start the application
	if err := cmd.Start(); err != nil {
		return nil, "", err
	}

	// Start a goroutine to read the application's output
	go func() {
		for {
			line, err := outputBuf.ReadString('\n')
			if err != nil {
				time.Sleep(100 * time.Millisecond)
				continue
			}
			t.Log("BFF output:", strings.TrimSpace(line))

		}
	}()

	env.Config.BFFServer.Host = "host.docker.internal"

	time.Sleep(2 * time.Second)

	return cmd, env.Config.BFFServer.Port, nil
}

func runContractTest(t *testing.T, env *tests.TestEnvironment) {
	printHeaderr(t, 4, "Starting tests")
	testLogs, err := tests.RunTestContainer(env)
	if err != nil {
		t.Fatalf("Could not run test container: %s", err)
	}

	// Print test outcomes
	t.Log("Test Results:")
	t.Log(testLogs)
}

func tearDownAll(t *testing.T, env *tests.TestEnvironment) {
	if bffServiceCmd != nil {
		if err := bffServiceCmd.Process.Kill(); err != nil {
			t.Logf("Failed to kill BFF process: %v", err)
		}
		if err := bffServiceCmd.Wait(); err != nil {
			t.Logf("Failed to wait for BFF process: %v", err)
		}
	}
	if env.KafkaServiceContainer != nil {
		err := tests.VerifyKafkaExpectations(env)
		if err != nil {
			t.Logf("Kafka expectations were not met: %s", err)
			t.Fail()
		}
		if err := env.KafkaServiceContainer.Terminate(env.Ctx); err != nil {
			t.Logf("Failed to terminate Kafka container: %v", err)
		}
	}
	if env.DomainServiceContainer != nil {
		if err := env.DomainServiceContainer.Terminate(env.Ctx); err != nil {
			t.Logf("Failed to terminate stub container: %v", err)
		}
	}
}

func printHeaderr(t *testing.T, stepNum int, title string) {
	t.Log("")
	t.Logf("======== STEP %d =========", stepNum)
	t.Log(title)
	t.Log("=========================")
	t.Log("")
}
