package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sorame/common"
	"sorame/database"
	"sorame/common/middleware"
	"sorame/handler"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const defaultPort = "3000"

func main() {
	// Ensure the "logs" directory exists; create it if it doesn't
	logDir := "logs"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.Mkdir(logDir, 0755)
		if err != nil {
			log.Fatalf("Failed to create log directory: %v", err)
		}
	}

	// Create a new log file with a timestamp in its name
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	logFileName := filepath.Join(logDir, fmt.Sprintf("sorame_service_%s.log", timestamp))
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	// Set up multi-writer for logging to both terminal and log file
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)

	log.Println(`
	███████╗ ██████╗ ██████╗  █████╗ ███╗   ███╗███████╗
	██╔════╝██╔═══██╗██╔══██╗██╔══██╗████╗ ████║██╔════╝
	███████╗██║   ██║██████╔╝███████║██╔████╔██║█████╗  
	╚════██║██║   ██║██╔══██╗██╔══██║██║╚██╔╝██║██╔══╝  
	███████║╚██████╔╝██║  ██║██║  ██║██║ ╚═╝ ██║███████╗
	╚══════╝ ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝
	`)

	// Load environment variables
	common.LoadEnv()

	// Initialize Redis
	redisClient := database.InitRedis()
	handler.SetLinkRepo(redisClient)

	// Initialize the router
	router := mux.NewRouter()

	// Apply the logging middleware
	router.Use(middleware.Logging)
	router.Use(middleware.JsonLogging)

	// API insert link
	router.HandleFunc("/api/v1/link", handler.InsertLink).Methods("POST")

	// API get link
	router.HandleFunc("/api/v1/link/{shareID}", handler.GetLink).Methods("GET")

	// Status endpoint for health checks
	router.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodHead {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodHead)

	// Serve static files (HTML, CSS, JS)
	router.PathPrefix("/").Handler(handler.ServeStatic("./static"))

	// Custom NotFoundHandler to serve 404.html page
	router.NotFoundHandler = http.HandlerFunc(handler.NotFoundHandler)

	// Set up CORS middleware with desired options
	c := cors.New(cors.Options{
		// AllowedOrigins defines which origins are permitted to access the resources.
		// "*" allows all origins; in production, specify the exact origins.
		AllowedOrigins: []string{"*"},
		// AllowCredentials indicates whether the request can include user credentials like cookies.
		AllowCredentials: true,
		// Add all necessary headers for CORS
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{
			"Authorization", "Content-Type", "Accept",
			"Origin", "User-Agent", "Cache-Control",
		},
		// Allow browsers to cache preflight requests for 1 hour (in seconds)
		MaxAge: 3600,
		// Expose these headers to the browser
		ExposedHeaders: []string{"Content-Length", "Content-Type", "X-Auth-Token"},
		// Debug mode for development
		Debug: false,
	})

	// Wrap the router with the CORS middleware
	handler := c.Handler(router)

	// Get the server port from environment variables or use the default
	port := os.Getenv("SORAME_SERVICE_PORT")
	if port == "" {
		port = defaultPort
	}

	// Start the server
	log.Println("🚀 Sorame service is running on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
