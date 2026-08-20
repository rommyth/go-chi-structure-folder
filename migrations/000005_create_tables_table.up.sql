CREATE TABLE IF NOT EXISTS tables(
    id BIGSERIAL PRIMARY KEY,
    number_of_guests INTEGER NOT NULL CHECK (number_of_guests < 0),
    table_number INTEGER NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_tables_id ON tables(id);
