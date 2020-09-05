package main

import (
	"time"

	"github.com/jmoiron/sqlx"
)

func fetchEntriesFromComments(comments []Comment) (map[int]Entry, error) {
	var entryIDs []int
	for _, v := range comments {
		entryIDs = append(entryIDs, v.EntryID)
	}
	query := "SELECT id,user_id,private,created_at,title FROM entries WHERE id IN (?)"
	inQuery, inArgs, err := sqlx.In(query, entryIDs)
	if err != nil {
		return nil, err
	}
	rows, err := db.Query(inQuery, inArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	res := map[int]Entry{}
	for rows.Next() {
		var entryID, userID, private int
		var title string
		var createdAt time.Time
		checkErr(rows.Scan(&entryID, &userID, &private, &createdAt, &title))

		res[entryID] = Entry{entryID, userID, private == 1, title, "", createdAt}
	}
	return res, nil
}
