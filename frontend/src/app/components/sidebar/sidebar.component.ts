import { Component, inject, signal } from "@angular/core";
import {
  ConversationService,
  Conversation,
} from "../../services/conversation.service";
import { DatePipe } from "@angular/common";
import { FormsModule } from "@angular/forms";

@Component({
  selector: "app-sidebar",
  templateUrl: "./sidebar.component.html",
  styleUrls: ["./sidebar.component.scss"],
  imports: [DatePipe, FormsModule],
})
export class SidebarComponent {
  conv = inject(ConversationService);
  editingId = signal<string | null>(null);
  editName = signal("");

  startEdit(c: Conversation, event: Event): void {
    event.stopPropagation();
    this.editingId.set(c.id);
    this.editName.set(c.name);
  }

  commitEdit(): void {
    const id = this.editingId();
    if (id) this.conv.renameConversation(id, this.editName());
    this.editingId.set(null);
  }

  cancelEdit(): void {
    this.editingId.set(null);
  }

  delete(id: string, event: Event): void {
    event.stopPropagation();
    this.conv.deleteConversation(id);
  }
}
