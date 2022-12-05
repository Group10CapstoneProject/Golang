package repository

import (
	"context"
	"strings"

	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/Group10CapstoneProject/Golang/utils/myerrors"
	"gorm.io/gorm"
)

type memberRepositoryImpl struct {
	db *gorm.DB
}

// CreateMember implements MemberRepository
func (r *memberRepositoryImpl) CreateMember(body *model.Member, ctx context.Context) error {
	err := r.db.WithContext(ctx).Create(body).Error
	return err
}

// CreateMemberType implements MemberRepository
func (r *memberRepositoryImpl) CreateMemberType(body *model.MemberType, ctx context.Context) error {
	err := r.db.WithContext(ctx).Create(body).Error
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			if err := r.CheckMemberTypeIsDeleted(body); err == nil {
				return nil
			}
			return myerrors.ErrDuplicateRecord
		}
		return err
	}
	return nil
}

// CheckMemberTypeIsDeleted implements MemberRepository
func (r *memberRepositoryImpl) CheckMemberTypeIsDeleted(body *model.MemberType) error {
	memberType := model.MemberType{}
	err := r.db.Where("name = ?", body.Name).First(&model.MemberType{}).Error
	if err == nil {
		return myerrors.ErrDuplicateRecord
	}
	err = r.db.Unscoped().Where("name = ?", body.Name).First(&memberType).Update("deleted_at", nil).Error
	if err != nil {
		return err
	}
	body.ID = memberType.ID

	if err := r.UpdateMemberType(body, context.Background()); err != nil {
		return err
	}
	return nil
}

// DeleteMember implements MemberRepository
func (r *memberRepositoryImpl) DeleteMember(body *model.Member, ctx context.Context) error {
	res := r.db.WithContext(ctx).Delete(body)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return myerrors.ErrRecordNotFound
	}
	return nil
}

// DeleteMemberType implements MemberRepository
func (r *memberRepositoryImpl) DeleteMemberType(body *model.MemberType, ctx context.Context) error {
	res := r.db.WithContext(ctx).Delete(body)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return myerrors.ErrRecordNotFound
	}
	return nil
}

// FindMemberById implements MemberRepository
func (r *memberRepositoryImpl) FindMemberById(id uint, ctx context.Context) (*model.Member, error) {
	member := model.Member{}
	err := r.db.WithContext(ctx).Where("id = ?", id).
		Preload("User").
		Preload("MemberType").
		Preload("PaymentMethod").
		First(&member).Error
	return &member, err
}

// FindMemberByUser implements MemberRepository
func (r *memberRepositoryImpl) FindMemberByUser(userId uint, ctx context.Context) (*model.Member, error) {
	member := model.Member{}
	err := r.db.WithContext(ctx).Where("user_id = ? AND status = ?", userId, model.ACTIVE).
		Preload("User").
		Preload("MemberType").
		Preload("PaymentMethod").
		First(&member).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, myerrors.ErrRecordNotFound
		}
		return nil, err
	}
	return &member, nil
}

// FindMemberTypes implements MemberRepository
func (r *memberRepositoryImpl) FindMemberTypes(ctx context.Context) ([]model.MemberType, error) {
	memberTypes := []model.MemberType{}
	err := r.db.WithContext(ctx).Find(&memberTypes).Error
	return memberTypes, err
}

// FindMemberTypeById implements MemberRepository
func (r *memberRepositoryImpl) FindMemberTypeById(id uint, ctx context.Context) (*model.MemberType, error) {
	memberType := model.MemberType{}
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&memberType).Error
	return &memberType, err
}

// FindMembers implements MemberRepository
func (r *memberRepositoryImpl) FindMembers(page *model.Pagination, ctx context.Context) ([]model.Member, int, error) {
	members := []model.Member{}
	var count int64
	offset := (page.Limit * page.Page) - page.Limit

	query := r.db.WithContext(ctx).Model(&model.Member{})
	if page.Q != "" {
		query.Where("users.name LIKE ? OR users.email LIKE ? OR MemberType.name", "%"+page.Q+"%", "%"+page.Q+"%", "%"+page.Q+"%")
	}
	err := query.
		Preload("User").
		Preload("MemberType").
		Count(&count).
		Offset(offset).
		Limit(page.Limit).
		Find(&members).
		Error

	return members, int(count), err
}

// UpdateMember implements MemberRepository
func (r *memberRepositoryImpl) UpdateMember(body *model.Member, ctx context.Context) error {
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

// UpdateMemberType implements MemberRepository
func (r *memberRepositoryImpl) UpdateMemberType(body *model.MemberType, ctx context.Context) error {
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

func NewMemberRepository(database *gorm.DB) MemberRepository {
	return &memberRepositoryImpl{
		db: database,
	}
}
