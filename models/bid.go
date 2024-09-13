package models

type Bid struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	TenderID    string `json:"tenderId"`
	AuthorID    string `json:"authorId"`
	Version     int    `json:"version"`
	CreatedAt   string `json:"createdAt"`
}
