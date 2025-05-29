CREATE TABLE tb_account (
  pk_st_id VARCHAR(36) PRIMARY KEY,
  st_first_name VARCHAR(50) NOT NULL,
  st_last_name VARCHAR(50) NOT NULL,
  st_cpf VARCHAR(11) UNIQUE NOT NULL,
  st_email VARCHAR(100) UNIQUE NOT NULL,
  st_password VARCHAR(255) NOT NULL,
  it_number SERIAL UNIQUE,
  dt_created_at TIMESTAMP NOT NULL
);
