ALTER TABLE comments ADD INDEX user_id_idx(user_id, created_at);

-- ALTER TABLE comments DROP INDEX user_id_idx;

INSERT INTO comment_targets(comment_id, user_id,target_user_id,created_at) SELECT c.id, c.user_id, e.user_id, c.created_at FROM comments c INNER JOIN entries e ON e.id = c.entry_id;
