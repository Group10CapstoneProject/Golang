package repository

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/model"
)

type MemberRepository interface {
	// member
	CreateMember(body *model.Member, ctx context.Context) (*model.Member, error)
	FindMembers(page *model.Pagination, ctx context.Context) ([]model.Member, int, error)
	FindMemberById(id uint, ctx context.Context) (*model.Member, error)
	FindMemberByUser(userId uint, ctx context.Context) (*model.Member, error)
	ReadMembers(body *model.Member, ctx context.Context) ([]model.Member, error)
	UpdateMember(body *model.Member, ctx context.Context) error
	DeleteMember(body *model.Member, ctx context.Context) error
	MemberInactive(body model.Member, ctx context.Context) error

	// member type
	CreateMemberType(body *model.MemberType, ctx context.Context) error
	FindMemberTypes(ctx context.Context) ([]model.MemberType, error)
	FindMemberTypeById(id uint, ctx context.Context) (*model.MemberType, error)
	CheckMemberTypeIsDeleted(body *model.MemberType) error
	UpdateMemberType(body *model.MemberType, ctx context.Context) error
	DeleteMemberType(body *model.MemberType, ctx context.Context) error
}
