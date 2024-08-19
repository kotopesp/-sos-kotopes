import {Component, Input} from "@angular/core";

@Component({
  selector: 'chat-type-button',
  standalone: true,
  templateUrl: "chat-type-button.component.html",
  styles: `
    .chat-type-icon {
      margin-right: 20px;
    }

    .counter {
      display: flex;
      align-items: center;
      justify-content: center;
      width: 30px;
      height: 30px;
      background: white;
      opacity: 70%;
      font-size: 18px;
      border-radius: 20px;
      margin-right: 15px;
    }

    .chat-type-button {
      display: flex;
      align-items: center;
      height: 100px;
      width: 100%;
      font-size: 20px;
      border-radius: 100px;
      padding-left: 30px;
      margin-bottom: 10px;
    }
  `
})
export class ChatTypeButtonComponent {
  @Input({required: true})
  label = ''

  @Input({required: true})
  color = ''

  @Input()
  counter = 0

  @Input()
  icon = "allChats.svg"
}
