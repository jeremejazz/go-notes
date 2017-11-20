# A simple Fileserver

We will be creating a simple static file server. This can be useful for hosting static files.


## The basic web server setup

We'll import the packages we need for our webserver

```go
package main 

import (
    "net/http"
    "os"
)

func main(){
    dir,_ := os.Getwd()
    http.ListenAndServe(":3000", http.FileServer(http.Dir(dir)))

}
```

Running this code and going to http://localhost:8080 will give you the current directory listing


## Accepting Command-line Arguments

To make things flexible we can add some parameters in case we need to change the port and path.
Go comes with a library that parses command-line parameter flags. 


The full code: 

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