import * as grpc from "@grpc/grpc-js";
import * as protoLoader from "@grpc/proto-loader";
import path from "path";

const PROTO_PATH = path.join(__dirname, "../../../../../grpc/proto/rag_service.proto");
const RAG_ENGINE_ADDRESS = process.env.RAG_ENGINE_ADDRESS || "localhost:50051";

const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
});

const protoDescriptor = grpc.loadPackageDefinition(packageDefinition) as any;
const lorekeeper = protoDescriptor.lorekeeper;

let client: any = null;

export function getGrpcClient() {
  if (!client) {
    client = new lorekeeper.LoreKeeperRAG(
      RAG_ENGINE_ADDRESS,
      grpc.credentials.createInsecure(),
    );
    console.log(
      `[gRPC Client] Connected to RAG Engine at ${RAG_ENGINE_ADDRESS}`,
    );
  }
  return client;
}

export interface RuleQueryRequest {
  question: string;
  campaign_id: string;
}

export interface RuleContextResponse {
  rules: string[];
  source: string;
  relevance_score: number;
}

export function searchRules(
  query: RuleQueryRequest,
): Promise<RuleContextResponse> {
  return new Promise((resolve, reject) => {
    const client = getGrpcClient();

    console.log(`[gRPC Client] Calling SearchRules with: "${query.question}"`);

    client.SearchRules(
      query,
      (error: grpc.ServiceError | null, response: RuleContextResponse) => {
        if (error) {
          console.error("[gRPC Client] SearchRules error:", error.message);
          reject(error);
          return;
        }
        resolve(response);
      },
    );
  });
}
