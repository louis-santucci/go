import {Component, OnInit} from '@angular/core';
import {FormControl, FormGroup, Validators} from "@angular/forms";
import {StorageService} from "../../services/storage.service";
import {LoggerService} from "../../services/logger.service";
import {UserService} from "../../services/user.service";
import {ToastLevel} from "../../models/toast-level";
import {ActivatedRoute, Router} from "@angular/router";
import {AlertService} from "../../services/alert.service";
import {RoutingUtils} from "../../utils/routing-utils";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  public loginFormGroup = new FormGroup({
    email: new FormControl('', [Validators.email, Validators.required]),
    password: new FormControl('', [Validators.required])
  })


  public isLoggedIn = false;
  public isSuccessful = false;
  private returnUrl?: string;

  constructor(private route: ActivatedRoute,
              private storageService: StorageService,
              private logger: LoggerService,
              private userService: UserService,
              private alertService: AlertService,
              private router: Router) {
  }

  ngOnInit(): void {
    if (this.storageService.isLoggedIn()) {
      this.isLoggedIn = true;
      RoutingUtils.goToHomepage(this.router);
    }
    this.returnUrl = this.route.snapshot.queryParams['returnUrl'] || '/';
  }

  login(): void {
    let hasErrors = false;
    if (this.loginFormGroup.value.email === null || this.loginFormGroup.value.email === '') {
      this.logger.toast(ToastLevel.WARN, "Email cannot be empty", "Register User Error");
      hasErrors = true;
    }
    if (this.loginFormGroup.value.password === null || this.loginFormGroup.value.password === '') {
      this.logger.toast(ToastLevel.WARN, "Password cannot be empty", "Register User Error");
      hasErrors = true;
    }
    if (!hasErrors) {
      this.userService.login(<string>this.loginFormGroup.value.email, <string>this.loginFormGroup.value.password)
        .subscribe({
          next: data => {
            this.storageService.saveUser(data.token, data.user.email);
            this.router.navigate([this.returnUrl])
              .then(() => {
                window.location.reload();
              });
          },
          error: err => {
            this.isSuccessful = false;
            this.alertService.error('ERROR: ' + err.error.error);
          },
          complete: () => this.logger.info('LoginUser() DONE')
        });
    }
  }
}
