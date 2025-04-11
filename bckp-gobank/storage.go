package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateUser(*User) error
	DeleteUser(int) error
	UpdateUser(*User) error
	GetUsers() ([]*User, error)
	GetUserByID(int) (*User, error)
	GetUserByCPF(string) (*User, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=scarelli dbname=db_gobank password=neymardemoicano2011 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.createUserTable()
}

func (s *PostgresStore) createUserTable() error {
	query := `
	create table if not exists tb_user (
		pk_it_id serial primary key,
		st_first_name varchar(50) not null,
		st_last_name varchar(50) not null,
		st_cpf varchar(11) unique not null,
		st_email varchar(100) unique not null,
		st_password varchar(255) not null,
		it_number serial unique,
		db_balance decimal(10,2) not null,
		dt_created_at timestamp not null
	);`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateUser(acc *User) error {
	query := `
	insert into tb_user 
		(st_first_name, st_last_name, st_cpf, st_email, st_password, it_number, db_balance, dt_created_at)
	values 
		($1, $2, $3, $4, $5, $6, $7, $8);`

	resp, err := s.db.Query(
		query,
		acc.FirstName,
		acc.LastName,
		acc.CPF,
		acc.Email,
		acc.Password,
		acc.Number,
		acc.Balance,
		acc.CreatedAt,
	)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)
	return nil
}

func (s *PostgresStore) UpdateUser(*User) error {
	return nil
}

func (s *PostgresStore) DeleteUser(id int) error {
	_, err := s.db.Query("delete from tb_user where pk_it_id = $1", id)
	return err
}

func (s *PostgresStore) GetUserByID(id int) (*User, error) {
	rows, err := s.db.Query("select pk_it_id, st_first_name, st_last_name, st_cpf, st_email, it_number, db_balance, dt_created_at from tb_user where pk_it_id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoUser(rows)
	}

	return nil, fmt.Errorf("user %d not found", id)
}

func (s *PostgresStore) GetUserByCPF(cpf string) (*User, error) {
	rows, err := s.db.Query("select st_cpf, st_password from tb_user where st_cpf = $1", cpf)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanIntoLogin(rows)
	}

	return nil, fmt.Errorf("user %s not found", cpf)
}

func (s *PostgresStore) GetUsers() ([]*User, error) {
	rows, err := s.db.Query("select * from tb_user;")
	if err != nil {
		return nil, err
	}

	users := []*PublicUser{}
	for rows.Next() {
		user, err := scanIntoUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func scanIntoLogin(rows *sql.Rows) (*User, error) {
	user := new(User)
	err := rows.Scan(
		&user.CPF,
		&user.Password,
	)

	return user, err
}

func scanIntoUser(rows *sql.Rows) (*PublicUser, error) {
	user := new(PublicUser)
	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.CPF,
		&user.Email,
		&user.Number,
		&user.Balance,
		&user.CreatedAt)

	return user, err
}
