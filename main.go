package main

import(
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/sparrc/go-ping"
)

type Category struct {
    Rtts []string		`json: "rtts_latency"`
    Avg  string			`json: "average_latency"`
    Max string			`json: "maximum_latency"`
    Min  string			`json: "average_latency"`
}

var timeoutSec int
var message    string
var logFile *os.File

func Getping(w http.ResponseWriter, r *http.Request) {
	dest := r.FormValue("dest")
	pinger, err := ping.NewPinger(dest)
	pinger.SetPrivileged(true)
	pinger.Timeout = 3 * time.Second
	if err != nil {
	    w.Write([]byte(err.Error()))
	}
	pinger.Count = 3
	pinger.Run() // blocks until finished
	stats := pinger.Statistics() // get send/receive/rtt stats
	fmt.Println(stats)

	avg := stats.AvgRtt
	Savg := fmt.Sprint(avg)
    max := stats.MaxRtt
    Smax := fmt.Sprint(max)
    min := stats.MinRtt
    Smin := fmt.Sprint(min)
    pack := stats.PacketLoss


    // masukkan rtts dalam bentuk []time.duration ke rtts dalam bentuk []string
    rttsString := []string{}
    for _, rtt := range(stats.Rtts) {
    	rttsString = append(rttsString, rtt.String())
    }

    // Create an instance of the Box struct.
    category := Category{
    	Rtts: rttsString,
        Avg: Savg,
        Max: Smax,
        Min: Smin,
    }
    // Create JSON from the instance data.
    // ... Ignore errors.
    b, _ := json.MarshalIndent(category, "", "    ")
    // Convert bytes to string.
    // s := string(b)
    // fmt.Println(s)
	w.Write(b)	

}

func main() {
	r := mux.NewRouter()
	// List of Routes
	r.HandleFunc("/ping", Getping).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", "8080"), r))
}
