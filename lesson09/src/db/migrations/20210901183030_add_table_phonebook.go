package migrations

import "github.com/wshaman/migrate"

func init() {
	migrate.RegisterSQL(20210901183030,
		"lunarnuts <n.kuandyk1995@gmail.com>",
		"add_table_phonebook",
		`-- create table phonebook
		CREATE TABLE groups (
			group_id SERIAL PRIMARY KEY,
			gname VARCHAR(64)
		);
		CREATE TABLE phonebook (
			id SERIAL PRIMARY KEY,
			name VARCHAR(64),
			phone VARCHAR(64),
			group_id int,
			FOREIGN KEY (group_id) REFERENCES groups(group_id)
		);`,
		`-- drop tables
		DROP TABLE IF EXISTS phonebook;
		DROP TABLE IF EXISTS groups;`,
	)
}
