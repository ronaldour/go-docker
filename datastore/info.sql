-- Database: books_database

-- DROP DATABASE books_database;

CREATE DATABASE books_database
    WITH OWNER = demo
       ENCODING = 'UTF8'
       CONNECTION LIMIT = -1;

\connect books_database 
-- Table: books

-- DROP TABLE books;

CREATE TABLE books
(
 id serial NOT NULL,
 name character varying NOT NULL,
 author character varying,
 pages integer,
 publication_date date,
 CONSTRAINT pk_books PRIMARY KEY (id )
)
WITH (
 OIDS=FALSE
);