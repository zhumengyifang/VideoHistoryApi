create table historyInfo
(
    id         int primary key auto_increment,
    openId     char(36) not null unique,
    authorName varchar(100),
    created_at datetime,
    updated_at datetime,
    deleted_at datetime
);