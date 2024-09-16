package repository

import (
	"log"
	"tender-service/models"
	"time"

	"github.com/google/uuid"
)

func CreateTender(tender *models.Tender) error {
	query := `INSERT INTO tender (id, name, description, servicetype, status, organizationid, version, createdat) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	tender.ID = uuid.New().String()
	tender.Status = "Created"
	tender.CreatedAt = time.Now().Format(time.RFC3339)
	tender.Version = 1

	_, err := GetDB().Exec(query, tender.ID, tender.Name, tender.Description, tender.ServiceType, tender.Status, tender.OrganizationID, tender.Version, tender.CreatedAt)
	if err != nil {
		log.Printf("Ошибка при создании тендера: %v", err)
		return err
	}

	return nil
}

func GetTenders(limit, offset int, serviceTypes []string) ([]models.Tender, error) {
	query := "SELECT id, name, description, service_type, status, organization_id, version, created_at FROM tenders"
	args := []interface{}{}

	if len(serviceTypes) > 0 {
		query += " WHERE service_type IN ("
		for i := range serviceTypes {
			if i > 0 {
				query += ","
			}
			query += "?"
			args = append(args, serviceTypes[i])
		}
		query += ")"
	}

	query += " ORDER BY name LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tenders []models.Tender
	for rows.Next() {
		var tender models.Tender
		if err := rows.Scan(&tender.ID, &tender.Name, &tender.Description, &tender.ServiceType, &tender.Status, &tender.OrganizationID, &tender.Version, &tender.CreatedAt); err != nil {
			return nil, err
		}
		tenders = append(tenders, tender)
	}

	return tenders, nil
}

func UpdateTenderStatus(tenderID, status string) error {
	query := `UPDATE tenders SET status = $1 WHERE id = $2`

	_, err := GetDB().Exec(query, status, tenderID)
	if err != nil {
		return err
	}

	return nil
}

func GetTendersByUsername(username string, limit, offset int) ([]models.Tender, error) {
	query := `
		SELECT tender.id, tender.name, tender.description, tender.service_type, tender.status, tender.organization_id, tender.version, tender.created_at
		FROM tender
		WHERE tender.organization_id = (
			SELECT organization_responsible.organization_id
			FROM organization_responsible
			JOIN employee ON organization_responsible.employee_id = employee.id
			WHERE employee.username = $1
		)
		ORDER BY tender.name
		LIMIT $2 OFFSET $3
	`

	rows, err := GetDB().Query(query, username, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tenders []models.Tender
	for rows.Next() {
		var tender models.Tender
		err := rows.Scan(&tender.ID, &tender.Name, &tender.Description, &tender.ServiceType, &tender.Status, &tender.OrganizationID, &tender.Version, &tender.CreatedAt)
		if err != nil {
			return nil, err
		}
		tenders = append(tenders, tender)
	}

	return tenders, nil
}
