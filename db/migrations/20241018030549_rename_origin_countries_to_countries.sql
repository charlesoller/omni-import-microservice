-- +goose Up
-- +goose StatementBegin

ALTER TABLE origin_countries
RENAME TO countries;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE countries
RENAME TO origin_countries;

-- +goose StatementEnd
