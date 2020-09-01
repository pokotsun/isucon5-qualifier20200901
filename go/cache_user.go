package main

import (
	"encoding/json"
	"fmt"
)

func getUserByIDKey(userID int) string {
	return fmt.Sprintf("USER-USERID-%d", userID)
}

func setUserToCacheByID(user User) {
	key := getUserByIDKey(user.ID)
	data, err := json.Marshal(user)
	if err != nil {
		logger.Errorf("json Marshal Err on User: %s", err)
	}

	err = cacheClient.SingleSet(key, data)
	if err != nil {
		logger.Errorf("Failed to Set Cache User: %s", err)
	}
}

func getUserFromCacheByID(userID int) (user User, err error) {
	key := getUserByIDKey(userID)
	data, err := cacheClient.SingleGet(key)
	if err != nil {
		logger.Errorf("Failed to Get Cache Of User: %s", err)
		return user, err
	}
	if err = json.Unmarshal(data, &user); err != nil {
		logger.Errorf("Failed to UnMarshal User: %s", err)
		return user, err
	}
	return user, nil
}
