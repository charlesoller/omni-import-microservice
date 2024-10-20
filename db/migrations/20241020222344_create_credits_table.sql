-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS cast_members (
    id SERIAL PRIMARY KEY,
    cast_id INT NOT NULL DEFAULT 0,
    character VARCHAR(255) NOT NULL DEFAULT '',
    credit_id VARCHAR(255) NOT NULL DEFAULT '',
    gender SMALLINT NOT NULL DEFAULT 0,
    adult BOOLEAN NOT NULL DEFAULT false,
    known_for_department VARCHAR(255) NOT NULL DEFAULT '',
    name VARCHAR(255) NOT NULL DEFAULT '',
    original_name VARCHAR(255) NOT NULL DEFAULT '',
    popularity FLOAT NOT NULL DEFAULT 0,
    profile_path VARCHAR(255) NOT NULL DEFAULT '',
    "order" INT NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS crew_members (
    id SERIAL PRIMARY KEY,
    credit_id VARCHAR(255) NOT NULL DEFAULT '',
    department VARCHAR(255) NOT NULL DEFAULT '',
    job VARCHAR(255) NOT NULL DEFAULT '',
    gender SMALLINT NOT NULL DEFAULT 0,
    adult BOOLEAN NOT NULL DEFAULT false,
    known_for_department VARCHAR(255) NOT NULL DEFAULT '',
    name VARCHAR(255) NOT NULL NOT NULL DEFAULT '',
    original_name VARCHAR(255) NOT NULL DEFAULT '',
    popularity FLOAT NOT NULL DEFAULT 0,
    profile_path VARCHAR(255) NOT NULL DEFAULT ''
);

CREATE TABLE IF NOT EXISTS credits (
  id INT PRIMARY KEY,
  FOREIGN KEY (id) REFERENCES movies(id) ON DELETE CASCADE 
);

CREATE TABLE IF NOT EXISTS credits_cast_member (
    credit_id INT NOT NULL,
    cast_id INT NOT NULL,
    FOREIGN KEY (credit_id) REFERENCES credits(id) ON DELETE CASCADE,
    FOREIGN KEY (cast_id) REFERENCES cast_members(id) ON DELETE CASCADE,
    PRIMARY KEY (credit_id, cast_id)
);

CREATE TABLE IF NOT EXISTS credits_crew_member (
    credit_id INT NOT NULL,
    crew_id INT NOT NULL,
    FOREIGN KEY (credit_id) REFERENCES credits(id) ON DELETE CASCADE,
    FOREIGN KEY (crew_id) REFERENCES crew_members(id) ON DELETE CASCADE,
    PRIMARY KEY (credit_id, crew_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS crew_members;
DROP TABLE IF EXISTS cast_members;
DROP TABLE IF EXISTS credits;
DROP TABLE IF EXISTS credits_cast_member;
DROP TABLE IF EXISTS credits_crew_member;
-- +goose StatementEnd
