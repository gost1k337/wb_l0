CREATE TABLE IF NOT EXISTS orders
(
    id               SERIAL PRIMARY KEY,
    order_uid        VARCHAR(255),
    track_number     VARCHAR(255),
    entry            VARCHAR(255),
    deliveries       JSONB,
    payments         JSONB,
    items            JSONB,
    locale           VARCHAR(10),
    customer_id      VARCHAR(255),
    delivery_service VARCHAR(255),
    shardkey         VARCHAR(10),
    sm_id            INT,
    date_created     TIMESTAMP,
    oof_shard        VARCHAR(10)
);
