create
    definer = root@`%` procedure insertTestData(IN userCount int, IN videoCount int)
begin
    declare i int default 0;
    declare j int default 0;

    declare newOpenid char(36);
    declare queryUserId int;
    while userCount > i
        do
            set newOpenid = uuid();
            insert into History.historyInfo(openId, authorName, created_at, updated_at)
            values (newOpenid, newOpenid, now(), now());
            select id into queryUserId from History.historyInfo where openId = newOpenid limit 1;
            while videoCount > j
                do
                    insert into History.videoHistories(userId, videoId, useTime, title, coverUrl, updateCount, isDel,
                                                       submitDateTime, created_at, updated_at)
                    values (queryUserId, uuid(), RAND() * 10000, '1,2,3', 'www.baidu.com', 0, 0, now(), now(), now());
                    set j = j + 1;
                end while;
            set i = i + 1;
            set j = 0;
        end while;
end;

