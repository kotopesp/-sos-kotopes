package seeker

import (
	"context"
	"errors"
	"github.com/kotopesp/sos-kotopes/internal/core"
	"github.com/kotopesp/sos-kotopes/pkg/logger"
)

type service struct {
	seekersStore core.SeekersStore
}

func New(seekersStore core.SeekersStore) core.SeekersService {
	return &service{seekersStore: seekersStore}
}

func (s *service) CreateSeeker(ctx context.Context, seeker core.Seeker, equipment core.Equipment) (core.Seeker, error) {
	if _, err := s.GetSeeker(ctx, seeker.UserID); !errors.Is(err, core.ErrSeekerNotFound) {
		logger.Log().Error(ctx, core.ErrSeekerExists.Error())
		return core.Seeker{}, core.ErrSeekerExists
	}

	return s.seekersStore.CreateSeeker(ctx, seeker, equipment)
}

func (s *service) GetSeeker(ctx context.Context, id int) (core.Seeker, error) {
	seeker, err := s.seekersStore.GetSeeker(ctx, id)
	if err != nil {
		logger.Log().Error(ctx, core.ErrSeekerNotFound.Error())
		return core.Seeker{}, core.ErrSeekerNotFound
	}

	if seeker.IsDeleted == true {
		logger.Log().Error(ctx, core.ErrSeekerDeleted.Error())
		return core.Seeker{}, core.ErrSeekerDeleted
	}

	return seeker, nil
}

func (s *service) UpdateSeeker(ctx context.Context, updateSeeker core.UpdateSeeker) (core.Seeker, error) {
	updates := make(map[string]interface{})

	if updateSeeker.Description != nil {
		updates["description"] = *updateSeeker.Description
	}

	if updateSeeker.Location != nil {
		updates["location"] = *updateSeeker.Location
	}

	if updateSeeker.Price != nil {
		updates["price"] = *updateSeeker.Price
	}

	if updateSeeker.HaveCar != nil {
		updates["have_car"] = *updateSeeker.HaveCar
	}

	if len(updates) == 0 {
		logger.Log().Error(ctx, core.ErrEmptyUpdateRequest.Error())
		return core.Seeker{}, core.ErrEmptyUpdateRequest
	}

	getSeeker, err := s.GetSeeker(ctx, *updateSeeker.UserID)
	if err != nil {
		logger.Log().Error(ctx, core.ErrSeekerNotFound.Error())
		return core.Seeker{}, core.ErrSeekerNotFound
	}

	if getSeeker.IsDeleted == true {
		logger.Log().Error(ctx, core.ErrSeekerDeleted.Error())
		return core.Seeker{}, core.ErrSeekerDeleted
	}

	return s.seekersStore.UpdateSeeker(ctx, getSeeker.ID, updates)
}
