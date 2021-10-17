package server

import (
	"astro/auto"
	"astro/config"
	"astro/router"
	"fmt"
	"log"
	"net/http"
	"os"
)

func Run() {
	config.Load()
	fmt.Println("config file loaded")
	auto.Load()
	fmt.Println("DB loaded")
	port := os.Getenv("PORT") 
	fmt.Printf("\n\tListening.......[::]:%s \n", port)
	Listen(port)
}

func Listen(port string) {
	r := router.New()
	fmt.Println("server started")
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), r)
	if err != nil {
		fmt.Println(err)
		fmt.Println("server stopping ........")
		log.Fatal("error is : ", err)
	}
	
}
