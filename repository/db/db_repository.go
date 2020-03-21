package db

import (
	"github.com/gocql/gocql"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/oauth-api/domain/accesstoken"
	"gitlab.com/aubayaml/aubayaml-go/bookstore/utils-go/errors"
)

//NewRepository create new repo
func NewRepository(s *gocql.Session) CassandraRepository {
	return &dbRepository{
		Session: s,
	}
}

//CassandraRepository interface to cassandra
type CassandraRepository interface {
	GetByID(string) (*accesstoken.AccessToken, errors.RestErr)
	Create(accesstoken.AccessToken) errors.RestErr
	UpdateExpirationTime(accesstoken.AccessToken) errors.RestErr
}

type dbRepository struct {
	Session *gocql.Session
}

const (
	queryGetAccessToken = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token = ?"
)

func (r *dbRepository) GetByID(ID string) (*accesstoken.AccessToken, errors.RestErr) {
	var result accesstoken.AccessToken
	if err := r.Session.Query(queryGetAccessToken, ID).Scan(&result.Token, &result.UserID, &result.ClientID, &result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NotFoundError("no access token for given ID")
		}
		return nil, errors.InternalServerError("database error", err)
	}

	return &result, nil
}

const (
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?)"
)

func (r *dbRepository) Create(at accesstoken.AccessToken) errors.RestErr {
	if err := r.Session.Query(queryCreateAccessToken, at.Token, at.UserID, at.ClientID, at.Expires).Exec(); err != nil {
		return errors.InternalServerError("database error", err)
	}

	return nil
}

const (
	queryUpdateExpires = "UPDATE access_tokens SET expires = ? WHERE access_token = ?"
)

func (r *dbRepository) UpdateExpirationTime(at accesstoken.AccessToken) errors.RestErr {
	if err := r.Session.Query(queryUpdateExpires, at.Expires, at.Token).Exec(); err != nil {
		return errors.InternalServerError("database error", err)
	}

	return nil
}
