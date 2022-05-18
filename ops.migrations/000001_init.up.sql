CREATE TABLE IF NOT EXISTS users(
    id VARCHAR PRIMARY KEY NOT NULL,
    name VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS accounts(
    id VARCHAR PRIMARY KEY NOT NULL,
    user_id VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    description VARCHAR,
    type VARCHAR NOT NULL,
    balance_value BIGINT NOT NULL,
    balance_currency VARCHAR NOT NULL
);

CREATE INDEX idx_accounts_id_user_id ON accounts(id, user_id);

CREATE INDEX idx_accounts_user_id_name ON accounts(id, user_id);

CREATE TABLE IF NOT EXISTS transactions(
    id VARCHAR PRIMARY KEY NOT NULL,
    account_id VARCHAR NOT NULL,
    amount BIGINT NOT NULL,
    currency VARCHAR NOT NULL
);

CREATE INDEX idx_transactions_id_account_id ON transactions(id, account_id);