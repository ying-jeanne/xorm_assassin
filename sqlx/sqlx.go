package sqlx_example

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Team struct {
	ID        int
	Name      string         `db:"name"`
	OrgID     int            `db:"org_id"`
	CreatedAt time.Time      `db:"created"`
	UpdatedAt time.Time      `db:"updated"`
	Email     sql.NullString `db:"email"`
}

func insertTeam(db *sqlx.DB, team Team) error {
	// this is a transaction, a transaction should start with MustBegin, then end by commit
	// Inside of grafana code, instead of put session into the context, we can put the transaction into the context
	// tx := db.MustBegin()
	// tx.MustExec("INSERT INTO team (name, org_id, created, updated, email) VALUES ($1, $2, $3, $4, $5)", "wangyxxx", 0, time.Now(), time.Now(), "w.x@gmail.com")
	// named exec allows the user to insert into table with a struct object
	_, err := db.NamedExec("INSERT INTO team (name, org_id, created, updated, email) VALUES (:name, :org_id, :created, :updated, :email)", team)
	// tx.Commit()
	// here we would insert manually the fields instead of go struct, it is how sqlx is working today.
	// from the same project tho, there is a helper which helps us to build go struct https://github.com/jmoiron/modl/blob/master/modl_test.go
	// it could insert go struct instead of field by field
	return err
}

func getTeam(db *sqlx.DB, name string) (bool, Team, error) {
	// teams := []Team{}
	// err := db.Select(&teams, "SELECT * FROM team ORDER BY name ASC")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// here to cast the type into go struct directly we still need the tag for the fields that are not having the same name
	// but the error is threw correctly when it could not found the corresponding field, so it is still better than xorm?
	// for _, team := range teams {
	// 	fmt.Printf("%#v\n", team)
	// }

	// get one single result
	team1 := Team{}
	found := true
	err := db.Get(&team1, db.Rebind("SELECT * FROM team WHERE name=?"), name)
	if err == sql.ErrNoRows {
		found = false
		err = nil
	}
	return found, team1, err
	// for the field that could be nullable, we need to explicitely set it to type sql.NullString, otherwise we will get error when the field is not found
}

func updateTeam(db *sqlx.DB, team Team) error {
	_, err := db.NamedExec(`UPDATE team SET name=:name WHERE id =:id`, team)
	return err
}

func deleteTeam(db *sqlx.DB, name string) error {
	// the db.Rebind is to translate the ? to different presentation in different database type
	_, err := db.Exec(db.Rebind("DELETE FROM team WHERE name=?"), name)
	return err
	// if err != nil {
	// return 0, err
	// }
	// ids, err := rows.SliceScan()
	// if len(ids) > 0 {
	// 	return ids[0].(int), err
	// }
	// return 0, err
}

func InitLib(driver, file string) *sqlx.DB {
	// set engine of sqlite3 here
	db, err := sqlx.Connect(driver, file)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Create_team_table_postgres(db *sqlx.DB) {
	schema := `
		DROP TABLE IF EXISTS "team";
		CREATE TABLE "team" (
			"id" SERIAL PRIMARY KEY NOT NULL,
			"name" VARCHAR(190) NOT NULL,
			"org_id" BIGINT NOT NULL,
			"created" TIMESTAMP NOT NULL DEFAULT now(),
			"updated" TIMESTAMP NOT NULL DEFAULT now(),
			"email" VARCHAR(190)
		);
	`
	if _, err := db.Exec(schema); err != nil {
		panic(err)
	}
}

func Create_team_table_mysql(db *sqlx.DB) {
	dropTable := "DROP TABLE IF EXISTS `team`;"
	createTable := "CREATE  TABLE `team` (" +
		"`id` INT NOT NULL AUTO_INCREMENT ," +
		"`name` VARCHAR(190) NOT NULL," +
		"`org_id` BIGINT NOT NULL ," +
		"`created` DATETIME NOT NULL ," +
		"`updated` DATETIME NOT NULL ," +
		"`email` VARCHAR(190)," +
		"PRIMARY KEY (`id`) );"

	if _, err := db.Exec(dropTable); err != nil {
		panic(err)
	}
	if _, err := db.Exec(createTable); err != nil {
		panic(err)
	}
}

func Senario(i int, db *sqlx.DB) {
	num := strconv.Itoa(i)
	team1 := Team{Name: "mynamee" + num, OrgID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	err := insertTeam(db, team1)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("the inserted team %+v\n", team1)
	_, team3, err := getTeam(db, team1.Name)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("the get team %+v\n", team3)
	team2 := Team{ID: team3.ID, OrgID: 0, Name: "princess" + num}
	err = updateTeam(db, team2)
	if err != nil {
		log.Fatal(err)
	}

	_, team4, err := getTeam(db, team2.Name)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("the updated team %+v\n", team4)

	err = deleteTeam(db, team4.Name)
	if err != nil {
		log.Fatal(err)
	}

	_, _, err = getTeam(db, team4.Name)
	if err != nil {
		log.Fatal(err)
	}
}
