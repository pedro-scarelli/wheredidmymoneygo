CREATE TABLE tb_account (
  pk_it_id SERIAL PRIMARY KEY,
  st_first_name VARCHAR(50) NOT NULL,
  st_last_name VARCHAR(50) NOT NULL,
  st_cpf VARCHAR(11) UNIQUE NOT NULL,
  st_email VARCHAR(100) UNIQUE NOT NULL,
  st_password VARCHAR(255) NOT NULL,
  it_number SERIAL UNIQUE,
  db_balance DECIMAL(10,2) NOT NULL,
  dt_created_at TIMESTAMP NOT NULL
);
