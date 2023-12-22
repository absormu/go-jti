package entity

type PhoneData struct {
	ID         int64    `json:"id"`
	Number     string   `json:"number"`
	Type       int      `json:"type"`
	Provider   Provider `json:"provider"`
	IsDeleted  int      `json:"is_deleted"`
	CreatedAt  string   `json:"created_at"`
	CreatedBy  string   `json:"created_by"`
	ModifiedAt string   `json:"modified_at"`
	ModifiedBy string   `json:"modified_by"`
}
