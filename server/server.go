package server

import (
	"astro/auto"
	"astro/config"
	"astro/router"
	"fmt"
	"log"
	"net/http"
)

func Run() {
	config.Load()
	fmt.Println("config file loaded")
	auto.Load()
	fmt.Println("DB loaded")
	// fmt.Printf("\n\tListening.......[::]:%d \n", config.PORT)
	Listen(8081)
}

func Listen(port int) {
	r := router.New()
	fmt.Println("server started")
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), r)
	if err != nil {
		fmt.Println(err)
		fmt.Println("server stopping ........")
		log.Fatal("error is : ", err)
	}
	
}
