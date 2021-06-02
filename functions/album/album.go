package album

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aofiee/album/utils"
	"github.com/joho/godotenv"
)

var errEnv = godotenv.Load(".env")

func sendJSON(v interface{}, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(&v)
	if err != nil {
		log.Println(err)
	}
	w.Write(jsonData)
}
func SearchAlbums(w http.ResponseWriter, r *http.Request) {
	if errEnv != nil {
		panic(errEnv)
	}
	var params struct {
		Artist string `json:"artist_name"`
		Track  string `json:"track"`
	}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		log.Fatal(err)
		return
	}
	if params.Artist == "" && params.Track == "" {
		var msg struct {
			Msg string `json:"msg"`
		}
		msg.Msg = "please input params."
		sendJSON(&msg, w, r)
		return
	}
	data := getJSON("https://api.discogs.com/database/search?artist=" + params.Artist + "&track=" + params.Track)
	var result utils.AlbumsJson
	json.Unmarshal(data, &result)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(showData(&result, w, r))
	if err != nil {
		log.Println(err)
	}
	w.Write(jsonData)
}
func showData(result *utils.AlbumsJson, w http.ResponseWriter, r *http.Request) []utils.Results {
	var rs []utils.Results
	for _, v := range result.Results {
		t := trackList(v.ResourceURL)
		rs = append(rs, utils.Results{
			Thumb:   v.Thumb,
			Artist:  v.Title,
			Country: v.Country,
			Year:    v.Year,
			Genres:  strings.Join(v.Genre, ","),
			Tracks:  t,
		})
	}
	return rs
}
func trackList(url string) []utils.TrackAlbum {
	var t []utils.TrackAlbum
	data := getJSON(url)
	var result utils.TrackJson
	json.Unmarshal(data, &result)
	for _, v := range result.Tracklist {
		t = append(t, utils.TrackAlbum{
			TrackName: v.Title,
		})
	}
	return t
}
func noRedirect(req *http.Request, via []*http.Request) error {
	return errors.New("Don't redirect!")
}
func getJSON(link string) []byte {
	client := &http.Client{
		CheckRedirect: noRedirect,
	}
	discogs := "Discogs key=" + os.Getenv("KEY") + ", secret=" + os.Getenv("SECRET")
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		log.Println("err", req)
	}
	req.Header.Add("Authorization", discogs)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Println("error ", err)
	}
	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		log.Println("error", err)
	} else {
		return body
	}
	return []byte{}
}
