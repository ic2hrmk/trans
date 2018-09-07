package web

import "net/http"

func (wds *WebDashboardServer) videoStreamHandler(w http.ResponseWriter, r *http.Request) {
	wds.videoStream.ServeHTTP(w, r)
}
