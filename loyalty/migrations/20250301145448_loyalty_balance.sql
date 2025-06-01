-- +goose Up
-- +goose StatementBegin
CREATE TABLE loyalty_balance (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id INT NOT NULL,
    count INT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE loyalty_balance;
-- +goose StatementEnd
