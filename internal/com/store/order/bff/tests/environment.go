package tests

import (
	"context"
	"specmatic-order-bff-grpc-go/internal/com/store/order/bff/config"

	"github.com/testcontainers/testcontainers-go"
)

type TestEnvironment struct {
	Ctx                      context.Context
	DomainServiceContainer   testcontainers.Container
	DomainServiceDynamicPort string
	KafkaServiceContainer    testcontainers.Container
	KafkaServiceDynamicPort  string
	KafkaDynamicAPIPort      string
	BffServiceContainer      testcontainers.Container
	BffServiceDynamicPort    string
	DockerNetwork            *testcontainers.DockerNetwork
	Config                   *config.Config
}
