-- +goose Up
-- +goose StatementBegin
ALTER TABLE "orders"
    ADD COLUMN pick_up_point_id INTEGER;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "orders"
    DROP COLUMN pick_up_point_id;
-- +goose StatementEnd
