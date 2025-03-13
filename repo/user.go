package repo

import (
	"context"
	"database/sql"
	"errors"
	"warehouse-management-system/entity"
	"warehouse-management-system/sentinel"
)

type UserRepo interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	InsertUser(ctx context.Context, req *entity.InsertUser) (int, error)
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	res := entity.User{}

	query := `
	SELECT
		u.id,
		u.email,
		u.password,
		u.role,
		u.created_at,
		u.updated_at,
		u.deleted_at
	FROM
		users u
	WHERE
		u.email = $1;
	`

	tx := extractTx(ctx)
	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, email).Scan(&res.ID, &res.Email, &res.Password, &res.Role, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt)
	} else {
		err = r.db.QueryRowContext(ctx, query, email).Scan(&res.ID, &res.Email, &res.Password, &res.Role, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt)
	}
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = sentinel.ErrNotFound
		}
		return nil, err
	}

	return &res, nil
}

func (r *userRepo) InsertUser(ctx context.Context, req *entity.InsertUser) (int, error) {
	query := `
	INSERT INTO
		users
	(
		email,
		password,
		role,
		created_at,
		updated_at
	)
	VALUES
	(
		$1,
		$2,
		$3,
		NOW(),
		NOW()
	)
	RETURNING
		id`

	var ID int
	tx := extractTx(ctx)
	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, req.Email, req.Password, req.Role).Scan(&ID)
	} else {
		err = r.db.QueryRowContext(ctx, query, req.Email, req.Password, req.Role).Scan(&ID)
	}
	if err != nil {
		return -1, err
	}

	return ID, nil
}
