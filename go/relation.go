package main

import "time"

func FetchFriendMap(userID int) (map[int]time.Time, error) {
	res := map[int]time.Time{}
	rows, err := db.Query(`SELECT another, created_at FROM relations WHERE one = ?`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var anotherID int
		var createdAt time.Time

		if err := rows.Scan(anotherID, createdAt); err != nil {
			return nil, err
		}
		res[anotherID] = createdAt
	}
	return res, nil
}
