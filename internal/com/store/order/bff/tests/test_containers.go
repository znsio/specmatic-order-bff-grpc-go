package tests

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func StartDomainService(env *TestEnvironment) (testcontainers.Container, string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current directory: %v", err)
	}

	port, err := nat.NewPort("tcp", env.Config.Backend.Port)
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
			env.DockerNetwork.Name,
		},
		NetworkAliases: map[string][]string{
			env.DockerNetwork.Name: {"order-api-mock"},
		},
		WaitingFor: wait.ForLog("Stub server is running"),
	}

	stubContainer, err := testcontainers.GenericContainer(env.Ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, "", err
	}

	domainServicePort, err := stubContainer.MappedPort(env.Ctx, port)
	if err != nil {
		return nil, "", err
	}

	return stubContainer, domainServicePort.Port(), nil
}

func StartKafkaMock(env *TestEnvironment) (testcontainers.Container, string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, "", fmt.Errorf("Error getting current directory: %v", err)

	}

	port, err := nat.NewPort("tcp", env.Config.KafkaService.Port)
	if err != nil {
		return nil, "", fmt.Errorf("invalid port number: %w", err)
	}

	fmt.Println("EXposing ports at =====> : ", port.Port())
	fmt.Println("API port at : ", env.Config.KafkaService.ApiPort)

	req := testcontainers.ContainerRequest{
		Name:         "specmatic-kafka",
		Image:        "znsio/specmatic-kafka-trial",
		ExposedPorts: []string{port.Port() + "/tcp", env.Config.KafkaService.ApiPort + "/tcp"},
		Networks: []string{
			env.DockerNetwork.Name,
		},
		NetworkAliases: map[string][]string{
			env.DockerNetwork.Name: {"specmatic-kafka"},
		},
		Env: map[string]string{
			"KAFKA_EXTERNAL_HOST": env.Config.KafkaService.Host,
			"KAFKA_EXTERNAL_PORT": env.Config.KafkaService.Port,
		},
		Cmd: []string{"virtualize", "--mock-server-api-port=" + env.Config.KafkaService.ApiPort},
		Mounts: testcontainers.Mounts(
			testcontainers.BindMount(filepath.Join(pwd, "specmatic.yaml"), "/usr/src/app/specmatic.yaml"),
		),
		WaitingFor: wait.ForLog("Listening on topics: (product-queries)").WithStartupTimeout(2 * time.Minute),
	}

	kafkaC, err := testcontainers.GenericContainer(env.Ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		fmt.Printf("Error starting Kafka mock container: %v", err)
	}

	mappedPort, err := kafkaC.MappedPort(env.Ctx, port)
	if err != nil {
		fmt.Printf("Error getting mapped port for Kafka mock: %v", err)
	}

	fmt.Println("Mapped ports at =====> : ", mappedPort.Port())

	return kafkaC, mappedPort.Port(), nil
}

func StartBFFService(t *testing.T, env *TestEnvironment) (testcontainers.Container, string, error) {

	port, err := nat.NewPort("tcp", env.Config.BFFServer.Port)
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
			"DOMAIN_SERVER_PORT": env.Config.Backend.Port,
			"DOMAIN_SERVER_HOST": "order-api-mock",
			"KAFKA_PORT":         env.Config.KafkaService.Port,
			"KAFKA_HOST":         "specmatic-kafka",
		},
		Networks: []string{
			env.DockerNetwork.Name,
		},
		NetworkAliases: map[string][]string{
			env.DockerNetwork.Name: {"bff-service"},
		},
		ExposedPorts: []string{env.Config.BFFServer.Port + "/tcp"},
		WaitingFor:   wait.ForLog("Starting gRPC server"),
	}

	bffContainer, err := testcontainers.GenericContainer(env.Ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, "", err
	}

	bffPort, err := bffContainer.MappedPort(env.Ctx, port)
	if err != nil {
		return nil, "", err
	}

	return bffContainer, bffPort.Port(), nil
}

func RunTestContainer(env *TestEnvironment) (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current directory: %v", err)
	}

	bffPortInt, err := strconv.Atoi(env.Config.BFFServer.Port)
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
			env.DockerNetwork.Name,
		},
		WaitingFor: wait.ForLog("Passed Tests:"),
	}

	testContainer, err := testcontainers.GenericContainer(env.Ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return "", err
	}
	defer testContainer.Terminate(env.Ctx)

	// Streaming testing logs to terminal
	logReader, err := testContainer.Logs(env.Ctx)
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
