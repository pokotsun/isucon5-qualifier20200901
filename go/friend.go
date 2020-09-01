package main

import "time"

func FetchFriendDict(userID int) (map[int]time.Time, error) {
	res := map[int]time.Time{}
	relations := []Relation{}
	err := db.Select(&relations, "SELECT another, created_at FROM relations WHERE one = ?", userID)
	if err != nil {
		return nil, err
	}
	for _, v := range relations {
		res[v.Another] = v.CreatedAt
	}
	return res, nil
}

func isFriendInDict(friendDict map[int]time.Time, targetID int) bool {
	_, ok := friendDict[targetID]
	return ok
}

func isPermitted(friendDict map[int]time.Time, selfID, targetID int) bool {
	if selfID == targetID {
		return true
	}
	return isFriendInDict(friendDict, targetID)
}

func FetchFriendComments(userID int) (comments []Comment, err error) {
	err = db.Select(&comments, "SELECT * FROM comments WHERE user_id IN (SELECT another FROM relations WHERE one = ?) ORDER BY created_at DESC LIMIT 10", userID)
	return
}
