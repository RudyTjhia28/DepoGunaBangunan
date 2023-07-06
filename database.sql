-- development.login_history definition

-- Drop table

-- DROP TABLE development.login_history;

CREATE TABLE development.login_history (
	id serial4 NOT NULL,
	username varchar(255) NOT NULL,
	login_time timestamptz NOT NULL,
	login_success bool NOT NULL,
	CONSTRAINT login_history_pkey PRIMARY KEY (id)
);

-- development.products definition

-- Drop table

-- DROP TABLE development.products;

CREATE TABLE development.products (
	id serial4 NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	"name" varchar(255) NULL,
	description varchar(255) NULL,
	price int4 NULL,
	quantity int4 NULL,
	CONSTRAINT products_pkey PRIMARY KEY (id)
);

-- development.users definition

-- Drop table

-- DROP TABLE development.users;

CREATE TABLE development.users (
	id serial4 NOT NULL,
	created_at timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	username varchar(255) NULL,
	"password" varchar NULL,
	fullname varchar NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id),
	CONSTRAINT users_username_key UNIQUE (username)
);

INSERT INTO development.products (id, created_at, updated_at, deleted_at, "name", description, price, quantity) VALUES(1, NULL, NULL, NULL, 'Product 1', 'Description 1', 10, 50);
INSERT INTO development.products (id, created_at, updated_at, deleted_at, "name", description, price, quantity) VALUES(2, NULL, NULL, NULL, 'Product 2', 'Description 2', 20, 100);
INSERT INTO development.products (id, created_at, updated_at, deleted_at, "name", description, price, quantity) VALUES(3, NULL, NULL, NULL, 'Product 3', 'Description 3', 15, 75);
INSERT INTO development.products (id, created_at, updated_at, deleted_at, "name", description, price, quantity) VALUES(4, NULL, NULL, NULL, 'Product 4', 'Description 4', 30, 25);
INSERT INTO development.products (id, created_at, updated_at, deleted_at, "name", description, price, quantity) VALUES(5, NULL, NULL, NULL, 'Product 5', 'Description 5', 25, 60);
INSERT INTO development.products (id, created_at, updated_at, deleted_at, "name", description, price, quantity) VALUES(6, NULL, NULL, NULL, 'Product 6', 'Description 6', 5, 200);
INSERT INTO development.products (id, created_at, updated_at, deleted_at, "name", description, price, quantity) VALUES(7, NULL, NULL, NULL, 'Product 7', 'Description 7', 12, 80);
INSERT INTO development.products (id, created_at, updated_at, deleted_at, "name", description, price, quantity) VALUES(8, NULL, NULL, NULL, 'Product 8', 'Description 8', 8, 150);
INSERT INTO development.products (id, created_at, updated_at, deleted_at, "name", description, price, quantity) VALUES(9, NULL, NULL, NULL, 'Product 9', 'Description 9', 18, 90);
INSERT INTO development.products (id, created_at, updated_at, deleted_at, "name", description, price, quantity) VALUES(10, NULL, NULL, NULL, 'Product 10', 'Description 10', 22, 40);
INSERT INTO development.products (id, created_at, updated_at, deleted_at, "name", description, price, quantity) VALUES(11, NULL, NULL, NULL, 'Product 91', 'Description 91', 14, 120);
INSERT INTO development.products (id, created_at, updated_at, deleted_at, "name", description, price, quantity) VALUES(12, NULL, NULL, NULL, 'Product 92', 'Description 92', 9, 70);
INSERT INTO development.products (id, created_at, updated_at, deleted_at, "name", description, price, quantity) VALUES(13, NULL, NULL, NULL, 'Product 93', 'Description 93', 11, 180);
INSERT INTO development.products (id, created_at, updated_at, deleted_at, "name", description, price, quantity) VALUES(14, NULL, NULL, NULL, 'Product 94', 'Description 94', 16, 35);
INSERT INTO development.products (id, created_at, updated_at, deleted_at, "name", description, price, quantity) VALUES(15, NULL, NULL, NULL, 'Product 95', 'Description 95', 27, 55);
INSERT INTO development.products (id, created_at, updated_at, deleted_at, "name", description, price, quantity) VALUES(16, NULL, NULL, NULL, 'Product 96', 'Description 96', 13, 100);
INSERT INTO development.products (id, created_at, updated_at, deleted_at, "name", description, price, quantity) VALUES(17, NULL, NULL, NULL, 'Product 97', 'Description 97', 7, 250);
INSERT INTO development.products (id, created_at, updated_at, deleted_at, "name", description, price, quantity) VALUES(18, NULL, NULL, NULL, 'Product 98', 'Description 98', 19, 85);
INSERT INTO development.products (id, created_at, updated_at, deleted_at, "name", description, price, quantity) VALUES(19, NULL, NULL, NULL, 'Product 99', 'Description 99', 21, 120);
INSERT INTO development.products (id, created_at, updated_at, deleted_at, "name", description, price, quantity) VALUES(20, NULL, NULL, NULL, 'Product 100', 'Description 100', 29, 30);

