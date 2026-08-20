CREATE TABLE IF NOT EXISTS order_items(
    id BIGSERIAL PRIMARY KEY,
    order_id BIGINT NOT NULL,
    food_id BIGINT NOT NULL,
    quantity INTEGER NOT NULL CHECK (quantity > 0),
    unit_price NUMERIC(10,2) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_order_items_order FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
    CONSTRAINT fk_order_items_food FOREIGN KEY (food_id) REFERENCES foods(id) ON DELETE RESTRICT
);

CREATE INDEX idx_order_items ON order_items(id);
CREATE INDEX idx_orders_items_id ON order_items(order_id);
CREATE INDEX idx_orders_items_food ON order_items(food_id);
