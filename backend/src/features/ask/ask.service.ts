import { GoogleGenerativeAI } from "@google/generative-ai";
import { searchRules } from "../../shared/grpc/grpcClient";

export interface AskInput {
  question: string;
  campaignId: string;
}

export interface AskResult {
  answer: string;
  context: {
    rulesUsed: number;
    source: string;
    relevanceScore: number;
    retrievedRules: string[];
  };
  pipeline: string;
}

export async function ask({ question, campaignId }: AskInput): Promise<AskResult> {
  const ruleContext = await searchRules({ question, campaign_id: campaignId });

  const systemPrompt = `You are the Dungeon Master's assistant for D&D 5e.
Use ONLY the following official rules as your knowledge base to answer questions.
Be concise, precise, and reference specific rules when answering.

Retrieved Rules Context:
${ruleContext.rules.map((r, i) => `[Rule ${i + 1}] ${r}`).join("\n\n")}

Source: ${ruleContext.source}
Relevance Score: ${(ruleContext.relevance_score * 100).toFixed(0)}%`;

  const genAI = new GoogleGenerativeAI(process.env.GEMINI_API_KEY!);
  const model = genAI.getGenerativeModel({ model: "gemini-2.5-flash-lite" });
  const result = await model.generateContent([systemPrompt, question]);
  const answer = result.response.text();

  return {
    answer,
    context: {
      rulesUsed: ruleContext.rules.length,
      source: ruleContext.source,
      relevanceScore: ruleContext.relevance_score,
      retrievedRules: ruleContext.rules,
    },
    pipeline: "Angular → Backend (Node.js) → gRPC → RAG Engine (Go) → Gemini",
  };
}
