package rest

import (
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/federicoleon/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	rest.FlushMockups()

	httpCode, _ := strconv.Atoi(os.Getenv("test.oauth.users.repository.timeout.resp.code"))

	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          os.Getenv("test.oauth.users.repository.url"),
		ReqBody:      `{"email":"email@gmail.com","password":"lero lero"}`,
		RespHTTPCode: httpCode,
	})

	repo := usersRepository{}

	user, err := repo.LoginUser("email@gmail.com", "lero lero")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "invalid rest client response when trying to login user", err.Message())
}

func TestLoginUserInvalidErrInterface(t *testing.T) {
	rest.FlushMockups()

	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          os.Getenv("test.oauth.users.repository.url"),
		ReqBody:      `{"email":"email@gmail.com","password":"lero lero"}`,
		RespHTTPCode: 400,
		RespBody:     `{"message": "invalid login credentials", "status":"404", "error": "not_found"}`,
	})

	repo := usersRepository{}

	user, err := repo.LoginUser("email@gmail.com", "lero lero")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "invalid error interface when trying to login user", err.Message())
}

func TestLoginUserInvalidCredentials(t *testing.T) {
	rest.FlushMockups()

	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          os.Getenv("test.oauth.users.repository.url"),
		ReqBody:      `{"email":"email@gmail.com","password":"lero lero"}`,
		RespHTTPCode: 404,
		RespBody:     `{"message": "invalid login credentials", "status":404, "error": "not_found"}`,
	})

	repo := usersRepository{}

	user, err := repo.LoginUser("email@gmail.com", "lero lero")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status())
	assert.EqualValues(t, "invalid login credentials", err.Message())

}
func TestLoginUserInvalidUserJsonResponse(t *testing.T) {
	rest.FlushMockups()

	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          os.Getenv("test.oauth.users.repository.url"),
		ReqBody:      `{"email":"email@gmail.com","password":"lero lero"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": "1", "first_name": "FirstName", "last_name": "LastName", "email": "email@gmail.com"}`,
	})

	repo := usersRepository{}

	user, err := repo.LoginUser("email@gmail.com", "lero lero")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "invalid user interface when trying to login user", err.Message())
}
func TestLoginUserNoError(t *testing.T) {
	rest.FlushMockups()

	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          os.Getenv("test.oauth.users.repository.url"),
		ReqBody:      `{"email":"email@gmail.com","password":"lero lero"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": 1, "first_name": "FirstName", "last_name": "LastName", "email": "email@gmail.com"}`,
	})

	repo := usersRepository{}

	user, err := repo.LoginUser("email@gmail.com", "lero lero")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1, user.ID)
	assert.EqualValues(t, "FirstName", user.FirstName)
	assert.EqualValues(t, "LastName", user.LastName)
	assert.EqualValues(t, "email@gmail.com", user.Email)
}
