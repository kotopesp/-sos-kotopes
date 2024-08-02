import { Component } from '@angular/core';
import { NgClass } from "@angular/common";

@Component({
  selector: 'find-status-flag',
  standalone: true,
  imports: [NgClass],
  templateUrl: './find-status-flag.component.html',
  styleUrl: './find-status-flag.component.scss'
})
export class FindStatusFlagComponent {
  // here we get data from the backend, ex. we get initialState = "Найден"
  initialStateArr = ["Пропал", "Найден", "Ищет дом"]
  initialState: string = this.initialStateArr[Math.floor(Math.random() * 3)]

  ngOnInit() {
    this.setFlagClass();
  }

  petStatusClass: string = ""
  buttonLabel: string = ""
  setFlagClass() {
    switch(this.initialState) {
      case "Пропал": {
        this.buttonLabel = "Пропал"
        this.petStatusClass = "lost_status"
        break
      }

      case "Найден": {
        this.buttonLabel = "Найден"
        this.petStatusClass = "found-home_status"
        break
      }

      case "Ищет дом": {
        this.buttonLabel = "Ищет дом"
        this.petStatusClass = "looking-for-home_status"
        break
      }
    }
  }

}
