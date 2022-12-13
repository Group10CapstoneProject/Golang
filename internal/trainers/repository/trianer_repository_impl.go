package repository

import (
	"context"
	"strings"

	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/Group10CapstoneProject/Golang/utils/myerrors"
	"gorm.io/gorm"
)

type trainerRepositoryImpl struct {
	db *gorm.DB
}

// CheckSkillIsDeleted implements TrainerRepository
func (r *trainerRepositoryImpl) CheckSkillIsDeleted(body *model.Skill) error {
	skill := model.Skill{}
	err := r.db.Where("name = ?", body.Name).First(&model.Skill{}).Error
	if err == nil {
		return myerrors.ErrDuplicateRecord
	}
	err = r.db.Unscoped().Where("name = ?", body.Name).First(&skill).Update("deleted_at", nil).Error
	if err != nil {
		return err
	}
	body.ID = skill.ID

	if err := r.UpdateSkill(body, context.Background()); err != nil {
		return err
	}
	return nil
}

// CreateSkill implements TrainerRepository
func (r *trainerRepositoryImpl) CreateSkill(body *model.Skill, ctx context.Context) error {
	err := r.db.WithContext(ctx).Create(body).Error
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			if err := r.CheckSkillIsDeleted(body); err == nil {
				return nil
			}
			return myerrors.ErrDuplicateRecord
		}
		return err
	}
	return nil
}

// CreateTrainer implements TrainerRepository
func (r *trainerRepositoryImpl) CreateTrainer(body *model.Trainer, ctx context.Context) error {
	err := r.db.WithContext(ctx).Create(body).Error
	if err != nil {
		return err
	}
	return nil
}

// CreateTrainerBooking implements TrainerRepository
func (r *trainerRepositoryImpl) CreateTrainerBooking(body *model.TrainerBooking, ctx context.Context) error {
	err := r.db.WithContext(ctx).Create(body).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteSkill implements TrainerRepository
func (r *trainerRepositoryImpl) DeleteSkill(body *model.Skill, ctx context.Context) error {
	check := model.Skill{}
	res := r.db.WithContext(ctx).Preload("TrainerSkill").First(&check, body)
	if res.Error != nil {
		return res.Error
	}
	if len(check.TrainerSkill) != 0 {
		return myerrors.ErrRecordIsUsed
	}
	res.Delete(body)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return myerrors.ErrRecordNotFound
	}
	return nil
}

// DeleteTrainer implements TrainerRepository
func (r *trainerRepositoryImpl) DeleteTrainer(body *model.Trainer, ctx context.Context) error {
	check := model.Trainer{}
	res := r.db.WithContext(ctx).Preload("TrainerBooking").First(&check, body)
	if res.Error != nil {
		return res.Error
	}
	if len(check.TrainerBooking) != 0 {
		return myerrors.ErrRecordIsUsed
	}
	res.Delete(body)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return myerrors.ErrRecordNotFound
	}
	return nil
}

// DeleteTrainerBooking implements TrainerRepository
func (r *trainerRepositoryImpl) DeleteTrainerBooking(body *model.TrainerBooking, ctx context.Context) error {
	res := r.db.WithContext(ctx).Delete(body)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return myerrors.ErrRecordNotFound
	}
	return nil
}

// FindSkillById implements TrainerRepository
func (r *trainerRepositoryImpl) FindSkillById(id uint, ctx context.Context) (*model.Skill, error) {
	skill := model.Skill{}
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&skill).Error
	if err != nil {
		return nil, err
	}
	return &skill, nil
}

// FindSkills implements TrainerRepository
func (r *trainerRepositoryImpl) FindSkills(ctx context.Context) ([]model.Skill, error) {
	skills := []model.Skill{}
	err := r.db.WithContext(ctx).Find(&skills).Error
	if err != nil {
		return nil, err
	}
	return skills, nil
}

// FindTrainerBookingById implements TrainerRepository
func (r *trainerRepositoryImpl) FindTrainerBookingById(id uint, ctx context.Context) (*model.TrainerBooking, error) {
	trainerBooking := model.TrainerBooking{}
	err := r.db.WithContext(ctx).Where("id = ?", id).
		Preload("User").
		Preload("Trainer").
		Preload("Trainer.TrainerSkill").
		Preload("PaymentMethod").
		First(&trainerBooking).Error
	if err != nil {
		return nil, err
	}
	return &trainerBooking, nil
}

// FindTrainerBookingByUser implements TrainerRepository
func (r *trainerRepositoryImpl) FindTrainerBookingByUser(userId uint, ctx context.Context) ([]model.TrainerBooking, error) {
	trainerBookings := []model.TrainerBooking{}
	err := r.db.WithContext(ctx).Where("user_id = ?", userId).
		Preload("User").
		Preload("Trainer").
		Find(&trainerBookings).Error
	if err != nil {
		return nil, err
	}
	return trainerBookings, nil
}

// FindTrainerBookings implements TrainerRepository
func (r *trainerRepositoryImpl) FindTrainerBookings(page *model.Pagination, ctx context.Context) ([]model.TrainerBooking, int, error) {
	trainerBooking := []model.TrainerBooking{}
	var count int64
	offset := (page.Limit * page.Page) - page.Limit

	query := r.db.WithContext(ctx).Model(&model.TrainerBooking{}).
		Joins("LEFT JOIN users ON users.id = trainer_bookings.user_id").
		Joins("LEFT JOIN trainer ON trainer.id = trainer_bookings.trainer_id")
	if page.Q != "" {
		query.Where("users.name LIKE ? OR users.email LIKE ? OR trainer.name LIKE ?", "%"+page.Q+"%", "%"+page.Q+"%", "%"+page.Q+"%")
	}
	err := query.
		Preload("User").
		Preload("trainer").
		Count(&count).
		Offset(offset).
		Limit(page.Limit).
		Find(&trainerBooking).
		Error

	if err != nil {
		return nil, 0, err
	}
	return trainerBooking, int(count), nil
}

// FindTrainerById implements TrainerRepository
func (r *trainerRepositoryImpl) FindTrainerById(id uint, ctx context.Context) (*model.Trainer, error) {
	panic("unimplemented")
}

// FindTrainers implements TrainerRepository
func (r *trainerRepositoryImpl) FindTrainers(ctx context.Context) ([]model.Trainer, error) {
	panic("unimplemented")
}

// ReadTrainerBooking implements TrainerRepository
func (r *trainerRepositoryImpl) ReadTrainerBooking(cond *model.TrainerBooking, ctx context.Context) ([]model.TrainerBooking, error) {
	panic("unimplemented")
}

// UpdateSkill implements TrainerRepository
func (r *trainerRepositoryImpl) UpdateSkill(body *model.Skill, ctx context.Context) error {
	panic("unimplemented")
}

// UpdateTrainer implements TrainerRepository
func (r *trainerRepositoryImpl) UpdateTrainer(body *model.Trainer, ctx context.Context) error {
	panic("unimplemented")
}

// UpdateTrainerBooking implements TrainerRepository
func (r *trainerRepositoryImpl) UpdateTrainerBooking(body *model.TrainerBooking, ctx context.Context) error {
	panic("unimplemented")
}

func NewTrainerRepository(database *gorm.DB) TrainerRepository {
	return &trainerRepositoryImpl{
		db: database,
	}
}
