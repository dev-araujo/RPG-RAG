package server

import (
	"context"
	"log"
	"strings"

	pb "github.com/lorekeeper/grpc/pb"
)

type RagServer struct {
	pb.UnimplementedLoreKeeperRAGServer
}

func (s *RagServer) SearchRules(ctx context.Context, req *pb.RuleQuery) (*pb.RuleContext, error) {
	log.Printf("[RAG Engine] Received query: %q (campaign: %s)", req.Question, req.CampaignId)

	questionLower := strings.ToLower(req.Question)
	var matchedRules []string
	var bestScore float32 = 0.0

	for _, entry := range dndRulesDB {
		score := float32(0)
		for _, kw := range entry.keywords {
			if strings.Contains(questionLower, kw) {
				score += 1.0 / float32(len(entry.keywords))
			}
		}
		if score > 0 {
			matchedRules = append(matchedRules, entry.rule)
			if score > bestScore {
				bestScore = score
			}
		}
	}

	if len(matchedRules) == 0 {
		matchedRules = []string{
			"D&D 5e CORE RULE: The most important rule is the Golden Rule — the DM can override any rule to serve the fun of the table. When in doubt, the DM makes the call.",
			"D&D 5e STRUCTURE: The game is structured around Ability Scores (STR, DEX, CON, INT, WIS, CHA), Proficiency Bonus (scales with level), and the d20 roll as the resolution mechanic for most actions.",
		}
		bestScore = 0.3
	}

	log.Printf("[RAG Engine] Found %d relevant rules (score: %.2f)", len(matchedRules), bestScore)

	return &pb.RuleContext{
		Rules:          matchedRules,
		Source:         "D&D 5e SRD (Systems Reference Document) - Mock Vector Store",
		RelevanceScore: bestScore,
	}, nil
}
