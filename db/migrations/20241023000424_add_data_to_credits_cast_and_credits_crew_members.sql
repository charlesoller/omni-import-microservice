-- +goose Up
-- +goose StatementBegin
ALTER TABLE credits_cast_member
  ADD COLUMN character TEXT NOT NULL DEFAULT '',
  ADD COLUMN "order" INT NOT NULL DEFAULT 0;

ALTER TABLE credits_crew_member
  ADD COLUMN department VARCHAR(255) NOT NULL DEFAULT '',
  ADD COLUMN job VARCHAR(255) NOT NULL DEFAULT '';

UPDATE credits_cast_member
SET character = (SELECT character FROM cast_members WHERE credits_cast_member.cast_id = cast_members.id),
    "order" = (SELECT "order" FROM cast_members WHERE credits_cast_member.cast_id = cast_members.id);

UPDATE credits_crew_member
SET job = (SELECT job FROM crew_members WHERE credits_crew_member.crew_id = crew_members.id),
    department = (SELECT department FROM crew_members WHERE credits_crew_member.crew_id = crew_members.id);

ALTER TABLE cast_members
  DROP COLUMN character,
  DROP COLUMN "order";

ALTER TABLE crew_members
  DROP COLUMN job,
  DROP COLUMN department;

ALTER TABLE credits_cast_member
  RENAME COLUMN cast_id TO member_id;

ALTER TABLE credits_crew_member 
  RENAME COLUMN crew_id TO member_id;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE credits_cast_member
  RENAME COLUMN member_id TO cast_id;

ALTER TABLE credits_crew_member 
  RENAME COLUMN member_id TO crew_id;

ALTER TABLE cast_members
  ADD COLUMN character VARCHAR(255) NOT NULL DEFAULT '',
  ADD COLUMN "order" INT NOT NULL DEFAULT 0;

ALTER TABLE crew_members
  ADD COLUMN job VARCHAR(255) NOT NULL DEFAULT '',
  ADD COLUMN department VARCHAR(255) NOT NULL DEFAULT '';
-- +goose StatementEnd
