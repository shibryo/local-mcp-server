package app

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/shibryo/local-mcp-server/internal/infra"
)

const (
	GitHub_Owner = "shibryo"
	GitHub_Repo  = "local-mcp-server"
)

func NewIdeaForMCPHandler(github *infra.GitHub) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		title, ok := request.Params.Arguments["title"].(string)
		if !ok {
			return nil, errors.New("title must be a string")
		}
		description, ok := request.Params.Arguments["description"].(string)
		if !ok {
			return nil, errors.New("description must be a string")
		}
		tags, ok := request.Params.Arguments["tags"].(string)
		if !ok {
			return nil, errors.New("tags must be a string")
		}
		tagsArray := []string{}
		// Split tags by comma
		for _, tag := range strings.Split(tags, ",") {
			if trimmed := strings.TrimSpace(tag); trimmed != "" {
				tagsArray = append(tagsArray, trimmed)
			}
		}

		// Create issue on GitHub
		err := github.CreateIssue(GitHub_Owner, GitHub_Repo, title, description, &tagsArray)
		if err != nil {
			return nil, fmt.Errorf("failed to create issue: %w", err)
		}

		return mcp.NewToolResultText("New idea for MCP has been created successfully."), nil
	}
}
