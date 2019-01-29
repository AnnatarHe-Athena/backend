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
	ID         int           `json:"id" db:"id"`
	Img        string        `json:"img" db:"img"`
	Text       string        `json:"text" db:"text"`
	Permission int           `json:"permission" db:"premission"`
	Cate       int           `json:"cate" db:"cate"`
	FromID     string        `json:"from_id" db:"from_id"`
	FromURL    string        `json:"from_url" db:"from_url"`
	CreatedAt  time.Time     `json:"createdAt" db:"createdat"`
	CreatedBy  sql.NullInt64 `json:"createdBy" db:"createdby"`
	UpdatedAt  time.Time     `json:"updatedAt" db:"updatedat"`
	Content    string        `json:"content" db:"content"`
	Likes      int           `json:"likes" db:"likes"`
	Md5        string        `json:"md5" db:"md5"`
}

type Cells []*Cell

const fieldQuery = "SELECT id, text, content, img, cate, premission, from_url, from_id, createdat FROM cells WHERE "

func CellsFetchAll(cate, row, offset, permission int32) (Cells, error) {
	var err error

	cellList := []*Cell{}

	if permission == 3 {
		err = DBInstance.Select(&cellList, "SELECT * FROM cells WHERE premission=$1 ORDER BY id DESC LIMIT $2 OFFSET $3", permission, row, offset)
	} else {
		err = DBInstance.Select(&cellList, "SELECT * FROM cells WHERE cate=$1 AND premission=$2 ORDER BY id DESC LIMIT $3 OFFSET $4", cate, permission, row, offset)
	}

	log.Println(cellList, err)

	return cellList, err

	// result := GetCellsFromRows(rows)
	// return result, nil
}

func GetCellsFromRows(rows *sql.Rows) (result Cells) {
	for rows.Next() {
		var id, cate, permission int
		var text, content, img, fromID, fromURL string
		var createdAt time.Time
		if err := rows.Scan(&id, &text, &content, &img, &cate, &permission, &fromURL, &fromID, &createdAt); err != nil {
			utils.ErrorLog(err)
			return
		}
		result = append(result, &Cell{
			ID:         id,
			Img:        img,
			Text:       text,
			Content:    content,
			Permission: permission,
			Cate:       cate,
			FromID:     fromID,
			FromURL:    fromURL,
			CreatedAt:  createdAt,
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

	defer stat.Close()
	for _, cell := range cs {
		var id int
		err := stat.QueryRow(cell.Img, cell.Text, cell.Cate, cell.Permission, cell.FromID, cell.FromURL).Scan(&id)
		cell.ID = id
		log.Println(cell)
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

	defer stat.Close()
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
		CreatedAt:  int64(cell.CreatedAt.Second()),
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
