package controller

import "github.com/labstack/echo/v4"

type MemberController interface {
	// member
	CreateMember(c echo.Context) error
	CreateMemberAdmin(c echo.Context) error
	GetMembers(c echo.Context) error
	GetMemberDetail(c echo.Context) error
	GetMemberUser(c echo.Context) error
	UpdateMember(c echo.Context) error
	SetStatusMember(c echo.Context) error
	MemberPayment(c echo.Context) error
	DeleteMember(c echo.Context) error
	CancelMember(c echo.Context) error

	// member type
	CreateMemberType(c echo.Context) error
	GetMemberTypes(c echo.Context) error
	GetMemberTypeDetail(c echo.Context) error
	UpdateMemberType(c echo.Context) error
	DeleteMemberType(c echo.Context) error
}
