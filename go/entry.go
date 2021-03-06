package main

import (
	"strconv"

	"github.com/jmoiron/sqlx"
)

func FetchEntryDictByComments(comments []Comment) (map[int]Entry, error) {
	res := map[int]Entry{}
	var entryIds []string
	for _, v := range comments {
		entryIds = append(entryIds, strconv.Itoa(v.EntryID))
	}
	query := "SELECT * FROM entries WHERE id in (?)"
	inQuery, inArgs, err := sqlx.In(query, entryIds)
	if err != nil {
		return nil, err
	}
	var entries []Entry
	err = db.Select(&entries, inQuery, inArgs...)
	if err != nil {
		return nil, err
	}

	for _, v := range entries {
		res[v.ID] = v
	}
	return res, nil
}
