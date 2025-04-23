package main

import (
	"fmt"

	"github.com/mark3labs/mcp-go/server"
	"github.com/shibryo/local-mcp-server/internal/mymcp"
)

func main() {
	// Create MCP server
	s := server.NewMCPServer(
		"Demo 🚀",
		"1.0.0",
	)

	// Add tools
	mymcp.AddHelloTool(s)
	mymcp.AddNewIdeaForMCPTool(s)

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
