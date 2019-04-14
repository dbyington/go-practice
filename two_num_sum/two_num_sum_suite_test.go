package two_num_sum

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTwoNumSum(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TwoNumSum Suite")
}
