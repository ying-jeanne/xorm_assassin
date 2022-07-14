package team

// benchmark go performance http://www.inanzzz.com/index.php/post/yz8n/using-golang-bench-benchstat-and-benchcmp-to-measure-performance

import (
	"testing"

	goqu_example "github.com/ying-jeanne/xorm_assassin/goqu"
	sqlx_example "github.com/ying-jeanne/xorm_assassin/sqlx"
	xorm_example "github.com/ying-jeanne/xorm_assassin/xorm"
)

func BenchmarkGoquDBExecution(b *testing.B) {
	db := goqu_example.InitLib(file)
	for i := 0; i < b.N; i++ {
		team_service_goqu(i, db)
	}
}

func BenchmarkSqlxDBExecution(b *testing.B) {
	db := sqlx_example.InitLib(file)
	for i := 0; i < b.N; i++ {
		team_service_sqlx(i, db)
	}
}

func BenchmarkXormDBExecution(b *testing.B) {
	db := xorm_example.InitLib(file)
	for i := 0; i < b.N; i++ {
		team_service_xorm(i, db)
	}
}
