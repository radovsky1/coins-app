package main

import (
	"coins-app/internal/es"
	"coins-app/internal/es/kafka"
	"coins-app/internal/service"
	"coins-app/internal/service/webapi"
	"coins-app/internal/storage"
	"coins-app/internal/storage/psql"
	"coins-app/internal/transport/rest"
	"coins-app/internal/transport/rest/handler"
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := initConfig(); err != nil {
		logrus.Fatal("error initializing configs: ", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatal("error loading env variables: ", err.Error())
	}

	db, err := psql.NewPostgresDB(psql.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatal("failed to initialize db: ", err.Error())
	}

	binanceWebAPI := webapi.NewBinanceWebAPI(webapi.BinanceWebAPIConfig{
		APIKey:    os.Getenv("BINANCE_API_KEY"),
		APISecret: os.Getenv("BINANCE_API_SECRET"),
	})
	kafkaWriter := kafka.NewKafkaWriter(kafka.Config{
		Address: viper.GetString("kafka.address"),
		Topic:   viper.GetString("kafka.topic"),
	})
	msgBroker := es.NewKafkaMessageBroker(kafkaWriter)
	storages := storage.NewStoragePostgres(db)
	services := service.NewService(storages, binanceWebAPI, msgBroker)
	handlers := handler.NewHandler(services)

	srv := rest.NewServer(viper.GetString("port"), handlers.InitRoutes())
	go func() {
		if err := srv.Run(); err != nil {
			logrus.Fatal("error occurred while running http server: ", err.Error())
		}
	}()

	logrus.Print("Coins app started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Coins app shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Error("error occurred on server shutting down: ", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Error("error occurred on db connection close: ", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
