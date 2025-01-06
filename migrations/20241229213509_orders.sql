-- +goose Up
CREATE TABLE delivery (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    phone VARCHAR NOT NULL,
    zip VARCHAR NOT NULL,
    city VARCHAR NOT NULL,
    address VARCHAR NOT NULL,
    region VARCHAR NOT NULL,
    email VARCHAR NOT NULL
);

CREATE TABLE payment (
    id SERIAL PRIMARY KEY,
    transaction VARCHAR UNIQUE NOT NULL,
    request_id VARCHAR UNIQUE,
    currency VARCHAR NOT NULL,
    provider VARCHAR NOT NULL,
    amount INTEGER NOT NULL,
    payment_dt INTEGER UNIQUE NOT NULL,
    bank VARCHAR NOT NULL,
    delivery_cost INTEGER NOT NULL,
    goods_total INTEGER NOT NULL,
    custom_fee INTEGER NOT NULL
);

CREATE TABLE item (
    id SERIAL PRIMARY KEY,
    chrt_id INTEGER UNIQUE NOT NULL,
    track_number VARCHAR NOT NULL,
    price INTEGER NOT NULL,
    rid VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    sale INTEGER NOT NULL,
    size VARCHAR NOT NULL,
    total_price INTEGER NOT NULL,
    nm_id INTEGER UNIQUE NOT NULL,
    brand VARCHAR NOT NULL,
    status INTEGER NOT NULL
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    order_uid VARCHAR UNIQUE NOT NULL,
    track_number VARCHAR UNIQUE NOT NULL,
    entry VARCHAR NOT NULL,
    delivery_id INTEGER REFERENCES delivery(id),
    payment_id INTEGER REFERENCES payment(id),
    locale VARCHAR NOT NULL,
    internal_signature VARCHAR,
    customer_id VARCHAR NOT NULL,
    delivery_service VARCHAR NOT NULL,
    shardkey VARCHAR NOT NULL,
    sm_id INTEGER NOT NULL,
    date_created TIMESTAMP NOT NULL DEFAULT NOW(),
    oof_shard VARCHAR NOT NULL
);

CREATE TABLE orders_and_items (
    id SERIAL PRIMARY KEY,
    order_id INTEGER REFERENCES orders(id),
    item_id INTEGER REFERENCES item(id)
);

-- +goose Down
DROP TABLE delivery;
DROP TABLE payment;
DROP TABLE item;
DROP TABLE orders;
DROP TABLE orders_and_items;