create database exemplo_db;

\connect exemplo_db

DROP TABLE IF EXISTS postages;

create table if not exists postages (
    id serial not null
        constraint pk_postages primary key,
    title varchar(64) not null,
    message varchar(64) not null
);

INSERT INTO
    postages (title, message)
VALUES ('LIVRO: Admirável mundo novo', 'messagem'),
       ('LIVRO: A hora da estrela', 'menssagem');

DROP TABLE IF EXISTS users;

create table if not exists users (
    id serial not null
        constraint pk_users primary key,
    name varchar(64) not null,
    email varchar(64) not null,
    password varchar(64) not null
);

INSERT INTO
    users (name, email, password)
VALUES ('teste', 'teste', 'teste');

INSERT INTO
    users (name, email, password)
     VALUES   ('teste', 'teste', 'teste');

select id,
       title,
       message
    from postages
    where title like '%ens%';

-- select id,
--        name,
--        count,
--        price
--     from postages
--     where upper(name) like upper('%ens%')
--         and price > 10;

select Id, title, message from postages;

select Id, name, email, password from users;

-- update postages set count = 13 where id = 7;