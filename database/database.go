package database

import "yeni/ent"

var client *ent.Client
var err error

type postgreySQL struct{}
type MySQL struct{}
type SQLite struct{}

func Connection() {
	var database MySQL
	database.conn()
}
