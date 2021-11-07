package main

import (
	"encoding/json"
	"fmt"
	"github.com/pion/webrtc"
	"io"
	"log"
	"net/http"
)

var answer webrtc.SessionDescription
var offer webrtc.SessionDescription
var answerR io.ReadCloser
var offerR io.ReadCloser

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./view/index.html")
}

func doSignaling(w http.ResponseWriter, r *http.Request) {

	for offer.SDP == "" {

	}
	log.Println("offer sdp ", offer.SDP)
	buf := make([]byte, 64)
	for {
		n, err := offerR.Read(buf)
		w.Write(buf[:n])
		if err != nil || n == 0 {
			w.WriteHeader(200)
			log.Println(w)
			return
		}

	}
	log.Println("cancel")
}
func Answer(w http.ResponseWriter, r *http.Request) {
	answerR = r.Body
	log.Println("11111")
	if err := json.NewDecoder(r.Body).Decode(&answer); err != nil {
		panic(err)
	}
	log.Println(answer)

}
func SDPHandler(w http.ResponseWriter, r *http.Request) {
	offerR = r.Body
	if err := json.NewDecoder(r.Body).Decode(&offer); err != nil {
		panic(err)
	}

	log.Println("sdp", offer, r.Host)

	for answer.SDP == "" {

	}
	log.Println("answer sdp ", offer.SDP)
	buf := make([]byte, 64)
	for {
		n, err := answerR.Read(buf)
		if err != nil {
			return
		}
		w.Write(buf[:n])
	}

}
func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/answer", SDPHandler)
	http.HandleFunc("/offer", doSignaling)
	http.HandleFunc("/addanswer", Answer)
	go func() {

		fmt.Println("Open http://localhost:8080 to access this demo")
		panic(http.ListenAndServe(":8080", nil))
	}()
	fmt.Println("Open http://localhost:8081 to access this demo")
	panic(http.ListenAndServe(":8081", nil))
}
