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
	"testing"
	"time"

	"github.com/schollz/progressbar/v3"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var bffDockerfile = "Dockerfile"

func TestIntegration(t *testing.T) {
	ctx := context.Background()

	// Step 1: Start stub container
	printHeader(1, "Starting Domain Stub")
	bar := startProgressBar()
	stubContainer, stubPort, err := startStubContainer(ctx)
	bar.Finish()
	if err != nil {
		t.Fatalf("Could not start stub container: %s", err)
	}

	defer stubContainer.Terminate(ctx)

	// Step 2: Start BFF container
	printHeader(1, "Starting BFF app")
	bffContainer, bffPort, err := startBFFContainer(ctx, stubPort)
	if err != nil {
		t.Fatalf("Could not start BFF container: %s", err)
	}
	defer bffContainer.Terminate(ctx)

	// Step 3: Run test container
	printHeader(1, "Starting tests")
	testLogs, err := runTestContainer(ctx, bffPort)
	if err != nil {
		t.Fatalf("Could not run test container: %s", err)
	}

	// Print test outcomes
	fmt.Println("Test Results:")
	fmt.Println(testLogs)
}

func startStubContainer(ctx context.Context) (testcontainers.Container, string, error) {

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current directory: %v", err)
	}

	req := testcontainers.ContainerRequest{
		Image:        "znsio/specmatic-grpc-trial",
		ExposedPorts: []string{"9000/tcp"},
		Cmd:          []string{"stub"},
		Mounts: testcontainers.Mounts(
			testcontainers.BindMount(filepath.Join(pwd, "specmatic.yaml"), "/usr/src/app/specmatic.yaml"),
		),
		WaitingFor: wait.ForLog("Stub server is running"),
		// WaitingFor:   wait.ForListeningPort("9000/tcp"),
	}

	stubContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, "", err
	}

	stubPort, err := stubContainer.MappedPort(ctx, "9000")
	if err != nil {
		return nil, "", err
	}

	return stubContainer, stubPort.Port(), nil
}

func startBFFContainer(ctx context.Context, stubPort string) (testcontainers.Container, string, error) {
	// Build the image using Docker CLI
	bffImageName := "specmatic-order-bff-grpc-go"
	buildCmd := exec.Command("docker", "build", "-t", bffImageName, "-f", bffDockerfile, ".")
	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr
	if err := buildCmd.Run(); err != nil {
		return nil, "", fmt.Errorf("could not build BFF image: %w", err)
	}

	// Capture both stdout and stderr
	var out bytes.Buffer
	buildCmd.Stdout = &out
	buildCmd.Stderr = &out

	if err := buildCmd.Run(); err != nil {
		fmt.Println("Docker build output:")
		fmt.Println(out.String())
		return nil, "", fmt.Errorf("could not build BFF image: %w", err)
	}

	// Using the built image
	req := testcontainers.ContainerRequest{
		Image: bffImageName,
		Env: map[string]string{
			"STUB_SERVER_PORT": stubPort,
		},
		ExposedPorts: []string{"8080/tcp"},
		WaitingFor:   wait.ForListeningPort("8080/tcp"),
	}

	fmt.Println("Container created")

	bffContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, "", err
	}

	bffPort, err := bffContainer.MappedPort(ctx, "8080")
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

	req := testcontainers.ContainerRequest{
		Image: "znsio/specmatic-grpc-trial",
		Cmd:   []string{"test", "--port=:" + bffPort},
		Mounts: testcontainers.Mounts(
			testcontainers.BindMount(filepath.Join(pwd, "specmatic.yaml"), "/specmatic.yaml"),
		),
		Env: map[string]string{
			"BACKEND_PORT": bffPort,
		},
		WaitingFor: wait.ForLog("Tests completed"),
		// WaitingFor:   wait.ForListeningPort("9000/tcp"),
	}

	testContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return "", err
	}
	defer testContainer.Terminate(ctx)

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

func printHeader(stepNum int, title string) {
	fmt.Println("")
	fmt.Printf("======== STEP %d =========\n", stepNum)
	fmt.Println(title)
	fmt.Println("=========================")
	fmt.Println("")
}

func startProgressBar() *progressbar.ProgressBar {
	// Create a progress bar
	bar := progressbar.NewOptions(-1, // Use -1 for an indefinite progress bar
		progressbar.OptionSetDescription("..."),
		progressbar.OptionSpinnerType(22),
		progressbar.OptionFullWidth(),
		progressbar.OptionClearOnFinish(),
	)

	// Update the progress bar in a separate goroutine
	go func() {
		for {
			bar.Add(1)
			time.Sleep(100 * time.Millisecond) // Adjust the speed as necessary
		}
	}()

	return bar
}
