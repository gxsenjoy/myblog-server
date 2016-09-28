package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGrpcGateway(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GrpcGateway Suite")
}
