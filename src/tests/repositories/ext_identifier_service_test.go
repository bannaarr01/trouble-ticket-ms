package repositories_test

import (
	"trouble-ticket-ms/src/models"
	"trouble-ticket-ms/src/repositories"
	"trouble-ticket-ms/src/tests/mocks"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ExtIdentifierRepository", func() {
	var (
		mockRepo *mocks.MockExtIdentifierRepository
		repo     repositories.ExtIdentifierRepository
	)

	BeforeEach(func() {
		mockRepo = &mocks.MockExtIdentifierRepository{}
		repo = mockRepo
	})

	Describe("Save", func() {
		It("should save an external identifier successfully", func() {
			// Arrange
			ticket := models.TroubleTicket{}
			typeModel := models.Type{}

			extIdentifier := models.NewExternalIdentifier(ticket.ID, &models.CreateExternalIdentifierDTO{
				TypeID: typeModel.ID,
			}, models.SetField("CreatedBy", "test-user"))

			mockRepo.SaveFunc = func(extId *models.ExternalIdentifier) error {
				extId.ID = 1 // DB auto-increment
				return nil
			}

			err := repo.Save(&extIdentifier)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(extIdentifier.ID).To(Equal(uint64(1)))
		})
	})

	Describe("FindByTicket", func() {
		It("should find and return external identifiers by ticket ID successfully", func() {
			ticketID := uint64(1)
			expectedExtIdentifiers := []models.ExternalIdentifier{
				{Owner: "Ginkgo", TroubleTicketID: ticketID, TypeID: 1},
			}

			mockRepo.FindByTicketFunc = func(extIdentifiers *[]models.ExternalIdentifier, ticketId uint64) error {
				*extIdentifiers = expectedExtIdentifiers
				return nil
			}

			var extIdentifiers []models.ExternalIdentifier

			err := repo.FindByTicket(&extIdentifiers, ticketID)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(extIdentifiers).Should(HaveLen(1))
			Expect(extIdentifiers[0].TroubleTicketID).To(Equal(ticketID))
		})
	})

	Describe("Remove", func() {
		It("should remove an external identifier successfully", func() {
			extIdentifierID := uint64(1)

			mockRepo.RemoveFunc = func(id uint64) error {
				return nil
			}

			err := repo.Remove(extIdentifierID)

			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})
