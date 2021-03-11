create table currency (
    id serial,
    name varchar(30) not null,
    title varchar(10) not null,

    PRIMARY KEY(id)
);

create table orders (
    id serial,
    currency_id int not null,
    rate decimal,
    createdAt timestamp without time zone,

    PRIMARY KEY(id),
    constraint fk_order_currency
        foreign key (currency_id)
        references currency(id)
)