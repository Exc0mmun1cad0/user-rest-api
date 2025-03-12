package app

import (
	"test-api-task/internal/repository/user/postgres"
	"test-api-task/internal/service/userservice"

	"github.com/jmoiron/sqlx"
)

type Container struct {
	pgsqlx *sqlx.DB
}

func NewContainer(pgSqlxConn *sqlx.DB) *Container {
	return &Container{
		pgsqlx: pgSqlxConn,
	}
}

func (c *Container) GetPgsqlx() *sqlx.DB {
	return c.pgsqlx
}

func (c *Container) GetUserRepo() *postgres.Repo {
	return postgres.NewRepo(c.GetPgsqlx())
}

func (c *Container) GetUserService() *userservice.Service {
	return userservice.NewService(c.GetUserRepo())
}
