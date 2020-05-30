package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/skvoch/galley/internal/galley/model"
)

func NewBoardRepo(db *sqlx.DB) *BoardRepo {
	return &BoardRepo{
		db: db,
	}
}

type BoardRepo struct {
	db *sqlx.DB
}

func (b *BoardRepo) Change(task *model.Task) error {
	_, err := b.db.Exec("UPDATE public.board SET title=$1, description=$2, status=$3, urgency=$4 WHERE id = $5;",
		task.Title, task.Description, task.Status, task.Urgency, task.ID)

	if err != nil {
		return err
	}

	return err
}

func (b *BoardRepo) Create(task *model.Task) error {
	_, err := b.db.Exec("INSERT INTO public.board(title, description, status, urgency) VALUES ($1, $2, $3, $4);",
		task.Title, task.Description, task.Status, task.Urgency)

	if err != nil {
		return err
	}

	return err
}

func (b *BoardRepo) ReadAll() ([]*model.Task, error) {

	rows, err := b.db.Query("SELECT id, title, description, status, urgency FROM public.board;")

	if err != nil {
		return nil, err
	}

	out := make([]*model.Task, 0)

	for rows.Next() {
		task := &model.Task{}

		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.Urgency)

		if err != nil {
			return nil, err
		}
		out = append(out, task)
	}

	return out, nil
}
