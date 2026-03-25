import {
  Component,
  OnInit,
  ViewChild,
  ElementRef,
  AfterViewChecked,
  inject,
  signal,
} from "@angular/core";
import { LoreService, AskResponse } from "../../services/lore.service";
import { NgClass, SlicePipe, DatePipe } from "@angular/common";
import { FormsModule } from "@angular/forms";

interface ChatMessage {
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

@Component({
  selector: "app-master-panel",
  templateUrl: "./master-panel.component.html",
  styleUrls: ["./master-panel.component.scss"],
  imports: [FormsModule, DatePipe],
})
export class MasterPanelComponent implements OnInit, AfterViewChecked {
  @ViewChild("chatScroll") chatScroll!: ElementRef;

  campaignId = `campaign-${Date.now()}`;
  question = signal("");
  messages = signal<ChatMessage[]>([]);
  isAskingRules = signal(false);
  statusText = signal("Sistema pronto");

  private loreService = inject(LoreService);

  ngOnInit(): void {
    this.messages.update((msgs) => [
      ...msgs,
      {
        role: "lorekeeper",
        content:
          "Bem-vindo, Mestre! Sou o LoreKeeper, seu oráculo de regras de D&D 5e. Pergunte-me qualquer coisa sobre as regras e eu consultarei os tomos para você.",
        timestamp: new Date(),
      },
    ]);
  }

  ngAfterViewChecked(): void {
    this.scrollToBottom();
  }

  askQuestion(): void {
    const q = this.question().trim();
    if (!q || this.isAskingRules()) return;

    this.question.set("");

    this.messages.update((msgs) => [
      ...msgs,
      {
        role: "master",
        content: q,
        timestamp: new Date(),
      },
    ]);

    this.isAskingRules.set(true);
    this.statusText.set("Consultando oráculo...");

    this.loreService
      .ask({ question: q, campaignId: this.campaignId })
      .subscribe({
        next: (response: AskResponse) => {
          this.messages.update((msgs) => [
            ...msgs,
            {
              role: "lorekeeper",
              content: response.answer,
              timestamp: new Date(),
              context: response.context,
              pipeline: response.pipeline,
            },
          ]);
          this.isAskingRules.set(false);
          this.statusText.set("Sistema pronto");
        },
        error: (err) => {
          this.messages.update((msgs) => [
            ...msgs,
            {
              role: "lorekeeper",
              content: `Erro de conexão: ${err.error?.error || err.message}. Verifique os serviços.`,
              timestamp: new Date(),
            },
          ]);
          this.isAskingRules.set(false);
          this.statusText.set("Erro no sistema");
        },
      });
  }

  onKeydown(event: KeyboardEvent): void {
    if (event.key === "Enter" && !event.shiftKey) {
      event.preventDefault();
      this.askQuestion();
    }
  }

  updateQuestion(value: string): void {
    this.question.set(value);
  }

  private scrollToBottom(): void {
    try {
      this.chatScroll.nativeElement.scrollTop =
        this.chatScroll.nativeElement.scrollHeight;
    } catch {}
  }
}
