import {Component, OnInit} from '@angular/core';
import {FormControl, Validators} from "@angular/forms";
import {StorageService} from "../../services/storage.service";
import {LoggerService} from "../../services/logger.service";
import {UserService} from "../../services/user.service";
import {ToastLevel} from "../../models/toast-level";
import {Router} from "@angular/router";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  public email = new FormControl('', [Validators.email, Validators.required]);
  public password = new FormControl('', [Validators.required]);

  public hide = true;

  public isLoggedIn = false;
  public isSuccessful = false;
  public errorMessage = '';
  public loggedToken?: string;

  constructor(private storageService: StorageService,
              private logger: LoggerService,
              private userService: UserService,
              private router: Router) {
  }

  ngOnInit(): void {
    if (this.storageService.isLoggedIn()) {
      this.isLoggedIn = true;
      this.loggedToken = this.storageService.getUserToken();
    }
  }

  login(): void {
    let hasErrors = false;
    if (this.email.value === null || this.email.value === '') {
      this.logger.toast(ToastLevel.WARN, "Email cannot be empty", "Register User Error");
      hasErrors = true;
    }
    if (this.password.value === null || this.password.value === '') {
      this.logger.toast(ToastLevel.WARN, "Password cannot be empty", "Register User Error");
      hasErrors = true;
    }
    if (!hasErrors) {
      this.userService.login(<string>this.email.value, <string>this.password.value)
        .subscribe({
          next: data => {
            this.storageService.saveUser(data.token, data.email);
            this.goToHomepage();
          },
          error: err => {
            this.errorMessage = err.error.message;
            this.isSuccessful = false;
          },
          complete: () => this.logger.info('LoginUser() DONE')
        });
    }
  }

  private goToHomepage(): void {
    this.router.navigateByUrl("/")
  }
}
