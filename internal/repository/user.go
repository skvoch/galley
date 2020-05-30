package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/skvoch/galley/internal/model"
)

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

type UserRepo struct {
	db *sqlx.DB
}

func (u *UserRepo) Exist(user *model.User) (bool, error) {

	rows, err := u.db.Query("SELECT hash FROM public.users WHERE hash = $1;",
		user.Hash,
	)

	if err != nil {
		return false, err
	}

	return rows.Next(), nil
}

func (u *UserRepo) ReadAll() ([]*model.User, error) {

	rows, err := u.db.Query("SELECT hash, first_name, second_name FROM public.users;")

	if err != nil {
		return nil, err
	}

	out := make([]*model.User, 0)

	for rows.Next() {
		user := &model.User{}

		err = rows.Scan(&user.Hash, &user.FirstName, &user.SecondName)

		if err != nil {
			return nil, err
		}
		out = append(out, user)
	}

	return out, nil
}

func (u *UserRepo) Create(user *model.User) error {

	_, err := u.db.Exec("INSERT INTO public.users(hash, first_name, second_name) VALUES ($1, $2, $3);",
		user.Hash, user.FirstName, user.SecondName)

	if err != nil {
		return err
	}

	return err
}
