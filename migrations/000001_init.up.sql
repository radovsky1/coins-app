CREATE TABLE users
(
    id         bigserial PRIMARY KEY,
    name       varchar     NOT NULL,
    username   varchar     NOT NULL,
    password   varchar     NOT NULL,
    created_at timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE accounts
(
    id       bigserial PRIMARY KEY,
    user_id  bigint,
    balance  bigint,
    currency varchar NOT NULL
);

CREATE TABLE transfers
(
    id              bigserial PRIMARY KEY,
    from_account_id bigint,
    to_account_id   bigint,
    amount          bigint   NOT NULL,
    currency        varchar NOT NULL
);

CREATE INDEX ON accounts (user_id);

COMMENT ON COLUMN transfers.amount IS 'must be positive';

ALTER TABLE accounts
    ADD FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE transfers
    ADD FOREIGN KEY (from_account_id) REFERENCES accounts (id);

ALTER TABLE transfers
    ADD FOREIGN KEY (to_account_id) REFERENCES accounts (id);
