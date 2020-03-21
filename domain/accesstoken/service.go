package accesstoken

import (
	"gitlab.com/aubayaml/aubayaml-go/bookstore/oauth-api/domain/users"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/utils-go/errors"
)

//UsersRepository rest api access
type UsersRepository interface {
	LoginUser(string, string) (*users.User, errors.RestErr)
}

//Repository interface
type Repository interface {
	GetByID(string) (*AccessToken, errors.RestErr)
	Create(AccessToken) errors.RestErr
	UpdateExpirationTime(AccessToken) errors.RestErr
}

//Service interface
type Service interface {
	GetByID(string) (*AccessToken, errors.RestErr)
	Create(LoginRequest) (*AccessToken, errors.RestErr)
	UpdateExpirationTime(AccessToken) errors.RestErr
}

type service struct {
	repository   Repository
	userRestRepo UsersRepository
}

func (s *service) GetByID(id string) (*AccessToken, errors.RestErr) {
	at := &AccessToken{
		Token: id,
	}

	if err := at.ValidateToken(); err != nil {
		return nil, err
	}

	token, err := s.repository.GetByID(at.Token)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *service) Create(lr LoginRequest) (*AccessToken, errors.RestErr) {
	if err := lr.Validate(); err != nil {
		return nil, err
	}
	user, err := s.userRestRepo.LoginUser(lr.AuthID, lr.AuthSecret)
	if err != nil {
		return nil, err
	}

	at := GetNewAccessToken(user.ID)
	at.Generate()

	if err := s.repository.Create(at); err != nil {
		return nil, err
	}

	return &at, nil
}

func (s *service) UpdateExpirationTime(at AccessToken) errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(at)
}

//NewService Return new service with given repository
func NewService(repo Repository, userRestRepo UsersRepository) Service {
	return &service{
		repository:   repo,
		userRestRepo: userRestRepo,
	}
}
