create table if not exists user_profile
(
    created_at         timestamp             default CURRENT_TIMESTAMP,
    updated_at         timestamp             default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    deleted_at         bigint                default 0,
    id                 bigint PRIMARY KEY,
    name               varchar(64)  not null default '',
    gender             tinyint      not null default 0,
    portrait           varchar(255) not null default '',
    openid             varchar(128) not null default '',
    session_key        varchar(128) not null default '',
    current_subject_id bigint       not null default 0,
    study_total_day    int          not null default 0,
    study_last_time    timestamp    not null default '1970-01-01 00:00:01',
    study_num          int          not null default 0,
    unique (openid)
) ENGINE = INNODB;

create table if not exists subject
(
    created_at  timestamp             default CURRENT_TIMESTAMP,
    updated_at  timestamp             default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    deleted_at  bigint                default 0,
    id          bigint PRIMARY KEY,
    name        varchar(128) not null default '',
    description text,
    total       int                   default 0,
    unique (name)
) ENGINE = INNODB;

create table if not exists knowledge
(
    created_at  timestamp             default CURRENT_TIMESTAMP,
    updated_at  timestamp             default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    deleted_at  bigint                default 0,
    id          bigint PRIMARY KEY,
    subject_id  bigint       not null default 0,
    name        varchar(128) not null default '',
    description text,
    other       varchar(512) not null default '{}',
    unique (subject_id, name)
) ENGINE = INNODB;

create table if not exists study_record
(
    created_at   timestamp       default CURRENT_TIMESTAMP,
    updated_at   timestamp       default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    deleted_at   bigint          default 0,
    id           bigint PRIMARY KEY,
    uid          bigint not null default 0,
    subject_id   bigint not null default 0,
    knowledge_id bigint not null default 0,
    unique (uid, subject_id, knowledge_id)
) ENGINE = INNODB;