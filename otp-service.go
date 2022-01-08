package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"otp/interfaces"
	"otp/routes"
	"otp/sharedinfrastructure/helper"
	"otp/sharedinfrastructure/persistence"
)

func init() {
	fileName := "log/otp-service.log"
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(f)
}

func main() {
	address, port, mode, dbhost, dbname := helper.LoadConfig()
	otp, err := persistence.ConnectDB(dbhost, dbname)
	if err != nil {
		fmt.Println(err)
	}

	otpEndPoint := interfaces.NewOtp(otp.Otp)
	fmt.Println("App running on " + address + ":" + port)
	if mode == "dev" {
		r := routes.SetupRouter(port, address, otpEndPoint)
		http.ListenAndServe(":"+port, r)
	}

}
