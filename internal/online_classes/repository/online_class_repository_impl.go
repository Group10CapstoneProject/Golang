package repository

import (
	"context"
	"strings"

	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/Group10CapstoneProject/Golang/utils/myerrors"
	"gorm.io/gorm"
)

type onlineClassRepositoryImpl struct {
	db *gorm.DB
}

// CheckOnlineClassCategoryIsDeleted implements OnlineClassRepository
func (r *onlineClassRepositoryImpl) CheckOnlineClassCategoryIsDeleted(body *model.OnlineClassCategory) error {
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

	if err := r.UpdateOnlineClassCategory(body, context.Background()); err != nil {
		return err
	}
	return nil
}

// CreateOnlineClass implements OnlineClassRepository
func (r *onlineClassRepositoryImpl) CreateOnlineClass(body *model.OnlineClass, ctx context.Context) error {
	err := r.db.WithContext(ctx).Create(body).Error
	if err != nil {
		if strings.Contains(err.Error(), "Error 1452:") {
			return myerrors.ErrForeignKey(err)
		}
		return err
	}
	return nil
}

// CreateOnlineClassBooking implements OnlineClassRepository
func (r *onlineClassRepositoryImpl) CreateOnlineClassBooking(body *model.OnlineClassBooking, ctx context.Context) (*model.OnlineClassBooking, error) {
	err := r.db.WithContext(ctx).Create(body).Error
	if err != nil {
		if strings.Contains(err.Error(), "Error 1452:") {
			return nil, myerrors.ErrForeignKey(err)
		}
		return nil, err
	}
	return body, nil
}

// CreateOnlineClassCategory implements OnlineClassRepository
func (r *onlineClassRepositoryImpl) CreateOnlineClassCategory(body *model.OnlineClassCategory, ctx context.Context) error {
	err := r.db.WithContext(ctx).Create(body).Error
	if err != nil {
		if strings.Contains(err.Error(), "Error 1062:") {
			if err := r.CheckOnlineClassCategoryIsDeleted(body); err == nil {
				return nil
			}
			return myerrors.ErrDuplicateRecord
		}
		return err
	}
	return nil
}

