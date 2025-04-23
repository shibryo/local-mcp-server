package mymcp

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func AddWeatherTool(s *server.MCPServer) {
	// Add the weather tool to the server
	weatherTool := mcp.NewTool("weather",
		mcp.WithDescription("Get the weather for a location"),
		mcp.WithString("location:(e.g. 久留米の ID (400040))",
			mcp.Required(),
		),
	)

	// Add tool handler
	s.AddTool(weatherTool, app.WeatherHandler())
}
