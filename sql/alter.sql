ALTER TABLE comments ADD INDEX user_id_idx(user_id, created_at);

-- ALTER TABLE comments DROP INDEX user_id_idx;

INSERT INTO comment_targets(comment_id, user_id,target_user_id,created_at) SELECT c.id, c.user_id, e.user_id, c.created_at FROM comments c INNER JOIN entries e ON e.id = c.entry_id;

-- DROP TABLE IF EXISTS date_footprints;
CREATE TABLE  date_footprints (
  `user_id` int NOT NULL,
  `owner_id` int NOT NULL,
  `created_at` date NOT NULL,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY(`user_id`, `owner_id`, `created_at`));

INSERT INTO date_footprints(user_id, owner_id, created_at, updated_at) SELECT user_id, owner_id, DATE(created_at) AS date, MAX(created_at) AS updated FROM footprints GROUP BY user_id, owner_id, DATE(created_at);
ALTER TABLE date_footprints ADD INDEX idx_updated_at(user_id, updated_at);
