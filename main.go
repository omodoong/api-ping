package main

import(
	"fmt"
	"net/http"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

var timeoutSec int
var message    string
var logFile *os.File

func Getping(w http.ResponseWriter, r *http.Request) {
	dest := r.FormValue("dest")
	fmt.Println("destionation : ", dest)
	fmt.Fprintln(w, "Report Connection...")
	fmt.Fprint(w, "destionation : ")
	w.Write([]byte(dest))

	out, _ := exec.Command("ping", dest, "-c 3", "-i 3", "-w 3").Output()
	
	if strings.Contains(string(out), "Destination Host Unreachable") {
    	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "Status No Connected")
    	fmt.Fprintln(w, "")
    	fmt.Fprintln(w, "Status Not Connected")
		} else {
    		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "Status Connected")
    		fmt.Fprintln(w, "")
    		fmt.Fprintln(w, "Status Connected")

	}
	fmt.Fprintln(w, "time :", time.Now().Format("Mon Jan 2 15:04:05 MST 2006"))		
}

func main() {

	r := mux.NewRouter()
	// List of Routes
	r.HandleFunc("/ping", Getping).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", "8080"), r))

}
