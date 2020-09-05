ALTER TABLE comments ADD INDEX user_id_idx(user_id, created_at);

-- ALTER TABLE comments DROP INDEX user_id_idx;

-- DROP TABLE IF EXISTS date_footprints;
CREATE TABLE  date_footprints (
  `user_id` int NOT NULL,
  `owner_id` int NOT NULL,
  `created_at` date NOT NULL,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY(`user_id`, `owner_id`, `created_at`));

INSERT INTO date_footprints(user_id, owner_id, created_at, updated_at) SELECT user_id, owner_id, DATE(created_at) AS date, MAX(created_at) AS updated FROM footprints GROUP BY user_id, owner_id, DATE(created_at);
ALTER TABLE date_footprints ADD INDEX idx_updated_at(updated_at);

-- ALTER TABLE entries DROP title text;
-- ALTER TABLE entries DROP content text;
ALTER TABLE entries ADD COLUMN title text;
ALTER TABLE entries ADD COLUMN content text;
UPDATE entries SET title = SUBSTRING_INDEX(body, "\n", 1), content = SUBSTRING_INDEX(body, "\n", 2);
