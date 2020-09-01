package main

import "github.com/jmoiron/sqlx"

func FetchCommentsForMe(userID int) (comments []Comment, err error) {
	var entryIDs []int
	rows, err := db.Query("SELECT id FROM entries WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var entryID int
		err = rows.Scan(&entryID)
		if err != nil {
			return
		}

		entryIDs = append(entryIDs, entryID)
	}
	logger.Infow("FetchCommentsForMe", "entryIDs", entryIDs)
	query := "SELECT * FROM comments WHERE entry_id IN (?) ORDER BY created_at DESC LIMIT 10"
	inQuery, inArgs, err := sqlx.In(query, entryIDs)
	if err != nil {
		return nil, err
	}
	err = db.Select(&comments, inQuery, inArgs...)
	return
}
