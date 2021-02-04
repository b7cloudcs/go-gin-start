package util

import (
	"fmt"
	"go-gin-start/ent"
	"log"

	_ "github.com/lib/pq"
)

// 全局var
var DBC *ent.Client

// DbBootstrap 初始化
func DbBootstrap() {

	// 建立链接
	client, err := ent.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			Cfg.Section("database").Key("db_host").String(),
			Cfg.Section("database").Key("db_port").String(),
			Cfg.Section("database").Key("db_user").String(),
			Cfg.Section("database").Key("db_name").String(),
			Cfg.Section("database").Key("db_password").String(),
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	// run the auto migration tool.
	if err := client.Schema.Create(Ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// sets
	DBC = client

}
