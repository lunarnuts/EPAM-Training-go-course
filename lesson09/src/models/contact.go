package models

import (
	"database/sql"
	"fmt"
	"log"
)

type Contact struct {
	ID      uint64 `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Phone   string `json:"phone" db:"phone"`
	GroupId uint64 `json:"groupId" db:"group_id"`
	Group   string `json:"group" db:"gname"`
}

func ListContacts(db *sql.DB) ([]Contact, error) {
	rows, err := db.Query(
		`select p.id, p.name, p.phone, g.name from phonebook p
		left join groups g on g.id=c.group_id`)
	if err != nil {
		return nil, err
	}
	return rowsToContacts(rows)
}

func SelectContact(db *sql.DB, id uint64) (Contact, error) {
	row := db.QueryRow(
		`select p.id, p.name, p.phone, g.name from phonebook p
		left join groups g on g.id=c.group_id where id =$1`, id)
	c := &Contact{}
	var s *string
	if err := row.Scan(&c.ID, &c.Name, &c.Phone, &s); err != nil {
		return Contact{}, fmt.Errorf("failed to fetch user (scan)")
	}
	if s == nil {
		*s = ""
	}
	c.Group = *s
	return *c, nil
}

func rowsToContacts(rows *sql.Rows) ([]Contact, error) {
	contacts := make([]Contact, 0)
	for rows.Next() {
		c := &Contact{}
		var s *string
		if err := rows.Scan(&c.ID, &c.Name, &c.Phone, &s); err != nil {
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

func InsertContact(db *sql.DB, c *Contact) error {
	var id int
	query := "insert into phonebook (name, phone, group_id) values ($1,$2,$3) returning id"
	if err := db.QueryRow(query, c.Name, c.Phone, c.GroupId).Scan(&id); err != nil {
		log.Printf("Error occured: %v", err)
		return fmt.Errorf("failed to insert into contacts")
	}
	c.ID = uint64(id)
	return nil
}

func UpdateContact(db *sql.DB, c *Contact) error {
	query := "update phonebook set name=$1, phone=$2, group_id=$3 where id=$4;"
	if _, err := db.Exec(query, c.Name, c.Phone, c.GroupId, c.ID); err != nil {
		return fmt.Errorf("failed to update contact")
	}
	return nil
}

func DeleteContact(db *sql.DB, id uint64) error {
	query := "delete from phonebook where id=$1;"
	if _, err := db.Exec(query, id); err != nil {
		return fmt.Errorf("failed to delete contact")
	}
	return nil
}

func AssignContactToGroup(db *sql.DB, id uint64, gid uint64) error {
	c, err := SelectContact(db, id)
	if err != nil {
		return fmt.Errorf("failed to fetch user (scan)")
	}
	c.GroupId = gid
	UpdateContact(db, &c)
	return nil
}
