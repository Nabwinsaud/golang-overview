-- Active: 1714124661583@@127.0.0.1@5432@godb
CREATE TABLE students (
    id int not null PRIMARY KEY, name VARCHAR(200), age int, email VARCHAR(100)
)

ALTER TABLE students DROP COLUMN id

ALTER TABLE students ADD COLUMN id serial PRIMARY KEY