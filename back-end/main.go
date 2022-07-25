package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx context.Context = context.Background()
var redisDB *redis.Client

func applyCORS(w *http.ResponseWriter){
    (*w).Header().Set("Content-Type", "application/json")
    (*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
    (*w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Accept-Encoding, Content-Length")
}

// GET @ /stats
func stats(w http.ResponseWriter, r *http.Request){
    if r.Method != "OPTIONS"{
        return
    }else if r.Method != "GET"{
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    applyCORS(&w)
}

func main(){
    DB_URI := os.Getenv("DB_URI")
    DB_PASSWORD := os.Getenv("DB_PASSWORD")
    for{
        redisDB = redis.NewClient(&redis.Options{
            Addr: DB_URI,
            Password: DB_PASSWORD,
            DB: 0,
        })
        log.Println("Testing redis client...");
        err := redisDB.Ping(ctx).Err()
        if err != nil{
            log.Println("Redis client error:", err, "(sleeping for 30s)")
            _ = redisDB.Close()
            time.Sleep(30 * time.Second)
        }else{
            break
        }
    }

    http.Handle("/stats", http.HandlerFunc(stats))
    PORT := os.Getenv("PORT")
    if PORT == ""{
        PORT = "8080"
    }
    log.Printf("Starting server on: %s\n", PORT)
    log.Print(http.ListenAndServe(":" + PORT, nil))
}
