import { Component } from '@angular/core';
import {RouterLink} from "@angular/router";
import {NgForOf} from "@angular/common";

@Component({
  selector: 'app-header',
  standalone: true,
  imports: [
    RouterLink,
    NgForOf
  ],
  templateUrl: './header.component.html',
  styleUrl: './header.component.scss'
})
export class HeaderComponent {
  headerItems = [
    {
      path: '',
      text: 'Как я могу помочь?',
      className: 'header__how-to-help'
    },
    {
      path: 'ads',
      text: 'Объявления',
      className: 'header__ads'
    },
    {
      path: 'login',
      text: 'Войти',
      className: 'header__login'
    },
  ]
}