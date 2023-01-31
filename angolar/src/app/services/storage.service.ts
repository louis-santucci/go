import {Injectable} from '@angular/core';
import {LoggerService} from "./logger.service";

const TOKEN_KEY = "ANGOLAR_TOKEN_KEY";
const USER_EMAIL = "ANGOLAR_USER_EMAIL";

@Injectable({
  providedIn: 'root'
})
export class StorageService {

  constructor(private logger: LoggerService) {
  }

  clean(): void {
    window.sessionStorage.clear();
    this.logger.log('Cleaned local storage');
  }

  public saveUser(token: string, email: string): void {
    window.sessionStorage.removeItem(TOKEN_KEY);
    window.sessionStorage.setItem(TOKEN_KEY, JSON.stringify(token));
    window.sessionStorage.setItem(USER_EMAIL, JSON.stringify(email));
  }

  public getUserToken(): string | undefined {
    const token = window.sessionStorage.getItem(TOKEN_KEY);
    if (token) {
      return JSON.parse(token);
    }

    return undefined;
  }

  public getUserEmail(): string | undefined {
    const email = window.sessionStorage.getItem(USER_EMAIL);
    if (email) {
      return JSON.parse(email);
    }

    return undefined;
  }

  public isLoggedIn(): boolean {
    const user = window.sessionStorage.getItem(TOKEN_KEY);
    const email = window.sessionStorage.getItem(USER_EMAIL);
    return user !== null && email !== null;
  }


}
