package filters

import (
	"github.com/emicklei/go-restful"
	"log"
)

func LogRequest(
	request *restful.Request,
	response *restful.Response,
	chain *restful.FilterChain,
) {
	log.Printf(" %8s | %s", request.Request.Method, request.SelectedRoutePath())
	chain.ProcessFilter(request, response)
}
