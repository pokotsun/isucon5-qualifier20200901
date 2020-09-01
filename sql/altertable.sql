ALTER TABLE footprints ADD INDEX user_id_idx(user_id);
ALTER TABLE footprints ADD UNIQUE INDEX user_id_and_owner_id_idx(user_id, owner_id);
INSERT INTO tmp_footprints(user_id, owner_id, created_at, created_date_at) SELECT user_id, owner_id, MAX(created_at), date(created_at) AS created_date_at FROM footprints GROUP BY user_id, owner_id, DATE(created_at)
