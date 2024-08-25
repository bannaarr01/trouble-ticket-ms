package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"trouble-ticket-ms/src/controllers"
	"trouble-ticket-ms/src/models"
	"trouble-ticket-ms/src/tests/mocks"
)

var _ = Describe("AttachmentController", func() {
	var (
		mockService *mocks.MockAttachmentService
		controller  controllers.AttachmentController
		context     *gin.Context
		recorder    *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		mockService = &mocks.MockAttachmentService{}
		controller = controllers.NewAttachmentController(mockService)
		recorder = httptest.NewRecorder()
		context, _ = gin.CreateTestContext(recorder)
	})

	Describe("Remove", func() {
		BeforeEach(func() {
			context.Params = gin.Params{gin.Param{Key: "ref", Value: "test-ref"}}
		})

		It("should remove the attachment successfully", func() {
			mockService.RemoveFunc = func(ref string) error {
				return nil
			}

			controller.Remove(context)

			Expect(recorder.Code).To(Equal(http.StatusOK))
			var response map[string]string
			json.Unmarshal(recorder.Body.Bytes(), &response)
			Expect(response["message"]).To(Equal("ok"))
		})

		It("should return an error if removal fails", func() {
			mockService.RemoveFunc = func(ref string) error {
				return errors.New("removal error")
			}

			controller.Remove(context)

			Expect(recorder.Code).To(Equal(http.StatusInternalServerError))
			var response map[string]string
			json.Unmarshal(recorder.Body.Bytes(), &response)
			Expect(response["message"]).To(Equal("removal error"))
		})
	})

	Describe("FindByTicket", func() {
		BeforeEach(func() {
			context.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
		})

		It("should find attachments by ticket successfully", func() {
			mockAttachments := []models.AttachmentDTO{{Ref: "test-ref"}}
			mockService.FindByTicketFunc = func(id uint64) ([]models.AttachmentDTO, error) {
				return mockAttachments, nil
			}

			controller.FindByTicket(context)

			Expect(recorder.Code).To(Equal(http.StatusOK))
			var response map[string]interface{}
			json.Unmarshal(recorder.Body.Bytes(), &response)
			Expect(response["data"]).To(HaveLen(1))
		})

		It("should return an error if finding attachments fails", func() {
			mockService.FindByTicketFunc = func(id uint64) ([]models.AttachmentDTO, error) {
				return nil, errors.New("find error")
			}

			controller.FindByTicket(context)

			Expect(recorder.Code).To(Equal(http.StatusInternalServerError))
			var response map[string]string
			json.Unmarshal(recorder.Body.Bytes(), &response)
			Expect(response["message"]).To(ContainSubstring("find error"))
		})
	})

	Describe("FindOne", func() {
		BeforeEach(func() {
			context.Params = gin.Params{gin.Param{Key: "ref", Value: "test-ref"}}
		})

		It("should find one attachment successfully", func() {
			mockAttachment := models.AttachmentDTO{Ref: "test-ref"}
			mockService.FindOneFunc = func(ref string) (*models.AttachmentDTO, error) {
				return &mockAttachment, nil
			}

			controller.FindOne(context)

			Expect(recorder.Code).To(Equal(http.StatusOK))
			var response map[string]interface{}
			json.Unmarshal(recorder.Body.Bytes(), &response)
			data, ok := response["data"].(map[string]interface{})
			Expect(ok).To(BeTrue())
			Expect(data).To(HaveKeyWithValue("ref", "test-ref"))
		})

		It("should return an error if finding one attachment fails", func() {
			mockService.FindOneFunc = func(ref string) (*models.AttachmentDTO, error) {
				return nil, errors.New("find one error")
			}

			controller.FindOne(context)

			Expect(recorder.Code).To(Equal(http.StatusInternalServerError))
			var response map[string]string
			json.Unmarshal(recorder.Body.Bytes(), &response)
			Expect(response["message"]).To(Equal("find one error"))
		})
	})

	Describe("Upload", func() {
		BeforeEach(func() {
			context.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
			context.Set("user", &models.Claims{})

			// Mock file upload
			buf := new(bytes.Buffer)
			mw := multipart.NewWriter(buf)
			w, _ := mw.CreateFormFile("file", "test.txt")
			w.Write([]byte("test content"))
			mw.Close()

			req, _ := http.NewRequest("POST", "/", buf)
			req.Header.Set("Content-Type", mw.FormDataContentType())

			context.Request = req
		})

		It("should upload an attachment successfully", func() {
			file, header, _ := context.Request.FormFile("file")
			context.Set("file", file)
			context.Set("fileHeader", header)

			mockAttachment := models.AttachmentDTO{Ref: "uploaded-ref"}
			mockService.SaveFunc = func(id uint64, user *models.Claims, file *multipart.File, header *multipart.FileHeader) (*models.AttachmentDTO, error) {
				return &mockAttachment, nil
			}

			controller.Upload(context)

			Expect(recorder.Code).To(Equal(http.StatusOK))
			var response map[string]interface{}
			json.Unmarshal(recorder.Body.Bytes(), &response)
			data, ok := response["data"].(map[string]interface{})
			Expect(ok).To(BeTrue())
			Expect(data).To(HaveKeyWithValue("ref", "uploaded-ref"))
			Expect(data).To(HaveKey("mime_type"))
			Expect(data).To(HaveKey("original_name"))
			Expect(data).To(HaveKey("size"))
			Expect(data).To(HaveKey("name"))
			Expect(data).To(HaveKey("description"))
			Expect(data).To(HaveKey("created_by"))
			Expect(data).To(HaveKey("type"))
			Expect(data).To(HaveKey("path"))
			Expect(data).To(HaveKey("url"))
			Expect(data).To(HaveKey("created_at"))
		})

		It("should return an error if upload fails", func() {
			file, header, _ := context.Request.FormFile("file")
			context.Set("file", file)
			context.Set("fileHeader", header)

			mockService.SaveFunc = func(id uint64, user *models.Claims, file *multipart.File, header *multipart.FileHeader) (*models.AttachmentDTO, error) {
				return nil, errors.New("upload error")
			}

			controller.Upload(context)

			Expect(recorder.Code).To(Equal(http.StatusInternalServerError))
			var response map[string]string
			json.Unmarshal(recorder.Body.Bytes(), &response)
			Expect(response["message"]).To(Equal("upload error"))
		})
	})
})
