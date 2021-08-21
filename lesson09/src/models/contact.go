package models

import (
	"database/sql"
	"fmt"
)

type Contact struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Phone string `db:"phone"`
	Group string `db:"group"`
}

func ContactList(db *sql.DB) ([]Contact, error) {
	rows, err := db.Query(
		`select c.id, c.name, c.phone, g.name as group from contacts c
		left join groups g on c.group_id=c.id`)
	if err != nil {
		return nil, err
	}
	return rowsToContacts(rows)
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

func ContactSave(db *sql.DB, c *Contact) error {
	if c.ID == 0 {
		return insertContact(db, c)
	}
	return updateContact(db, c)
}

func insertContact(db *sql.DB, c *Contact) error {
	var id int64
	query := "insert into contacts (name, phone) values ($1,$2) returning id"
	if err := db.QueryRow(query, c.Name, c.Phone).Scan(&id); err != nil {
		return fmt.Errorf("failed to insert into contacts")
	}
	c.ID = int(id)
	return nil
}

func updateContact(db *sql.DB, c *Contact) error {
	query := "update contacts set name=$1, phone=$2 where id=$3;"
	if _, err := db.Exec(query, c.Name, c.Phone, c.ID); err != nil {
		return fmt.Errorf("failed to update contact")
	}
	return nil
}

func assignContactToGroup() {}
