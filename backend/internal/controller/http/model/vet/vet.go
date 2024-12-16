package vet

import (
	"github.com/kotopesp/sos-kotopes/internal/controller/http/model/pagination"
	"github.com/kotopesp/sos-kotopes/internal/controller/http/model/user"
	"time"
)

type VetsCreate struct {
	UserID             int     `json:"user_id" validate:"required,min=0"`
	IsOrganization     bool    `json:"is_organization" validate:"required"`
	Education          *string `json:"education"`
	OrgName            *string `json:"org_name"`
	Location           *string `json:"location" validate:"required"`
	Price              float64 `json:"price"`
	OrgEmail           *string `json:"org_email"`
	InnNumber          *string `json:"inn_number"`
	IsRemoteConsulting bool    `json:"column:is_remote_consulting"`
	IsInpatient        bool    `json:"is_inpatient"`
	Description        *string `json:"description"`
}

type VetsUpdate struct {
	ID                 int     `json:"id"`
	UserID             int     `json:"user_id"`
	IsOrganization     bool    `json:"is_organization"`
	Education          *string `json:"education"`
	OrgName            *string `json:"org_name"`
	Location           *string `json:"location"`
	Price              float64 `json:"price"`
	OrgEmail           *string `json:"org_email"`
	InnNumber          *string `json:"inn_number"`
	IsRemoteConsulting bool    `json:"column:is_remote_consulting"`
	IsInpatient        bool    `json:"is_inpatient"`
	Description        *string `json:"description"`
}

type VetsResponse struct {
	ID                 int       `json:"id"`
	UserID             int       `json:"user_id"`
	IsOrganization     bool      `json:"is_organization"`
	Education          *string   `json:"education"`
	OrgName            *string   `json:"org_name"`
	Location           *string   `json:"location"`
	Price              float64   `json:"price"`
	OrgEmail           *string   `json:"org_email"`
	InnNumber          *string   `json:"inn_number"`
	IsRemoteConsulting bool      `gorm:"column:is_remote_consulting"`
	IsInpatient        bool      `json:"is_inpatient"`
	Description        *string   `json:"description"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type VetsResponseWithUser struct {
	Vet  VetsResponse
	User user.ResponseUser
}

type VetsResponseWithMeta struct {
	Meta pagination.Pagination  `json:"meta"`
	Data []VetsResponseWithUser `json:"payload"`
}

type GetAllVetsParams struct {
	SortBy    *string  `query:"sort_by"`
	SortOrder *string  `query:"sort_order"`
	Location  *string  `query:"location"`
	MinRating *float64 `query:"min_rating" validate:"omitempty,gte=1,lte=5"`
	MaxRating *float64 `query:"max_rating" validate:"omitempty,gte=1,lte=5"`
	MinPrice  *float64 `query:"min_price" validate:"omitempty,gte=0"`
	MaxPrice  *float64 `query:"max_price" validate:"omitempty,gte=0"`
	Limit     *int     `query:"limit" validate:"omitempty,gt=0"`
	Offset    *int     `query:"offset" validate:"omitempty,gte=0"`
}
