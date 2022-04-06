package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

var _ = Describe("UserLogin", func() {
	var page *agouti.Page

	BeforeEach(func() {
		var err error
		page, err = agoutiDriver.NewPage(agouti.Browser("chrome"))
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})

	It("should manage user authentication", func() {
		By("redirecting the user to the login from the home page", func() {
			Expect(page.Navigate("http://dummy.railflow.io/")).To(Succeed())
			Expect(page).To(HaveURL("http://dummy.railflow.io/signin"))
		})
	})

	It("should manage user authentication", func() {
		By("redirecting the user to the login from the home page", func() {
			Expect(page.Navigate("http://dummy.railflow.io/")).To(Succeed())
			Expect(page).To(HaveURL("http://dummy.railflow.io/signin"))
		})
	})
})
