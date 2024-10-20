-- +goose Up
-- +goose StatementBegin
ALTER TABLE languages
  DROP COLUMN id;
ALTER TABLE languages
  ADD PRIMARY KEY (iso_639_1);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE languages
  ADD COLUMN id SERIAL PRIMARY KEY;
ALTER TABLE languages
  DROP CONSTRAINT countries_pkey;
-- +goose StatementEnd
