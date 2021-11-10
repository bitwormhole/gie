package vo

import "github.com/bitwormhole/gie/server/web/dto"

type AgentLink struct {
	BaseVO
	Agent dto.AgentLink
	// Message     dto.AgentMessage
	MessageList []*dto.AgentMessage
}
