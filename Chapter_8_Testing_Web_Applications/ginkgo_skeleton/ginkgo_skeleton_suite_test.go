package ginkgo_skeleton_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGinkgoSkeleton(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GinkgoSkeleton Suite")
}
