CREATE TABLE tb_movement (
  pk_st_id VARCHAR(36) PRIMARY KEY,
  it_value INTEGER NOT NULL,
  it_recurrence SMALLINT NOT NULL,
  st_type VARCHAR(7) NOT NULL,
  st_description VARCHAR(255) NOT NULL,
  dt_created_at TIMESTAMP NOT NULL,
  st_account_id VARCHAR(36),
  CONSTRAINT fk_account_movement FOREIGN KEY (st_account_id) REFERENCES tb_account (pk_st_id)
);
