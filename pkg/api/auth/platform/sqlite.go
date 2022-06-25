package platform

import (
	root "github.com/DarioCalovic/secretify"
	utildb "github.com/DarioCalovic/secretify/pkg/util/db"
)

type SQLite struct{}

func NewSQLiteAuthUserRepository() *SQLite {
	return &SQLite{}
}

// Create creates a new user on database
func (s SQLite) Create(db utildb.DB, srv root.User) (root.User, error) {
	sqlitedb := db.(*utildb.SQLiteDB)
	err := sqlitedb.DBConn.Create(&srv).Error
	return srv, err
}

// View returns single user by ID
func (s SQLite) View(db utildb.DB, id int) (root.User, error) {
	var user root.User
	return user, nil
}

// FindByUsername queries for single user by username
func (s SQLite) FindByUsername(db utildb.DB, uname string) (root.User, error) {
	sqlitedb := db.(*utildb.SQLiteDB)
	var user root.User
	err := sqlitedb.DBConn.Where("username = ?", uname).First(&user).Error
	if err != nil {
		return root.User{}, err
	}
	return user, err
}

// FindByToken queries for single user by token
func (s SQLite) FindByToken(db utildb.DB, token string) (root.User, error) {
	var user root.User
	// TODO : Well do it
	return user, nil
}

// Update
func (s SQLite) Update(db utildb.DB, user root.User) error {
	sqlitedb := db.(*utildb.SQLiteDB)
	return sqlitedb.DBConn.Model(&root.User{}).Where("id = ?", user.ID).Updates(user).Error
}
