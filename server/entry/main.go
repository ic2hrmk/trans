package entry

import (
	"log"
	"net/http"

	"trans/server/webapi"
	_ "trans/server/controller"
)



func main() {
	server := &http.Server{Addr: ":8080", Handler: webapi.InitRestContainer()}
	log.Fatal(server.ListenAndServe())
}
