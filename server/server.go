package server

import "github.com/bhumong/dbo-test/config"

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run(config.GetString("server.port"))
}
