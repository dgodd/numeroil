package numeroil_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestNumeroil(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Numeroil Suite")
}
