package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGrpcServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GrpcServer Suite")
}
