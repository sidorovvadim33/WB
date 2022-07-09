-- name: create-deliveries
create table delivery
(
    delivery_name varchar unique,
    phone         varchar(20),
    zip           varchar(30),
    city          varchar(30),
    address       varchar(30),
    region        varchar(30),
    email         varchar(30)
);

-- name: create-payments
create table payment
(
    payment_transaction varchar unique,
    request_id          varchar(30),
    currency            varchar(30),
    payment_provider    varchar(30),
    amount              int,
    payment_dt          int,
    bank                varchar(30),
    delivery_cost       int,
    goods_total         int,
    custom_fee          int
);

-- name: create-orders
create table orders
(
    order_uid          varchar primary key,
    track_number       varchar(25) NOT NULL,
    entry              varchar(25),
    delivery           varchar,
    payment            varchar,
    locale             varchar(3),
    internal_signature text,
    customer_id        varchar,
    delivery_service   varchar(30),
    shardkey           varchar(10),
    sm_id              int,
    date_created       timestamptz,
    oof_shard          varchar(5),
    constraint delivery_fk foreign key (delivery) references delivery (delivery_name),
    constraint payment_fk foreign key (payment) references payment (payment_transaction)
);

-- name: create-items
create table item
(
    chrt_id      bigint,
    order_uid    varchar,
    track_number varchar(30),
    price        int,
    rid          varchar(30),
    item_name    varchar(30),
    sale         int,
    size         int,
    total_price  int,
    nm_id        bigint,
    brand        varchar(30),
    status       int,
    primary key (chrt_id),
    foreign key (order_uid)
        references orders (order_uid)
        on delete cascade
);

-- name: insert-delivery
insert into delivery (delivery_name, phone, zip, city, address, region, email)
values (?, ?, ?, ?, ?, ?, ?);

insert into payment (payment_transaction, request_id, currency, payment_provider, amount, payment_dt, bank,
                     delivery_cost, goods_total, custom_fee)
values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

select order_uid from orders;