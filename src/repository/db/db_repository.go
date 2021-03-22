package db

import (
	"errors"

	"github.com/Kungfucoding23/bookstore_oauth_api/src/clients/cassandra"
	"github.com/Kungfucoding23/bookstore_oauth_api/src/domain/access_token"
	"github.com/Kungfucoding23/bookstore_utils-go/rest_errors"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_token WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_token(access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	queryUpdateExpires     = "UPDATE access_token SET expires=? WHERE access_token=?;"
)

func NewRepository() DBRepository {
	return &dbRepository{}
}

type DBRepository interface {
	GetByID(string) (*access_token.AccessToken, rest_errors.RestErr)
	Create(access_token.AccessToken) rest_errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) rest_errors.RestErr
}

type dbRepository struct{}

func (r *dbRepository) GetByID(id string) (*access_token.AccessToken, rest_errors.RestErr) {
	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(&result.Token, &result.UserID, &result.ClientID, &result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, rest_errors.NewNotFoundError("no access token found with given id")
		}
		return nil, rest_errors.NewInternalServerError("internal error", errors.New("database error"))
	}
	return &result, nil
}

func (r *dbRepository) Create(at access_token.AccessToken) rest_errors.RestErr {
	if err := cassandra.GetSession().Query(queryCreateAccessToken, at.Token, at.UserID, at.ClientID, at.Expires).Exec(); err != nil {
		return rest_errors.NewInternalServerError("internal error", errors.New("database error"))
	}
	return nil
}

func (r *dbRepository) UpdateExpirationTime(at access_token.AccessToken) rest_errors.RestErr {
	if err := cassandra.GetSession().Query(queryUpdateExpires, at.Expires, at.Token).Exec(); err != nil {
		return rest_errors.NewInternalServerError("internal error", errors.New("database error"))
	}
	return nil
}
