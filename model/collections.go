package model

import (
	"log"

	"github.com/douban-girls/backend/utils"
)

// Collection is map to db
type Collection struct {
	ID    int `json:"id" db:"id"`
	Cell  int `json:"cell" db:"cell"`
	Owner int `json:"owner" db:"owner"`
}

// Collections is a list of collection
type Collections []*Collection

// Save will save the map that cell.id to user.id
func (cs Collections) Save() error {
	stat, err := DBInstance.Prepare("INSERT INTO collections(cell, owner) VALUES($1, $2) RETURNING id")
	if err != nil {
		return err
	}
	for _, collection := range cs {
		var id int
		var _c Collection

		// 如果已经存在这条收藏记录，则不再插入
		if err := DBInstance.Get(&_c, "SELECT id from collections WHERE cell=$1 AND owner=$2", collection.Cell, collection.Owner); err != nil {
			log.Println(err)
			continue
		}

		if id != 0 {
			continue
		}

		err := stat.QueryRow(collection.Cell, collection.Owner).Scan(&id)
		collection.ID = id
		if err != nil {
			return err
		}
	}
	return nil
}

func FetchUserCollectionBy(id, from, size int) (Cells, error) {
	rows, err := DBInstance.Query("SELECT cells.id, cells.text, cells.content, cells.img, cells.cate, cells.premission, cells.from_url, cells.from_id, cells.createdAt FROM cells, collections WHERE collections.cell = cells.id AND collections.owner = $1 OFFSET $2 LIMIT $3", id, from, size)

	if err != nil {
		utils.ErrorLog(err)
		return nil, err
	}
	defer rows.Close()
	collections := GetCellsFromRows(rows)

	return collections, err
}
