create table user
(
    id   varchar(255) not null comment 'ユニークなID'
        primary key,
    name varchar(255) not null comment 'ユーザ名'
);
