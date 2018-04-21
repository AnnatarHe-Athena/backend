package model

import (
	"database/sql"
	"encoding/base64"
	"log"
	"time"

	"github.com/douban-girls/backend/proto"
	"github.com/douban-girls/backend/utils"
)

type Cell struct {
	ID         int    `json:"id"`
	Img        string `json:"img"`
	Text       string `json:"text"`
	Permission int    `json:"permission"`
	Cate       int    `json:"cate"`
	FromID     string `json:"from_id"`
	FromURL    string `json:"from_url"`
	CreatedAt  int64  `json:"createdAt"`
	Content    string `json:"content"`
	Md5        string `json:"md5"`
}

type Cells []*Cell

const fieldQuery = "SELECT id, text, content, img, cate, premission, from_url, from_id, createdat FROM cells WHERE "

func CellsFetchAll(cate, row, offset, permission int32) (Cells, error) {
	var rows *sql.Rows
	var err error
	if permission == 3 {
		rows, err = DBInstance.Query(fieldQuery+"premission=$1 ORDER BY id DESC LIMIT $2 OFFSET $3", permission, row, offset)
	} else {
		rows, err = DBInstance.Query(fieldQuery+"cate=$1 AND premission=$2 ORDER BY id DESC LIMIT $3 OFFSET $4", cate, permission, row, offset)
	}

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	result := GetCellsFromRows(rows)
	return result, nil
}

func GetCellsFromRows(rows *sql.Rows) (result Cells) {
	for rows.Next() {
		var id, cate, permission int
		var text, content, img, fromID, fromURL, createdAt string
		if err := rows.Scan(&id, &text, &content, &img, &cate, &permission, &fromURL, &fromID, &createdAt); err != nil {
			utils.ErrorLog(err)
			return
		}
		createdAtUnix := utils.Timestamp(createdAt)
		result = append(result, &Cell{
			ID:         id,
			Img:        img,
			Text:       text,
			Content:    content,
			Permission: permission,
			Cate:       cate,
			FromID:     fromID,
			FromURL:    fromURL,
			CreatedAt:  createdAtUnix,
		})
	}
	result.EncodeImageURL()
	return
}
func (cs Cells) Save() error {
	stat, err := DBInstance.Prepare("INSERT INTO cells(img, text, cate, premission, from_id, from_url) VALUES($1, $2, $3, $4, $5, $6) ON CONFLICT (img) DO NOTHING RETURNING id")
	if err != nil {
		utils.ErrorLog(err)
		return err
	}
	for _, cell := range cs {
		var id int
		err := stat.QueryRow(cell.Img, cell.Text, cell.Cate, cell.Permission, cell.FromID, cell.FromURL).Scan(&id)
		cell.ID = id
		log.Println(*cell)
		if err != nil {
			utils.ErrorLog(err)
			return err
		}
	}
	return nil
}

func (cell *Cell) Save() error {

	stat, err := DBInstance.Prepare("INSERT INTO cells(img, text, cate, premission, md5, from_id, from_url) VALUES($1, $2, $3, $4, $5, $6, $7) ON CONFLICT (img) DO NOTHING RETURNING id")
	if err != nil {
		utils.ErrorLog(err)
		return err
	}
	var id int
	e := stat.QueryRow(cell.Img, cell.Text, cell.Cate, cell.Permission, cell.Md5, cell.FromID, cell.FromURL).Scan(&id)
	if e != nil {
		utils.ErrorLog(err)
		return err
	}
	cell.ID = id
	return nil
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
	return &proto.CellItem{
		Id:         int32(cell.ID),
		Img:        cell.Img,
		Text:       cell.Text,
		Content:    cell.Content,
		Permission: int32(cell.Permission),
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

func (this *Cell) Remove(reallyDestroy bool) bool {
	var rows *sql.Rows
	var err error
	if reallyDestroy {
		rows, err = DBInstance.Query("DELETE FROM cells WHERE id=$1", this.ID)
	} else {
		rows, err = DBInstance.Query("UPDATE cells SET premission=3, updatedat=$1 WHERE id=$2", time.Now(), this.ID)
	}
	if err != nil {
		log.Println("error occean when cell to hide or remove", err)
		return false
	}
	defer rows.Close()
	return true
}
