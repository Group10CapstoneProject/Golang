package service

import (
	"context"

	"github.com/Group10CapstoneProject/Golang/constans"
	"github.com/Group10CapstoneProject/Golang/internal/dashboard/dto"
	memberRepo "github.com/Group10CapstoneProject/Golang/internal/members/repository"
	offlineClassRepo "github.com/Group10CapstoneProject/Golang/internal/offline_classes/repository"
	onlineClassRepo "github.com/Group10CapstoneProject/Golang/internal/online_classes/repository"
	trainerRepo "github.com/Group10CapstoneProject/Golang/internal/trainers/repository"
	userRepo "github.com/Group10CapstoneProject/Golang/internal/users/repository"
	"github.com/Group10CapstoneProject/Golang/model"
)

type dashboardServiceImpl struct {
	userRepository         userRepo.UserRepository
	memberRepository       memberRepo.MemberRepository
	offlineClassRepository offlineClassRepo.OfflineClassRepository
	onlineClassRepository  onlineClassRepo.OnlineClassRepository
	trainerRepository      trainerRepo.TrainerRepository
}

// GetDashboard implements DashboardService
func (s *dashboardServiceImpl) GetDashboard(ctx context.Context) (*dto.DashboardDto, error) {
	_, user, err := s.userRepository.FindUsers(&model.Pagination{Limit: 100000}, constans.Role_user, ctx)
	if err != nil {
		return nil, err
	}
	memberData, err := s.memberRepository.ReadMembers(&model.Member{Status: model.ACTIVE}, ctx)
	if err != nil {
		return nil, err
	}
	member := len(memberData)
	memberTypeData, err := s.memberRepository.FindMemberTypes(ctx)
	if err != nil {
		return nil, err
	}
	memberType := len(memberTypeData)
	offlineClassData, err := s.offlineClassRepository.FindOfflineClasses(&model.OfflineClass{}, "", ctx)
	if err != nil {
		return nil, err
	}
	offlineClass := len(offlineClassData)
	onlineClassData, err := s.onlineClassRepository.FindOnlineClasses("", ctx)
	if err != nil {
		return nil, err
	}
	onlineClass := len(onlineClassData)
	trainerData, err := s.trainerRepository.FindTrainers(&model.Trainer{}, "", "", ctx)
	if err != nil {
		return nil, err
	}
	trainer := len(trainerData)
	result := dto.DashboardDto{
		TotalUser:         user,
		TotalMember:       member,
		TotalMemberType:   memberType,
		TotalOfflineClass: offlineClass,
		TotalOnlineClass:  onlineClass,
		TotalTrainer:      trainer,
	}
	return &result, nil
}

func NewDashboardService(userRepository userRepo.UserRepository, memberRepository memberRepo.MemberRepository, offlineClassRepository offlineClassRepo.OfflineClassRepository, onlineClassRepository onlineClassRepo.OnlineClassRepository, trainerRepository trainerRepo.TrainerRepository) DashboardService {
	return &dashboardServiceImpl{
		userRepository:         userRepository,
		memberRepository:       memberRepository,
		offlineClassRepository: offlineClassRepository,
		onlineClassRepository:  onlineClassRepository,
		trainerRepository:      trainerRepository,
	}
}
