package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/Group10CapstoneProject/Golang/utils/myerrors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type memberRepositoryImpl struct {
	db *gorm.DB
}

// CreateMember implements MemberRepository
func (r *memberRepositoryImpl) CreateMember(body *model.Member, ctx context.Context) (*model.Member, error) {
	err := r.db.WithContext(ctx).Create(body).Error
	if err != nil {
		if strings.Contains(err.Error(), "Error 1452:") {
			return nil, myerrors.ErrForeignKey(err)
		}
		return nil, err
	}
	fmt.Println(body.ID)
	return body, nil
}

// CreateMemberType implements MemberRepository
func (r *memberRepositoryImpl) CreateMemberType(body *model.MemberType, ctx context.Context) error {
	err := r.db.WithContext(ctx).Create(body).Error
	if err != nil {
		if strings.Contains(err.Error(), "Error 1062:") {
			return myerrors.ErrDuplicateRecord
		}
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
	check := model.MemberType{}
	res := r.db.WithContext(ctx).Preload("Member").First(&check, body)
	if res.Error != nil {
		return res.Error
	}
	if len(check.Member) != 0 {
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

	query := r.db.WithContext(ctx).Model(&model.Member{}).Joins("LEFT JOIN users ON users.id = members.user_id").Joins("LEFT JOIN member_types ON member_types.id = members.member_type_id")
	if page.Q != "" {
		query.Where("users.name LIKE ? OR users.email LIKE ? OR member_types.name LIKE ?", "%"+page.Q+"%", "%"+page.Q+"%", "%"+page.Q+"%")
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
	if body.Status == model.ACTIVE {
		err := r.MemberInactive(*body, ctx)
		if err != nil {
			return err
		}
	}
	res := r.db.WithContext(ctx).Model(body).Updates(body)
	err := res.Error
	if err != nil {
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

// MemberInactive implements MemberRepository
func (r *memberRepositoryImpl) MemberInactive(body model.Member, ctx context.Context) error {
	err := r.db.WithContext(ctx).
		Model(&body).
		First(&body).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return myerrors.ErrRecordNotFound
		}
		return err
	}
	err = r.db.WithContext(ctx).
		Model(&model.Member{}).
		Where("user_id = ? AND status = ?", body.UserID, model.ACTIVE).
		Update("status", model.INACTIVE).
		Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateMemberType implements MemberRepository
func (r *memberRepositoryImpl) UpdateMemberType(body *model.MemberType, ctx context.Context) error {
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

// ReadMembers implements MemberRepository
func (r *memberRepositoryImpl) ReadMembers(body *model.Member, ctx context.Context) ([]model.Member, error) {
	members := []model.Member{}
	err := r.db.WithContext(ctx).
		Model(&model.Member{}).
		Preload(clause.Associations).
		Find(&members, body).
		Order("updated_at DESC").Error
	if err != nil {
		return nil, err
	}
	return members, nil
}

func NewMemberRepository(database *gorm.DB) MemberRepository {
	return &memberRepositoryImpl{
		db: database,
	}
}
