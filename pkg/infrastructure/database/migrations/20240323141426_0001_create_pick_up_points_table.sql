-- +goose Up
-- +goose StatementBegin
CREATE TABLE "pick_up_points"
(
    id           SERIAL PRIMARY KEY NOT NULL,
    name         VARCHAR(50)        NOT NULL,
    phone_number VARCHAR(20)        NOT NULL,
    region       VARCHAR(50)        NOT NULL,
    city         VARCHAR(50)        NOT NULL,
    street       VARCHAR(50)        NOT NULL,
    house_num    VARCHAR(10)        NOT NULL,

    CONSTRAINT pick_up_points_name_unique UNIQUE (name)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE pick_up_points;
-- +goose StatementEnd
