-- +goose Up

CREATE TYPE auth_provider AS ENUM (
    'NONE',
    'GOOGLE',
    'FACEBOOK',
    'GITHUB'
);

CREATE TYPE user_role AS ENUM (
    'CLIENT',
    'ADMIN'
);

CREATE TABLE IF NOT EXISTS "user" (
    id UUID PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    nickname VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    auth_provider auth_provider NOT NULL,
    user_role user_role NOT NULL,
    created_at TIMESTAMP NOT NULL,
    modified_at TIMESTAMP NOT NULL
);
