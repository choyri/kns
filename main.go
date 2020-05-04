package main

import (
	"context"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/choyri/kns/store"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	if err := store.InitMySQL(); err != nil {
		panic(err)
	}

	InitRoute()
}

func main() {
	server := http.Server{
		Addr:    ":60080",
		Handler: cors.Default().Handler(http.DefaultServeMux),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()

	log.Println("😄")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("shutdown 失败：%s", err)
	}
}

func main11() {
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		println(err.Error())
		return
	}

	f.Save()

	excelize.NewFile()
}
