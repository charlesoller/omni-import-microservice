-- +goose Up
-- +goose StatementBegin
ALTER TABLE countries
  DROP COLUMN id;
ALTER TABLE countries
  ADD PRIMARY KEY (iso_3166_1);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE countries
  ADD COLUMN id SERIAL PRIMARY KEY;
ALTER TABLE countries
  DROP CONSTRAINT countries_pkey;
-- +goose StatementEnd
