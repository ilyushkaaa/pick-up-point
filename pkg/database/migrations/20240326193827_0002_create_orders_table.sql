-- +goose Up
-- +goose StatementBegin
CREATE TABLE "orders"
(
    id                      INTEGER PRIMARY KEY NOT NULL,
    client_id               INTEGER             NOT NULL,
    weight                  FLOAT               NOT NULL,
    price                   FLOAT               NOT NULL,
    package_type            VARCHAR,
    storage_expiration_date TIMESTAMP           NOT NULL,
    order_issue_date        TIMESTAMP,
    is_returned             BOOLEAN             NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "orders";
-- +goose StatementEnd

