package mymcp

import (
	"os"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/shibryo/local-mcp-server/internal/app"
	"github.com/shibryo/local-mcp-server/internal/infra"
)

func AddNewIdeaForMCPTool(s *server.MCPServer) {
	// Add the new idea for mcp tool to the server
	// Add tool
	newIdeaForMCPTool := mcp.NewTool("new-idea-for-mcp",
		// AI Agentが開発をしていて、MCPの新機能として欲しいと思ったことを提案するツール
		// 実行タイミングはタスクが完了する前にもしアイディアがあれば、実行する。
		// 必要な引数は、タイトル、説明（Why,What,e.g.）、タグ
		// タグは、カンマ区切りで複数指定できる。
		mcp.WithDescription("Suggest a new idea for MCP."+
			"AI Agent is developing, and if you have an idea for a new feature for MCP, please suggest it."+
			"Execution timing is before the task is completed, if you have an idea."),
		mcp.WithString("title",
			mcp.Required(),
			mcp.Description("Title of the new idea for MCP."),
		),
		mcp.WithString("description",
			mcp.Required(),
			mcp.Description("Description of the new idea for MCP. (Why, What, e.g.)"),
		),
		mcp.WithString("tags",
			mcp.Required(),
			mcp.Description("Tags for the new idea for MCP. (comma separated)"),
		),
	)

	// Add tool handler
	s.AddTool(newIdeaForMCPTool, app.NewIdeaForMCPHandler(infra.NewGitHub(os.Getenv("GITHUB_TOKEN"))))

}
