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

type suiteUserRepository struct {
	suite.Suite
	mock       sqlmock.Sqlmock
	repository *userRepositoryImpl
}

func (s *suiteUserRepository) SetupSuite() {
	db, mocking, _ := sqlmock.New()

	dbGorm, _ := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	s.mock = mocking
	s.repository = &userRepositoryImpl{
		db: dbGorm,
	}
}

func (s *suiteUserRepository) TearDown() {
	s.mock = nil
	s.repository = nil
}

func (s *suiteUserRepository) TestCreateUser() {
	testCase := []struct {
		Name        string
		Body        model.User
		ExpectedErr error
		MockReturn  error
	}{
		{
			Name: "success",
			Body: model.User{
				Name:     "test",
				Email:    "test@gmail.com",
				Password: "testtest",
				Role:     2,
			},
			ExpectedErr: nil,
			MockReturn:  nil,
		},
		{
			Name: "error",
			Body: model.User{
				Name:     "test",
				Email:    "test@gmail.com",
				Password: "testtest",
				Role:     2,
			},
			ExpectedErr: errors.New("other error"),
			MockReturn:  errors.New("other error"),
		},
		{
			Name: "duplicate email",
			Body: model.User{
				Name:     "test",
				Email:    "test@gmail.com",
				Password: "testtest",
				Role:     2,
			},
			ExpectedErr: myerrors.ErrEmailAlredyExist,
			MockReturn:  errors.New("Duplicate entry"),
		},
	}

	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			s.SetupSuite()

			// set mock result or error
			s.mock.ExpectBegin()
			db := s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`email`,`password`,`role`,`member_id`,`session_id`) VALUES (?,?,?,?,?,?,?,?,?)"))
			if v.ExpectedErr != nil {
				db.WillReturnError(v.MockReturn)
				s.mock.ExpectRollback()
			} else {
				db.WillReturnResult(sqlmock.NewResult(1, 1))
				s.mock.ExpectCommit()
			}

			err := s.repository.CreateUser(&v.Body, context.Background())

			s.Equal(v.ExpectedErr, err)

			s.TearDown()
		})
	}
}

func (s *suiteUserRepository) TestCheckUserIsEmpty() {
	testCase := []struct {
		Name        string
		ExpectedErr error
		ExpectedRes bool
		CheckRes    sqlmock.Rows
		CheckErr    error
	}{
		{
			Name:        "success",
			ExpectedErr: nil,
			ExpectedRes: false,
			CheckRes: *sqlmock.NewRows([]string{"id", "name", "email", "password"}).
				AddRow(1, "test", "test@gmail.com", "123456"),
			CheckErr: nil,
		},
		{
			Name:        "is empty",
			ExpectedErr: nil,
			ExpectedRes: true,
			CheckRes:    *sqlmock.NewRows([]string{"id", "name", "email", "password"}),
			CheckErr:    gorm.ErrRecordNotFound,
		},
		{
			Name:        "error",
			ExpectedErr: errors.New("error"),
			ExpectedRes: false,
			CheckRes:    *sqlmock.NewRows([]string{"id", "name", "email", "password"}),
			CheckErr:    errors.New("error"),
		},
	}

	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			s.SetupSuite()

			// set mock result or error
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1")).
				WillReturnError(v.CheckErr).
				WillReturnRows(&v.CheckRes)

			res, err := s.repository.CheckUserIsEmpty(context.Background())

			s.Equal(v.ExpectedErr, err)
			s.Equal(v.ExpectedRes, res)

			s.TearDown()
		})
	}
}

func TestSuiteUserRepository(t *testing.T) {
	suite.Run(t, new(suiteUserRepository))
}
