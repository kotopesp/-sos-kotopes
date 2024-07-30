package role

import (
	"context"
	"errors"
	"github.com/kotopesp/sos-kotopes/internal/core"
	"github.com/kotopesp/sos-kotopes/pkg/postgres"
	"gorm.io/gorm"
	"time"
)

type store struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) core.RoleStore {
	return &store{pg}
}

const Seeker = "seeker"
const Keeper = "keeper"
const Vet = "vet"

func (s *store) GetUserRoles(ctx context.Context, id int) (roles map[string]core.Role, err error) {
	roles = make(map[string]core.Role)

	tableNames := []string{"seekers", "keepers", "vets"}
	roleNames := []string{Seeker, Keeper, Vet}

	for i, name := range tableNames {
		var role core.Role
		key := roleNames[i]
		if err = s.DB.WithContext(ctx).Table(name).
			Where("user_id = ?", id).
			First(&role).Error; err == nil {
			roles[key] = role
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	return roles, nil
}

func (s *store) GiveRoleToUser(ctx context.Context, id int, givenRole core.GivenRole) (addedRole core.Role, roleName string, err error) {

	tx := s.DB.Begin()
	if tx.Error != nil {
		return core.Role{}, "", tx.Error
	}

	defer func() {
		if r := recover(); r != nil || err != nil {
			tx.Rollback()
		}
	}()

	var user core.User
	if err = tx.Table("users").First(&user, "id = ?", id).Error; err != nil {
		return core.Role{}, "", core.ErrNoSuchUser
	}

	now := time.Now()

	switch givenRole.Name {
	case Seeker:
		seeker := core.Role{
			UserID:      id,
			Description: givenRole.Description,
			CreatedAt:   now,
			UpdatedAt:   now,
		}
		if err := tx.Table("seekers").Create(&seeker).Error; err != nil {
			return core.Role{}, "", err
		} else {
			roleName = Seeker
			addedRole = seeker
		}
	case Keeper:
		keeper := core.Role{
			UserID:      id,
			Description: givenRole.Description,
			CreatedAt:   now,
			UpdatedAt:   now,
		}
		if err := tx.Table("keepers").Create(&keeper).Error; err != nil {
			return core.Role{}, "", err
		} else {
			roleName = Keeper
			addedRole = keeper
		}
	case Vet:
		vet := core.Role{
			UserID:      id,
			Description: givenRole.Description,
			CreatedAt:   now,
			UpdatedAt:   now,
		}
		if err := tx.Table("vets").Create(&vet).Error; err != nil {
			return core.Role{}, "", err
		} else {
			roleName = Vet
			addedRole = vet
		}
	default:
		return core.Role{}, "", core.ErrInvalidRole
	}

	return addedRole, roleName, tx.Commit().Error
}
func (s *store) DeleteUserRole(ctx context.Context, id int, roleName string) (err error) {

	tx := s.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	var user core.User
	if err = s.DB.Table("users").First(&user, "id = ?", id).Error; err != nil {
		return errors.New("user not found")
	}

	switch roleName {
	case Seeker:
		var seeker core.Seeker
		err = tx.Table("seekers").Where("user_id = ?", id).Delete(seeker).Error
	case Keeper:
		var keeper core.Keeper
		err = tx.Table("keepers").Where("user_id = ?", id).Delete(keeper).Error
	case Vet:
		var vet core.Vet
		err = tx.Table("vets").Where("user_id = ?", id).Delete(vet).Error
	default:
		tx.Rollback()
		return errors.New("invalid roleName name")
	}

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error

}

// check for user existing
func (s *store) UpdateUserRole(ctx context.Context, id int, updateRole core.UpdateRole) (err error) {
	tx := s.DB.Begin()

	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}

	updates := make(map[string]interface{})
	if updateRole.Description != nil {
		updates["description"] = *updateRole.Description
	}

	if len(updates) == 0 {
		return errors.New("no fields to update")
	} else {
		updates["updated_at"] = time.Now()
	}
	roleName := updateRole.Name
	switch roleName {
	case Seeker:
		result := tx.Table("seekers").Where("id = ?", id).Updates(updates)
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
	case Keeper:
		result := tx.Table("keepers").Where("id = ?", id).Updates(updates)
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
	case Vet:
		result := tx.Table("vets").Where("id = ?", id).Updates(updates)
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
	default:
		tx.Rollback()
		return errors.New("invalid updateRole name")
	}

	return tx.Commit().Error
}
