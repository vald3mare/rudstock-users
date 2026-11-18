-- +goose Up
CREATE TABLE user_profiles (
    user_id UUID PRIMARY KEY,
    name    TEXT NOT NULL DEFAULT '',
    phone   TEXT NOT NULL DEFAULT ''
);

CREATE TABLE user_addresses (
    id         UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id    UUID        NOT NULL,
    city       TEXT        NOT NULL CHECK (city <> ''),
    street     TEXT        NOT NULL CHECK (street <> ''),
    house      TEXT        NOT NULL CHECK (house <> ''),
    flat       TEXT        NOT NULL DEFAULT '',
    comment    TEXT        NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX user_addresses_owner_idx ON user_addresses (user_id, created_at DESC, id DESC);

CREATE TABLE orders (
    id         UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id    UUID,
    number     TEXT        NOT NULL UNIQUE,
    status     TEXT        NOT NULL,
    total      BIGINT      NOT NULL CHECK (total >= 0), -- копейки
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX orders_owner_idx ON orders (user_id, created_at DESC, id DESC)
    WHERE user_id IS NOT NULL;

CREATE TABLE order_items (
    order_id UUID    NOT NULL REFERENCES orders (id) ON DELETE CASCADE,
    position INTEGER NOT NULL CHECK (position > 0),
    card_id  UUID    NOT NULL,
    quantity INTEGER NOT NULL CHECK (quantity > 0),
    PRIMARY KEY (order_id, position)
);

-- +goose Down
DROP TABLE order_items;
DROP TABLE orders;
DROP TABLE user_addresses;
DROP TABLE user_profiles;
