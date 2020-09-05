ALTER TABLE comments ADD INDEX user_id_idx(user_id, created_at);

-- ALTER TABLE comments DROP INDEX user_id_idx;

ALTER TABLE entries ADD COLUMN title text;
ALTER TABLE entries ADD COLUMN content text;
UPDATE entries SET title = SUBSTRING_INDEX(body, "\n", 1), content = SUBSTRING_INDEX(body, "\n", 2);
-- ALTER TABLE entries DROP title text;
-- ALTER TABLE entries DROP content text;
