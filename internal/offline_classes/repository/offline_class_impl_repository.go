package repository

import (
	"context"
	"strings"

	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/Group10CapstoneProject/Golang/utils/myerrors"
	"gorm.io/gorm"
)

type offlineClassRepositoryImpl struct {
	db *gorm.DB
}

// CreateOfflineClass implements OfflineClassRepository
func (r *offlineClassRepositoryImpl) CreateOfflineClass(body *model.OfflineClass, ctx context.Context) error {
	err := r.db.WithContext(ctx).Create(body).Error
	return err
}

// CreateOfflineClassBooking implements OfflineClassRepository
func (r *offlineClassRepositoryImpl) CreateOfflineClassBooking(body *model.OfflineClassBooking, ctx context.Context) error {
	err := r.db.WithContext(ctx).Create(body).Error
	return err
}

// CreateOfflineClassCategory implements OfflineClassRepository
func (r *offlineClassRepositoryImpl) CreateOfflineClassCategory(body *model.OfflineClassCategory, ctx context.Context) error {
	err := r.db.WithContext(ctx).Create(body).Error
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			if err := r.CheckOfflineClassCategoryIsDeleted(body); err == nil {
				return nil
			}
			return myerrors.ErrDuplicateRecord
		}
		return err
	}
	return nil
}

// CheckOfflineClassCategoryIsDeleted implements OnlineClassRepository
func (r *offlineClassRepositoryImpl) CheckOfflineClassCategoryIsDeleted(body *model.OfflineClassCategory) error {
	onlineClassCategory := model.OnlineClassCategory{}
	err := r.db.Where("name = ?", body.Name).First(&model.OnlineClassCategory{}).Error
	if err == nil {
		return myerrors.ErrDuplicateRecord
	}
	err = r.db.Unscoped().Where("name = ?", body.Name).First(&onlineClassCategory).Update("deleted_at", nil).Error
	if err != nil {
		return err
	}
	body.ID = onlineClassCategory.ID

	if err := r.UpdateOfflineClassCategory(body, context.Background()); err != nil {
		return err
	}
	return nil
}

// DeleteOfflineClass implements OfflineClassRepository
func (r *offlineClassRepositoryImpl) DeleteOfflineClass(body *model.OfflineClass, ctx context.Context) error {
	res := r.db.WithContext(ctx).Delete(body)
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
	res := r.db.WithContext(ctx).Delete(body)
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
		First(&offlineClass).Error
	return &offlineClass, err
}

// FindOfflineClassCategoryById implements OfflineClassRepository
func (r *offlineClassRepositoryImpl) FindOfflineClassCategoryById(id uint, ctx context.Context) (*model.OfflineClassCategory, error) {
	offlineClassCategory := model.OfflineClassCategory{}
	err := r.db.WithContext(ctx).Where("id = ?", id).
		Preload("OfflineCLass").
		First(&offlineClassCategory).Error
	return &offlineClassCategory, err
}

// FindOfflineClassCategorys implements OfflineClassRepository
func (r *offlineClassRepositoryImpl) FindOfflineClassCategories(cond *model.OfflineClassCategory, ctx context.Context) ([]model.OfflineClassCategory, error) {
	offlineClassCategorys := []model.OfflineClassCategory{}
	err := r.db.WithContext(ctx).Model(&model.OfflineClassCategory{}).
		Find(&offlineClassCategorys).
		Error

	return offlineClassCategorys, err
}

// FindOfflineClasses implements OfflineClassRepository
func (r *offlineClassRepositoryImpl) FindOfflineClasses(cond *model.OfflineClass, ctx context.Context) ([]model.OfflineClass, error) {
	offlineClasses := []model.OfflineClass{}
	err := r.db.WithContext(ctx).Model(&model.OfflineClass{}).
		Find(&offlineClasses, cond).
		Error

	return offlineClasses, err
}

// ReadOfflineClassBookings implements OfflineClassRepository
func (r *offlineClassRepositoryImpl) ReadOfflineClassBookings(cond *model.OfflineClassBooking, ctx context.Context) ([]model.OfflineClassBooking, error) {
	offlineClassBooking := []model.OfflineClassBooking{}
	err := r.db.WithContext(ctx).Find(&offlineClassBooking, cond).Error
	return offlineClassBooking, err
}

// UpdateOfflineClass implements OfflineClassRepository
func (r *offlineClassRepositoryImpl) UpdateOfflineClass(body *model.OfflineClass, ctx context.Context) error {
	res := r.db.WithContext(ctx).Model(body).Updates(body)
	if res.Error != nil {
		if strings.Contains(res.Error.Error(), "Duplicate entry") {
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
		if strings.Contains(res.Error.Error(), "Duplicate entry") {
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
		if strings.Contains(res.Error.Error(), "Duplicate entry") {
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
	res := r.db.WithContext(ctx).Model(body)
	if operation == "increment" {
		res.Update("slot_booked", "slot_booked + 1")
	} else if operation == "decrement" {
		res.Update("slot_booked", "slot_booked - 1")
	}
	if res.Error != nil {
		if strings.Contains(res.Error.Error(), "Duplicate entry") {
			return myerrors.ErrDuplicateRecord
		}
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
