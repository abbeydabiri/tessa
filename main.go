package main

//package TESSA

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"tessa/api"
	"tessa/blockchain"
	"tessa/config"
	"tessa/utils"
)

func main() {

	// flag.StringVar(&blockchain.ETHAddress, "to", "", "Debug flag forces no cache")
	// flag.Float64Var(&blockchain.ETHAmount, "amt", 0.0, "Debug flag forces no cache")
	// flag.Parse()

	utils.Logger("")
	config.Init(nil) //Init Config.yaml

	go blockchain.EthClientDial("")
	api.StartRouter()
	//
}

//Start ...
func Start(TIMEZONE, VERSION, COOKIE, DBPATH, OS, OSPATH, ADDRESS string) {
	//OS e.g "ios" or "android"
	//PATH e.g "/sdcard/com.sample.app/"
	var yaml = []byte(fmt.Sprintf(`timezone: %v
version: %v
cookie: %v
db: %v
os: %v
path: %v
address: %v
encryption_keys:
  public: public.pem
  private: private.pem
`, TIMEZONE, VERSION, COOKIE, DBPATH, OS, OSPATH, ADDRESS))

	utils.Logger(OSPATH)
	config.Init(yaml) //Init Config.yaml
	go api.StartRouter()
}

//IsRunning ...
func Status(url string) bool {
	client := &http.Client{Timeout: time.Duration(time.Second)}
	if _, err := client.Get(url); err == nil {
		return true
	}
	return false
}

//Stop ...
func Stop() {
	sMessage := "stopping service @ " + config.Get().Address
	println(sMessage)
	log.Println(sMessage)
	os.Exit(1)
}
