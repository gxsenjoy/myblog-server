package myblog_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMyblogServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MyblogServer Suite")
}
