package main

var (
	UserIDDict          = map[int]User{}
	UserAccountNameDict = map[string]User{}
)

func InitUserCache() (err error) {
	UserIDDict = map[int]User{}
	UserAccountNameDict = map[string]User{}

	var users []User
	err = db.Select(&users, "SELECT id, account_name, nick_name, email FROM users")
	if err != nil {
		return
	}
	for _, v := range users {
		UserIDDict[v.ID] = v
		UserAccountNameDict[v.AccountName] = v
	}
	return nil
}
