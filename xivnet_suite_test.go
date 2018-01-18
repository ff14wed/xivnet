package xivnet_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestXIVNet(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "XIVNet Suite")
}
