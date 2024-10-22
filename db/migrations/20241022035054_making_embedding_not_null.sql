-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION generate_zero_vector(length INT) 
RETURNS float8[] AS $$
BEGIN
    RETURN ARRAY(SELECT 0.0 FROM generate_series(1, length));
END;
$$ LANGUAGE plpgsql;

UPDATE movies
SET embedding = generate_zero_vector(384)
WHERE embedding IS NULL;

ALTER TABLE movies
  ALTER COLUMN embedding SET DEFAULT generate_zero_vector(384),
  ALTER COLUMN embedding SET NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE movies
  ALTER COLUMN embedding DROP DEFAULT,
  ALTER COLUMN embedding DROP NOT NULL;

DROP FUNCTION IF EXISTS generate_zero_vector();
-- +goose StatementEnd
