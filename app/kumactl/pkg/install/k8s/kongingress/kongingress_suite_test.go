package kongingress_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestKIC(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kong Ingress Suite")
}
