package services_test

import (
	"errors"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"net/http/httptest"
	"trouble-ticket-ms/src/models"
	"trouble-ticket-ms/src/services"
	"trouble-ticket-ms/src/tests/mocks"
)

var _ = Describe("ExtIdentifierService", func() {
	var (
		mockRepository *mocks.MockExtIdentifierRepository
		service        services.ExtIdentifierService
		recorder       *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		mockRepository = &mocks.MockExtIdentifierRepository{}
		service = services.NewExtIdentifierService(mockRepository, services.AppDependencies{})
		recorder = httptest.NewRecorder()
		_, _ = gin.CreateTestContext(recorder)
	})

	Describe("Create", func() {
		It("should create and return the external identifier DTO successfully", func() {
			mockRepository.SaveFunc = func(externalIdentifier *models.ExternalIdentifier) error {
				return nil
			}

			authUserName := "testUser"
			troubleTicketId := uint64(1)
			createDTO := &models.CreateExternalIdentifierDTO{}

			dto, err := service.Create(authUserName, troubleTicketId, createDTO)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(dto).ShouldNot(BeNil())
		})

		It("should return an error if saving the external identifier fails", func() {
			mockRepository.SaveFunc = func(externalIdentifier *models.ExternalIdentifier) error {
				return errors.New("save error")
			}

			authUserName := "testUser"
			troubleTicketId := uint64(1)
			createDTO := &models.CreateExternalIdentifierDTO{}

			dto, err := service.Create(authUserName, troubleTicketId, createDTO)
			Expect(err).Should(HaveOccurred())
			Expect(dto).Should(BeNil())
		})
	})

	Describe("FindByTicket", func() {
		It("should find and return external identifiers by ticket ID successfully", func() {
			mockIdentifiers := []models.ExternalIdentifier{
				{Owner: "GinkGo", TroubleTicketID: 1},
			}
			mockRepository.FindByTicketFunc = func(extIdentifiers *[]models.ExternalIdentifier, ticketID uint64) error {
				*extIdentifiers = mockIdentifiers
				return nil
			}

			dtos, err := service.FindByTicket(1)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(dtos).Should(HaveLen(len(mockIdentifiers)))
		})

		It("should return an error if finding external identifiers fails", func() {
			mockRepository.FindByTicketFunc = func(extIdentifiers *[]models.ExternalIdentifier, ticketID uint64) error {
				return errors.New("find error")
			}

			dtos, err := service.FindByTicket(1)
			Expect(err).Should(HaveOccurred())
			Expect(dtos).Should(BeNil())
		})
	})

	Describe("Remove", func() {
		It("should remove the external identifier successfully", func() {
			mockRepository.RemoveFunc = func(extIdentifierID uint64) error {
				return nil
			}

			err := service.Remove(1)
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("should return an error if removing the external identifier fails", func() {
			mockRepository.RemoveFunc = func(extIdentifierID uint64) error {
				return errors.New("remove error")
			}

			err := service.Remove(1)
			Expect(err).Should(HaveOccurred())
		})
	})
})
