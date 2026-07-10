package svc

import (
	"database/sql"
	"fmt"

	"github.com/Yooz-1999/agentforge/apps/core-rpc/internal/config"
	"github.com/Yooz-1999/agentforge/apps/core-rpc/internal/repository"
	_ "github.com/go-sql-driver/mysql"
)

type ServiceContext struct {
	Config   config.Config
	DB       *sql.DB
	UserRepo *repository.UserRepository
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := sql.Open("mysql", c.MySQL.DataSource)
	if err != nil {
		panic(fmt.Sprintf("open mysql failed: %v", err))
	}

	if err := db.Ping(); err != nil {
		panic(fmt.Sprintf("ping mysql failed: %v", err))
	}

	return &ServiceContext{
		Config:   c,
		DB:       db,
		UserRepo: repository.NewUserRepository(db),
	}
}
