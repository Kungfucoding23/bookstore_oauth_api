package db

import (
	"github.com/Kungfucoding23/bookstore_oauth_api/src/clients/cassandra"
	"github.com/Kungfucoding23/bookstore_oauth_api/src/domain/access_token"
	"github.com/Kungfucoding23/bookstore_oauth_api/src/utils/errors"
)

func NewRepository() DBRepository {
	return &dbRepository{}
}

type DBRepository interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct{}

func (r *dbRepository) GetByID(string) (*access_token.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()
	//TODO: implement get access token from cassandra db
	return nil, errors.NewInternalServerError("database connection not implemented yet")
}
