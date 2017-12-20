package model

import (
	"database/sql"
	"encoding/base64"
	"time"

	"github.com/douban-girls/backend/proto"
	"github.com/douban-girls/backend/utils"
)

type Cell struct {
	ID         int    `json:"id"`
	Img        string `json:"img"`
	Text       string `json:"text"`
	Premission int    `json:"premission"`
	Cate       int    `json:"cate"`
	FromID     string `json:"from_id"`
	FromURL    string `json:"from_url"`
	CreatedAt  int64  `json:"createdAt"`
}

type Cells []*Cell

func CellsFetchAll(cate, row, offset, premission int32) (Cells, error) {
	rows, err := DBInstance.Query("SELECT id, text, img, cate, premission, from_url, from_id, createdat FROM cells WHERE cate=$1 AND premission=$2 ORDER BY id DESC LIMIT $3 OFFSET $4", cate, premission, row, offset)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	result := GetCellsFromRows(rows)
	return result, nil
}

func GetCellsFromRows(rows *sql.Rows) (result Cells) {
	for rows.Next() {
		var id, cate, premission int
		var text, img, fromID, fromURL, createdAt string
		if err := rows.Scan(&id, &text, &img, &cate, &premission, &fromURL, &fromID, &createdAt); err != nil {
			utils.ErrorLog(err)
			return
		}
		createdAtUnix := getTimestamp(createdAt)
		result = append(result, &Cell{
			ID:         id,
			Img:        img,
			Text:       text,
			Premission: premission,
			Cate:       cate,
			FromID:     fromID,
			FromURL:    fromURL,
			CreatedAt:  createdAtUnix,
		})
	}
	result.EncodeImageURL()
	return
}

func getTimestamp(createdBy string) (createdAtUnix int64) {
	timestamp, err := time.Parse(time.RFC3339, createdBy)
	if err != nil {
		createdAtUnix = time.Now().Unix()
		utils.ErrorLog(err)
	} else {
		createdAtUnix = timestamp.Unix()
	}
	return
}

func (cs Cells) EncodeImageURL() {
	for _, cell := range cs {
		cell.EncodeImageURL()
	}
}

func (cell *Cell) EncodeImageURL() {
	cell.Img = base64.StdEncoding.EncodeToString([]byte(cell.Img))
}

func (cell Cell) ConvertToProtoType() *proto.CellItem {
	// TODO
	return &proto.CellItem{
		Id:         int32(cell.ID),
		Img:        cell.Img,
		Text:       cell.Text,
		Premission: int32(cell.Premission),
		Cate:       int32(cell.Cate),
		FromID:     cell.FromID,
		FromURL:    cell.FromURL,
		CreatedAt:  cell.CreatedAt,
	}
}

func (cs Cells) ConvertToProtoType() (result []*proto.CellItem) {
	for _, item := range cs {
		result = append(result, item.ConvertToProtoType())
	}
	return
}
