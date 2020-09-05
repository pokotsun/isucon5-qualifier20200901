package main

import (
	"database/sql"
	"time"
)

func FetchFriendMap(userID int) (map[int]time.Time, error) {
	res, ok := RelationDict[userID]
	if !ok {
		return nil, sql.ErrNoRows
	}
	return res, nil
}

func FetchFriendIDs(m map[int]time.Time) []int {
	res := []int{}
	for id, _ := range m {
		res = append(res, id)
	}
	return res
}
