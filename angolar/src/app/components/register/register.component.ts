import {Component, OnInit} from '@angular/core';
import {UserService} from "../../services/user.service";
import {FormBuilder, FormControl, Validators} from "@angular/forms";
import {LoggerService} from "../../services/logger.service";
import {ToastLevel} from "../../models/toast-level";
import {Router} from "@angular/router";
import {AlertService} from "../../services/alert.service";

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {

  public email = new FormControl('', [Validators.required, Validators.email]);
  public password = new FormControl('', [Validators.required, Validators.minLength(8)]);
  public name = new FormControl('', [Validators.required]);

  public hide = true;

  constructor(private userService: UserService,
              private formBuilder: FormBuilder,
              private logger: LoggerService,
              private alertService: AlertService,
              private router: Router) {
  }

  ngOnInit() {
  }

  createUser() {
    let hasErrors = false;
    if (this.email.value === null || this.email.value === '') {
      this.logger.toast(ToastLevel.WARN, "Email cannot be empty", "Register User Error");
      hasErrors = true;
    }
    if (this.name.value === null || this.name.value === '') {
      this.logger.toast(ToastLevel.WARN, "Name cannot be empty", "Register User Error");
      hasErrors = true;
    }
    if (this.password.value === null || this.password.value === '') {
      this.logger.toast(ToastLevel.WARN, "Password cannot be empty", "Register User Error");
      hasErrors = true;
    }
    if (!hasErrors) {
      this.userService.register(<string>this.email.value, <string>this.name.value, <string>this.password.value)
        .subscribe({
          next: data => {
            this.alertService.success('Registration successful', true);
            this.goToLoginPage(data.data.email);
          },
          error: err => {
            this.alertService.error('ERROR: ' + err.error.error);
          },
          complete: () => this.logger.info('RegisterUser() DONE')
        });
    }
  }

  private goToLoginPage(email: string): void {
    this.router.navigateByUrl("/login").then(() => {
      this.logger.toast(ToastLevel.SUCCESS, 'New user ' + email + ' created', 'Register User SUCCESS');
    })
  }
}
