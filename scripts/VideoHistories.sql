create table videoHistories
(
    id          int primary key auto_increment,
    userId      int      not null,
    videoId     char(36) not null unique,
    useTime     int,
    title       varchar(255),
    coverUrl    varchar(255),
    updateCount int,
    isDel       bool,
    submitDate  time,
    createdAt  time,
    updatedAt  time,
    foreign key (userId) references users (id)
);