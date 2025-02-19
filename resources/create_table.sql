create table dim_user_di(
    device_id     varchar(255),
    user_id       varchar(255),
    brand         varchar(255),
    mobile_model  varchar(255),
    launch_time   timestamp,
    register_time timestamp
);
create index device_idx on dim_user_di(device_id);

create table dim_joke_di(
    joke_id     varchar(255),
    content     text,
    category    varchar(255),
    type        varchar(255),
    setup       text,
    delivery    text,
    language    varchar(255),
    source      varchar(255),
    create_date date
);
create index joke_idx on dim_joke_di(joke_id);

create table dwd_joke_act_rt(
    device_id varchar(255),
    joke_id   varchar(255),
    act_type  varchar(255),
    act_time  timestamp
);
create index device_idx on dwd_joke_act_rt(device_id);
create index joke_idx on dwd_joke_act_rt(joke_id);