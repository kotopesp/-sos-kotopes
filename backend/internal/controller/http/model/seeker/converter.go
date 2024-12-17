package seeker

import "github.com/kotopesp/sos-kotopes/internal/core"

func equipmentChecker(s string) bool {
	return s != ""
}

func (seeker *CreateSeeker) GetEquipment() (core.Equipment, error) {
	if len(seeker.Equipment) != 5 {
		return core.Equipment{}, core.ErrInvalidEquipment
	}

	return core.Equipment{
		HaveMetalCage:   equipmentChecker(seeker.Equipment[0]),
		HavePlasticCage: equipmentChecker(seeker.Equipment[1]),
		HaveNet:         equipmentChecker(seeker.Equipment[2]),
		HaveLadder:      equipmentChecker(seeker.Equipment[3]),
		HaveOther:       seeker.Equipment[4],
	}, nil
}

func (seeker *CreateSeeker) ToCoreSeeker() core.Seeker {
	return core.Seeker{
		UserID:           seeker.UserID,
		AnimalType:       seeker.AnimalType,
		Description:      seeker.Description,
		Location:         seeker.Location,
		EquipmentRental:  seeker.EquipmentRental,
		HaveCar:          seeker.HaveCar,
		Price:            seeker.Price,
		WillingnessCarry: seeker.WillingnessCarry,
	}
}

func (seeker *UpdateSeeker) ToCoreUpdateSeeker() core.UpdateSeeker {
	return core.UpdateSeeker{
		UserID:           seeker.UserID,
		AnimalType:       seeker.AnimalType,
		Description:      seeker.Description,
		Location:         seeker.Location,
		EquipmentRental:  seeker.EquipmentRental,
		HaveCar:          seeker.HaveCar,
		Price:            seeker.Price,
		WillingnessCarry: seeker.WillingnessCarry,
	}
}

func ToResponseSeeker(seeker *core.Seeker) ResponseSeeker {
	if seeker == nil {
		return ResponseSeeker{}
	}
	return ResponseSeeker{
		ID:               seeker.ID,
		UserID:           seeker.UserID,
		AnimalType:       seeker.AnimalType,
		Location:         seeker.Location,
		EquipmentRental:  seeker.EquipmentRental,
		Equipment:        seeker.EquipmentID,
		Description:      seeker.Description,
		HaveCar:          seeker.HaveCar,
		Price:            seeker.Price,
		WillingnessCarry: seeker.WillingnessCarry,
	}
}
