create table orders
(
    id   serial
        constraint orders_pk
            primary key,
    data text
);

alter table orders
    owner to postgres;

create unique index orders_id_uindex
    on orders (id);
