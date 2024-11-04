package core

import (
	"context"
	"errors"
	"time"
)

type Veterinary struct {
	ID               int        `gorm:"column:id"`
	UserID           int        `gorm:"column:user_id"`
	Entity           string     `gorm:"column:entity"` // individual or legal entity
	Username         string     `gorm:"column:username"`
	Firstname        *string    `gorm:"column:firstname"`
	Lastname         *string    `gorm:"column:lastname"`
	Patronymic       *string    `gorm:"column:patronymic"`
	Education        *string    `gorm:"column:education"`
	OrgName          *string    `gorm:"column:org_name"`
	Location         *string    `gorm:"column:location"`
	OrgEmail         *string    `gorm:"column:org_email"`
	InnNumber        *string    `gorm:"column:inn_number"`
	RemoteConsulting bool       `gorm:"column:remote_consulting"`
	Inpatient        bool       `gorm:"column:inpatient"`
	Description      *string    `gorm:"column:description"`
	IsDeleted        bool       `gorm:"column:is_deleted"`
	CreatedAt        time.Time  `gorm:"column:created_at"`
	UpdatedAt        time.Time  `gorm:"column:updated_at"`
	DeletedAt        *time.Time `gorm:"column:deleted_at"`
}

type UpdateVeterinary struct {
	Username         *string `gorm:"column:username"`
	Firstname        *string `gorm:"column:firstname"`
	Lastname         *string `gorm:"column:lastname"`
	Patronymic       *string `gorm:"column:patronymic"`
	Education        *string `gorm:"column:education"`
	OrgName          *string `gorm:"column:org_name"`
	Location         *string `gorm:"column:location"`
	Price            float64 `gorm:"column:price"`
	OrgEmail         *string `gorm:"column:org_email"`
	InnNumber        *string `gorm:"column:Inn_number"`
	RemoteConsulting bool    `gorm:"column:remote_consuting"`
	Inpatient        bool    `gorm:"column:inpatient"`
	Description      *string `gorm:"column:description"`
}

type VeterinaryDetails struct {
	Veterinary Veterinary
	User       User
}

type GetAllVeterinaryParams struct {
	SortBy    *string
	SortOrder *string
	Location  *string
	MinRating *float64
	MaxRating *float64
	MinPrice  *float64
	MaxPrice  *float64
	Limit     *int
	Offset    *int
}

type VeterinaryStore interface {
	UpdateByID(ctx context.Context, update UpdateVeterinary) (updatedVeterinary Veterinary, err error)
	GetAll(ctx context.Context, params GetAllVeterinaryParams) ([]Veterinary, error)
	GetByID(ctx context.Context, id int) (veterinary Veterinary, err error)
	GetByOrgName(ctx context.Context, orgName string) (veterinary Veterinary, err error)
	Create(ctx context.Context, veterinary Veterinary) error
	DeleteByID(ctx context.Context, id int) error
}

type VeterinaryService interface {
	UpdateByID(ctx context.Context, userID int, veterinary UpdateVeterinary) (VeterinaryDetails, err error)
	GetAll(ctx context.Context, params GetAllVeterinaryParams) ([]VeterinaryDetails, error)
	GetByID(ctx context.Context, id int) (veterinary Veterinary, err error)
	GetByOrgName(ctx context.Context, orgName string) (veterinary Veterinary, err error)
	Create(ctx context.Context, veterinary Veterinary) error
	DeleteByID(ctx context.Context, id int, userID int) error
}

// errors
var (
	ErrNoSuchVeterinary = errors.New("veterinary does not exist")
)

// TableName table name in db for gorm
func (Veterinary) TableName() string {
	return "Veterinaries"
}
