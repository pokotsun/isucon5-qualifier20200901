package main

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

type User struct {
	ID          int    `db:"id"`
	AccountName string `db:"account_name"`
	NickName    string `db:"nick_name"`
	Email       string `db:"email"`
	PassHash    string `db:"passhash"`
	Salt        string `db:"salt"`
}

type Profile struct {
	UserID    int            `db:"user_id"`
	FirstName string         `db:"first_name"`
	LastName  string         `db:"last_name"`
	Sex       string         `db:"sex"`
	Birthday  mysql.NullTime `db:"birthday"`
	Pref      string         `db:"pref"`
	UpdatedAt time.Time      `db:"updated_at"`
}

type Entry struct {
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	Private   bool      `db:"private"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
}

type Comment struct {
	ID        int       `db:"id"`
	EntryID   int       `db:"entry_id"`
	UserID    int       `db:"user_id"`
	Comment   string    `db:"comment"`
	CreatedAt time.Time `db:"created_at"`
}

type Friend struct {
	ID        int       `db:"id"`
	CreatedAt time.Time `db:"created_at"`
}

type Footprint struct {
	UserID    int       `db:"user_id"`
	OwnerID   int       `db:"owner_id"`
	CreatedAt time.Time `db:"created_at"`
	Updated   time.Time `db:"updated"`
}

type Relation struct {
	ID        int       `db:"id"`
	One       int       `db:"one"`
	Another   int       `db:"another"`
	CreatedAt time.Time `db:"created_at"`
}
