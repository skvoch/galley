package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/skvoch/galley/internal/galley/model"
	"strconv"
)

func NewClicksRepo(db *sqlx.DB) *ClicksRepo {
	return &ClicksRepo{
		db: db,
	}
}

type ClicksRepo struct {
	db *sqlx.DB
}

func (c *ClicksRepo) ReadLast(count int) ([]*model.ClickStats, error) {
	countStr := strconv.Itoa(count)

	rows, err := c.db.Query("SELECT hash, count FROM public.clicks ORDER BY id DESC LIMIT " + countStr + ";")

	out := make([]*model.ClickStats, 0)
	for rows.Next() {
		stats := &model.ClickStats{}

		err = rows.Scan(&stats.Hash, &stats.Count)

		if err != nil {
			return nil, err
		}
		out = append(out, stats)
	}

	return out, nil
}

func (c *ClicksRepo) Add(stats *model.ClickStats) error {
	_, err := c.db.Exec("INSERT INTO public.clicks(hash, count) VALUES ($1, $2) ;",
		stats.Hash, stats.Count)

	if err != nil {
		return err
	}

	return err
}
