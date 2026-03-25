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

var dndRulesDB = []struct {
	keywords []string
	rule     string
}{
	{
		keywords: []string{"attack", "hit", "roll", "attack roll"},
		rule:     "ATTACK ROLLS: When you make an attack roll, roll a d20 and add the relevant modifier. For melee weapon attacks, use STR modifier. For ranged weapon attacks, use DEX modifier. Add your proficiency bonus if you're proficient with the weapon. A roll of 20 is a Critical Hit, dealing double damage dice.",
	},
	{
		keywords: []string{"spell", "casting", "spellcasting", "magic"},
		rule:     "SPELLCASTING: To cast a spell, you must expend a spell slot of the spell's level or higher. Your spell save DC = 8 + proficiency bonus + spellcasting ability modifier. Concentration spells require a Constitution saving throw (DC 10 or half damage taken) when damaged.",
	},
	{
		keywords: []string{"saving throw", "save", "constitution", "strength", "dexterity", "wisdom", "intelligence", "charisma"},
		rule:     "SAVING THROWS: Roll a d20 + relevant ability modifier + proficiency bonus (if proficient). Common DCs: Easy (10), Medium (15), Hard (20), Very Hard (25), Nearly Impossible (30). Advantage/Disadvantage applies before the roll.",
	},
	{
		keywords: []string{"initiative", "combat", "turn", "round", "action"},
		rule:     "COMBAT INITIATIVE: At the start of combat, roll d20 + DEX modifier. On your turn you can: Move up to your speed, take one Action (Attack, Cast Spell, Dash, Disengage, Dodge, Help, Hide, Ready, Use Object), one Bonus Action (if available), and one Free Action. Reactions can be taken outside your turn.",
	},
	{
		keywords: []string{"death", "dying", "unconscious", "death saving throw", "stabilize"},
		rule:     "DEATH SAVING THROWS: When you drop to 0 HP, you fall unconscious. On each of your turns, roll d20: 10+ is a success, 9 or less is a failure. 3 successes = stable, 3 failures = dead. Rolling 1 = 2 failures. Rolling 20 = regain 1 HP. Taking damage adds 1 failure.",
	},
	{
		keywords: []string{"skill", "check", "ability check", "proficiency"},
		rule:     "ABILITY CHECKS: Roll d20 + ability modifier + proficiency bonus (if proficient). The DM sets the DC. Common skills: Athletics (STR), Acrobatics/Stealth/Sleight of Hand (DEX), Arcana/History/Investigation/Nature/Religion (INT), Animal Handling/Insight/Medicine/Perception/Survival (WIS), Deception/Intimidation/Performance/Persuasion (CHA).",
	},
	{
		keywords: []string{"rest", "short rest", "long rest", "heal", "recovery", "hit dice"},
		rule:     "RESTING: Short Rest (1 hour): Spend Hit Dice to recover HP (roll die + CON modifier). Long Rest (8 hours): Recover all HP, recover half your total Hit Dice (min 1), regain all spell slots. You can only benefit from one long rest per 24 hours.",
	},
	{
		keywords: []string{"condition", "frightened", "paralyzed", "poisoned", "stunned", "prone", "grappled"},
		rule:     "CONDITIONS: Frightened: Disadvantage on ability checks and attacks while source is visible, can't move closer. Paralyzed: Incapacitated, auto-fail STR/DEX saves, attacks have Advantage, hits within 5ft are Critical Hits. Prone: Disadvantage on attacks, melee attacks against you have Advantage, ranged attacks have Disadvantage. Costs half movement to stand up.",
	},
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
