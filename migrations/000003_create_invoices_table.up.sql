CREATE TABLE IF NOT EXISTS invoices(
    id BIGSERIAL PRIMARY KEY,
    order_id BIGINT NOT NULL,
    payment_method VARCHAR(20) CHECK (payment_method IN ('CASH','CREDIT')),
    payment_status VARCHAR(20) CHECK (payment_status IN ('CANCELED','PENDING', 'PAID')),
    payment_due_date TIMESTAMPTZ NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_invoices_id ON invoices(id);
