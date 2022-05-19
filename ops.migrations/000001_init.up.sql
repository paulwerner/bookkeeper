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
    balance_currency VARCHAR NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT uc_accounts_user_id_name UNIQUE(user_id, name)
);

CREATE INDEX idx_accounts_id_user_id ON accounts(id, user_id);

CREATE INDEX idx_accounts_user_id ON accounts(user_id);

CREATE TABLE IF NOT EXISTS transactions(
    id VARCHAR PRIMARY KEY NOT NULL,
    account_id VARCHAR NOT NULL,
    description VARCHAR,
    amount BIGINT NOT NULL,
    currency VARCHAR NOT NULL,
    FOREIGN KEY (account_id) REFERENCES accounts(id)
);

CREATE INDEX idx_transactions_id_account_id ON transactions(id, account_id);