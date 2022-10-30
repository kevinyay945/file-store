package pcloud

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPcloud(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pcloud Suite")
}
