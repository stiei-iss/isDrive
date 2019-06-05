package init_config

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
)

type databaseConfig struct {
	IP string	`json:"ip"`
	Port string	`json:"port"`
	User string	`json:"user"`
	Pass string	`json:"password"`
}

func TestInitConfig(t *testing.T){
	RunWebServer()
}

var dbConfig databaseConfig

func initConfigHandler(w http.ResponseWriter, r *http.Request) {
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
}

func WriteConfig(config []byte) {
	cfgFile, err := os.Create("configuration.json")
	defer cfgFile.Close()

	if err != nil{
		panic(err)
	}
	cfgFile.Write(config)
}

func RunWebServer(){
	http.HandleFunc("/init", initConfigHandler)
	http.ListenAndServe(":8219", nil)
}

