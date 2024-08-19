import {Component, inject, Input, signal} from '@angular/core';
import {NgIf} from "@angular/common";
import {FormControl, FormGroup, ReactiveFormsModule, Validators} from "@angular/forms";
import {RegisterService} from "../../../../services/register-service/register.service";


@Component({
  selector: 'app-register-overlay',
  standalone: true,
  imports: [
    NgIf,
    ReactiveFormsModule
  ],
  templateUrl: './register-overlay.component.html',
  styleUrl: './register-overlay.component.scss'
})
export class RegisterOverlayComponent {
  @Input() isRegisterOverlay: any;
  isPasswordVisible = signal<boolean>(false);
  isPasswordVisibleRepeat = signal<boolean>(false);

  formRegister = new FormGroup({
    username: new FormControl(null, Validators.required),
    password: new FormControl(null, Validators.required),
  })

  registerService = inject(RegisterService)

  onSubmit() {

    if (this.formRegister.valid) {

      //@ts-ignore
      this.registerService.registration(this.formRegister.value)
        .subscribe(res => console.log(res));
    }
  }
}
