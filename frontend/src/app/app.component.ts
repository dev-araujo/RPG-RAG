import { Component } from "@angular/core";
import { MasterPanelComponent } from "./components/master-panel/master-panel.component";
import { SidebarComponent } from "./components/sidebar/sidebar.component";

@Component({
  selector: "app-root",
  templateUrl: "./app.component.html",
  styleUrls: ["./app.component.scss"],
  imports: [MasterPanelComponent, SidebarComponent],
})
export class AppComponent {
  title = "LoreKeeper";
}
