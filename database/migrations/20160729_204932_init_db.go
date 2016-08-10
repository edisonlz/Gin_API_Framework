package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InitDb_20160729_204932 struct {
	migration.Migration
}

//bee migrate -driver='mysql' -conn='root:@tcp(127.0.0.1:3306)/go_platform'

// DO NOT MODIFY
func init() {
	m := &InitDb_20160729_204932{}
	m.Created = "20160729_204932"
	migration.Register("InitDb_20160729_204932", m)
}

// Run the migrations
func (m *InitDb_20160729_204932) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *InitDb_20160729_204932) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
