-- +goose Up
CREATE TABLE delivery (
    id INT PRIMARY KEY,
    name VARCHAR NOT NULL,
    phone VARCHAR NOT NULL,
    zip VARCHAR NOT NULL,
    city VARCHAR NOT NULL,
    address VARCHAR NOT NULL,
    region VARCHAR NOT NULL,
    email VARCHAR NOT NULL
);

CREATE TABLE payment (
    id INT PRIMARY KEY,
    transaction VARCHAR UNIQUE NOT NULL,
    request_id VARCHAR UNIQUE,
    currency VARCHAR NOT NULL,
    provider VARCHAR NOT NULL,
    amount INT NOT NULL,
    payment_dt INT UNIQUE NOT NULL,
    bank VARCHAR NOT NULL,
    delivery_cost INT NOT NULL,
    goods_total INT NOT NULL,
    custom_fee INT NOT NULL
);

CREATE TABLE item (
    id INT PRIMARY KEY,
    chrt_id INT UNIQUE NOT NULL,
    track_number VARCHAR NOT NULL,
    price INT NOT NULL,
    rid VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    sale INT NOT NULL,
    size VARCHAR NOT NULL,
    total_price INT NOT NULL,
    nm_id INT UNIQUE NOT NULL,
    brand VARCHAR NOT NULL,
    status INT NOT NULL
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    order_uid VARCHAR UNIQUE NOT NULL,
    track_number VARCHAR UNIQUE NOT NULL,
    entry VARCHAR NOT NULL,
    delivery_id INT REFERENCES delivery(id),
    payment_id INT REFERENCES payment(id),
    locale VARCHAR NOT NULL,
    internal_signature VARCHAR,
    customer_id VARCHAR NOT NULL,
    delivery_service VARCHAR NOT NULL,
    shardkey VARCHAR NOT NULL,
    sm_id INT NOT NULL,
    date_created TIMESTAMP NOT NULL DEFAULT NOW(),
    oof_shard VARCHAR NOT NULL
);

CREATE TABLE orders_and_items (
    id INT PRIMARY KEY,
    order_id INT REFERENCES orders(id),
    item_id INT REFERENCES item(id)
);

-- +goose Down
DROP TABLE delivery;
DROP TABLE payment;
DROP TABLE item;
DROP TABLE orders;
DROP TABLE orders_and_items;