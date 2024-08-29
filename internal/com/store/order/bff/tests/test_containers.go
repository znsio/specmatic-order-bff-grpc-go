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
	"github.com/go-resty/resty/v2"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/tidwall/gjson"
)

var specmaticGrpcImage = "znsio/specmatic-grpc-trial:0.0.7"

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
		Image:        specmaticGrpcImage,
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

	mappedApiPort, err := kafkaC.MappedPort(env.Ctx, nat.Port(env.Config.KafkaService.ApiPort))
	if err != nil {
		fmt.Printf("Error getting API server port: %v", err)
	} else {
		env.KafkaDynamicAPIPort = mappedApiPort.Port()
	}

	// Get the host IP
	kafkaAPIHost, err := kafkaC.Host(env.Ctx)
	if err != nil {
		return nil, "", fmt.Errorf("Error getting host IP: %v", err)
	}
	env.KafkaAPIHost = kafkaAPIHost

	if err := SetKafkaExpectations(env); err != nil {
		fmt.Printf("failed to set Kafka expectations ==== : %v", err)
	}

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
		Image: specmaticGrpcImage,
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

func SetKafkaExpectations(env *TestEnvironment) error {
	client := resty.New()

	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(fmt.Sprintf(`[{"topic": "product-queries", "count": %d}]`, env.ExpectedMessageCount)).
		Post(fmt.Sprintf("http://%s:%s/_expectations", env.KafkaAPIHost, env.KafkaDynamicAPIPort))

	return err
}

func VerifyKafkaExpectations(env *TestEnvironment) error {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		Post(fmt.Sprintf("http://%s:%s/_expectations/verifications", env.KafkaAPIHost, env.KafkaDynamicAPIPort))

	if err != nil {
		return err
	}

	if !gjson.GetBytes(resp.Body(), "success").Bool() {
		return fmt.Errorf("Kafka mock verification failed: %v", gjson.GetBytes(resp.Body(), "errors").Array())
	}

	fmt.Println("Kafka mock expectations were met successfully.")
	return nil
}
