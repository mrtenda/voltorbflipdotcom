package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"

	"goji.io"
	"goji.io/pat"

	"github.com/mrtenda/voltorbflipdotcom/server/voltorbflip"
	"github.com/zenazn/goji/graceful"
	"os"
)

const (
	staticContentPathKey = "VFLIP_STATIC_CONTENT_PATH"
)

type SolveApiRequest struct {
	Tiles       voltorbflip.VfPSolBoard
	BoardTotals voltorbflip.VfBoardTotals
}

type SolveApiResponse struct {
	Tiles          [5][5]voltorbflip.VfPSolTile
	IsPossible     bool
	IsWon          bool
	SafestPosition voltorbflip.VfBoardPosition
	Safety         float32
}

type SolveApiHandler struct{}

func (_ SolveApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request SolveApiRequest
	requestBytes, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(requestBytes, &request)

	isPossible, isWon, tiles, safestPosition, safety := voltorbflip.Solve(&request.BoardTotals, request.Tiles)

	apiResponse := SolveApiResponse{
		IsPossible:     isPossible,
		IsWon:          isWon,
		Tiles:          tiles,
		SafestPosition: safestPosition,
		Safety:         safety}
	b, _ := json.Marshal(apiResponse)

	w.Write(b)
}

func main() {
	mux := goji.NewMux()

	mux.Handle(pat.Post("/api/solve"), http.TimeoutHandler(SolveApiHandler{}, 15*time.Second, "timed out"))

	var staticContentPath string
	if os.Getenv(staticContentPathKey) == "" {
		staticContentPath, _ = filepath.Abs("./jekyll-site/_site")
	} else {
		staticContentPath, _ = filepath.Abs(os.Getenv(staticContentPathKey))
	}

	mux.Handle(pat.Get("/*"), http.FileServer(http.Dir(staticContentPath)))

	server := &graceful.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 15,
	}
	server.ListenAndServe()
}
