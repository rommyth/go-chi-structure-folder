CREATE TABLE IF NOT EXISTS orders(
    id BIGSERIAL PRIMARY KEY,
    table_id BIGINT NOT NULL,
    notes TEXT NULL,
    order_date TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_orders_table FOREIGN KEY (table_id) REFERENCES tables(id) ON DELETE CASCADE
);

CREATE INDEX idx_orders_id ON orders(id);
