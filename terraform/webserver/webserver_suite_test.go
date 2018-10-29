package webserver_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestWebserver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Webserver Suite")
}
