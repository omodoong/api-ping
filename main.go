package main

import(
	"fmt"
	"net/http"
	"log"
	"os/exec"
	"strings"

	"github.com/gorilla/mux"
)

var timeoutSec int

func Getping(w http.ResponseWriter, r *http.Request) {
	dest := r.FormValue("dest")
	fmt.Println("destionation : ", dest)
	fmt.Fprintln(w, "Report Connection...")
	fmt.Fprint(w, "destionation : ")
	w.Write([]byte(dest))

	out, _ := exec.Command("ping", dest, "-c 3", "-i 3", "-w 10").Output()
	if strings.Contains(string(out), "Destination Host Unreachable") {
    	fmt.Println("Unreachable")
    	fmt.Fprintln(w, "   (Unreachable)")
		} else {
    		fmt.Println("Status Connected")
    		fmt.Fprintln(w, "  (Status Conencted)")
	}		
}

func main() {

	r := mux.NewRouter()
	// List of Routes
	r.HandleFunc("/ping", Getping).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", "8080"), r))

}
