package xorm_example

import (
	"log"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

type Team struct {
	ID        int       `xorm:"'id' pk autoincr"`
	Name      string    `xorm:"name"`
	OrgID     int       `xorm:"org_id"`
	CreatedAt time.Time `xorm:"'created'"`
	UpdatedAt time.Time `xorm:"'updated'"`
	Email     string
}

func insertTeam(e *xorm.Engine, team1 Team) error {
	// for insert, xorm is actually checking that all the filled field has a corresponding column name, if not, error out
	// not set field would be filled directly with default value when creation
	_, err := e.Insert(&team1)
	return err
}

func getTeam(e *xorm.Engine, name string) (bool, Team, error) {
	var team1 Team
	found, err := e.Where("name=?", name).Get(&team1)
	return found, team1, err
}

func updateTeam(e *xorm.Engine, team Team) error {
	_, err := e.ID(team.ID).Update(team)
	return err
}

func deleteTeam(e *xorm.Engine, name string) error {
	_, err := e.Exec("DELETE FROM team WHERE name=?", name)
	return err
}

func InitLib(driver, file string) *xorm.Engine {
	// set engine of sqlite3 here
	engine, err := xorm.NewEngine(driver, file)
	if err != nil {
		log.Fatal(err)
	}
	engine.SetTableMapper(names.GonicMapper{})
	return engine
}

func Create_team_table_postgres(db *xorm.Engine) {
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

func Create_team_table_mysql(db *xorm.Engine) {
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

func Senario(i int, engine *xorm.Engine) {
	num := strconv.Itoa(i)
	team1 := Team{Name: "mynamee" + num, OrgID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	err := insertTeam(engine, team1)
	// fmt.Printf("the inserted team %+v\n", team1)
	if err != nil {
		log.Fatal(err)
	}

	_, team3, err := getTeam(engine, team1.Name)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("the get team %+v\n", team3)

	// here it is very confusing, since xorm omit the default value OrgID 0 here, we would have to
	// force the update by calling .AllCols().Update or .Cols("org_id").Update
	// but if we put .AllCols() we have to put all the fields that are mandatory so it is also not convinient
	team2 := Team{ID: team3.ID, OrgID: 0, Name: "princess" + num}
	err = updateTeam(engine, team2)
	if err != nil {
		log.Fatal(err)
	}

	_, team4, err := getTeam(engine, team2.Name)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("the updated team %+v\n", team4)

	err = deleteTeam(engine, team4.Name)
	if err != nil {
		log.Fatal(err)
	}

	_, _, err = getTeam(engine, team4.Name)
	if err != nil {
		log.Fatal(err)
	}
}
