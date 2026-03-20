package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// SSE service for streaming bridge events from PostgreSQL to frontend clients
func main() {
	// DATABASE_URL is expected in format:
	// postgres://user:password@host:port/dbname?sslmode=disable
	dbURL := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	defer db.Close()

	// NOTE:
	// sql.Open does NOT establish a real connection immediately.
	// Consider db.Ping() here in production to fail fast.

	r := gin.Default()

	// CORS configuration (frontend access control)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Endpoint to initiate monitoring (currently just logs intent)
	r.POST("/api/bridge/deposit", func(c *gin.Context) {
		var req struct {
			TxHash string `json:"tx_hash"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "invalid request"})
			return
		}

		// This does NOT enqueue anything - only logs
		log.Printf("Monitoring requested for tx: %s", req.TxHash)

		c.JSON(200, gin.H{
			"status":  "monitoring_started",
			"tx_hash": req.TxHash,
		})
	})

	// SSE endpoint: streams status updates for a given transaction
	r.GET("/api/bridge/events/:txHash", func(c *gin.Context) {
		rawTxHash := c.Param("txHash")

		// Normalize hash to match DB format (assumes DB stores with 0x prefix)
		searchHash := "0x" + rawTxHash

		// SSE required headers
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")

		// Disable buffering (important for nginx / reverse proxies)
		c.Writer.Header().Set("X-Accel-Buffering", "no")

		// Polling interval (acts as pseudo-stream)
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()

		log.Printf("SSE: started stream for tx: %s (searching as: %s)", rawTxHash, searchHash)

		for {
			select {
			// Client disconnected - stop loop
			case <-c.Request.Context().Done():
				log.Printf("SSE: client disconnected for tx %s", rawTxHash)
				return

			case <-ticker.C:
				var status string

				// Query latest status from DB
				err := db.QueryRow(
					"SELECT status FROM bridge_events WHERE tx_hash = $1",
					searchHash,
				).Scan(&status)

				if err != nil {
					if err == sql.ErrNoRows {
						// No record yet - emit default "detected" state
						fmt.Fprintf(c.Writer,
							"data: {\"status\":\"detected\",\"tx_hash\":\"%s\"}\n\n",
							rawTxHash,
						)
						c.Writer.Flush()
						continue
					}

					// Any other DB error - terminate stream
					log.Printf("DB Error: %v", err)
					return
				}

				// Send SSE event (JSON payload)
				fmt.Fprintf(c.Writer,
					"data: {\"status\":\"%s\", \"relay_tx\":\"%s\"}\n\n",
					status,
					rawTxHash,
				)
				c.Writer.Flush()

				// Terminal states - close stream
				if status == "relayed" || status == "failed" {
					log.Printf("SSE: finished for %s with status %s", rawTxHash, status)
					return
				}
			}
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on :%s", port)

	r.Run(":" + port)
}
