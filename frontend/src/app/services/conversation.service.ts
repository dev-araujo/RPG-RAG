import { Injectable, signal, computed } from "@angular/core";

export interface ChatMessage {
  role: "master" | "lorekeeper";
  content: string;
  timestamp: Date;
  context?: {
    rulesUsed: number;
    relevanceScore: number;
    source: string;
    retrievedRules: string[];
  };
  pipeline?: string;
}

export interface Conversation {
  id: string;
  name: string;
  messages: ChatMessage[];
  createdAt: Date;
  updatedAt: Date;
}

const STORAGE_KEY = "grimoire_conversations";
const ACTIVE_KEY = "grimoire_active_id";

function welcomeMessage(): ChatMessage {
  return {
    role: "lorekeeper",
    content:
      "Bem-vindo, Mestre! Sou o LoreKeeper, seu oráculo de regras de D&D 5e. Pergunte-me qualquer coisa sobre as regras e eu consultarei os tomos para você.",
    timestamp: new Date(),
  };
}

@Injectable({ providedIn: "root" })
export class ConversationService {
  conversations = signal<Conversation[]>([]);
  activeId = signal<string>("");

  activeConversation = computed(() =>
    this.conversations().find((c) => c.id === this.activeId())
  );

  constructor() {
    this.load();
    if (this.conversations().length === 0) {
      this.createConversation();
    }
  }

  private load(): void {
    try {
      const raw = localStorage.getItem(STORAGE_KEY);
      if (raw) {
        const data = JSON.parse(raw) as any[];
        const parsed: Conversation[] = data.map((c) => ({
          ...c,
          createdAt: new Date(c.createdAt),
          updatedAt: new Date(c.updatedAt),
          messages: (c.messages as any[]).map((m) => ({
            ...m,
            timestamp: new Date(m.timestamp),
          })),
        }));
        this.conversations.set(parsed);
      }
      const activeId = localStorage.getItem(ACTIVE_KEY);
      if (activeId && this.conversations().some((c) => c.id === activeId)) {
        this.activeId.set(activeId);
      } else if (this.conversations().length > 0) {
        this.activeId.set(this.conversations()[0].id);
      }
    } catch {
    }
  }

  private persist(): void {
    try {
      localStorage.setItem(STORAGE_KEY, JSON.stringify(this.conversations()));
      localStorage.setItem(ACTIVE_KEY, this.activeId());
    } catch {}
  }

  createConversation(): void {
    const id = `conv-${Date.now()}`;
    const conv: Conversation = {
      id,
      name: "New Scroll",
      messages: [welcomeMessage()],
      createdAt: new Date(),
      updatedAt: new Date(),
    };
    this.conversations.update((list) => [conv, ...list]);
    this.activeId.set(id);
    this.persist();
  }

  selectConversation(id: string): void {
    this.activeId.set(id);
    localStorage.setItem(ACTIVE_KEY, id);
  }

  renameConversation(id: string, name: string): void {
    const trimmed = name.trim();
    if (!trimmed) return;
    this.conversations.update((list) =>
      list.map((c) => (c.id === id ? { ...c, name: trimmed } : c))
    );
    this.persist();
  }

  deleteConversation(id: string): void {
    const remaining = this.conversations().filter((c) => c.id !== id);
    this.conversations.set(remaining);
    if (this.activeId() === id) {
      if (remaining.length > 0) {
        this.activeId.set(remaining[0].id);
      } else {
        this.createConversation();
        return;
      }
    }
    this.persist();
  }

  saveMessages(messages: ChatMessage[]): void {
    const id = this.activeId();
    this.conversations.update((list) =>
      list.map((c) =>
        c.id === id ? { ...c, messages, updatedAt: new Date() } : c
      )
    );
    this.persist();
  }

  autoName(firstQuestion: string): void {
    const id = this.activeId();
    const conv = this.conversations().find((c) => c.id === id);
    if (conv?.name === "New Scroll") {
      const name =
        firstQuestion.length > 32
          ? firstQuestion.slice(0, 32) + "…"
          : firstQuestion;
      this.renameConversation(id, name);
    }
  }
}
