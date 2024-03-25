package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Openig struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Link     string
	Salary   int64
	Remote   bool
}

type OpenigResponse struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deteledAt,omitempty"`
	Role      string    `json:"role"`
	Company   string    `json:"company"`
	Location  string    `json:"location"`
	Link      string    `json:"link"`
	ID        uint      `json:"id"`
	Salary    int64     `json:"salary"`
	Remote    bool      `json:"remote"`
}
