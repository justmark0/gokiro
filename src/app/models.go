package app

import (
	"database/sql"
	"fmt"
)

// DB is a sql connection variable
var DB *sql.DB

// StoringItem is a structure in database which storing one user item
type StoringItem struct {
	ID     int
	IsFile bool
	Data   string
}

//// Env is a environment variable for any request
// type Env struct {
//	db     *sql.DB
//	logger *log.Logger
//}

// CreateStorageItem is storing given data to db and returns id
func CreateStorageItem(isFile bool, data string) (int, error) {
	result, err := DB.Exec("insert into items (`isFile`, `data`) values (?, ?);", isFile, data)
	if err != nil {
		// TODO add logging
		return 0, err
	}
	affected, _ := result.RowsAffected()
	return int(affected), nil
}

// GetStorageItem Returns storage item with such id or empty storing item with error
func GetStorageItem(id int) (StoringItem, error) {
	row := DB.QueryRow("select idFile, data from items as i where i.id=$1;", id)
	item := &StoringItem{}
	switch err := row.Scan(&item.IsFile, &item.Data); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return StoringItem{}, fmt.Errorf("Not found")
	case nil:
		return *item, nil
	default:
		return StoringItem{}, err
	}
}
