package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
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
    AccessToken string `json:"access_token"`
    TokenType string `json:"token_type"`
    ExpiresIn int64 `json:"expires_in"`
    RefreshToken string `json:"refresh_token"`
}

type statsResp struct{
    Href string `json:"href"`
    Items []string `json:"items"`
    Limit uint64 `json:"limit"`
    Next string `json:"next"`
    Offset uint64 `json:"offset"`
    Previous string `json:"previous"`
    Total uint64 `json:"total"`

}

var ctx context.Context = context.Background()
var redisDB *redis.Client
var CLIENT_SECRET string

var SPOTIFY_TOKEN_URL string = "https://accounts.spotify.com/api/token"
var SPOTIFY_BASE_URL string = "https://api.spotify.com/v1/me/top/"
var CLIENT_ID string
var REDIRECT_URI string

func applyCORS(w *http.ResponseWriter){
    (*w).Header().Set("Content-Type", "application/json")
    (*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:5713 https://api.spotify.com")
    (*w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Accept-Encoding, Content-Length")
}

func getArtists(token string)(statsResp, error){
    req, err := http.NewRequest("GET", SPOTIFY_BASE_URL + "artists?limit=20&time_range=long_term", nil)
    req.Header.Add("Authorization", "Bearer " + token)
    req.Header.Add("Content-Type", "application/json")
    if err != nil {
        return statsResp{}, err
    }

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return statsResp{}, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200{
        return statsResp{}, errors.New("resp not 200")
    }

    readBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return statsResp{}, err
    }

    var respJson statsResp
    err = json.Unmarshal([]byte(readBody), &respJson)
    if err != nil {
        return statsResp{}, err
    }

    return respJson, nil
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

    // exchange code for token
    var encrpytedToken = base64.StdEncoding.EncodeToString([]byte(CLIENT_ID + ":" + CLIENT_SECRET))
    req, err := http.NewRequest("POST", SPOTIFY_TOKEN_URL, strings.NewReader(params.Encode()))
    req.Header.Add("Authorization", "Basic " + encrpytedToken)
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
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
    if err != nil{
        log.Println("Error unmarshaling token:", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    artists, err := getArtists(spotifyRes.AccessToken)
    if err != nil{
        log.Println("Error getting artists:", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(artists)
}

func main(){
    DB_URI := os.Getenv("DB_URI")
    DB_PASSWORD := os.Getenv("DB_PASSWORD")
    CLIENT_SECRET = os.Getenv("CLIENT_SECRET")
    if CLIENT_SECRET == ""{
        log.Println("CLIENT_SECRET env var has to be set")
        return
    }

    REDIRECT_URI = os.Getenv("REDIRECT_URI")
    if REDIRECT_URI == ""{
        REDIRECT_URI = "https://spotify-stats-gray.vercel.app/stats"
    }

    CLIENT_ID = os.Getenv("CLIENT_ID")
    if CLIENT_ID == ""{
        log.Println("CLIENT_ID env var has to be set")
        return
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
