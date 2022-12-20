package model

type StatusType string

const (
	PENDING  StatusType = "PENDING"
	WAITING  StatusType = "WAITING"
	ACTIVE   StatusType = "ACTIVE"
	INACTIVE StatusType = "INACTIVE"
	REJECT   StatusType = "REJECT"
	DONE     StatusType = "DONE"
	CANCEL   StatusType = "CANCEL"
)
