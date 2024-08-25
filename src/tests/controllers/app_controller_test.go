package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"trouble-ticket-ms/src/controllers"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("AppController", func() {
	var (
		appController controllers.AppController
		w             *httptest.ResponseRecorder
		c             *gin.Context
	)

	BeforeEach(func() {
		appController = controllers.NewAppController()
		w = httptest.NewRecorder()
		c, _ = gin.CreateTestContext(w)
	})

	Describe("Index", func() {
		It("should return a welcome message", func() {
			appController.Index(c)

			Expect(w.Code).To(Equal(http.StatusOK))

			var response map[string]string
			err := json.Unmarshal(w.Body.Bytes(), &response)
			Expect(err).To(BeNil())
			Expect(response["message"]).To(Equal("Trouble Ticket API"))
		})
	})
})