// DeleteOnlineClass implements OnlineClassRepository
func (r *onlineClassRepositoryImpl) DeleteOnlineClass(body *model.OnlineClass, ctx context.Context) error {
	check := model.OnlineClass{}
	res := r.db.WithContext(ctx).Preload("OnlineClassBooking").First(&check, body)
	if res.Error != nil {
		return res.Error
	}
	if len(check.OnlineClassBooking) != 0 {
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

// DeleteOnlineClassBooking implements OnlineClassRepository
func (r *onlineClassRepositoryImpl) DeleteOnlineClassBooking(body *model.OnlineClassBooking, ctx context.Context) error {
	res := r.db.WithContext(ctx).Delete(body)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return myerrors.ErrRecordNotFound
	}
	return nil
}

// DeleteOnlineClassCategory implements OnlineClassRepository
func (r *onlineClassRepositoryImpl) DeleteOnlineClassCategory(body *model.OnlineClassCategory, ctx context.Context) error {
	check := model.OnlineClassCategory{}
	res := r.db.WithContext(ctx).Preload("OnlineClass").First(&check, body)
	if res.Error != nil {
		return res.Error
	}
	if len(check.OnlineClass) != 0 {
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

// FindOnlineClassBookingById implements OnlineClassRepository
func (r *onlineClassRepositoryImpl) FindOnlineClassBookingById(id uint, ctx context.Context) (*model.OnlineClassBooking, error) {
	onlineClassBooking := model.OnlineClassBooking{}
	err := r.db.WithContext(ctx).Where("id = ?", id).
		Preload("User").
		Preload("OnlineClass").
		Preload("OnlineClass.OnlineClassCategory").
		Preload("PaymentMethod").
		First(&onlineClassBooking).Error
	return &onlineClassBooking, err
}

// FindOnlineClassBookingByUser implements OnlineClassRepository
func (r *onlineClassRepositoryImpl) FindOnlineClassBookingByUser(userId uint, ctx context.Context) ([]model.OnlineClassBooking, error) {
	onlineClassBooking := []model.OnlineClassBooking{}
	err := r.db.WithContext(ctx).Where("user_id = ?", userId).
		Preload("User").
		Preload("OnlineClass").
		Preload("OnlineClass.OnlineClassCategory").
		Preload("PaymentMethod").
		Find(&onlineClassBooking).Error
	if err != nil {
		return nil, err
	}
	return onlineClassBooking, nil
}

// FindOnlineClassBookings implements OnlineClassRepository
func (r *onlineClassRepositoryImpl) FindOnlineClassBookings(page *model.Pagination, ctx context.Context) ([]model.OnlineClassBooking, int, error) {
	onlineClassBooking := []model.OnlineClassBooking{}
	var count int64
	offset := (page.Limit * page.Page) - page.Limit

	query := r.db.WithContext(ctx).Model(&model.OnlineClassBooking{}).Joins("LEFT JOIN users ON users.id = online_class_bookings.user_id").Joins("LEFT JOIN online_classes ON online_classes.id = online_class_bookings.online_class_id")
	if page.Q != "" {
		query.Where("users.name LIKE ? OR users.email LIKE ? OR online_classes.title LIKE ?", "%"+page.Q+"%", "%"+page.Q+"%", "%"+page.Q+"%")
	}
	err := query.
		Preload("User").
		Preload("OnlineClass").
		Preload("OnlineClass.OnlineClassCategory").
		Count(&count).
		Offset(offset).
		Limit(page.Limit).
		Find(&onlineClassBooking).
		Error

	return onlineClassBooking, int(count), err
}

// FindOnlineClassById implements OnlineClassRepository
func (r *onlineClassRepositoryImpl) FindOnlineClassById(id uint, ctx context.Context) (*model.OnlineClass, error) {
	onlineClass := model.OnlineClass{}
	err := r.db.WithContext(ctx).Where("id = ?", id).
		Preload("OnlineClassCategory").
		Preload("Trainer").
		Preload("OnlineClassCategory.OnlineClass").
		First(&onlineClass).Error
	return &onlineClass, err
}

// FindOnlineClassCategoryById implements OnlineClassRepository
func (r *onlineClassRepositoryImpl) FindOnlineClassCategoryById(id uint, ctx context.Context) (*model.OnlineClassCategory, error) {
	onlineClassCategory := model.OnlineClassCategory{}
	err := r.db.WithContext(ctx).Where("id = ?", id).Preload("OnlineClass").First(&onlineClassCategory).Error
	return &onlineClassCategory, err
}

// FindOnlineClassCategorys implements OnlineClassRepository
func (r *onlineClassRepositoryImpl) FindOnlineClassCategories(ctx context.Context) ([]model.OnlineClassCategory, error) {
	onlineClassCategories := []model.OnlineClassCategory{}
	err := r.db.WithContext(ctx).Preload("OnlineClass").Find(&onlineClassCategories).Error
	return onlineClassCategories, err
}

// FindOnlineClasss implements OnlineClassRepository
func (r *onlineClassRepositoryImpl) FindOnlineClasses(ctx context.Context) ([]model.OnlineClass, error) {
	onlineClasses := []model.OnlineClass{}
	err := r.db.WithContext(ctx).Preload("OnlineClassCategory").Find(&onlineClasses).Error
	return onlineClasses, err
}

// UpdateOnlineClass implements OnlineClassRepository
func (r *onlineClassRepositoryImpl) UpdateOnlineClass(body *model.OnlineClass, ctx context.Context) error {
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

// UpdateOnlineClassBooking implements OnlineClassRepository
func (r *onlineClassRepositoryImpl) UpdateOnlineClassBooking(body *model.OnlineClassBooking, ctx context.Context) error {
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

// UpdateOnlineClassCategory implements OnlineClassRepository
func (r *onlineClassRepositoryImpl) UpdateOnlineClassCategory(body *model.OnlineClassCategory, ctx context.Context) error {
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

// ReadOnlineClassBooking implements OnlineClassRepository
func (r *onlineClassRepositoryImpl) ReadOnlineClassBooking(cond *model.OnlineClassBooking, ctx context.Context) ([]model.OnlineClassBooking, error) {
	onlineClassBooking := []model.OnlineClassBooking{}
	err := r.db.WithContext(ctx).
		Model(&model.OnlineClassBooking{}).
		Preload("User").
		Preload("OnlineClass").
		Find(&onlineClassBooking, cond).
		Order("updated_at DESC").Error
	if err != nil {
		return nil, err
	}
	return onlineClassBooking, nil
}

func NewOnlineClassRepository(database *gorm.DB) OnlineClassRepository {
	return &onlineClassRepositoryImpl{
		db: database,
	}
}
