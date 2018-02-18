package main

import(
	"fmt"
	"net/http"
	"log"
	"os/exec"
	"strconv"

	"github.com/gorilla/mux"
)

var timeoutSec int

func Getping(w http.ResponseWriter, r *http.Request) {
	dest := r.FormValue("dest")
	fmt.Println("destionation : ", dest)
	fmt.Fprint(w, "test ping destionation : ")
	w.Write([]byte(dest))

	exec.Command("ping", "-c", "1", "-w", strconv.Itoa(timeoutSec), dest)
}

func main() {

	r := mux.NewRouter()
	// List of Routes
	r.HandleFunc("/ping", Getping).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", "8080"), r))

}
