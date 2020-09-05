ALTER TABLE comments ADD INDEX user_id_idx(user_id, created_at);

-- ALTER TABLE comments DROP INDEX user_id_idx;
