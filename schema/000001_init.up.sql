CREATE TABLE users
(
    id                  serial                                      not null unique,
    balance             int,
    services            varchar(255)
);

CREATE TABLE users_transactions
(
    id                  serial                                      not null unique,
    user_id             int references users (id)                   not null,
    netto               int                                         not null,
    cashflow            boolean,
    source_or_purpose   varchar(255)                                not null,
    transTime           timestamp                                   not null default NOW()
);
