-- +goose Up
-- +goose StatementBegin
CREATE TABLE l0
(
    id varchar PRIMARY KEY NOT NULL ,
    data json NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE l0
-- +goose StatementEnd
