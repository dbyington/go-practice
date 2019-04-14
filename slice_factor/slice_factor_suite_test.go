package slice_factor

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSliceFactor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SliceFactor Suite")
}
