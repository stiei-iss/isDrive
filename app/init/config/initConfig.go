package config

import (
	"encoding/json"
	"fmt"
	"github.com/stiei-iss/isDrive/database/conn"
	"net/http"
	"os"
)

type databaseConfig struct {
	IP string	`json:"ip"`
	Port string	`json:"port"`
	User string	`json:"user"`
	Pass string	`json:"password"`
}


var dbConfig databaseConfig

func InitConfigHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("------")
	dbConfig = databaseConfig{
		r.FormValue("database-ip"),
		r.FormValue("database-port"),
		r.FormValue("database-user"),
		r.FormValue("database-passwd"),
		}
	configData, _ := json.MarshalIndent(dbConfig,"","\t")
	WriteConfig(configData)

	fmt.Printf("%s\n", string(configData))
	add := "mongodb://"+dbConfig.IP+":"+dbConfig.Port
	fmt.Println(add)
	conn.Connect(add)
}

func WriteConfig(config []byte) {
	cfgFile, err := os.Create("configuration.json")
	defer cfgFile.Close()

	if err != nil{
		panic(err)
	}
	cfgFile.Write(config)
}


