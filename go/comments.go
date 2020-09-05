package main

import (
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

func fetchEntriesFromComments(comments []Comment) (map[int]Entry, error) {
	var entryIDs []int
	for _, v := range comments {
		entryIDs = append(entryIDs, v.EntryID)
	}
	query := "SELECT * FROM entries WHERE id IN (?)"
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
		var body string
		var createdAt time.Time
		checkErr(rows.Scan(&entryID, &userID, &private, &body, &createdAt))
		res[entryID] = Entry{entryID, userID, private == 1, strings.SplitN(body, "\n", 2)[0], strings.SplitN(body, "\n", 2)[1], createdAt}
	}
	return res, nil
}
