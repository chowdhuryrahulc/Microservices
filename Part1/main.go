package video1

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


// curl localhost:8080 -d "Nic": works same as hitting the endpoint with a http request from browser or postman
// 		-d "Nic" means we send data as string named "Nic"

func main(){
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request){
		log.Println("Hello World")				// logs everything
		d, err := ioutil.ReadAll(r.Body)
		if err != nil{
			//! Case 1 (long process)
			// rw.WriteHeader: helps us write the respnse code (eg: 200)
			// http pkg: has all the status codes defined as constants
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Ooops"))	// Writing error msg

			//! Case 2 (short process, all 2 satz in 1 line, either case1 or case2)
			http.Error(rw, "Ooops", http.StatusBadGateway)

			return
		}
		fmt.Fprintf(rw, "Hello %s", d)
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye World")
	})

	// you can name port with a pirticular ip address like localhost:8080 or 127.1.0.0.8080. 
	// Or you can open port for every ip address. Like :8080
	http.ListenAndServe(":8080", nil)	// nil is the hanler
}