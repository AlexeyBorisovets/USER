package repository

import (
	"github.com/AlexeyBorisovets/USER/internal/model"

	"context"

	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

// PRepository p
type PRepository struct {
	Pool *pgxpool.Pool
}

// CreateUser add user to db
func (p *PRepository) CreateUser(ctx context.Context, user *model.User) (string, error) {
	newID := uuid.New().String()
	_, err := p.Pool.Exec(ctx, "insert into users(id,name,password,balance,usertype) values($1,$2,$3,$4,$5)",
		newID, &user.Name, &user.Password, &user.Balance, &user.UserType)
	if err != nil {
		log.Errorf("database error with create user: %v", err)
		return "", err
	}
	return newID, nil
}

// GetUserByID select user by id
func (p *PRepository) GetUserByID(ctx context.Context, userId string) (*model.User, error) {
	u := model.User{}
	err := p.Pool.QueryRow(ctx, "select id,name,password,balance,usertype from users where id=$1", userId).Scan(
		&u.ID, &u.Name, &u.Password, &u.Balance, &u.UserType)
	if err != nil {
		if err == pgx.ErrNoRows {
			return &model.User{}, fmt.Errorf("user with this id doesnt exist: %v", err)
		}
		log.Errorf("database error, select by id: %v", err)
		return &model.User{}, err
	}
	return &u, nil
}

// GetAllUsers select all users from db
func (p *PRepository) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	rows, err := p.Pool.Query(ctx, "select id,name,password,balance,usertype from users")
	if err != nil {
		log.Errorf("database error with select all users, %v", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user := model.User{}
		err = rows.Scan(&user.ID, &user.Name, &user.Password, &user.Balance, &user.UserType)
		if err != nil {
			log.Errorf("database error with select all users, %v", err)
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

// DeleteUser delete user by id
func (p *PRepository) DeleteUser(ctx context.Context, id string) error {
	a, err := p.Pool.Exec(ctx, "delete from users where id=$1", id)
	if a.RowsAffected() == 0 {
		return fmt.Errorf("user with this id doesnt exist")
	}
	if err != nil {
		if err == pgx.ErrNoRows {
			return fmt.Errorf("user with this id doesnt exist: %v", err)
		}
		log.Errorf("error with delete user %v", err)
		return err
	}
	return nil
}

// UpdateUser update parameters for user
func (p *PRepository) UpdateUser(ctx context.Context, id string, user *model.User) error {
	a, err := p.Pool.Exec(ctx, "update users set name=$1,password=$2,balance=$3,usertype=$4 where id=$5",
		&user.Name, &user.Password, &user.Balance, &user.UserType, id)
	if a.RowsAffected() == 0 {
		return fmt.Errorf("user with this id doesnt exist")
	}
	if err != nil {
		log.Errorf("error with update user %v", err)
		return err
	}
	return nil
}

// UpdateAuth logout, delete refresh token
func (p *PRepository) UpdateAuth(ctx context.Context, id string, refreshToken string) error {
	a, err := p.Pool.Exec(ctx, "update users set refreshToken=$1 where id=$2", refreshToken, id)
	if a.RowsAffected() == 0 {
		return fmt.Errorf("user with this id doesnt exist")
	}
	if err != nil {
		log.Errorf("error with update user %v", err)
		return err
	}
	return nil
}

// SelectByIDAuth get id and refresh token by id
func (p *PRepository) SelectByIDAuth(ctx context.Context, id string) (model.User, error) {
	user := model.User{}
	err := p.Pool.QueryRow(ctx, "select id,refreshToken from users where id=$1", id).Scan(&user.ID, &user.RefreshToken)

	if err != nil /*err==no-records*/ {
		if err == pgx.ErrNoRows {
			return model.User{}, fmt.Errorf("user with this id doesnt exist: %v", err)
		}
		log.Errorf("database error, select by id: %v", err)
		return model.User{}, err /*p, fmt.errorf("user with this id doesnt exist")*/
	}
	return user, nil
}

//----

// GetUserByUserType select user by usertype
func (p *PRepository) GetUserByUserType(ctx context.Context, usertype string) ([]*model.User, error) {
	var users []*model.User
	rows, err := p.Pool.Query(ctx, "select id,name,password,balance,usertype from users where usertype=$1", usertype)
	if err != nil {
		log.Errorf("database error with select users by usertype, %v", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user := model.User{}
		err = rows.Scan(&user.ID, &user.Name, &user.Password, &user.Balance, &user.UserType)
		if err != nil {
			log.Errorf("database error with select users by usertype, %v", err)
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

// GetUserByID select balance by id
func (p *PRepository) GetBalanceByID(ctx context.Context, userId string) (uint, error) {
	user := model.User{}
	err := p.Pool.QueryRow(ctx, "select balance from users where id=$1", userId).Scan(&user.Balance)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, fmt.Errorf("user with this id doesnt exist(func: balance by id): %v", err)
		}
		log.Errorf("database error, balance by id: %v", err)
		return 0, err
	}
	balance := user.Balance
	return balance, nil
}

// UpdateBalance update balance by id
func (p *PRepository) UpdateBalance(ctx context.Context, userid string, balance uint) error {
	a, err := p.Pool.Exec(ctx, "update users set balance=$1 where id=$2", balance, userid)
	if a.RowsAffected() == 0 {
		return fmt.Errorf("user with this id doesnt exist")
	}
	if err != nil {
		log.Errorf("error with update user balance %v", err)
		return err
	}
	return nil
}
