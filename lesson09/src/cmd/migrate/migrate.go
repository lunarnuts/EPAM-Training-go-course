package migrate

import (
	"fmt"

	_ "github.com/lib/pq"

	"github.com/lunarnuts/go-course/tree/lesson09/src/db"
	_ "github.com/lunarnuts/go-course/tree/lesson09/src/db/migrations"

	"github.com/wshaman/migrate"

	DB "github.com/wshaman/course-db/src/db/db"
)

func MigrateUp(setup db.DBSetup) error {
	dbObj, err := DB.New(DB.DBSetup(setup))
	if err != nil {
		return fmt.Errorf("Migrate: Failed to connect to DB: %v", err)
	}
	defer dbObj.Close()
	err = migrate.Up(dbObj)
	if err != nil {
		return fmt.Errorf("Migration up unsuccessful: %v", err)
	}
	return nil
}

func MigrateDown(setup db.DBSetup) error {
	dbObj, err := DB.New(DB.DBSetup(setup))
	if err != nil {
		return fmt.Errorf("Migrate: Failed to connect to DB: %v", err)
	}
	err = migrate.Down(dbObj)
	if err != nil {
		return fmt.Errorf("Migration down unsuccessful: %v", err)
	}
	return nil
}
