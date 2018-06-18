package model

import "github.com/douban-girls/backend/utils"

type Collection struct {
	ID    int `json:"id"`
	Cell  int `json:"cell"`
	Owner int `json:"owner"`
}

type Collections []*Collection

func (cs Collections) Save() error {
	stat, err := DBInstance.Prepare("INSERT INTO collections(cell, owner) VALUES($1, $2) RETURNING id")
	if err != nil {
		return err
	}
	for _, collection := range cs {
		var id int
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
