package app

import (
	"context"
	"errors"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/shibryo/local-mcp-server/internal/infra"
)

func WeatherHandler(weatherAPI infra.WeatherAPI) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		location, ok := request.Params.Arguments["location"].(string)
		if !ok {
			return nil, errors.New("location must be a string")
		}
		if len(loacation) != 6 {

		// Simulate fetching weather data
		weatherData := fmt.Sprintf("The weather in %s is sunny with a temperature of 25°C.", location)

		return mcp.NewToolResultText(weatherData), nil
	}
}
