create table category (
    id          bigserial primary key,
    name        text unique,
    is_visible bool,
    created_at timestamp,
    updated_at timestamp
);

create table pet (
    id          bigserial primary key,
    name        text,
    category_id bigint,
    created_at timestamp,
    updated_at timestamp,

    constraint fk_category_id foreign key(category_id) references  category(id)
);
