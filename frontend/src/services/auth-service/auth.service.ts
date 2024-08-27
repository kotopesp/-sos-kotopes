import {inject, Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {environment} from "../../environments/environment";
import {FormControl, ɵFormGroupValue, ɵTypedOrUntyped} from "@angular/forms";
import {Observable} from "rxjs";

interface LoginResponse {
  token: string;
  user: {
    id: number;
    name: string;
    email: string;
  };
}

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  http = inject(HttpClient);
  baseApiUrl: string = environment.apiUrl;

  login(payload: ɵTypedOrUntyped<{ password: FormControl<null>; username: FormControl<null> }, ɵFormGroupValue<{
    password: FormControl<null>;
    username: FormControl<null>
  }>, string>): Observable<LoginResponse> {
    return this.http.post(
      `${this.baseApiUrl}auth/login`,
      payload
      )
  }
}
