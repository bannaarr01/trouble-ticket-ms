package repositories

import (
	"errors"
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
	FindAll(*models.Claims, *models.GetTroubleTicketQuery, *[]models.TroubleTicket) (int64, error)
	FindOne(uint64, *models.Claims) (*models.TroubleTicket, error)
	Remove(uint64, *models.Claims) error
	Update(*models.TroubleTicket, string) error
	FindAllFilter(*models.Filters) error
}

type troubleTicketRepository struct {
	db *db.DB
}

func (t *troubleTicketRepository) FindAll(
	authUser *models.Claims,
	query *models.GetTroubleTicketQuery,
	troubleTickets *[]models.TroubleTicket,
) (int64, error) {
	// Base query
	baseQuery := t.db.DB.
		//Debug().
		Model(&models.TroubleTicket{}).
		Joins("LEFT JOIN external_identifiers e ON e.trouble_ticket_id = trouble_tickets.id").
		Joins("LEFT JOIN related_parties r ON r.trouble_ticket_id = trouble_tickets.id").
		Joins("LEFT JOIN related_entities l on l.trouble_ticket_id = trouble_tickets.id").
		Joins("LEFT JOIN notes n on n.trouble_ticket_id = trouble_tickets.id")

	// Build WHERE clause
	whereClause := "trouble_tickets.deleted_at IS NULL"
	args := []interface{}{}

	if !isAdmin(authUser) {
		whereClause += " AND trouble_tickets.created_by = ?"
		args = append(args, authUser.PreferredUsername)
	}

	// Add filters
	if query.Ref != nil {
		whereClause += " AND trouble_tickets.ref = ?"
		args = append(args, utils.DerefPtr(query.Ref))
	}
	if query.Name != nil {
		whereClause += " AND trouble_tickets.name LIKE ?"
		args = append(args, "%"+utils.DerefPtr(query.Name)+"%")
	}
	if query.TypeID != nil {
		whereClause += " AND trouble_tickets.type_id = ?"
		args = append(args, utils.DerefPtr(query.TypeID))
	}
	if query.StatusID != nil {
		whereClause += " AND trouble_tickets.status_id = ?"
		args = append(args, utils.DerefPtr(query.StatusID))
	}
	if query.ChannelID != nil {
		whereClause += " AND trouble_tickets.channel_id = ?"
		args = append(args, utils.DerefPtr(query.ChannelID))
	}
	if query.SeverityID != nil {
		whereClause += " AND trouble_tickets.severity_id = ?"
		args = append(args, utils.DerefPtr(query.SeverityID))
	}
	if query.PriorityID != nil {
		whereClause += " AND trouble_tickets.priority_id = ?"
		args = append(args, utils.DerefPtr(query.PriorityID))
	}
	if query.ExternalIDOwner != nil {
		whereClause += " AND e.owner = ?"
		args = append(args, utils.DerefPtr(query.ExternalIDOwner))
	}
	if query.RelatedPartyEmail != nil {
		whereClause += " AND r.email = ?"
		args = append(args, utils.DerefPtr(query.RelatedPartyEmail))
	}
	if query.RelatedEntityRef != nil {
		whereClause += " AND l.ref = ?"
		args = append(args, utils.DerefPtr(query.RelatedEntityRef))
	}
	if query.NoteAuthor != nil {
		whereClause += " AND n.author = ?"
		args = append(args, utils.DerefPtr(query.NoteAuthor))
	}

	// Apply WHERE clause
	baseQuery = baseQuery.Where(whereClause, args...).Group("trouble_tickets.id")

	// Get total count Bf Applying Limit & Offset
	var totalCount int64
	if err := baseQuery.Count(&totalCount).Error; err != nil {
		return 0, err
	}

	// Apply ordering, limit, and offset
	result := baseQuery.
		Order("trouble_tickets.created_at DESC").
		Limit(int(query.Limit)).
		Offset(int(query.Offset)).
		Preload("Type").
		Preload("Status").
		Preload("Channel").
		Preload("Severity").
		Preload("Priority").
		Find(&troubleTickets)

	if result.Error != nil {
		return 0, result.Error
	}

	return totalCount, nil
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

		if err = utils.CheckRelatedRecordExists(tx, &models.Type{}, troubleTicket.TypeID, "id"); err != nil {
			return err
		}

		if err = utils.CheckRelatedRecordExists(tx, &models.Channel{}, troubleTicket.ChannelID, "id"); err != nil {
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

// FindOne You can only see the ticket if it was created by you, even if u have the ticket id Else you are an admin,
func (t *troubleTicketRepository) FindOne(ticketID uint64, authUser *models.Claims) (*models.TroubleTicket, error) {
	var troubleTicket models.TroubleTicket

	condition, args := buildCondition(ticketID, authUser, isAdmin(authUser))

	err := preloadAssociations(t.db.DB).Where(condition, args...).First(&troubleTicket).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("record does not exist")
	}
	return &troubleTicket, err
}

func (t *troubleTicketRepository) Remove(ticketID uint64, authUser *models.Claims) error {
	condition, args := buildCondition(ticketID, authUser, isAdmin(authUser))

	updateData := map[string]interface{}{
		"deleted_at": time.Now(),
		"deleted_by": authUser.PreferredUsername,
	}

	result := t.db.Model(&models.TroubleTicket{}).Where(condition, args...).Updates(updateData)

	if result.RowsAffected == 0 {
		return fmt.Errorf("ticket with ID:%v does not exist", ticketID)
	}

	if result.Error != nil {
		return result.Error
	}
	return nil
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

func isAdmin(authUser *models.Claims) bool {
	adminRoles := []string{"admin", "super_admin"}

	for _, role := range authUser.RealmAccess.Roles {
		if utils.Contains(adminRoles, role) {
			return true
		}
	}
	return false
}

func buildCondition(ticketID uint64, authUser *models.Claims, isAdmin bool) (string, []interface{}) {
	condition := "id = ? AND deleted_at IS NULL"
	args := []interface{}{ticketID}

	if !isAdmin {
		condition += " AND created_by = ?"
		args = append(args, authUser.PreferredUsername)
	}

	return condition, args
}
