package main

import "time"

type Relation struct {
	One       int       `db:"one"`
	Another   int       `db:"another"`
	CreatedAt time.Time `db:"created_at"`
}

var (
	RelationDict = map[int]map[int]time.Time{}
)

func InitRelationCache() error {
	var rel []Relation
	err := db.Select(&rel, "SELECT one, another, created_at FROM relations")
	if err != nil {
		return err
	}
	for _, v := range rel {
		if RelationDict[v.One] == nil {
			RelationDict[v.One] = map[int]time.Time{}
		}
		RelationDict[v.One][v.Another] = v.CreatedAt
	}
	return nil
}

func AddRelation(one, another int) {
	if RelationDict[one] == nil {
		RelationDict[one] = map[int]time.Time{}
	}
	if RelationDict[another] == nil {
		RelationDict[another] = map[int]time.Time{}
	}
	now := time.Now()
	RelationDict[one][another] = now
	RelationDict[another][one] = now
}
