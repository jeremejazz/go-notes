# A simple Fileserver



The full code

```go
package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))

}

func main() {
	var dir string
	port := flag.String("port", "3000", "port to serve HTTP on")
	path := flag.String("path", "", "path to serve")
	flag.Parse()

	if *path == "" {
		dir, _ = os.Getwd()
	} else {
		dir = *path
	}
	log.Println("Attempting to listen port:", *port)
	err := http.ListenAndServe(":"+*port, http.FileServer(http.Dir(dir)))
	log.Fatal(err)
}



```
source file: [fs.go](code/fs/fs.go)