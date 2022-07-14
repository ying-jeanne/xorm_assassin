package team

import (
	"github.com/doug-martin/goqu"
	"github.com/jmoiron/sqlx"
	goqu_example "github.com/ying-jeanne/xorm_assassin/goqu"
	sqlx_example "github.com/ying-jeanne/xorm_assassin/sqlx"
	xorm_example "github.com/ying-jeanne/xorm_assassin/xorm"
	"xorm.io/xorm"
)

const file string = "grafana.db"

func team_service_xorm(i int, db *xorm.Engine) {
	xorm_example.Senario(i, db)
}

func team_service_goqu(i int, db *goqu.Database) {
	goqu_example.Senario(i, db)
}

func team_service_sqlx(i int, db *sqlx.DB) {
	sqlx_example.Senario(i, db)
}
