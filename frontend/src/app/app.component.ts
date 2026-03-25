import { Component } from "@angular/core";
import { MasterPanelComponent } from "./components/master-panel/master-panel.component";

@Component({
  selector: "app-root",
  templateUrl: "./app.component.html",
  styleUrls: ["./app.component.scss"],
  imports: [MasterPanelComponent],
})
export class AppComponent {
  title = "LoreKeeper";
}
