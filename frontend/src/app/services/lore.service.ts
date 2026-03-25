import { Injectable, inject } from "@angular/core";
import { HttpClient } from "@angular/common/http";
import { Observable } from "rxjs";

export interface AskRequest {
  question: string;
  campaignId: string;
}

export interface RuleContext {
  rulesUsed: number;
  source: string;
  relevanceScore: number;
  retrievedRules: string[];
}

export interface AskResponse {
  answer: string;
  context: RuleContext;
  pipeline: string;
}

@Injectable({ providedIn: "root" })
export class LoreService {
  private apiUrl = "/api";
  private http = inject(HttpClient);

  ask(request: AskRequest): Observable<AskResponse> {
    return this.http.post<AskResponse>(`${this.apiUrl}/ask`, request);
  }
}
