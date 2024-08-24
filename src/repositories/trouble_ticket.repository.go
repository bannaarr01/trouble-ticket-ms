package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"math/rand"
	"time"
	"trouble-ticket-ms/src/db"
	"trouble-ticket-ms/src/enums"
	"trouble-ticket-ms/src/models"
	"trouble-ticket-ms/src/utils"
)

type TroubleTicketRepository interface {
	Create(string, *models.CreateTroubleTicketDTO) (*models.TroubleTicket, error)
	FindAll(*[]models.TroubleTicket) error
	FindOne(*models.TroubleTicket, string) error
	Remove(*models.TroubleTicket) error
	Update(*models.TroubleTicket, string) error
	FindAllFilter(*models.Filters) error
}

type troubleTicketRepository struct {
	db *db.DB
}

type DeterminedPriority struct {
	SeverityID uint64
	PriorityID uint64
}

func (t *troubleTicketRepository) Create(authUserName string, cDto *models.CreateTroubleTicketDTO) (*models.TroubleTicket, error) {
	var troubleTicket models.TroubleTicket

	err := t.db.Transaction(func(tx *gorm.DB) error {
		reference, err := generateUniqueReference(tx)
		if err != nil {
			return err
		}

		dP := determinePriority(cDto.ChannelID, cDto.TypeID)
		next3Days := time.Now().Add(3 * 24 * time.Hour)
		requestedResDate := cDto.RequestedResolutionDate

		troubleTicket = models.NewTroubleTicket(
			*cDto,
			reference,
			enums.AcknowledgedStatus,
			dP.PriorityID,
			dP.SeverityID,
			requestedResDate,
			&next3Days,
			models.SetField("CreatedBy", authUserName),
		)

		if err = preCheckRelatedRecords(tx, &troubleTicket); err != nil {
			return err
		}

		if err = tx.Create(&troubleTicket).Error; err != nil {
			return err
		}

		// Log status change
		statusChange := models.NewStatusChange(
			"trouble ticket created",
			troubleTicket.StatusID,
			troubleTicket.ID,
			models.SetField("CreatedBy", troubleTicket.CreatedBy),
		)

		if err = tx.Create(&statusChange).Error; err != nil {
			return err
		}

		// Preload the just created trouble ticket
		if err = preloadAssociations(tx).
			First(&troubleTicket, "id = ?", troubleTicket.ID).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &troubleTicket, nil
}

func (t *troubleTicketRepository) FindAllFilter(allFilter *models.Filters) error {
	if err := t.db.Find(&allFilter.Types).Error; err != nil {
		return err
	}
	if err := t.db.Find(&allFilter.Statuses).Error; err != nil {
		return err
	}
	if err := t.db.Find(&allFilter.Severities).Error; err != nil {
		return err
	}
	if err := t.db.Find(&allFilter.Channels).Error; err != nil {
		return err
	}
	if err := t.db.Find(&allFilter.Priorities).Error; err != nil {
		return err
	}
	if err := t.db.Find(&allFilter.Roles).Error; err != nil {
		return err
	}
	return nil
}

func (t *troubleTicketRepository) FindAll(troubleTickets *[]models.TroubleTicket) error {
	if err := preloadAssociations(t.db.DB).Find(troubleTickets).Error; err != nil {
		return err
	}

	return nil
}

func (t *troubleTicketRepository) FindOne(troubleTicket *models.TroubleTicket, id string) error {
	//TODO implement me
	panic("implement me")
}

func (t *troubleTicketRepository) Remove(troubleTicket *models.TroubleTicket) error {
	//TODO implement me
	panic("implement me")
}

func (t *troubleTicketRepository) Update(troubleTicket *models.TroubleTicket, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewTroubleTicketRepository(db *db.DB) TroubleTicketRepository {
	return &troubleTicketRepository{db}
}

func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = enums.Charset[rand.Intn(len(enums.Charset))]
	}
	return string(b)
}

func preloadAssociations(db *gorm.DB) *gorm.DB {
	return db.
		Preload("Type").
		Preload("Status").
		Preload("Channel").
		Preload("Severity").
		Preload("Priority").
		Preload("RelatedEntities").
		Preload("Attachments").
		Preload("Notes").
		Preload("StatusChanges", utils.NestedPreload("Status")).
		Preload("RelatedParties", utils.NestedPreload("Party", "Role")).
		Preload("ExternalIdentifiers", utils.NestedPreload("Type"))
}

func generateUniqueReference(db *gorm.DB) (string, error) {
	for {
		timestamp := time.Now().Unix()
		randomStr := generateRandomString(enums.Length)
		reference := fmt.Sprintf("TBT-%d-%s", timestamp, randomStr)

		var count int64
		if err := db.Model(&models.TroubleTicket{}).Where("ref = ?", reference).
			Count(&count).Error; err != nil {
			return "", err
		}

		if count == 0 {
			return reference, nil
		}
	}
}

// determinePriority can change logic
func determinePriority(channelId, typeID uint64) *DeterminedPriority {
	var determinedResult DeterminedPriority

	switch channelId {
	case enums.OperationChannel, enums.SalesChannel, enums.BillingChannel, enums.FinanceChannel:
		if typeID == enums.IncidentType {
			determinedResult = DeterminedPriority{enums.CriticalSeverity, enums.CriticalPriority}
		} else {
			determinedResult = DeterminedPriority{enums.MajorSeverity, enums.HighPriority}
		}
	case enums.SupportChannel:
		if typeID == enums.IncidentType {
			determinedResult = DeterminedPriority{enums.MajorSeverity, enums.CriticalPriority}
		} else {
			determinedResult = DeterminedPriority{enums.MajorSeverity, enums.HighPriority}
		}
	case enums.HRChannel:
		if typeID == enums.ComplainType {
			determinedResult = DeterminedPriority{enums.MajorSeverity, enums.HighPriority}
		} else {
			determinedResult = DeterminedPriority{enums.MinorSeverity, enums.MediumPriority}
		}
	default:
		switch typeID {
		case enums.RequestType:
			determinedResult = DeterminedPriority{enums.MajorSeverity, enums.MediumPriority}
		default:
			determinedResult = DeterminedPriority{enums.MinorSeverity, enums.LowPriority}
		}
	}

	return &determinedResult
}

func preCheckRelatedRecords(tx *gorm.DB, troubleTicket *models.TroubleTicket) error {
	var count int64
	if err := tx.Model(&models.Type{}).Where("id = ?", troubleTicket.TypeID).
		Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("type with ID %d does not exist", troubleTicket.TypeID)
	}

	if err := tx.Model(&models.Channel{}).Where("id = ?", troubleTicket.ChannelID).
		Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("channel with ID %d does not exist", troubleTicket.ChannelID)
	}

	return nil
}
