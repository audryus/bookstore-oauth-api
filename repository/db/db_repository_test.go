package db

import (
	"testing"

	"gitlab.com/aubayaml/aubayaml-go/bookstore/oauth-api/clients/cassandra"
)

func BenchmarkGetByID80(b *testing.B) {
	s := cassandra.GetSession()
	repo := &dbRepository{
		Session: s,
	}
	repo.GetByID("abc123")
}