INSERT INTO development.users (id, created_at, updated_at, deleted_at, username, "password", fullname) VALUES(1, NULL, NULL, NULL, 'john_doe', '7c6a180b36896a0a8c02787eeafb0e4c', 'john_doe');
INSERT INTO development.users (id, created_at, updated_at, deleted_at, username, "password", fullname) VALUES(2, NULL, NULL, NULL, 'alice_smith', '7c6a180b36896a0a8c02787eeafb0e4c', 'alice_smith');
INSERT INTO development.users (id, created_at, updated_at, deleted_at, username, "password", fullname) VALUES(3, NULL, NULL, NULL, 'robert_johnson', '7c6a180b36896a0a8c02787eeafb0e4c', 'robert_johnson');
INSERT INTO development.users (id, created_at, updated_at, deleted_at, username, "password", fullname) VALUES(4, NULL, NULL, NULL, 'emily_wilson', '7c6a180b36896a0a8c02787eeafb0e4c', 'emily_wilson');
INSERT INTO development.users (id, created_at, updated_at, deleted_at, username, "password", fullname) VALUES(5, NULL, NULL, NULL, 'michael_jones', '7c6a180b36896a0a8c02787eeafb0e4c', 'michael_jones');
INSERT INTO development.users (id, created_at, updated_at, deleted_at, username, "password", fullname) VALUES(6, NULL, NULL, NULL, 'sarah_brown', '7c6a180b36896a0a8c02787eeafb0e4c', 'sarah_brown');
INSERT INTO development.users (id, created_at, updated_at, deleted_at, username, "password", fullname) VALUES(7, NULL, NULL, NULL, 'david_clark', '7c6a180b36896a0a8c02787eeafb0e4c', 'david_clark');
INSERT INTO development.users (id, created_at, updated_at, deleted_at, username, "password", fullname) VALUES(8, NULL, NULL, NULL, 'jennifer_wright', '7c6a180b36896a0a8c02787eeafb0e4c', 'jennifer_wright');
INSERT INTO development.users (id, created_at, updated_at, deleted_at, username, "password", fullname) VALUES(9, NULL, NULL, NULL, 'matthew_davis', '7c6a180b36896a0a8c02787eeafb0e4c', 'matthew_davis');
INSERT INTO development.users (id, created_at, updated_at, deleted_at, username, "password", fullname) VALUES(10, NULL, NULL, NULL, 'olivia_taylor', '7c6a180b36896a0a8c02787eeafb0e4c', 'olivia_taylor');
INSERT INTO development.users (id, created_at, updated_at, deleted_at, username, "password", fullname) VALUES(11, NULL, NULL, NULL, 'samuel_thompson', '7c6a180b36896a0a8c02787eeafb0e4c', 'samuel_thompson');
INSERT INTO development.users (id, created_at, updated_at, deleted_at, username, "password", fullname) VALUES(12, NULL, NULL, NULL, 'lily_jackson', '7c6a180b36896a0a8c02787eeafb0e4c', 'lily_jackson');
INSERT INTO development.users (id, created_at, updated_at, deleted_at, username, "password", fullname) VALUES(13, NULL, NULL, NULL, 'benjamin_harris', '7c6a180b36896a0a8c02787eeafb0e4c', 'benjamin_harris');
INSERT INTO development.users (id, created_at, updated_at, deleted_at, username, "password", fullname) VALUES(14, NULL, NULL, NULL, 'amelia_martin', '7c6a180b36896a0a8c02787eeafb0e4c', 'amelia_martin');
INSERT INTO development.users (id, created_at, updated_at, deleted_at, username, "password", fullname) VALUES(15, NULL, NULL, NULL, 'daniel_white', '7c6a180b36896a0a8c02787eeafb0e4c', 'daniel_white');
