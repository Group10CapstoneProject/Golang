package dto

import (
	"github.com/Group10CapstoneProject/Golang/model"
)

// member type store and update request
type MemberTypeStoreRequest struct {
	Name               string `json:"name" validate:"required,name"`
	Price              uint   `json:"price" validate:"required,gte=0"`
	Description        string `json:"description,omitempty"`
	AccessOfflineClass bool   `json:"access_offline_class"`
	AccessOnlineClass  bool   `json:"access_online_class"`
	AccessTrainer      bool   `json:"access_trainer"`
	AccessGym          bool   `json:"access_gym"`
}

func (u *MemberTypeStoreRequest) ToModel() *model.MemberType {
	return &model.MemberType{
		Name:               u.Name,
		Price:              u.Price,
		Description:        u.Description,
		AccessOfflineClass: u.AccessOfflineClass,
		AccessOnlineClass:  u.AccessOnlineClass,
		AccessTrainer:      u.AccessTrainer,
		AccessGym:          u.AccessGym,
	}
}

type MemberTypeUpdateRequest struct {
	ID                 uint
	Name               string `json:"name,omitempty" validate:"omitempty,name"`
	Price              uint   `json:"price,omitempty" validate:"omitempty,gte=0"`
	Description        string `json:"description,omitempty"`
	AccessOfflineClass bool   `json:"access_offline_class,omitempty"`
	AccessOnlineClass  bool   `json:"access_online_class,omitempty"`
	AccessTrainer      bool   `json:"access_trainer,omitempty"`
	AccessGym          bool   `json:"access_gym,omitempty"`
}

func (u *MemberTypeUpdateRequest) ToModel() *model.MemberType {
	return &model.MemberType{
		ID:                 u.ID,
		Name:               u.Name,
		Price:              u.Price,
		Description:        u.Description,
		AccessOfflineClass: u.AccessOfflineClass,
		AccessOnlineClass:  u.AccessOnlineClass,
		AccessTrainer:      u.AccessTrainer,
		AccessGym:          u.AccessGym,
	}
}

// member type resource
type MemberTypeResource struct {
	ID                 uint   `json:"id"`
	Name               string `json:"name"`
	Price              uint   `json:"price"`
	Description        string `json:"description"`
	AccessOfflineClass bool   `json:"access_offline_class"`
	AccessOnlineClass  bool   `json:"access_online_class"`
	AccessTrainer      bool   `json:"access_trainer"`
	AccessGym          bool   `json:"access_gym"`
}

func (u *MemberTypeResource) FromModel(m *model.MemberType) {
	u.ID = m.ID
	u.Name = m.Name
	u.Price = m.Price
	u.Description = m.Description
	u.AccessOfflineClass = m.AccessOfflineClass
	u.AccessOnlineClass = m.AccessOnlineClass
	u.AccessTrainer = m.AccessOnlineClass
	u.AccessGym = m.AccessGym
}

type MemberTypeResources []MemberTypeResource

func (u *MemberTypeResources) FromModel(m []model.MemberType) {
	for _, each := range m {
		var resource MemberTypeResource
		resource.FromModel(&each)
		*u = append(*u, resource)
	}
}
