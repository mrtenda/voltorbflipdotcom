package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mrtenda/voltorbflipdotcom/server/voltorbflip"
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

func HandleRequest(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	// Only handle POST to /api/solve (or any path routed here)
	var solveReq SolveApiRequest
	err := json.Unmarshal([]byte(request.Body), &solveReq)
	if err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Invalid JSON request body",
		}, nil
	}

	isPossible, isWon, tiles, safestPosition, safety := voltorbflip.Solve(&solveReq.BoardTotals, solveReq.Tiles)

	apiResponse := SolveApiResponse{
		IsPossible:     isPossible,
		IsWon:          isWon,
		Tiles:          tiles,
		SafestPosition: safestPosition,
		Safety:         safety,
	}

	responseBytes, _ := json.Marshal(apiResponse)

	return events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusOK,
		Body:       string(responseBytes),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
