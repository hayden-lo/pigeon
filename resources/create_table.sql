create table dim_user_di(
    device_id varchar(255),
    user_id varchar(255),
    brand varchar(255),
    mobile_model varchar(255),
    launch_time timestamp,
    register_time timestamp
);

create table dim_joke_di(
    joke_id varchar(255),
    content varchar(255),
    create_date date
);

create table dwd_joke_act_rt(
    device_id varchar(255),
    joke_id varchar(255),
    act_type varchar(255),
    act_time timestamp
);