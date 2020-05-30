package repository

import (
	"github.com/jmoiron/sqlx"
	"os"
)

func New() (*Repository, error) {
	dbString := ""
	if len(os.Getenv("ENV_CLOUD")) > 0 {
		dbString = "host=/cloudsql/aerobic-smithy-271612:us-central1:galleydb sslmode=disable port=5432 user=postgres dbname=dev password=psql"
	} else {
		dbString = "host=35.232.150.118 port=5432 user=postgres dbname=dev password=psql"
	}

	db, err := sqlx.Connect("postgres", dbString)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Repository{
		db: db,
	}, nil
}

type Repository struct {
	db    *sqlx.DB
	board *BoardRepo
	users *UserRepo
}

func (r *Repository) Board() *BoardRepo {
	if r.board == nil {
		r.board = NewBoardRepo(r.db)
	}

	return r.board
}

func (r *Repository) Users() *UserRepo {
	if r.users == nil {
		r.users = NewUserRepo(r.db)
	}

	return r.users
}
