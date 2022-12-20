package repository

import (
	"context"
	"strings"

	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/Group10CapstoneProject/Golang/utils/myerrors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type offlineClassRepositoryImpl struct {
	db *gorm.DB
}

// CreateOfflineClass implements OfflineClassRepository
func (r *offlineClassRepositoryImpl) CreateOfflineClass(body *model.OfflineClass, ctx context.Context) error {
	err := r.db.WithContext(ctx).Create(body).Error
	if err != nil {
		if strings.Contains(err.Error(), "Error 1452:") {
			return myerrors.ErrForeignKey(err)
		}
		return err
	}
	return nil
}

// CreateOfflineClassBooking implements OfflineClassRepository
func (r *offlineClassRepositoryImpl) CreateOfflineClassBooking(body *model.OfflineClassBooking, ctx context.Context) (*model.OfflineClassBooking, error) {
	err := r.db.WithContext(ctx).Create(body).Error
	if err != nil {
		if strings.Contains(err.Error(), "Error 1452:") {
			return nil, myerrors.ErrForeignKey(err)
		}
		return nil, err
	}
	return body, nil
}

// CreateOfflineClassCategory implements OfflineClassRepository
func (r *offlineClassRepositoryImpl) CreateOfflineClassCategory(body *model.OfflineClassCategory, ctx context.Context) error {
	err := r.db.WithContext(ctx).Create(body).Error
	if err != nil {
		if strings.Contains(err.Error(), "Error 1062:") {
			return myerrors.ErrDuplicateRecord
		}
		return err
	}
	return nil
}

// DeleteOfflineClass implements OfflineClassRepository
func (r *offlineClassRepositoryImpl) DeleteOfflineClass(body *model.OfflineClass, ctx context.Context) error {
	check := model.OfflineClass{}
	res := r.db.WithContext(ctx).Preload("OfflineClassBooking").First(&check, body)
	if res.Error != nil {
		return res.Error
	}
	if len(check.OfflineClassBooking) != 0 {
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

// DeleteOfflineClassBooking implements OfflineClassRepository
func (r *offlineClassRepositoryImpl) DeleteOfflineClassBooking(body *model.OfflineClassBooking, ctx context.Context) error {
	res := r.db.WithContext(ctx).Delete(body)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return myerrors.ErrRecordNotFound
	}
	return nil
}

// DeleteOfflineClassCategory implements OfflineClassRepository
func (r *offlineClassRepositoryImpl) DeleteOfflineClassCategory(body *model.OfflineClassCategory, ctx context.Context) error {
	check := model.OfflineClassCategory{}
	res := r.db.WithContext(ctx).Preload("OfflineClass").First(&check, body)
	if res.Error != nil {
		return res.Error
	}
	if len(check.OfflineClass) != 0 {
		return myerrors.ErrRecordIsUsed
	}
	res.Unscoped().Delete(body)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return myerrors.ErrRecordNotFound
	}
	return nil
}

// FindOfflineClassBookingById implements OfflineClassRepository
func (r *offlineClassRepositoryImpl) FindOfflineClassBookingById(id uint, ctx context.Context) (*model.OfflineClassBooking, error) {
	offlineClassBooking := model.OfflineClassBooking{}
	err := r.db.WithContext(ctx).Where("id = ?", id).
		Preload("User").
		Preload("OfflineClass").
		Preload("OfflineClass.Trainer").
		Preload("OfflineClass.OfflineClassCategory").
		Preload("PaymentMethod").
		First(&offlineClassBooking).Error
	return &offlineClassBooking, err
}

// FindOfflineClassBookingByUser implements OfflineClassRepository
func (r *offlineClassRepositoryImpl) FindOfflineClassBookingByUser(userId uint, ctx context.Context) ([]model.OfflineClassBooking, error) {
	offlineClassBookings := []model.OfflineClassBooking{}
	var count int64
	err := r.db.WithContext(ctx).Where("user_id = ?", userId).
		Preload("User").
		Preload("OfflineClass").
		Preload("OfflineClass.Trainer").
		Preload("OfflineClass.OfflineClassCategory").
		Preload("PaymentMethod").
		Find(&offlineClassBookings).Count(&count).Error
	return offlineClassBookings, err
}

// FindOfflineClassBookings implements OfflineClassRepository
func (r *offlineClassRepositoryImpl) FindOfflineClassBookings(page *model.Pagination, cond *model.OfflineClassBooking, ctx context.Context) ([]model.OfflineClassBooking, int, error) {
	offlineClassBooking := []model.OfflineClassBooking{}
	var count int64
	offset := (page.Limit * page.Page) - page.Limit

	query := r.db.WithContext(ctx).Model(&model.OfflineClassBooking{}).Joins("LEFT JOIN users ON users.id = offline_class_bookings.user_id").Joins("LEFT JOIN offline_classes ON offline_classes.id = offline_class_bookings.offline_class_id")
	if page.Q != "" {
		query.Where("users.name LIKE ? OR users.email LIKE ? OR offline_classes.title LIKE ?", "%"+page.Q+"%", "%"+page.Q+"%", "%"+page.Q+"%")
	}
	err := query.
		Preload("User").
		Preload("OfflineClass").
		Preload("OfflineClass.OfflineClassCategory").
		Count(&count).
		Offset(offset).
		Limit(page.Limit).
		Find(&offlineClassBooking).
		Error

	return offlineClassBooking, int(count), err
}

// FindOfflineClassById implements OfflineClassRepository
func (r *offlineClassRepositoryImpl) FindOfflineClassById(id uint, ctx context.Context) (*model.OfflineClass, error) {
	offlineClass := model.OfflineClass{}
	err := r.db.WithContext(ctx).Where("id = ?", id).
		Preload("OfflineClassCategory").
		Preload("Trainer").
		Preload("OfflineClassCategory.OfflineClass").
		First(&offlineClass).Error
	return &offlineClass, err
}

// FindOfflineClassCategoryById implements OfflineClassRepository
func (r *offlineClassRepositoryImpl) FindOfflineClassCategoryById(id uint, ctx context.Context) (*model.OfflineClassCategory, error) {
	offlineClassCategory := model.OfflineClassCategory{}
	err := r.db.WithContext(ctx).Where("id = ?", id).
		Preload("OfflineClass").
		First(&offlineClassCategory).Error
	return &offlineClassCategory, err
}

// FindOfflineClassCategorys implements OfflineClassRepository
func (r *offlineClassRepositoryImpl) FindOfflineClassCategories(cond *model.OfflineClassCategory, ctx context.Context) ([]model.OfflineClassCategory, error) {
	offlineClassCategorys := []model.OfflineClassCategory{}
	err := r.db.WithContext(ctx).Model(&model.OfflineClassCategory{}).
		Preload("OfflineClass").
		Find(&offlineClassCategorys).
		Error

	return offlineClassCategorys, err
}

// FindOfflineClasses implements OfflineClassRepository
func (r *offlineClassRepositoryImpl) FindOfflineClasses(cond *model.OfflineClass, ctx context.Context) ([]model.OfflineClass, error) {
	offlineClasses := []model.OfflineClass{}
	err := r.db.WithContext(ctx).Model(&model.OfflineClass{}).Preload(clause.Associations).
		Find(&offlineClasses, cond).
		Error

	return offlineClasses, err
}

// ReadOfflineClassBookings implements OfflineClassRepository
func (r *offlineClassRepositoryImpl) ReadOfflineClassBookings(cond *model.OfflineClassBooking, ctx context.Context) ([]model.OfflineClassBooking, error) {
	offlineClassBooking := []model.OfflineClassBooking{}
	err := r.db.WithContext(ctx).
		Model(&model.OfflineClassBooking{}).
		Preload("User").
		Preload("OfflineClass").
		Preload("OfflineClass.Trainer").
		Find(&offlineClassBooking, cond).
		Order("updated_at DESC").Error
	if err != nil {
		return nil, err
	}
	return offlineClassBooking, nil
}

// UpdateOfflineClass implements OfflineClassRepository
func (r *offlineClassRepositoryImpl) UpdateOfflineClass(body *model.OfflineClass, ctx context.Context) error {
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

// UpdateOfflineClassBooking implements OfflineClassRepository
func (r *offlineClassRepositoryImpl) UpdateOfflineClassBooking(body *model.OfflineClassBooking, ctx context.Context) error {
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

// UpdateOfflineClassCategory implements OfflineClassRepository
func (r *offlineClassRepositoryImpl) UpdateOfflineClassCategory(body *model.OfflineClassCategory, ctx context.Context) error {
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

// OperationOfflineClassSlot implements OfflineClassRepository
func (r *offlineClassRepositoryImpl) OperationOfflineClassSlot(body *model.OfflineClass, operation string, ctx context.Context) error {
	res := r.db.WithContext(ctx).Model(body).First(body)
	if operation == "increment" {
		res.Update("slot_booked", body.SlotBooked+1)
	} else if operation == "decrement" {
		res.Update("slot_booked", body.SlotBooked-1)
	}
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return myerrors.ErrRecordNotFound
	}
	return nil
}

func NewOfflineClassRepository(database *gorm.DB) OfflineClassRepository {
	return &offlineClassRepositoryImpl{
		db: database,
	}
}
