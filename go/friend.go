package main

func FetchFriendDict(userID int) (map[int]bool, error) {
	res := map[int]bool{}
	relations := []Relation{}
	err := db.Select(&relations, "SELECT another FROM relations WHERE one = ?", userID)
	if err != nil {
		return nil, err
	}
	for _, v := range relations {
		res[v.Another] = true
	}
	return res, nil
}

func isFriendInDict(friendDict map[int]bool, targetID int) bool {
	_, ok := friendDict[targetID]
	return ok
}

func isPermitted(friendDict map[int]bool, selfID, targetID int) bool {
	if selfID == targetID {
		return true
	}
	return isFriendInDict(friendDict, targetID)
}

func FetchFriendComments(userID int) (comments []Comment, err error) {
	err = db.Select(&comments, "SELECT * FROM entries WHERE user_id IN (SELECT another FROM relations WHERE one = ?) LIMIT 10", userID)
	return
}
