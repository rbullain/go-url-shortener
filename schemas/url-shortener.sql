DROP TABLE IF EXISTS urlshortener;

CREATE TABLE urlshortener
(
    hash            varchar(6) collate utf8mb3_bin not null primary key,
    original_url    varchar(10240)                 not null,
    expiration_date timestamp                      not null
);

# CREATE EVENT IF NOT EXISTS `urlshortener_delete_expired`
#     ON SCHEDULE
#         EVERY 1 HOUR
#     DO
#     BEGIN
#         DELETE FROM urlshortener WHERE expiration_date < NOW();
#     END;