package models

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/lunarnuts/go-course/tree/lesson09/src/db"
)

type Contact struct {
	ID      uint64 `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Phone   string `json:"phone" db:"phone"`
	GroupId uint64 `json:"groupId" db:"group_id"`
	Group   string `json:"group" db:"gname"`
}

func ListContacts(db db.DBConn) ([]Contact, error) {
	rows, err := db.Query(context.Background(),
		`select p.id, p.name, p.phone, p.group_id, g.gname from phonebook p
		left join groups g on g.group_id=p.group_id`)
	if err != nil {
		return nil, err
	}
	return rowsToContacts(rows)
}

func SelectContact(db db.DBConn, id uint64) (Contact, error) {
	row := db.QueryRow(context.Background(),
		`select p.id, p.name, p.phone, g.name from phonebook p
		left join groups g on g.id=c.group_id where id =$1`, id)
	c := &Contact{}
	var s *string
	if err := row.Scan(&c.ID, &c.Name, &c.Phone, &c.GroupId, &s); err != nil {
		return Contact{}, fmt.Errorf("failed to fetch user (scan)")
	}
	if s == nil {
		*s = ""
	}
	c.Group = *s
	return *c, nil
}

func rowsToContacts(rows pgx.Rows) ([]Contact, error) {
	contacts := make([]Contact, 0)
	for rows.Next() {
		c := &Contact{}
		var s *string
		if err := rows.Scan(&c.ID, &c.Name, &c.Phone, &c.GroupId, &s); err != nil {
			return nil, fmt.Errorf("failed to fetch users (scan)")
		}
		if s == nil {
			*s = ""
		}
		c.Group = *s
		contacts = append(contacts, *c)
	}
	return contacts, nil
}

func InsertContact(db db.DBConn, c *Contact) error {
	var id int
	query := "insert into phonebook (name, phone, group_id) values ($1,$2,$3) returning id"
	if err := db.QueryRow(context.Background(), query, c.Name, c.Phone, c.GroupId).Scan(&id); err != nil {
		log.Printf("Error occured: %v", err)
		return fmt.Errorf("failed to insert into contacts")
	}
	c.ID = uint64(id)
	return nil
}

func UpdateContact(db db.DBConn, c *Contact) error {
	query := "update phonebook set name=$1, phone=$2, group_id=$3 where id=$4;"
	if _, err := db.Exec(context.Background(), query, c.Name, c.Phone, c.GroupId, c.ID); err != nil {
		return fmt.Errorf("failed to update contact")
	}
	return nil
}

func DeleteContact(db db.DBConn, id uint64) error {
	query := "delete from phonebook where id=$1;"
	if _, err := db.Exec(context.Background(), query, id); err != nil {
		return fmt.Errorf("failed to delete contact")
	}
	return nil
}

func AssignContactToGroup(db db.DBConn, id uint64, gid uint64) error {
	c, err := SelectContact(db, id)
	if err != nil {
		return fmt.Errorf("failed to fetch user (scan)")
	}
	c.GroupId = gid
	UpdateContact(db, &c)
	return nil
}
