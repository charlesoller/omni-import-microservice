-- +goose Up
-- +goose StatementBegin
ALTER TABLE cast_members
  ALTER COLUMN character TYPE TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE cast_members
  ALTER COLUMN character TYPE VARCHAR(255);
-- +goose StatementEnd
