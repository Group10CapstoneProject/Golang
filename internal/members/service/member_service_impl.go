package service

import (
	"bytes"
	"context"
	"errors"
	"io"
	"time"

	"github.com/Group10CapstoneProject/Golang/internal/members/dto"
	memberRepo "github.com/Group10CapstoneProject/Golang/internal/members/repository"
	notifRepo "github.com/Group10CapstoneProject/Golang/internal/notifications/repository"
	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/Group10CapstoneProject/Golang/utils/imgkit"
	"github.com/Group10CapstoneProject/Golang/utils/myerrors"
	"github.com/google/uuid"
)

type memberServiceImpl struct {
	memberRepository       memberRepo.MemberRepository
	notificationRepository notifRepo.NotificationRepository
	imagekitService        imgkit.ImagekitService
}

// CreateMember implements MemberService
func (s *memberServiceImpl) CreateMember(request *dto.MemberStoreRequest, ctx context.Context) error {
	member := request.ToModel()
	member.ExpiredAt = time.Now().Add(24 * time.Hour)
	member.Status = model.PENDING
	err := s.memberRepository.CreateMember(member, ctx)
	return err
}

// CreateMemberType implements MemberService
func (s *memberServiceImpl) CreateMemberType(request *dto.MemberTypeStoreRequest, ctx context.Context) error {
	memberType := request.ToModel()
	err := s.memberRepository.CreateMemberType(memberType, ctx)
	return err
}

// DeleteMember implements MemberService
func (s *memberServiceImpl) DeleteMember(id uint, ctx context.Context) error {
	member := model.Member{
		ID: id,
	}
	err := s.memberRepository.DeleteMember(&member, ctx)
	return err
}

// DeleteMemberType implements MemberService
func (s *memberServiceImpl) DeleteMemberType(id uint, ctx context.Context) error {
	memberType := model.MemberType{
		ID: id,
	}
	err := s.memberRepository.DeleteMemberType(&memberType, ctx)
	return err
}

// FindMemberById implements MemberService
func (s *memberServiceImpl) FindMemberById(id uint, ctx context.Context) (*dto.MemberDetailResource, error) {
	member, err := s.memberRepository.FindMemberById(id, ctx)
	if err != nil {
		return nil, err
	}
	if member.Status != model.INACTIVE && time.Now().After(member.ExpiredAt) {
		member.Status = model.INACTIVE
		body := model.Member{
			ID:     member.ID,
			Status: model.INACTIVE,
		}
		err := s.memberRepository.UpdateMember(&body, ctx)
		if err != nil {
			return nil, err
		}
	}
	var result dto.MemberDetailResource
	result.FromModel(member)
	return &result, nil
}

// FindMemberByUser implements MemberService
func (s *memberServiceImpl) FindMemberByUser(userId uint, ctx context.Context) (*dto.MemberResources, error) {
	member, err := s.memberRepository.FindMemberByUser(userId, ctx)
	if err != nil {
		return nil, err
	}
	var result dto.MemberResources
	result.FromModel(member)
	return &result, nil
}

// FindMemberTypes implements MemberService
func (s *memberServiceImpl) FindMemberTypes(ctx context.Context) (*dto.MemberTypeResources, error) {
	memberTypes, err := s.memberRepository.FindMemberTypes(ctx)
	if err != nil {
		return nil, err
	}
	var result dto.MemberTypeResources
	result.FromModel(memberTypes)
	return &result, nil
}

// FindMemberTypeById implements MemberService
func (s *memberServiceImpl) FindMemberTypeById(id uint, ctx context.Context) (*dto.MemberTypeResource, error) {
	memberType, err := s.memberRepository.FindMemberTypeById(id, ctx)
	if err != nil {
		return nil, err
	}
	var result dto.MemberTypeResource
	result.FromModel(memberType)
	return &result, nil
}

// FindMembers implements MemberService
func (s *memberServiceImpl) FindMembers(page *model.Pagination, ctx context.Context) (*dto.MemberResponses, error) {
	members, count, err := s.memberRepository.FindMembers(page, ctx)
	if err != nil {
		return nil, err
	}
	var result dto.MemberResources
	result.FromModel(members)

	response := dto.MemberResponses{
		Members: result,
		Page:    uint(page.Page),
		Limit:   uint(page.Limit),
		Count:   uint(count),
	}
	return &response, nil
}

// UpdateMember implements MemberService
func (s *memberServiceImpl) UpdateMember(request *dto.MemberUpdateRequest, ctx context.Context) error {
	member := request.ToModel()
	check, err := s.memberRepository.FindMemberById(request.ID, ctx)
	if err != nil {
		return err
	}
	if time.Now().After(check.ExpiredAt) {
		member.Status = model.INACTIVE
	}

	err = s.memberRepository.UpdateMember(member, ctx)
	return err
}

// UpdateMemberType implements MemberService
func (s *memberServiceImpl) UpdateMemberType(request *dto.MemberTypeUpdateRequest, ctx context.Context) error {
	memberType := request.ToModel()
	err := s.memberRepository.UpdateMemberType(memberType, ctx)
	return err
}

// SetStatusMember implements MemberService
func (s *memberServiceImpl) SetStatusMember(request *dto.SetStatusMember, ctx context.Context) error {
	member := request.ToModel()
	check, err := s.memberRepository.FindMemberById(request.ID, ctx)
	if err != nil {
		return err
	}

	if member.Status == model.ACTIVE && check.Status != model.ACTIVE && check.Status != model.INACTIVE {
		member.ExpiredAt = time.Now().Add(24 * 30 * time.Duration(check.Duration) * time.Hour)
		member.Code = uuid.New()
		member.ActivedAt = time.Now()
	} else if member.Status == model.REJECT && check.Status != model.REJECT {
		member.ExpiredAt = time.Now()
	}

	if time.Now().After(check.ExpiredAt) {
		member.Status = model.INACTIVE
	}

	err = s.memberRepository.UpdateMember(member, ctx)
	return err
}

// MemberPayment implements MemberService
func (s *memberServiceImpl) MemberPayment(request *dto.PaymMemberStoreRequest, ctx context.Context) error {
	// check member id
	id := request.ID
	member, err := s.memberRepository.FindMemberById(id, ctx)
	if err != nil {
		return err
	}
	if member.UserID != request.UserID {
		return myerrors.ErrPermission
	}
	if member.ProofPayment != "" {
		return errors.New("member already paid")
	}
	// create file buffer
	buf := bytes.NewBuffer(nil)

	if _, err := io.Copy(buf, request.File); err != nil {
		return err
	}
	url, err := s.imagekitService.Upload("member", buf.Bytes())
	if err != nil {
		return err
	}
	if url == "" {
		return myerrors.ErrFailedUpload
	}
	// update member
	body := model.Member{
		ID:           id,
		ProofPayment: url,
		ExpiredAt:    time.Now().Add(24 * time.Hour),
		Status:       model.WAITING,
	}
	s.memberRepository.UpdateMember(&body, ctx)
	if err != nil {
		return err
	}
	// push or create notification
	notif := model.Notification{
		UserID:          member.UserID,
		TransactionID:   id,
		TransactionType: "/members/",
		Title:           "Member",
	}
	if err := s.notificationRepository.CreateNotification(&notif, ctx); err != nil {
		return err
	}
	return nil
}

func NewMemberService(memberRepository memberRepo.MemberRepository, imagekit imgkit.ImagekitService, notificationRepository notifRepo.NotificationRepository) MemberService {
	return &memberServiceImpl{
		memberRepository:       memberRepository,
		notificationRepository: notificationRepository,
		imagekitService:        imagekit,
	}
}
