package go_practice

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoPractice(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoPractice Suite")
}
