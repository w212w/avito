package repository

import (
	"log"
	"tender-service/models"
	"time"

	"github.com/google/uuid"
)

func CreateBid(bid *models.Bid) error {
	query := `INSERT INTO bids (id, name, description, status, tender_id, author_id, author_type, version, created_at) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	bid.ID = uuid.New().String()
	bid.Status = "Created"
	bid.CreatedAt = time.Now().Format(time.RFC3339)
	bid.Version = 1

	_, err := GetDB().Exec(query, bid.ID, bid.Name, bid.Description, bid.Status, bid.TenderID, bid.AuthorID, bid.AuthorType, bid.Version, bid.CreatedAt)
	if err != nil {
		log.Printf("Ошибка при создании предложения: %v", err)
		return err
	}

	return nil
}

func CheckTenderExists(tenderID string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM tenders WHERE id = $1)`
	err := GetDB().QueryRow(query, tenderID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
