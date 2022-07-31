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
)

type statsReqBody struct{
    Code string `json:"code"`
    Stat int64 `json:"stat"`
}

type tokenResp struct{
    AccessToken string `json:"access_token"`
    TokenType string `json:"token_type"`
    ExpiresIn int64 `json:"expires_in"`
    RefreshToken string `json:"refresh_token"`
}

type artist struct{
    External_urls struct{
        Spotify string              `json:"spotify"`
    }                               `json:"external_urls"`
    Followers struct{
        //Href string                 `json:"href"`
        Total uint64                `json:"total"`
    }                               `json:"followers"`
    Genres []string                 `json:"genres"`
    //Id string                       `json:"id"`
    Images []struct{
        Height uint32               `json:"height"`
        Width uint32                `json:"width"`
        Url string                  `json:"url"`
    }                               `json:"images"`
    Name string                     `json:"name"`
    Popularity uint64               `json:"popularity"`
}

type track struct{
    Album struct{
        // Album_type string
        Artists []struct{
            External_urls struct{
                Spotify string      `json:"spotify"`
            }                       `json:"external_urls"`
            // Href string
            // Id string
            Name string             `json:"name"`
            // Type string
            // Uri string
        }                           `json:"artists"`
        Release_date string         `json:"release_date"`
        Images []struct{
            Height uint32           `json:"height"`
            Width uint32            `json:"width"`
            Url string              `json:"url"`
        }                           `json:"images"`
    }                               `json:"album"`
    // Available_markets []string
    External_urls struct{
        Spotify string              `json:"spotify"`
    }                               `json:"external_urls"`
    // Href string
    // Id string
    Name string                     `json:"name"`
    Popularity uint64               `json:"popularity"`
}

type statsRespArtists struct{
    //Href string `json:"href"`
    Items []artist `json:"items"`
    //Limit uint64 `json:"limit"`
    //Next string `json:"next"`
    //Offset uint64 `json:"offset"`
    //Previous string `json:"previous"`
    //Total uint64 `json:"total"`
}

type statsRespTrack struct{
    //Href string `json:"href"`
    Items []track `json:"items"`
    //Limit uint64 `json:"limit"`
    //Next string `json:"next"`
    //Offset uint64 `json:"offset"`
    //Previous string `json:"previous"`
    //Total uint64 `json:"total"`
}

type statsRespJSON struct{
    Artists []artist    `json:"artists"`
    Tracks []track      `json:"tracks"`
}

var ctx context.Context = context.Background()
var CLIENT_SECRET string

var SPOTIFY_TOKEN_URL string = "https://accounts.spotify.com/api/token"
var SPOTIFY_BASE_URL string = "https://api.spotify.com/v1/me/top/"
var CLIENT_ID string
var REDIRECT_URI string

func applyCORS(w *http.ResponseWriter){
    (*w).Header().Set("Content-Type", "application/json")
    (*w).Header().Set("Access-Control-Allow-Origin", "https://spotify-stats-gray.vercel.app")
    // (*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") // dev
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Accept-Encoding, Content-Length")
}

func getArtists(token string)(statsRespArtists, error){
    req, err := http.NewRequest("GET", SPOTIFY_BASE_URL + "artists?limit=10&time_range=long_term", nil)
    if err != nil {
        return statsRespArtists{}, err
    }
    req.Header.Add("Authorization", "Bearer " + token)
    req.Header.Add("Content-Type", "application/json")

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return statsRespArtists{}, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200{
        return statsRespArtists{}, errors.New("response from spotify: " + resp.Status)
    }

    readBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return statsRespArtists{}, err
    }

    var respJson statsRespArtists
    err = json.Unmarshal([]byte(readBody), &respJson)
    if err != nil {
        return statsRespArtists{}, err
    }

    // only three genres
    for i, artist := range respJson.Items{
        if len(artist.Genres) > 3 {
            respJson.Items[i].Genres = artist.Genres[0:3]
        }
    }

    return respJson, nil
}

func getSongs(token string)(statsRespTrack, error){
    req, err := http.NewRequest("GET", SPOTIFY_BASE_URL + "tracks?limit=15&time_range=long_term", nil)
    if err != nil {
        return statsRespTrack{}, nil
    }
    req.Header.Add("Authorization", "Bearer " + token)
    req.Header.Add("Content-Type", "application/json")

    resp, err := http.DefaultClient.Do(req)
    if err != nil{
        return statsRespTrack{}, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200{
        return statsRespTrack{}, errors.New("response from spotify: " + resp.Status)
    }

    readBody, err := ioutil.ReadAll(resp.Body)
    if err != nil{
        return statsRespTrack{}, err
    }

    var respJson statsRespTrack
    err = json.Unmarshal([]byte(readBody), &respJson)
    if err != nil{
        return statsRespTrack{}, err
    }

    // only three artists
    for i, track := range respJson.Items{
        if len(track.Album.Artists) > 3 {
            respJson.Items[i].Album.Artists = track.Album.Artists[0:3]
        }
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

    if resp.StatusCode != 200{
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    var spotifyRes tokenResp
    err = json.Unmarshal(body, &spotifyRes)
    if err != nil{
        log.Println("Error unmarshalling token:", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    artists, err := getArtists(spotifyRes.AccessToken)
    if err != nil{
        log.Println("Error getting artists:", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    tracks, err := getSongs(spotifyRes.AccessToken)
    if err != nil{
        log.Println("Error getting tracks:", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(statsRespJSON{Tracks: tracks.Items, Artists: artists.Items})
}

func main(){
    CLIENT_SECRET = os.Getenv("CLIENT_SECRET")
    if CLIENT_SECRET == ""{
        log.Println("CLIENT_SECRET env var has to be set")
        return
    }

    REDIRECT_URI = os.Getenv("REDIRECT_URI")
    if REDIRECT_URI == ""{
        REDIRECT_URI = "https://spotify-stats-gray.vercel.app/stats"
        // REDIRECT_URI = "http://localhost:5173/stats" // dev
    }

    CLIENT_ID = os.Getenv("CLIENT_ID")
    if CLIENT_ID == ""{
        log.Println("CLIENT_ID env var has to be set")
        return
    }

    http.Handle("/stats", http.HandlerFunc(stats))
    PORT := os.Getenv("PORT")
    if PORT == ""{
        PORT = "8080"
    }
    log.Printf("Starting server on: %s\n", PORT)
    log.Print(http.ListenAndServe(":" + PORT, nil))
}
