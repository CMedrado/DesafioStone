CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS accounts (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    cpf VARCHAR NOT NULL UNIQUE,
    secret VARCHAR NOT NULL,
    balance INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL
);
