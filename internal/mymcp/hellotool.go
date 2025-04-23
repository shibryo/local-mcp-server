package mymcp

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/shibryo/local-mcp-server/internal/app"
)

func AddHelloTool(s *server.MCPServer) {
	// Add the hello tool to the server
	// Add tool
	helloTool := mcp.NewTool("hello-world",
		mcp.WithDescription("Say hello to someone"),
		mcp.WithString("name",
			mcp.Required(),
		),
	)

	// Add tool handler
	s.AddTool(helloTool, app.HelleHandler())
}
