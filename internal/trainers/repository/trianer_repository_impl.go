package repository

import (
	"context"
	"strings"

	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/Group10CapstoneProject/Golang/utils/myerrors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

// CheckTrainerIsDeleted implements TrainerRepository
func (r *trainerRepositoryImpl) CheckTrainerIsDeleted(body *model.Trainer) error {
	trainer := model.Trainer{}
	err := r.db.Where("email = ?", body.Email).First(&model.Trainer{}).Error
	if err == nil {
		return myerrors.ErrDuplicateRecord
	}
	err = r.db.Unscoped().Where("email = ?", body.Email).First(&trainer).Update("deleted_at", nil).Error
	if err != nil {
		return err
	}
	body.ID = trainer.ID

	if err := r.UpdateTrainer(body, context.Background()); err != nil {
		return err
	}
	return nil
}

// CreateSkill implements TrainerRepository
func (r *trainerRepositoryImpl) CreateSkill(body *model.Skill, ctx context.Context) error {
	err := r.db.WithContext(ctx).Create(body).Error
	if err != nil {
		if strings.Contains(err.Error(), "Error 1062:") {
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
		if strings.Contains(err.Error(), "Error 1062:") {
			if err := r.CheckTrainerIsDeleted(body); err == nil {
				return nil
			}
			return myerrors.ErrDuplicateRecord
		}
		if strings.Contains(err.Error(), "Error 1452:") {
			return myerrors.ErrForeignKey(err)
		}
		return err
	}
	return nil
}

// CreateTrainerBooking implements TrainerRepository
func (r *trainerRepositoryImpl) CreateTrainerBooking(body *model.TrainerBooking, ctx context.Context) (*model.TrainerBooking, error) {
	var count int64
	var dailySlot int64
	date := body.Time.Format("2006-01-02")
	err := r.db.WithContext(ctx).Model(&model.TrainerBooking{}).Where("trainer_id = ? AND DATE(time) = ? AND status NOT IN (?,?,?)",
		body.TrainerID, date, model.CENCEL, model.INACTIVE, model.REJECT).
		Count(&count).Error
	if err != nil {
		return nil, err
	}
	err = r.db.WithContext(ctx).Model(&model.Trainer{}).Where("id = ?", body.TrainerID).Select("daily_slot").First(&dailySlot).Error
	if err != nil {
		return nil, err
	}
	if count >= dailySlot {
		return nil, myerrors.ErrTrainerIsFull
	}
	err = r.db.WithContext(ctx).Create(body).Error
	if err != nil {
		if strings.Contains(err.Error(), "Error 1452:") {
			return nil, myerrors.ErrForeignKey(err)
		}
		return nil, err
	}
	return body, nil
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
	err := r.db.WithContext(ctx).Find(&skills).Preload(clause.Associations).Error
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
		Preload("Trainer.TrainerSkill.Skill").
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
		Joins("LEFT JOIN trainers ON trainers.id = trainer_bookings.trainer_id")
	if page.Q != "" {
		query.Where("users.name LIKE ? OR users.email LIKE ? OR trainers.name LIKE ?", "%"+page.Q+"%", "%"+page.Q+"%", "%"+page.Q+"%")
	}
	err := query.
		Preload("User").
		Preload("Trainer").
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
	trainer := model.Trainer{}
	err := r.db.WithContext(ctx).Where("id = ?", id).
		Preload("TrainerSkill").
		Preload("TrainerBooking", "status = ?", model.ACTIVE).
		Preload("TrainerSkill.Skill").
		First(&trainer).Error
	if err != nil {
		return nil, err
	}
	return &trainer, nil
}

// FindTrainers implements TrainerRepository
func (r *trainerRepositoryImpl) FindTrainers(cond *model.Trainer, priceOrder string, date string, ctx context.Context) ([]model.Trainer, error) {
	traieners := []model.Trainer{}
	res := r.db.WithContext(ctx).Model(&model.Trainer{}).
		Preload("TrainerSkill").
		Preload("TrainerSkill.Skill")
	if priceOrder != "" {
		res.Order("price " + priceOrder)
	} else {
		res.Order("id DESC")
	}
	if date != "" {
		res.Where("daily_slot > (SELECT COUNT(a.id) FROM trainer_bookings a WHERE a.trainer_id = id AND DATE(a.time) = ? AND a.status NOT IN (?,?,?))",
			date, model.CENCEL, model.INACTIVE, model.REJECT)
	}
	if cond.Name != "" {
		res.Where("name LIKE ?", "%"+cond.Name+"%")
	}
	if cond.Gender != "" {
		res.Where("gender = ?", cond.Gender)
	}
	err := res.Find(&traieners).Error
	if err != nil {
		return nil, err
	}
	return traieners, err
}

// ReadTrainerBooking implements TrainerRepository
func (r *trainerRepositoryImpl) ReadTrainerBooking(cond *model.TrainerBooking, ctx context.Context) ([]model.TrainerBooking, error) {
	trainerBookings := []model.TrainerBooking{}
	err := r.db.WithContext(ctx).
		Model(&model.TrainerBooking{}).
		Preload(clause.Associations).
		Find(&trainerBookings, cond).
		Error
	if err != nil {
		return nil, err
	}
	return trainerBookings, nil
}

// UpdateSkill implements TrainerRepository
func (r *trainerRepositoryImpl) UpdateSkill(body *model.Skill, ctx context.Context) error {
	res := r.db.WithContext(ctx).Model(body).Updates(body)
	if res.Error != nil {
		if strings.Contains(res.Error.Error(), "Error 1062:") {
			return myerrors.ErrDuplicateRecord
		}
		return res.Error
	}
	if res.RowsAffected == 0 {
		return myerrors.ErrRecordNotFound
	}
	return nil
}

// UpdateTrainer implements TrainerRepository
func (r *trainerRepositoryImpl) UpdateTrainer(body *model.Trainer, ctx context.Context) error {
	res := r.db.Begin()
	res.WithContext(ctx).
		Model(&model.TrainerSkill{}).
		Where("trainer_id = ?", body.ID).
		Delete(&model.TrainerSkill{})
	err := res.WithContext(ctx).Model(body).Updates(body).Error
	res.Rollback()
	if err != nil {
		if strings.Contains(err.Error(), "Error 1062:") {
			return myerrors.ErrDuplicateRecord
		}
		if strings.Contains(err.Error(), "Error 1452:") {
			return myerrors.ErrForeignKey(err)
		}
		return err
	}
	res.Commit()
	return nil
}

// UpdateTrainerBooking implements TrainerRepository
func (r *trainerRepositoryImpl) UpdateTrainerBooking(body *model.TrainerBooking, ctx context.Context) error {
	res := r.db.WithContext(ctx).Model(body).Updates(body)
	err := res.Error
	if err != nil {
		if strings.Contains(err.Error(), "Error 1062:") {
			return myerrors.ErrDuplicateRecord
		}
		if strings.Contains(err.Error(), "Error 1452:") {
			return myerrors.ErrForeignKey(err)
		}
		return err
	}
	if res.RowsAffected == 0 {
		return myerrors.ErrRecordNotFound
	}
	return nil
}

func NewTrainerRepository(database *gorm.DB) TrainerRepository {
	return &trainerRepositoryImpl{
		db: database,
	}
}
