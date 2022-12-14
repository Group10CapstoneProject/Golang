package dto

import (
	"github.com/Group10CapstoneProject/Golang/model"
)

type HistoryResource struct {
	TransactionID uint             `json:"transaction_id"`
	Type          string           `json:"type"`
	Status        model.StatusType `json:"status"`
	CreatedAt     string           `json:"created_at"`
	ProductID     uint             `json:"product_id"`
	ProducName    string           `json:"product_name"`
	Total         uint             `json:"total"`
}

func (u *HistoryResource) FromModelMembers(m *model.Member) {
	u.TransactionID = m.ID
	u.Type = "member"
	u.Status = m.Status
	u.CreatedAt = m.CreatedAt.Format("2006-01-02 15:04:05")
	u.ProductID = m.MemberTypeID
	u.ProducName = m.MemberType.Name
	u.Total = m.Total
}

func (u *HistoryResource) FromModelOfflineClass(m *model.OfflineClassBooking) {
	u.TransactionID = m.ID
	u.Type = "offline_class"
	u.Status = m.Status
	u.CreatedAt = m.CreatedAt.Format("2006-01-02 15:04:05")
	u.ProductID = m.OfflineClassID
	u.ProducName = m.OfflineClass.Title
	u.Total = m.Total
}

func (u *HistoryResource) FromModelOnlineClass(m *model.OnlineClassBooking) {
	u.TransactionID = m.ID
	u.Type = "online_class"
	u.Status = m.Status
	u.CreatedAt = m.CreatedAt.Format("2006-01-02 15:04:05")
	u.ProductID = m.OnlineClassID
	u.ProducName = m.OnlineClass.Title
	u.Total = m.Total
}
