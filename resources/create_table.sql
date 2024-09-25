create table dim_joke_di(
    joke_id varchar(255),
    content varchar(255)
);

create table dwd_joke_act_rt(
    joke_id varchar(255),
    act_type varchar(255),
    act_time timestamp
);