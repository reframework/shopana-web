-- +goose Up

CREATE TABLE users (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v7(),
	first_name varchar(255) NOT NULL,
	last_name varchar(255) NOT NULL,
	email varchar(255) NOT NULL UNIQUE,
	phone_number varchar(255) UNIQUE,
	password varchar(255) NOT NULL,
	is_verified boolean NOT NULL DEFAULT false,
	is_blocked boolean NOT NULL DEFAULT false,
	timezone varchar(32) NOT NULL,
	language varchar(32) NOT NULL,
	invited_by uuid REFERENCES users(id),
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NOT NULL DEFAULT now()
);

CREATE TABLE app_config (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v7(),
	initialized boolean NOT NULL,
	initialized_at timestamp NOT NULL DEFAULT now(),
	app_version float NOT NULL DEFAULT 1.0
);

-- +goose Down
DROP TABLE users;
DROP TABLE invited_users;
DROP TABLE app_config;
