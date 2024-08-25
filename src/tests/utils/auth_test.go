package utils_test

import (
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"trouble-ticket-ms/src/utils"
)

var _ = Describe("Utils", func() {
	Describe("ExtractAuthTokenFromHeader", func() {
		var ctx *gin.Context
		var recorder *httptest.ResponseRecorder

		BeforeEach(func() {
			recorder = httptest.NewRecorder()
			ctx, _ = gin.CreateTestContext(recorder)
		})

		It("should return an error if no authorization header is provided", func() {
			// To Ensure the request has no Authorization header
			ctx.Request = httptest.NewRequest("GET", "/", nil)

			token, err := utils.ExtractAuthTokenFromHeader(ctx)
			// Check if token is nil and error is present
			Expect(token).To(BeNil())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("no authorization header provided"))
			Expect(recorder.Code).To(Equal(http.StatusUnauthorized))
		})

		It("should return an error if authorization header format is invalid", func() {
			ctx.Request = httptest.NewRequest("GET", "/", nil)
			ctx.Request.Header.Set("Authorization", "InvalidToken")
			token, err := utils.ExtractAuthTokenFromHeader(ctx)
			Expect(token).To(BeNil())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("invalid authorization header format"))
			Expect(recorder.Code).To(Equal(http.StatusUnauthorized))
		})

		It("should return the token if authorization header format is valid (Bearer case)", func() {
			expectedToken := "mytoken"
			ctx.Request = httptest.NewRequest("GET", "/", nil)
			ctx.Request.Header.Set("Authorization", "Bearer "+expectedToken)
			token, err := utils.ExtractAuthTokenFromHeader(ctx)
			Expect(token).NotTo(BeNil())
			Expect(*token).To(Equal(expectedToken))
			Expect(err).ToNot(HaveOccurred())
		})

		It("should return the token if authorization header format is valid (bearer case)", func() {
			expectedToken := "mytoken"
			ctx.Request = httptest.NewRequest("GET", "/", nil)
			ctx.Request.Header.Set("Authorization", "bearer "+expectedToken)
			token, err := utils.ExtractAuthTokenFromHeader(ctx)
			Expect(token).NotTo(BeNil())
			Expect(*token).To(Equal(expectedToken))
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Describe("Contains", func() {
		It("should return true if the element is in the slice", func() {
			slice := []string{"agent", "sysTest", "boom"}
			Expect(utils.Contains(slice, "sysTest")).To(BeTrue())
		})

		It("should return false if the element is not in the slice", func() {
			slice := []string{"agent", "sysTest", "boom"}
			Expect(utils.Contains(slice, "orange")).To(BeFalse())
		})

		It("should return false if the slice is empty", func() {
			slice := []string{}
			Expect(utils.Contains(slice, "sysTest")).To(BeFalse())
		})
	})
})
