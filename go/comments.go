package main

import "github.com/jmoiron/sqlx"

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
	var entries []Entry
	err = db.Select(&entries, inQuery, inArgs...)
	if err != nil {
		return nil, err
	}
	res := map[int]Entry{}
	for _, v := range entries {
		res[v.ID] = v
	}
	return res, nil
}
