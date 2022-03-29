package entity

import (
	"time"
)

type Payment struct {
	IDPayment       string    `json:"id_payment"`
	Status          string    `json:"status,omitempty"`
	Order           []Order   `json:"order,omitempty"`
	DeadlinePayment time.Time `json:"deadline_payment,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	DeletedAt       time.Time `json:"deleted_at,omitempty"`
}
