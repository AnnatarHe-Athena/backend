package model

import (
	pb "github.com/douban-girls/backend/proto"
	"github.com/douban-girls/backend/utils"
)

// Category deprecated 旧有的分类系统，之后会逐步被 tags 替代
type Category struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Src       int    `json:"src"`
	Count     int    `json:"count"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}
type Categories []Category

// FetchAllCategories deprecated 获取所有分类
func FetchAllCategories() (Categories, error) {
	rows, err := DBInstance.Query(`SELECT categories.id, categories.name, categories.src, count(cells.id), categories.createdat, categories.updatedat  FROM categories, cells WHERE categories.id = cells.cate GROUP BY categories.id;`)
	defer rows.Close()
	categories := []Category{}
	if err != nil {
		return categories, err
	}

	for rows.Next() {
		var id, src, count int
		var name, createdAt, updatedAt string
		rows.Scan(&id, &name, &src, &count, &createdAt, &updatedAt)
		category := Category{
			ID:        id,
			Name:      name,
			Src:       src,
			Count:     count,
			CreatedAt: utils.Timestamp(createdAt),
			UpdatedAt: utils.Timestamp(updatedAt),
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (this Category) Convert() *pb.CategoryItem {
	return &pb.CategoryItem{
		Id:        int32(this.ID),
		Name:      this.Name,
		Src:       int32(this.Src),
		Count:     int64(this.Count),
		CreatedAt: this.CreatedAt,
		UpdatedAt: this.UpdatedAt,
	}
}

func (categories Categories) Convert() (result []*pb.CategoryItem) {
	for _, c := range categories {
		result = append(result, c.Convert())
	}
	return
}
