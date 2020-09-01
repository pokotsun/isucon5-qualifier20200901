package main

import (
	"encoding/json"
	"fmt"
)

func getUserKeyByID(userID int) string {
	return fmt.Sprintf("USER-USERID-%d", userID)
}

func getUserKeyByEmail(email string) string {
	return fmt.Sprintf("USER-EMAIL-%s", email)
}

func setuserToCacheByKey(key string, user User) {
	data, err := json.Marshal(user)
	if err != nil {
		logger.Errorf("json Marshal Err on User: %s", err)
	}

	err = cacheClient.SingleSet(key, data)
	if err != nil {
		logger.Errorf("Failed to Set Cache User: %s", err)
	}
}

func setUserToCacheByID(user User) {
	key := getUserKeyByID(user.ID)
	setuserToCacheByKey(key, user)
}

func setUserToCacheByEmail(user User) {
	key := getUserKeyByEmail(user.Email)
	setuserToCacheByKey(key, user)
}

func getUserFromCacheByKey(key string) (user User, err error) {
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

func getUserFromCacheByID(userID int) (user User, err error) {
	key := getUserKeyByID(userID)
	return getUserFromCacheByKey(key)
}

func getUserFromCacheByEmail(email string) (User, error) {
	key := getUserKeyByEmail(email)
	return getUserFromCacheByKey(key)
}
