#CREATE DATABASE banking;
USE banking;

DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS customers;
DROP TABLE IF EXISTS users;

CREATE TABLE customers (
	customer_id int NOT NULL AUTO_INCREMENT,
    name varchar(100) NOT NULL,
    date_of_birth date NOT NULL,
    city varchar(100) NOT NULL,
    zipcode	varchar(10) NOT NULL,
    status tinyint NOT NULL DEFAULT 1,
    PRIMARY KEY (customer_id)
) ENGINE=InnoDB AUTO_INCREMENT=2000 DEFAULT CHARSET=latin1;

CREATE TABLE accounts (
	account_id int NOT NULL AUTO_INCREMENT,
    customer_id int NOT NULL references customers(customer_id),
    opening_date date NOT NULL,
    account_type varchar(20) NOT NULL,
    amount float NOT NULL,
    status tinyint NOT NULL DEFAULT 1,
    PRIMARY KEY (account_id)
) ENGINE=InnoDB AUTO_INCREMENT=2000 DEFAULT CHARSET=latin1;

CREATE TABLE transactions (
	transaction_id int NOT NULL AUTO_INCREMENT,
    account_id int NOT NULL references accounts(account_id),
    amount float NOT NULL,
    transaction_type varchar(20) NOT NULL,
    transaction_date date NOT NULL,
    PRIMARY KEY (transaction_id)
) ENGINE=InnoDB AUTO_INCREMENT=2000 DEFAULT CHARSET=latin1;

CREATE TABLE users (
	username varchar(20) primary key,
    password varchar(256) not null,
    role varchar(10) not null,
    customer_id int,
    created_on datetime default current_timestamp
);

INSERT INTO customers(name, date_of_birth, city, zipcode, status) VALUES
	('Pawel', '2000-01-01', 'Kielce', '25-314', 1),
    ('Jan', '2000-01-01', 'Kielce', '25-314', 1),
    ('Artur', '2000-01-01', 'Kielce', '25-314', 1),
    ('Bartosz', '2000-01-01', 'Kielce', '25-314', 0),
    ('Izabela', '2000-01-01', 'Kielce', '25-314', 1),
    ('Karolina', '2000-01-01', 'Kielce', '25-314', 1),
    ('Mateusz', '2000-01-01', 'Kielce', '25-314', 1),
    ('Wojciech', '2000-01-01', 'Kielce', '25-314', 0),
    ('Michal', '2000-01-01', 'Kielce', '25-314', 0),
    ('Krzysztof', '2000-01-01', 'Kielce', '25-314', 0);
    
INSERT INTO accounts(customer_id, opening_date, account_type, amount, status) VALUES
	(2000, '2020-01-01', 'saving', 6823.23, 1),
    (2001, '2020-01-01', 'checking', 3333.23, 1),
    (2004, '2020-01-01', 'saving', 16823.23, 1);
    
INSERT INTO users(username, password, role, customer_id) VALUES
	('siwonpawel', 'abc123', 'admin', null),
    ('2000', '123abc', 'user', 2000),
    ('2001', '123abc', 'user', 2001),
    ('2002', '123abc', 'user', 2001);