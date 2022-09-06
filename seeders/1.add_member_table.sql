drop table if exists members;
create table members (
    id  int auto_increment not null,
    name varchar(128) not null,
    role varchar(128) not null,
    primary key (`id`)
);

insert into members
    (name, role)
VALUES
    ('Trương Tấn Sang', 'Manager'),
    ('Nguyễn Thị Tường Vi', 'IT Comtor'),
    ('Đào Duy Thành', 'Developer'),
    ('Võ Trường Thanh', 'Developer'),
    ('Trần Thành An', 'Developer');