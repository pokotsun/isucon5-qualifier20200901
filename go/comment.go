package main

import (
	"encoding/json"
	"fmt"
)

const (
	COMMENT_KEY = "c:uid:"
)

func StoreLatestComments(userID int, comments []Comment) error {
	key := fmt.Sprintf("%s%d", COMMENT_KEY, userID)
	data, err := json.Marshal(comments)
	if err != nil {
		return err
	}
	return cacheClient.SingleSet(key, data)
}

func FetchLatestComments(userID int) ([]Comment, error) {
	key := fmt.Sprintf("%s%d", COMMENT_KEY, userID)
	var comments []Comment
	data, err := cacheClient.SingleGet(key)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &comments)
	return comments, err
}

func PurgeLatestComments(userID int) error {
	key := fmt.Sprintf("%s%d", COMMENT_KEY, userID)
	return cacheClient.SingleDelete(key)
}

func PushLatestComments(userID int, comment Comment) error {
	comments, err := FetchLatestComments(userID)
	if err != nil {
		return nil
	}
	var storedComments []Comment
	storedComments = append([]Comment{comment}, comments[0:len(comments)-2]...)
	return StoreLatestComments(userID, storedComments)
}
