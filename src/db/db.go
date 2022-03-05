package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type LastMigration struct {
	version string
}

func MigrateDB(db *sql.DB, databaseMigrationsPath string) {
	if databaseMigrationsPath == "" {
		databaseMigrationsPath = "./db/migrations"
	}

	files, err := ioutil.ReadDir(databaseMigrationsPath)
	sort.Slice(files, func(i, z int) bool {
		splitI := strings.Split(files[i].Name(), "-")
		iInts := convertVersionToInts(splitI[0])
		splitZ := strings.Split(files[z].Name(), "-")
		zInts := convertVersionToInts(splitZ[0])

		return doesNeedApplyMigration(iInts, zInts)
	})

	if err != nil {
		panic(err)
	}

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS " +
		`migrate_history("id" SERIAL PRIMARY KEY,` +
		`"version" varchar, comment varchar, created_at TIMESTAMP without time zone DEFAULT now())`); err != nil {
		panic(err)
	}

	res, _ := db.Query("SELECT version FROM migrate_history ORDER BY id DESC LIMIT 1")
	defer res.Close()
	lastMigration := &LastMigration{}
	if res.Next() == false {
		lastMigration.version = "0.0.0"
	} else if err = res.Scan(&lastMigration.version); err != nil {
		fmt.Println(lastMigration)
		panic(err)
	}

	lastMigrationParts := convertVersionToInts(lastMigration.version)

	for _, file := range files {
		if file.Name() == "main.go" {
			continue
		}

		split := strings.Split(file.Name(), "-")
		migrationVersion, migrationComment := convertVersionToInts(split[0]), split[1]
		migrationVersionString := split[0]
		if doesNeedApplyMigration(lastMigrationParts, migrationVersion) {
			var sqlFiles []os.FileInfo
			sqlFiles, err = ioutil.ReadDir(databaseMigrationsPath + "/" + migrationVersionString + "-" + migrationComment)
			if err != nil {
				panic(err)
			}

			fmt.Println("Migrate to " + migrationVersionString + "...")
			for _, sqlFile := range sqlFiles {
				sqlFileName := sqlFile.Name()
				if _, err := db.Exec(stringifySQL(databaseMigrationsPath + "/" + migrationVersionString + "-" + migrationComment + "/" + sqlFileName)); err != nil {
					panic(err)
				}
			}

			if _, err := db.Exec("INSERT INTO migrate_history (version, comment) VALUES ($1, $2)", migrationVersionString, migrationComment); err != nil {
				panic(err)
			}
		}
	}
	fmt.Println("Migration completed")
}

func convertVersionToInts(version string) [3]int {
	var err error
	var ints [3]int
	versionParts := strings.Split(version, ".")

	if len(versionParts) != 3 {
		return ints
	}
	for i := 0; i < 3; i++ {
		ints[i], err = strconv.Atoi(strings.TrimSpace(versionParts[i]))
		if err != nil {
			panic(err)
		}
	}

	return ints
}

func doesNeedApplyMigration(lastMigration [3]int, checkingMigration [3]int) bool {
	if checkingMigration[0] > lastMigration[0] {
		return true
	}

	if checkingMigration[0] == lastMigration[0] {
		if checkingMigration[1] > lastMigration[1] {
			return true
		}
		if checkingMigration[1] == lastMigration[1] {
			if checkingMigration[2] > lastMigration[2] {
				return true
			}
		}
	}

	return false
}

func stringifySQL(pathToSql string) string {
	sql, err := ioutil.ReadFile(pathToSql)

	if err == nil {
		return string(sql)
	}

	return err.Error()
}
