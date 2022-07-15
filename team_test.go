package team

// benchmark go performance http://www.inanzzz.com/index.php/post/yz8n/using-golang-bench-benchstat-and-benchcmp-to-measure-performance

import (
	"testing"

	goqu_example "github.com/ying-jeanne/xorm_assassin/goqu"
	sqlx_example "github.com/ying-jeanne/xorm_assassin/sqlx"
	xorm_example "github.com/ying-jeanne/xorm_assassin/xorm"
)

const file string = ":memory:"

func BenchmarkGoquDBExecution(b *testing.B) {
	db := goqu_example.InitLib("sqlite3", file)
	goqu_example.Create_team_table_sqlite(db)
	for i := 0; i < b.N; i++ {
		team_service_goqu(i, db)
	}
}

func BenchmarkSqlxDBExecution(b *testing.B) {
	db := sqlx_example.InitLib("sqlite3", file)
	sqlx_example.Create_team_table_sqlite(db)
	for i := 0; i < b.N; i++ {
		team_service_sqlx(i, db)
	}
}

func BenchmarkXormDBExecution(b *testing.B) {
	db := xorm_example.InitLib("sqlite3", file)
	xorm_example.Create_team_table_sqlite(db)
	for i := 0; i < b.N; i++ {
		team_service_xorm(i, db)
	}
}

// func TestGoquDBExecution(t *testing.T) {
// 	// here we need to start the grafana db in postgres and mysql before, then the test would be run
// 	// against sqlite3/postgres/mysql on create/update/get/delete to test the dialect on 3 different
// 	// DBs

// 	// From Grafana root repository run: make devenv sources=postgres_tests,mysql_tests
// 	db := goqu_example.InitLib("sqlite3", file)
// 	team_service_goqu(0, db)

// 	// For trouble shooting, connect to postgres with psql
// 	// psql -h localhost -p 5432 -U grafanatest -d grafanatest
// 	db = goqu_example.InitLib("postgres", "user=grafanatest password=grafanatest host=localhost port=5432 dbname=grafanatest sslmode=disable")
// 	goqu_example.Create_team_table_postgres(db)
// 	team_service_goqu(0, db)

// 	// TODO: the mysql test has some problem
// 	db = goqu_example.InitLib("mysql", "grafana:password@tcp(localhost:3306)/grafana_tests?collation=utf8mb4_unicode_ci?parseTime=true")
// 	goqu_example.Create_team_table_mysql(db)
// 	team_service_goqu(0, db)
// }

// func TestSQLXDBExecution(t *testing.T) {
// 	// here we need to start the grafana db in postgres and mysql before, then the test would be run
// 	// against sqlite3/postgres/mysql on create/update/get/delete to test the dialect on 3 different
// 	// DBs

// 	// From Grafana root repository run: make devenv sources=postgres_tests,mysql_tests
// 	db := sqlx_example.InitLib("sqlite3", file)
// 	team_service_sqlx(0, db)

// 	// For trouble shooting, connect to postgres with psql
// 	// psql -h localhost -p 5432 -U grafanatest -d grafanatest
// 	db = sqlx_example.InitLib("postgres", "user=grafanatest password=grafanatest host=localhost port=5432 dbname=grafanatest sslmode=disable")
// 	sqlx_example.Create_team_table_postgres(db)
// 	team_service_sqlx(0, db)

// 	// TODO: the mysql test has some problem
// 	db = sqlx_example.InitLib("mysql", "grafana:password@tcp(localhost:3306)/grafana_tests?collation=utf8mb4_unicode_ci?parseTime=true")
// 	sqlx_example.Create_team_table_mysql(db)
// 	team_service_sqlx(0, db)
// }

// func TestXormDBExecution(t *testing.T) {
// 	// here we need to start the grafana db in postgres and mysql before, then the test would be run
// 	// against sqlite3/postgres/mysql on create/update/get/delete to test the dialect on 3 different
// 	// DBs

// 	// From Grafana root repository run: make devenv sources=postgres_tests,mysql_tests
// 	db := xorm_example.InitLib("sqlite3", file)
// 	team_service_xorm(0, db)

// 	// For trouble shooting, connect to postgres with psql
// 	// psql -h localhost -p 5432 -U grafanatest -d grafanatest
// 	db = xorm_example.InitLib("postgres", "user=grafanatest password=grafanatest host=localhost port=5432 dbname=grafanatest sslmode=disable")
// 	xorm_example.Create_team_table_postgres(db)
// 	team_service_xorm(0, db)

// 	// TODO: the mysql test has some problem
// 	db = xorm_example.InitLib("mysql", "grafana:password@tcp(localhost:3306)/grafana_tests?collation=utf8mb4_unicode_ci")
// 	xorm_example.Create_team_table_mysql(db)
// 	team_service_xorm(0, db)
// }
