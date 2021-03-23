create table if not exists users (
    id serial primary key,
    first_name varchar(100) not null,
    last_name varchar(100) not null,
    full_name varchar(100) not null,
    createdAt timestamp without time zone default (now() at time zone 'utc'),
    updatedAt timestamp without time zone default (now() at time zone 'utc')
);