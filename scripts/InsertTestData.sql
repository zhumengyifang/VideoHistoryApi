create
    definer = root@`%` procedure insertTestData(IN userCount int, IN videoCount int)
begin
    declare i int default 0;
    declare j int default 0;

    declare openid char(36);
    declare userId int;
    while userCount > i
        do
            set openid = uuid();
            insert into History.historyInfo(openId, authorName, created_at, updated_at)
            values (openid, openid, now(), now());
            select id into userId from History.historyInfo where openId = openid limit 1;
            while videoCount > j
                do
                    insert into History.videoHistories(userId, videoId, useTime, title, coverUrl, updateCount, isDel,
                                                       submitDateTime, created_at, updated_at)
                    values (userId, uuid(), RAND() * 10000, '1,2,3', 'www.baidu.com', 0, 0, now(), now(), now());
                    set j = j + 1;
                end while;
            set j = 0;
            set i = i + 1;
        end while;
end;

