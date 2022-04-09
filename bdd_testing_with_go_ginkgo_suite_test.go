package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
)

func TestBddTestingWithGoGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BddTestingWithGoGinkgo Suite")
}

var agoutiDriver *agouti.WebDriver

var _ = BeforeSuite(func() {
	options := agouti.ChromeOptions("args", []string{
		"--headless",
	})
	// Choose a WebDriver:

	// agoutiDriver = agouti.PhantomJS()
	// agoutiDriver = agouti.Selenium()
	agoutiDriver = agouti.ChromeDriver(options)

	Expect(agoutiDriver.Start()).To(Succeed())
})

var _ = AfterSuite(func() {
	Expect(agoutiDriver.Stop()).To(Succeed())
})
