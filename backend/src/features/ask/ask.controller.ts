import { Request, Response } from "express";
import { ask } from "./ask.service";

export async function askController(req: Request, res: Response) {
  const { question, campaignId = "default-campaign" } = req.body;

  if (!question || typeof question !== "string") {
    return res
      .status(400)
      .json({ error: 'Field "question" is required and must be a string.' });
  }

  console.log(`\n[Backend] POST /api/ask — question: "${question}"`);

  try {
    const result = await ask({ question, campaignId });
    return res.json(result);
  } catch (error: any) {
    console.error("[Backend] Error processing /api/ask:", error.message);

    if (error.code === 14) {
      return res.status(503).json({
        error:
          "RAG Engine unavailable. Ensure the Go gRPC server is running on port 50051.",
        details: error.message,
      });
    }

    return res
      .status(500)
      .json({ error: "Internal server error", details: error.message });
  }
}
