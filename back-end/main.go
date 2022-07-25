package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

type statsReqBody struct{
    Code string `json:"code"`
    Stat int64 `json:"stat"`
    // stat: 0 - artist
    //       1 - songs
    //       2 - both
}

type tokenResp struct{
    Access_token string `json:"access_token"`
    Token_type string `json:"access_token"`
    Expires_in int64 `json:"access_token"`
    Refresh_token string `json:"access_token"`
}

var ctx context.Context = context.Background()
var redisDB *redis.Client
var CLIENT_SECRET string

var SPOTIFY_TOKEN_URL string = "https://accounts.spotify.com/api/token"
var CLIENT_ID string = "f5940d4d679948c5a33bfce4ad03ac50"
var REDIRECT_URI string

func applyCORS(w *http.ResponseWriter){
    (*w).Header().Set("Content-Type", "application/json")
    (*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
    (*w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Accept-Encoding, Content-Length")
}

// POST @ /stats
func stats(w http.ResponseWriter, r *http.Request){
    applyCORS(&w)
    if r.Method == "OPTIONS"{
        return
    }else if r.Method != "POST"{
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    // parse request body
    defer r.Body.Close()
    reqBody, err := ioutil.ReadAll(r.Body)
    if err != nil{
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    // unmarshal req body
    var reqBodyJson statsReqBody
    err = json.Unmarshal([]byte(reqBody), &reqBodyJson)
    if err != nil{
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    if reqBodyJson.Stat < 0 || reqBodyJson.Stat > 3{
        w.WriteHeader(http.StatusBadRequest)
        return
    }


    // generate params to be url encoded
    params := url.Values{}
    params.Add("grant_type", "authorization_code")
    params.Add("code", string(reqBodyJson.Code))
    params.Add("redirect_uri", REDIRECT_URI)
    log.Print(REDIRECT_URI)

    // exchange code for token
    var encrpytedToken = base64.StdEncoding.EncodeToString([]byte(CLIENT_ID + ":" + CLIENT_SECRET))
    req, err := http.NewRequest("POST", SPOTIFY_TOKEN_URL, strings.NewReader(params.Encode()))
    req.Header.Add("Authorization", "Basic " + encrpytedToken)
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    var spotifyRes tokenResp
    err = json.Unmarshal(body, &spotifyRes)
}

func main(){
    DB_URI := os.Getenv("DB_URI")
    DB_PASSWORD := os.Getenv("DB_PASSWORD")
    CLIENT_SECRET = os.Getenv("CLIENT_SECRET")
    if CLIENT_SECRET == ""{
        log.Println("CLIENT_SECERT env var has to be set")
        return
    }

    REDIRECT_URI = os.Getenv("REDIRECT_URI")
    if REDIRECT_URI == ""{
        REDIRECT_URI = "http://localhost:5173/stats"
    }

    // connecting to redis
    for{
        redisDB = redis.NewClient(&redis.Options{
            Addr: DB_URI,
            Password: DB_PASSWORD,
            DB: 0,
        })

        // testing redis client
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
