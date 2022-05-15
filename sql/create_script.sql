drop table if exists users cascade;

create table users (id serial primary key, email varchar);

drop table if exists products cascade;

create table products (
    id serial primary key,
    name varchar,
    price numeric(10, 2),
    quantity int
);

drop table if exists attributes cascade;

create table attributes (id serial primary key, name varchar);

drop table if exists attribute_values cascade;

create table attribute_values (
    id serial primary key,
    attribute_id int references attributes(id) on delete cascade,
    value varchar
);

drop table if exists product_attributes cascade;

create table product_attributes (
    id serial primary key,
    product_id int references products(id) on delete cascade,
    attribute_id int references attributes(id) on delete cascade,
    attribute_value_id int references attribute_values(id) on delete cascade,
    unique (product_id, attribute_id)
);

drop table if exists basket_items cascade;

create table basket_items (
    id serial primary key,
    user_id int references users(id) on delete cascade,
    product_id int references products(id) on delete cascade,
    quantity int,
    unique(user_id, product_id)
);

drop table if exists orders cascade;

drop type if exists order_status cascade;

create type order_status as enum ('pending', 'paid', 'expedited', 'cancelled');

create table orders (
    id serial primary key,
    user_id int references users(id) on delete cascade,
    address varchar,
    total numeric(10, 2),
    status order_status
);

drop table if exists order_items cascade;

create table order_items (
    id serial primary key,
    order_id int references orders(id) on delete cascade,
    product_id int references products(id),
    quantity int,
    unique(order_id, product_id)
);

drop table if exists search_log cascade;

create table search_log (
    id serial primary key, 
    query varchar,
    unique(query)
);

drop table if exists search_log_timestamps cascade;

create table search_log_timestamps (
    search_log_id int references search_log(id),
    searched_at timestamp
);

drop index if exists searched_at_idx;
create index searched_at_idx on search_log_timestamps (searched_at);

drop table if exists payments cascade;

create table payments (
    id serial primary key,
    order_id int references orders(id) on delete cascade,
    amount numeric(10, 2)
);