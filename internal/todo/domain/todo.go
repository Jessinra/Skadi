package domain

import "time"

type Todo struct {
	ID          uint64    `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	Text        string    `json:"text"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	UserID      uint64    `json:"userID"`
}
