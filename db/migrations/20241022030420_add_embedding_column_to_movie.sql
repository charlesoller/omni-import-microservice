-- +goose Up
-- +goose StatementBegin
ALTER TABLE movies
  ADD COLUMN embedding vector(384);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE movies
  DROP COLUMN embedding;
-- +goose StatementEnd
