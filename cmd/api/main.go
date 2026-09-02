package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/saasybyte/saasy-edge/db"
	"github.com/saasybyte/saasy-edge/db/sqlc"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "9090"
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	if dbHost == "" || dbPort == "" || dbName == "" || dbUser == "" || dbPassword == "" {
		log.Fatal("DB_HOST, DB_PORT, DB_NAME, DB_USER, DB_PASSWORD are required")
	}

	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	pool, err := db.Connect(context.Background(), databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	log.Println("connected to database")

	queries := sqlc.New(pool)
	handlers := InitializeHandlers(queries)

	r := gin.Default()

	corsOrigin := os.Getenv("WEB_URL")
	if corsOrigin == "" {
		corsOrigin = "http://localhost:4173"
	}

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{corsOrigin},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	routerGroup := r.Group("/api/v1")

	routerGroup.GET("/openapi", func(c *gin.Context) {
		c.File("api/openapi.yaml")
	})

	RegisterRoutes(routerGroup, handlers)

	// Start gRPC server
	go func() {
		listener, err := net.Listen("tcp", ":"+grpcPort)
		if err != nil {
			log.Fatalf("failed to listen on grpc port: %v", err)
		}

		grpcServer := grpc.NewServer()
		RegisterGRPCServices(grpcServer, handlers)
		reflection.Register(grpcServer) // TODO: disable during prod

		log.Printf("starting grpc server on :%s", grpcPort)
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve grpc: %v", err)
		}
	}()

	log.Printf("starting http server on :%s", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
