create table videoHistories
(
    id             int primary key auto_increment,
    userId         int      not null,
    videoId        char(36) not null,
    useTime        int,
    title          varchar(255),
    coverUrl       varchar(255),
    updateCount    int,
    isDel          bool,
    submitDateTime datetime,
    created_at     datetime,
    updated_at     datetime,
    deleted_at     datetime,
    foreign key (userId) references historyInfo (id)
);