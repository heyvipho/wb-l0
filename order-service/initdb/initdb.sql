create sequence models_id_seq
    as integer;

alter sequence models_id_seq owner to postgres;

create table orders
(
    id                 integer default nextval('models_id_seq'::regclass) not null
        constraint models_pk
            primary key,
    order_uid          varchar,
    track_number       varchar,
    entry              varchar,
    delivery           integer,
    payment            integer,
    items              integer[],
    locale             varchar,
    internal_signature varchar,
    customer_id        varchar,
    delivery_service   varchar,
    shardkey           varchar,
    sm_id              varchar,
    date_created       time,
    oof_shard          varchar
);

alter table orders
    owner to postgres;

alter sequence models_id_seq owned by orders.id;

create unique index models_id_uindex
    on orders (id);

create table deliveries
(
    id      serial
        constraint deliveries_pk
            primary key,
    name    varchar,
    phone   varchar,
    zip     varchar,
    city    varchar,
    address varchar,
    region  varchar,
    email   varchar
);

alter table deliveries
    owner to postgres;

create unique index deliveries_id_uindex
    on deliveries (id);

create table payments
(
    id            serial
        constraint payments_pk
            primary key,
    transaction   varchar,
    request_id    varchar,
    currency      varchar,
    provider      varchar,
    amount        integer,
    payment_dt    integer,
    bank          varchar,
    delivery_cost integer,
    goods_total   integer,
    custom_fee    integer
);

alter table payments
    owner to postgres;

create unique index payments_id_uindex
    on payments (id);

create table items
(
    id           serial
        constraint items_pk
            primary key,
    chrt_id      integer,
    track_number varchar,
    price        integer,
    rid          varchar,
    name         varchar,
    sale         integer,
    size         varchar,
    total_price  integer,
    nm_id        integer,
    brand        varchar,
    status       integer
);

alter table items
    owner to postgres;

create unique index items_id_uindex
    on items (id);
