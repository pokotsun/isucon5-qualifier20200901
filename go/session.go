package main

import "net/http"

func getCurrentUserID(w http.ResponseWriter, r *http.Request) int {
	session := getSession(w, r)
	userID, ok := session.Values["user_id"]
	if !ok {
		return 0
	}
	return userID.(int)
}
