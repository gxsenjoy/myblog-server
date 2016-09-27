package blogservice_posts_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPosts(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Posts Suite")
}
