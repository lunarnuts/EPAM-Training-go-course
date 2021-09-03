package migrations

import "github.com/wshaman/migrate"

func init() {
	migrate.RegisterSQL(20210901183528,
		"lunarnuts <n.kuandyk1995@gmail.com>",
		"insert_values_to_groups",
		`-- insert group names
		INSERT INTO groups(gname) 
			VALUES ('Employee'),('Customers'),('Managers');`,
		`-- delete group names
		DELETE FROM groups WHERE gname in ('Employee','Customers','Managers')`,
	)
}
