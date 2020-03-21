package rest

import (
	"encoding/json"
	"time"

	"github.com/federicoleon/golang-restclient/rest"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/oauth-api/domain/accesstoken"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/oauth-api/domain/users"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/utils-go/errors"
)

var (
	userClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8080",
		Timeout: 100 * time.Millisecond,
	}
)

//NewRepository create new repo
func NewRepository() UsersRepository {
	return &usersRepository{}
}

//UsersRepository rest api access
type UsersRepository interface {
	LoginUser(string, string) (*users.User, errors.RestErr)
}

type usersRepository struct {
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, errors.RestErr) {
	request := accesstoken.LoginRequest{
		AuthID:     email,
		AuthSecret: password,
	}
	response := userClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		return nil, errors.InternalServerError("invalid rest client response when trying to login user", errors.New("oAuth out of reach"))
	}
	if response.StatusCode > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.InternalServerError("invalid error interface when trying to login user", err)
		}
		return nil, restErr
	}
	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.InternalServerError("invalid user interface when trying to login user", err)
	}
	return &user, nil
}
