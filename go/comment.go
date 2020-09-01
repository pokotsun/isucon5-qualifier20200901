package main

import "github.com/jmoiron/sqlx"

func FetchCommentsForMe(userID int) (comments []Comment, err error) {
	var entryIDs []int
	err = db.Select(entryIDs, "SELECT id FROM entries WHERE user_id = ?", userID)
	if err != nil {
		return
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
