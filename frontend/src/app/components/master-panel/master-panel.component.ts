import {
  Component,
  OnInit,
  ViewChild,
  ElementRef,
  AfterViewChecked,
  inject,
  signal,
  effect,
  untracked,
} from "@angular/core";
import { LoreService, AskResponse } from "../../services/lore.service";
import { ConversationService, ChatMessage } from "../../services/conversation.service";
import { DatePipe } from "@angular/common";
import { FormsModule } from "@angular/forms";

@Component({
  selector: "app-master-panel",
  templateUrl: "./master-panel.component.html",
  styleUrls: ["./master-panel.component.scss"],
  imports: [FormsModule, DatePipe],
})
export class MasterPanelComponent implements AfterViewChecked {
  @ViewChild("chatScroll") chatScroll!: ElementRef;

  question = signal("");
  messages = signal<ChatMessage[]>([]);
  isAskingRules = signal(false);
  statusText = signal("Sistema pronto");

  private loreService = inject(LoreService);
  private convService = inject(ConversationService);

  constructor() {
    // Reload messages whenever the active conversation changes
    effect(() => {
      const id = this.convService.activeId();
      const conv = untracked(() =>
        this.convService.conversations().find((c) => c.id === id)
      );
      if (conv) {
        this.messages.set([...conv.messages]);
        this.isAskingRules.set(false);
        this.statusText.set("Sistema pronto");
      }
    });
  }

  ngAfterViewChecked(): void {
    this.scrollToBottom();
  }

  askQuestion(): void {
    const q = this.question().trim();
    if (!q || this.isAskingRules()) return;

    this.question.set("");

    // Auto-name the conversation on first question
    this.convService.autoName(q);

    this.messages.update((msgs) => [
      ...msgs,
      { role: "master", content: q, timestamp: new Date() },
    ]);
    this.convService.saveMessages(this.messages());

    this.isAskingRules.set(true);
    this.statusText.set("Consultando oráculo...");

    this.loreService
      .ask({ question: q, campaignId: this.convService.activeId() })
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
          this.convService.saveMessages(this.messages());
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
          this.convService.saveMessages(this.messages());
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
