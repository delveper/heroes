package nurepo

import (
	"context"
	"fmt"
	"sync"

	"github.com/delveper/heroes/cfg"
	"github.com/jackc/pgx/v4"
)

type Keeper struct {
	conn *pgx.Conn
	sync sync.RWMutex
}

func NewKeeper(opt *cfg.Options) (*Keeper, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		opt.Repo.Host, opt.Repo.Port, opt.Repo.UserName, opt.Repo.Password, opt.Repo.DbName)

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	return &Keeper{conn: conn}, nil
}
