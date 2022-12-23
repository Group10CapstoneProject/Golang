package repository

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Group10CapstoneProject/Golang/model"
	"github.com/Group10CapstoneProject/Golang/utils/myerrors"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type suiteMemberRepository struct {
	suite.Suite
	mock       sqlmock.Sqlmock
	repository *memberRepositoryImpl
}

func (s *suiteMemberRepository) SetupSuite() {
	db, mocking, _ := sqlmock.New()

	dbGorm, _ := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	s.mock = mocking
	s.repository = &memberRepositoryImpl{
		db: dbGorm,
	}
}

func (s *suiteMemberRepository) TearDown() {
	s.mock = nil
	s.repository = nil
}
func (s *suiteMemberRepository) TestCreateMember() {
	paymentId := uint(1)
	testCase := []struct {
		Name        string
		Body        model.Member
		ExpectedErr error
		MockReturn  error
	}{
		{
			Name: "Create member success",
			Body: model.Member{
				UserID:          1,
				MemberTypeID:    1,
				Duration:        1,
				PaymentMethodID: &paymentId,
				Total:           10000,
			},
			ExpectedErr: nil,
			MockReturn:  nil,
		},
		{
			Name: "Create member failed",
			Body: model.Member{
				UserID:          1,
				MemberTypeID:    1,
				Duration:        1,
				PaymentMethodID: &paymentId,
				Total:           10000,
			},
			ExpectedErr: errors.New("error"),
			MockReturn:  errors.New("error"),
		},
		{
			Name: "Create member invalid foreign key",
			Body: model.Member{
				UserID:          1,
				MemberTypeID:    1,
				Duration:        1,
				PaymentMethodID: &paymentId,
				Total:           10000,
			},
			ExpectedErr: errors.New("invalid `payment_methods` `id`"),
			MockReturn:  errors.New("Error 1452: Cannot add or update a child row: a foreign key constraint fails (`dev_gym_membership`.`members`, CONSTRAINT `fk_members_payment_method` FOREIGN KEY (`payment_method_id`) REFERENCES `payment_methods` (`id`))"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			s.SetupSuite()
			s.mock.ExpectBegin()
			db := s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `members` (`created_at`,`updated_at`,`deleted_at`,`user_id`,`member_type_id`,`expired_at`,`duration`,`status`,`proof_payment`,`payment_method_id`,`total`,`code`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"))
			if v.ExpectedErr != nil {
				db.WillReturnError(v.MockReturn)
				s.mock.ExpectRollback()
			} else {
				db.WillReturnResult(sqlmock.NewResult(1, 1))
				s.mock.ExpectCommit()
			}

			_, err := s.repository.CreateMember(&v.Body, context.Background())

			s.Equal(v.ExpectedErr, err)

			s.TearDown()
		})
	}
}
func (s *suiteMemberRepository) TestFindMembers() {
	paymentId := uint(1)
	testCase := []struct {
		Name                 string
		ExpectedErr          error
		ExpectedRes          []model.Member
		FindMemberErr        error
		FindMemberRes        *sqlmock.Rows
		PreloadUserErr       error
		PreloadUserRes       *sqlmock.Rows
		PreloadMemberTypeErr error
		PreloadMemberTypeRes *sqlmock.Rows
		Page                 model.Pagination
		Count                int
	}{
		{
			Name:        "Find members success",
			ExpectedErr: nil,
			ExpectedRes: []model.Member{
				{
					ID:     1,
					UserID: 1,
					User: model.User{
						ID:    1,
						Name:  "test",
						Email: "test@gmail.com",
					},
					MemberTypeID: 1,
					MemberType: model.MemberType{
						ID:    1,
						Name:  "test",
						Price: 10000,
					},
					Duration:        1,
					Status:          model.ACTIVE,
					PaymentMethodID: &paymentId,
					Total:           10000,
				},
			},
			FindMemberErr: nil,
			FindMemberRes: sqlmock.NewRows([]string{"id", "user_id", "member_type_id", "duration", "status", "payment_method_id", "total"}).
				AddRow(1, 1, 1, 1, model.ACTIVE, 1, 10000),
			PreloadUserErr: nil,
			PreloadUserRes: sqlmock.NewRows([]string{"id", "name", "email"}).
				AddRow(1, "test", "test@gmail.com"),
			PreloadMemberTypeErr: nil,
			PreloadMemberTypeRes: sqlmock.NewRows([]string{"id", "name", "price"}).
				AddRow(1, "test", 10000),
			Page: model.Pagination{
				Page:  1,
				Limit: 10,
				Q:     "",
			},
			Count: 1,
		},
		{
			Name:                 "Find members error",
			ExpectedErr:          errors.New("error"),
			ExpectedRes:          []model.Member{},
			FindMemberErr:        errors.New("error"),
			FindMemberRes:        sqlmock.NewRows([]string{"id", "user_id", "member_type_id", "duration", "status", "payment_method_id", "total"}),
			PreloadUserErr:       nil,
			PreloadUserRes:       sqlmock.NewRows([]string{"id", "name", "email"}),
			PreloadMemberTypeErr: nil,
			PreloadMemberTypeRes: sqlmock.NewRows([]string{"id", "name", "price"}),
			Page: model.Pagination{
				Page:  1,
				Limit: 10,
				Q:     "",
			},
			Count: 0,
		},
	}
	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			s.SetupSuite()

			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `members` LEFT JOIN users ON users.id = members.user_id LEFT JOIN member_types ON member_types.id = members.member_type_id WHERE `members`.`deleted_at` IS NULL")).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT `members`.`id`,`members`.`created_at`,`members`.`updated_at`,`members`.`deleted_at`,`members`.`user_id`,`members`.`member_type_id`,`members`.`expired_at`,`members`.`actived_at`,`members`.`duration`,`members`.`status`,`members`.`proof_payment`,`members`.`payment_method_id`,`members`.`total`,`members`.`code` FROM `members` LEFT JOIN users ON users.id = members.user_id LEFT JOIN member_types ON member_types.id = members.member_type_id WHERE `members`.`deleted_at` IS NULL ORDER BY id DESC LIMIT 10")).
				WillReturnRows(v.FindMemberRes).WillReturnError(v.FindMemberErr)

			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `member_types` WHERE `member_types`.`id` = ? AND `member_types`.`deleted_at` IS NULL")).
				WillReturnRows(v.PreloadMemberTypeRes).WillReturnError(v.PreloadMemberTypeErr)

			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`id` = ? AND `users`.`deleted_at` IS NULL")).
				WillReturnRows(v.PreloadUserRes).WillReturnError(v.PreloadUserErr)

			res, count, err := s.repository.FindMembers(&v.Page, context.Background())

			s.Equal(v.ExpectedErr, err)
			s.Equal(v.ExpectedRes, res)
			s.Equal(v.Count, count)

			s.TearDown()
		})
	}
}
func (s *suiteMemberRepository) TestFindMemberById() {
	paymentId := uint(1)
	testCase := []struct {
		Name                 string
		ID                   uint
		ExpectedErr          error
		ExpectedRes          *model.Member
		FindMemberErr        error
		FindMemberRes        *sqlmock.Rows
		PreloadUserErr       error
		PreloadUserRes       *sqlmock.Rows
		PreloadMemberTypeErr error
		PreloadMemberTypeRes *sqlmock.Rows
		PreloadPaymentErr    error
		PreloadPaymentRes    *sqlmock.Rows
	}{
		{
			Name:        "Find member by id success",
			ID:          1,
			ExpectedErr: nil,
			ExpectedRes: &model.Member{
				ID:     1,
				UserID: 1,
				User: model.User{
					ID:    1,
					Name:  "test",
					Email: "test@gmail.com",
				},
				MemberTypeID: 1,
				MemberType: model.MemberType{
					ID:    1,
					Name:  "test",
					Price: 10000,
				},
				PaymentMethodID: &paymentId,
				PaymentMethod: model.PaymentMethod{
					ID:   &paymentId,
					Name: "test",
				},
				Duration: 1,
				Status:   model.ACTIVE,
				Total:    10000,
			},
			FindMemberErr:        nil,
			FindMemberRes:        sqlmock.NewRows([]string{"id", "user_id", "member_type_id", "duration", "status", "payment_method_id", "total"}).AddRow(1, 1, 1, 1, model.ACTIVE, 1, 10000),
			PreloadUserErr:       nil,
			PreloadUserRes:       sqlmock.NewRows([]string{"id", "name", "email"}).AddRow(1, "test", "test@gmail.com"),
			PreloadMemberTypeErr: nil,
			PreloadMemberTypeRes: sqlmock.NewRows([]string{"id", "name", "price"}).AddRow(1, "test", 10000),
			PreloadPaymentErr:    nil,
			PreloadPaymentRes:    sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "test"),
		},
		{
			Name:                 "Find member by id failed",
			ID:                   1,
			ExpectedErr:          errors.New("error"),
			ExpectedRes:          nil,
			FindMemberErr:        errors.New("error"),
			FindMemberRes:        sqlmock.NewRows([]string{"id", "user_id", "member_type_id", "duration", "status", "payment_method_id", "total"}),
			PreloadUserErr:       nil,
			PreloadUserRes:       sqlmock.NewRows([]string{"id", "name", "email"}),
			PreloadMemberTypeErr: nil,
			PreloadMemberTypeRes: sqlmock.NewRows([]string{"id", "name", "price"}),
			PreloadPaymentErr:    nil,
			PreloadPaymentRes:    sqlmock.NewRows([]string{"id", "name"}),
		},
	}
	for _, v := range testCase {
		s.Run(v.Name, func() {
			s.SetupSuite()

			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `members` WHERE id = ? AND `members`.`deleted_at` IS NULL ORDER BY `members`.`id` LIMIT 1")).
				WillReturnRows(v.FindMemberRes).WillReturnError(v.FindMemberErr)

			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `member_types` WHERE `member_types`.`id` = ? AND `member_types`.`deleted_at` IS NULL")).
				WillReturnRows(v.PreloadMemberTypeRes).WillReturnError(v.PreloadMemberTypeErr)

			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `payment_methods` WHERE `payment_methods`.`id` = ? AND `payment_methods`.`deleted_at` IS NULL")).
				WillReturnRows(v.PreloadPaymentRes).WillReturnError(v.PreloadPaymentErr)

			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`id` = ? AND `users`.`deleted_at` IS NULL")).
				WillReturnRows(v.PreloadUserRes).WillReturnError(v.PreloadUserErr)

			res, err := s.repository.FindMemberById(v.ID, context.Background())

			s.Equal(v.ExpectedErr, err)
			s.Equal(v.ExpectedRes, res)

			s.TearDown()
		})
	}
}
func (s *suiteMemberRepository) TestFindMemberByUser() {

	paymentId := uint(1)
	testCase := []struct {
		Name                 string
		UserID               uint
		ExpectedErr          error
		ExpectedRes          *model.Member
		FindMemberErr        error
		FindMemberRes        *sqlmock.Rows
		PreloadUserErr       error
		PreloadUserRes       *sqlmock.Rows
		PreloadMemberTypeErr error
		PreloadMemberTypeRes *sqlmock.Rows
		PreloadPaymentErr    error
		PreloadPaymentRes    *sqlmock.Rows
	}{
		{
			Name:        "Find member by user success",
			UserID:      1,
			ExpectedErr: nil,
			ExpectedRes: &model.Member{
				ID:     1,
				UserID: 1,
				User: model.User{
					ID:    1,
					Name:  "test",
					Email: "test@gmail.com",
				},
				MemberTypeID: 1,
				MemberType: model.MemberType{
					ID:    1,
					Name:  "test",
					Price: 10000,
				},
				PaymentMethodID: &paymentId,
				PaymentMethod: model.PaymentMethod{
					ID:   &paymentId,
					Name: "test",
				},
				Duration: 1,
				Status:   model.ACTIVE,
				Total:    10000,
			},
			FindMemberErr:        nil,
			FindMemberRes:        sqlmock.NewRows([]string{"id", "user_id", "member_type_id", "duration", "status", "payment_method_id", "total"}).AddRow(1, 1, 1, 1, model.ACTIVE, 1, 10000),
			PreloadUserErr:       nil,
			PreloadUserRes:       sqlmock.NewRows([]string{"id", "name", "email"}).AddRow(1, "test", "test@gmail.com"),
			PreloadMemberTypeErr: nil,
			PreloadMemberTypeRes: sqlmock.NewRows([]string{"id", "name", "price"}).AddRow(1, "test", 10000),
			PreloadPaymentErr:    nil,
			PreloadPaymentRes:    sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "test"),
		},
		{
			Name:                 "Find member by user failed",
			UserID:               1,
			ExpectedErr:          errors.New("error"),
			ExpectedRes:          nil,
			FindMemberErr:        errors.New("error"),
			FindMemberRes:        sqlmock.NewRows([]string{"id", "user_id", "member_type_id", "duration", "status", "payment_method_id", "total"}),
			PreloadUserErr:       nil,
			PreloadUserRes:       sqlmock.NewRows([]string{"id", "name", "email"}),
			PreloadMemberTypeErr: nil,
			PreloadMemberTypeRes: sqlmock.NewRows([]string{"id", "name", "price"}),
			PreloadPaymentErr:    nil,
			PreloadPaymentRes:    sqlmock.NewRows([]string{"id", "name"}),
		},
	}
	for _, v := range testCase {
		s.Run(v.Name, func() {
			s.SetupSuite()

			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `members` WHERE (user_id = ? AND status = ?) AND `members`.`deleted_at` IS NULL ORDER BY `members`.`id` LIMIT 1")).
				WillReturnRows(v.FindMemberRes).WillReturnError(v.FindMemberErr)

			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `member_types` WHERE `member_types`.`id` = ? AND `member_types`.`deleted_at` IS NULL")).
				WillReturnRows(v.PreloadMemberTypeRes).WillReturnError(v.PreloadMemberTypeErr)

			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `payment_methods` WHERE `payment_methods`.`id` = ? AND `payment_methods`.`deleted_at` IS NULL")).
				WillReturnRows(v.PreloadPaymentRes).WillReturnError(v.PreloadPaymentErr)

			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`id` = ? AND `users`.`deleted_at` IS NULL")).
				WillReturnRows(v.PreloadUserRes).WillReturnError(v.PreloadUserErr)

			res, err := s.repository.FindMemberByUser(v.UserID, context.Background())

			s.Equal(v.ExpectedErr, err)
			s.Equal(v.ExpectedRes, res)

			s.TearDown()
		})
	}
}
func (s *suiteMemberRepository) TestReadMembers() {
	paymentId := uint(1)
	testCase := []struct {
		Name                 string
		Body                 *model.Member
		ExpectedErr          error
		ExpectedRes          []model.Member
		FindMemberErr        error
		FindMemberRes        *sqlmock.Rows
		PreloadUserErr       error
		PreloadUserRes       *sqlmock.Rows
		PreloadMemberTypeErr error
		PreloadMemberTypeRes *sqlmock.Rows
		PreloadPaymentErr    error
		PreloadPaymentRes    *sqlmock.Rows
	}{
		{
			Name: "Read members success",
			Body: &model.Member{
				ID:     1,
				UserID: 1,
			},
			ExpectedErr: nil,
			ExpectedRes: []model.Member{
				{
					ID:     1,
					UserID: 1,
					User: model.User{
						ID:    1,
						Name:  "test",
						Email: "test@gmail.com",
					},
					MemberTypeID: 1,
					MemberType: model.MemberType{
						ID:    1,
						Name:  "test",
						Price: 10000,
					},
					PaymentMethodID: &paymentId,
					PaymentMethod: model.PaymentMethod{
						ID:   &paymentId,
						Name: "test",
					},
					Duration: 1,
					Status:   model.ACTIVE,
					Total:    10000,
				},
			},
			FindMemberErr: nil,
			FindMemberRes: sqlmock.NewRows([]string{"id", "user_id", "member_type_id", "duration", "status", "payment_method_id", "total"}).
				AddRow(1, 1, 1, 1, model.ACTIVE, 1, 10000),
			PreloadUserErr: nil,
			PreloadUserRes: sqlmock.NewRows([]string{"id", "name", "email"}).
				AddRow(1, "test", "test@gmail.com"),
			PreloadMemberTypeErr: nil,
			PreloadMemberTypeRes: sqlmock.NewRows([]string{"id", "name", "price"}).
				AddRow(1, "test", 10000),
			PreloadPaymentErr: nil,
			PreloadPaymentRes: sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "test"),
		},
		{
			Name: "Read members error",
			Body: &model.Member{
				ID:     1,
				UserID: 1,
			},
			ExpectedErr:          errors.New("error"),
			ExpectedRes:          []model.Member(nil),
			FindMemberErr:        errors.New("error"),
			FindMemberRes:        sqlmock.NewRows([]string{"id", "user_id", "member_type_id", "duration", "status", "payment_method_id", "total"}),
			PreloadUserErr:       nil,
			PreloadUserRes:       sqlmock.NewRows([]string{"id", "name", "email"}),
			PreloadMemberTypeErr: nil,
			PreloadMemberTypeRes: sqlmock.NewRows([]string{"id", "name", "price"}),
			PreloadPaymentErr:    nil,
			PreloadPaymentRes:    sqlmock.NewRows([]string{"id", "name"}),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			s.SetupSuite()

			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `members` WHERE `members`.`id` = ? AND `members`.`user_id` = ? AND `members`.`deleted_at` IS NULL ORDER BY updated_at DESC")).
				WillReturnRows(v.FindMemberRes).WillReturnError(v.FindMemberErr)

			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `member_types` WHERE `member_types`.`id` = ? AND `member_types`.`deleted_at` IS NULL")).
				WillReturnRows(v.PreloadMemberTypeRes).WillReturnError(v.PreloadMemberTypeErr)

			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `payment_methods` WHERE `payment_methods`.`id` = ? AND `payment_methods`.`deleted_at` IS NULL")).
				WillReturnRows(v.PreloadPaymentRes).WillReturnError(v.PreloadPaymentErr)

			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`id` = ? AND `users`.`deleted_at` IS NULL")).
				WillReturnRows(v.PreloadUserRes).WillReturnError(v.PreloadUserErr)

			res, err := s.repository.ReadMembers(v.Body, context.Background())

			s.Equal(v.ExpectedErr, err)
			s.Equal(v.ExpectedRes, res)

			s.TearDown()
		})
	}
}
func (s *suiteMemberRepository) TestUpdateMember() {
	paymenId := uint(1)
	testCase := []struct {
		Name            string
		Body            *model.Member
		ExpectedErr     error
		UpdateMemberErr error
		RowAffacted     int64
	}{
		{
			Name: "Update member success",
			Body: &model.Member{
				ID:              1,
				UserID:          1,
				Total:           10000,
				PaymentMethodID: &paymenId,
			},
			ExpectedErr:     nil,
			UpdateMemberErr: nil,
			RowAffacted:     1,
		},
		{
			Name: "Update member error",
			Body: &model.Member{
				ID:              1,
				UserID:          1,
				Total:           10000,
				PaymentMethodID: &paymenId,
			},
			ExpectedErr:     errors.New("error"),
			UpdateMemberErr: errors.New("error"),
			RowAffacted:     0,
		},
		{
			Name: "Update member error foreign key",
			Body: &model.Member{
				ID:              1,
				UserID:          1,
				Total:           10000,
				PaymentMethodID: &paymenId,
			},
			ExpectedErr:     errors.New("invalid `payment_methods` `id`"),
			UpdateMemberErr: errors.New("Error 1452: Cannot add or update a child row: a foreign key constraint fails (`dev_gym_membership`.`members`, CONSTRAINT `fk_members_payment_method` FOREIGN KEY (`payment_method_id`) REFERENCES `payment_methods` (`id`))"),
			RowAffacted:     0,
		},
		{
			Name: "Update member no row effected",
			Body: &model.Member{
				ID:              1,
				UserID:          1,
				Total:           10000,
				PaymentMethodID: &paymenId,
			},
			ExpectedErr:     myerrors.ErrRecordNotFound,
			UpdateMemberErr: nil,
			RowAffacted:     0,
		},
	}
	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			s.SetupSuite()
			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `members` SET `updated_at`=?,`user_id`=?,`payment_method_id`=?,`total`=? WHERE `members`.`deleted_at` IS NULL AND `id` = ?")).
				WillReturnResult(sqlmock.NewResult(1, v.RowAffacted)).WillReturnError(v.UpdateMemberErr)
			if v.UpdateMemberErr == nil {
				s.mock.ExpectCommit()
			} else {
				s.mock.ExpectRollback()
			}

			err := s.repository.UpdateMember(v.Body, context.Background())

			s.Equal(v.ExpectedErr, err)

			s.TearDown()
		})
	}
}
func (s *suiteMemberRepository) TestDeleteMember() {
	testCase := []struct {
		Name            string
		Body            *model.Member
		ExpectedErr     error
		DeleteMemberErr error
		RowAffacted     int64
	}{
		{
			Name: "Delete member success",
			Body: &model.Member{
				ID: 1,
			},
			ExpectedErr:     nil,
			DeleteMemberErr: nil,
			RowAffacted:     1,
		},
		{
			Name: "Delete member error",
			Body: &model.Member{
				ID: 1,
			},
			ExpectedErr:     errors.New("error"),
			DeleteMemberErr: errors.New("error"),
			RowAffacted:     0,
		},
		{
			Name: "Delete member no row effected",
			Body: &model.Member{
				ID: 1,
			},
			ExpectedErr:     myerrors.ErrRecordNotFound,
			DeleteMemberErr: nil,
			RowAffacted:     0,
		},
	}
	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			s.SetupSuite()
			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `members` SET `deleted_at`=? WHERE `members`.`id` = ? AND `members`.`deleted_at` IS NULL")).
				WillReturnResult(sqlmock.NewResult(1, v.RowAffacted)).WillReturnError(v.DeleteMemberErr)
			if v.DeleteMemberErr == nil {
				s.mock.ExpectCommit()
			} else {
				s.mock.ExpectRollback()
			}

			err := s.repository.DeleteMember(v.Body, context.Background())

			s.Equal(v.ExpectedErr, err)

			s.TearDown()
		})
	}
}
func (s *suiteMemberRepository) TestMemberInactive() {
	testCase := []struct {
		Name              string
		Body              *model.Member
		ExpectedErr       error
		MemberInactiveErr error
		RowAffacted       int64
		CheckErr          error
	}{
		{
			Name: "Member inactive success",
			Body: &model.Member{
				ID: 1,
			},
			ExpectedErr:       nil,
			MemberInactiveErr: nil,
			CheckErr:          nil,
			RowAffacted:       1,
		},
		{
			Name: "Member inactive error",
			Body: &model.Member{
				ID: 1,
			},
			ExpectedErr:       errors.New("error"),
			MemberInactiveErr: errors.New("error"),
			RowAffacted:       0,
			CheckErr:          nil,
		},
		{
			Name: "Member inactive no row effected",
			Body: &model.Member{
				ID: 1,
			},
			ExpectedErr:       myerrors.ErrRecordNotFound,
			MemberInactiveErr: nil,
			RowAffacted:       0,
			CheckErr:          gorm.ErrRecordNotFound,
		},
		{
			Name: "Check error",
			Body: &model.Member{
				ID: 1,
			},
			ExpectedErr:       errors.New("error"),
			MemberInactiveErr: nil,
			RowAffacted:       0,
			CheckErr:          errors.New("error"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			s.SetupSuite()
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `members` WHERE `members`.`deleted_at` IS NULL AND `members`.`id` = ? ORDER BY `members`.`id` LIMIT 1")).WillReturnError(v.CheckErr).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "test"))

			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `members` SET `status`=?,`updated_at`=? WHERE (user_id = ? AND status = ?) AND `members`.`deleted_at` IS NULL")).
				WillReturnResult(sqlmock.NewResult(1, v.RowAffacted)).WillReturnError(v.MemberInactiveErr)
			if v.MemberInactiveErr == nil {
				s.mock.ExpectCommit()
			} else {
				s.mock.ExpectRollback()
			}

			err := s.repository.MemberInactive(*v.Body, context.Background())

			s.Equal(v.ExpectedErr, err)

			s.TearDown()
		})
	}
}
func (s *suiteMemberRepository) TestCreateMemberType() {
	testCase := []struct {
		Name        string
		Body        model.MemberType
		ExpectedErr error
		MockReturn  error
		CheckErr    error
	}{
		{
			Name: "Create member type success",
			Body: model.MemberType{
				Name:        "test",
				Description: "test",
				Price:       10000,
			},
			ExpectedErr: nil,
			MockReturn:  nil,
			CheckErr:    errors.New("error"),
		},
		{
			Name: "Create member failed",
			Body: model.MemberType{
				Name:        "test",
				Description: "test",
				Price:       10000,
			},
			ExpectedErr: errors.New("error"),
			MockReturn:  errors.New("error"),
			CheckErr:    errors.New("error"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			s.SetupSuite()
			s.mock.ExpectBegin()

			s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `member_types` (`created_at`,`updated_at`,`deleted_at`,`name`,`price`,`description`,`picture`,`access_offline_class`,`access_online_class`,`access_trainer`,`access_gym`) VALUES (?,?,?,?,?,?,?,?,?,?,?)")).WillReturnError(v.MockReturn).WillReturnResult(sqlmock.NewResult(1, 1))
			if v.MockReturn != nil {
				s.mock.ExpectRollback()
			} else {
				s.mock.ExpectCommit()
			}

			err := s.repository.CreateMemberType(&v.Body, context.Background())

			s.Equal(v.ExpectedErr, err)

			s.TearDown()
		})
	}
}
func (s *suiteMemberRepository) TestFindMemberTypes() {
	testCase := []struct {
		Name          string
		ExpectedErr   error
		ExpecResult   []model.MemberType
		MockReturn    error
		FindMemberRes *sqlmock.Rows
	}{
		{
			Name:        "Find member type success",
			ExpectedErr: nil,
			ExpecResult: []model.MemberType{
				{
					ID:   1,
					Name: "test",
				},
			},
			MockReturn:    nil,
			FindMemberRes: sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "test"),
		},
		{
			Name:          "Find member type failed",
			ExpectedErr:   errors.New("error"),
			ExpecResult:   nil,
			MockReturn:    errors.New("error"),
			FindMemberRes: sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "test"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			s.SetupSuite()
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `member_types` WHERE `member_types`.`deleted_at` IS NULL")).WillReturnError(v.MockReturn).WillReturnRows(v.FindMemberRes)

			result, err := s.repository.FindMemberTypes(context.Background())

			s.Equal(v.ExpectedErr, err)
			s.Equal(v.ExpecResult, result)

			s.TearDown()
		})
	}
}
func (s *suiteMemberRepository) TestFindMemberTypeById() {
	testCase := []struct {
		Name          string
		ID            uint
		ExpectedErr   error
		ExpecResult   *model.MemberType
		MockReturn    error
		FindMemberRes *sqlmock.Rows
	}{
		{
			Name:        "Find member type by id success",
			ID:          1,
			ExpectedErr: nil,
			ExpecResult: &model.MemberType{
				ID:   1,
				Name: "test",
			},
			MockReturn:    nil,
			FindMemberRes: sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "test"),
		},
		{
			Name:          "Find member type id failed",
			ID:            1,
			ExpectedErr:   errors.New("error"),
			ExpecResult:   nil,
			MockReturn:    errors.New("error"),
			FindMemberRes: sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "test"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			s.SetupSuite()
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `member_types` WHERE id = ? AND `member_types`.`deleted_at` IS NULL ORDER BY `member_types`.`id` LIMIT 1")).WillReturnError(v.MockReturn).WillReturnRows(v.FindMemberRes)

			result, err := s.repository.FindMemberTypeById(v.ID, context.Background())

			s.Equal(v.ExpectedErr, err)
			s.Equal(v.ExpecResult, result)

			s.TearDown()
		})
	}
}
func (s *suiteMemberRepository) TestUpdateMemberType() {
	testCase := []struct {
		Name                string
		Body                *model.MemberType
		ExpectedErr         error
		UpdateMemberTypeErr error
		RowAffacted         int64
	}{
		{
			Name: "Update member type success",
			Body: &model.MemberType{
				ID:   1,
				Name: "test",
			},
			ExpectedErr:         nil,
			UpdateMemberTypeErr: nil,
			RowAffacted:         1,
		},
		{
			Name: "Update member type error",
			Body: &model.MemberType{
				ID:   1,
				Name: "test",
			},
			ExpectedErr:         errors.New("error"),
			UpdateMemberTypeErr: errors.New("error"),
			RowAffacted:         0,
		},
		{
			Name: "Update member no row effected",
			Body: &model.MemberType{
				ID:   1,
				Name: "test",
			},
			ExpectedErr:         myerrors.ErrRecordNotFound,
			UpdateMemberTypeErr: nil,
			RowAffacted:         0,
		},
	}
	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			s.SetupSuite()
			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `member_types` SET `updated_at`=?,`name`=? WHERE `member_types`.`deleted_at` IS NULL AND `id` = ?")).
				WillReturnResult(sqlmock.NewResult(1, v.RowAffacted)).WillReturnError(v.UpdateMemberTypeErr)
			if v.UpdateMemberTypeErr == nil {
				s.mock.ExpectCommit()
			} else {
				s.mock.ExpectRollback()
			}

			err := s.repository.UpdateMemberType(v.Body, context.Background())

			s.Equal(v.ExpectedErr, err)

			s.TearDown()
		})
	}
}

func TestSuiteItemRepository(t *testing.T) {
	suite.Run(t, new(suiteMemberRepository))
}
