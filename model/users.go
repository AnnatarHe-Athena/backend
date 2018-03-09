package model

import (
	"database/sql"
	"errors"
	"log"

	pb "github.com/douban-girls/backend/proto"
	"github.com/douban-girls/backend/utils"
)

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Pwd       string `json:"-"`
	Avatar    string `json:"avatar"`
	Role      int    `json:"role"`
	Bio       string `json:"bio"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

const (
	fetchUserFieldPrefix = "SELECT id, email, name, pwd, avatar, role, bio, createdat, updatedat FROM users"
)

func (u *User) Auth() error {
	if u.Email == "" || u.Pwd == "" {
		return errors.New("auth fail. because no email or pwd field")
	}
	query := fetchUserFieldPrefix + " WHERE email=$1 AND pwd=$2"

	row := DBInstance.QueryRow(query, u.Email, u.Pwd)
	err := getUserDataStruct(row, u)
	if err != nil {
		log.Println(err)
		return errors.New("auth fail. email or password incorrect")
	}
	return nil
}

// Find a user by userID
func (u *User) Find() error {
	row := DBInstance.QueryRow(fetchUserFieldPrefix+" WHERE id=$1", u.ID)
	err := getUserDataStruct(row, u)
	if err != nil {
		log.Println(err)
		return errors.New("404: user not found")
	}
	return nil
}

func (u *User) Collections(offset, size int) (Cells, error) {
	return FetchUserCollectionBy(u.ID, offset, size)
}

func (u User) ConvertToProto() *pb.UserItem {
	return &pb.UserItem{
		Id:        int64(u.ID),
		Email:     u.Email,
		Pwd:       u.Pwd,
		Avatar:    u.Avatar,
		Name:      u.Name,
		Bio:       u.Bio,
		Role:      int32(u.Role),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}

}

func getUserDataStruct(row *sql.Row, u *User) error {
	var id, role int
	var email, name, pwd, avatar, bio, createdAt, updatedAt string
	err := row.Scan(&id, &email, &name, &pwd, &avatar, &role, &bio, &createdAt, &updatedAt)
	if err != nil {
		return err
	}
	u.ID = id
	u.Name = name
	u.Email = email
	u.Pwd = ""
	u.Avatar = avatar
	u.Role = role
	u.Bio = bio
	u.CreatedAt = utils.Timestamp(createdAt)
	u.UpdatedAt = utils.Timestamp(updatedAt)
	return nil
}

func getUsersInfoFrom(rows *sql.Rows) ([]*User, error) {
	defer rows.Close()
	var users []*User
	for rows.Next() {
		var id, role int
		var email, name, pwd, avatar, bio string
		if err := rows.Scan(&id, &email, &name, &pwd, &avatar, &role, &bio); err != nil {
			return nil, err
		}
		user := &User{
			ID:     id,
			Email:  email,
			Name:   name,
			Pwd:    pwd,
			Avatar: avatar,
			Role:   role,
			Bio:    bio,
		}
		users = append(users, user)
	}
	if len(users) == 0 {
		err := errors.New("no result")
		return nil, err
	}
	return users, nil
	// distinctedUsers := distinctUsers(users)
	// return distinctedUsers, nil
}
