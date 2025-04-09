package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*PublicAccount, error)
	GetAccountByID(int) (*PublicAccount, error)
	GetAccountByCPF(string) (*Account, error)
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
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {
	query := `
	create table if not exists tb_account (
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

func (s *PostgresStore) CreateAccount(acc *Account) error {
	query := `
	insert into tb_account 
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

func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	_, err := s.db.Query("delete from tb_account where pk_it_id = $1", id)
	return err
}

func (s *PostgresStore) GetAccountByID(id int) (*PublicAccount, error) {
	rows, err := s.db.Query("select pk_it_id, st_first_name, st_last_name, st_cpf, st_email, it_number, db_balance, dt_created_at from tb_account where pk_it_id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoAccount(rows)
	}

	return nil, fmt.Errorf("account %d not found", id)
}

func (s *PostgresStore) GetAccountByCPF(cpf string) (*Account, error) {
	rows, err := s.db.Query("select st_cpf, st_password from tb_account where st_cpf = $1", cpf)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanIntoLogin(rows)
	}

	return nil, fmt.Errorf("account %s not found", cpf)
}

func (s *PostgresStore) GetAccounts() ([]*PublicAccount, error) {
	rows, err := s.db.Query("select * from tb_account;")
	if err != nil {
		return nil, err
	}

	accounts := []*PublicAccount{}
	for rows.Next() {
		account, err := scanIntoAccount(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func scanIntoLogin(rows *sql.Rows) (*Account, error) {
	account := new(Account)
	err := rows.Scan(
		&account.CPF,
		&account.Password,
	)

	return account, err
}

func scanIntoAccount(rows *sql.Rows) (*PublicAccount, error) {
	account := new(PublicAccount)
	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.CPF,
		&account.Email,
		&account.Number,
		&account.Balance,
		&account.CreatedAt)

	return account, err
}
