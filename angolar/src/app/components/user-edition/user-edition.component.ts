import {Component, OnInit} from '@angular/core';
import {UserInfo} from "../../models/user-info";
import {UserService} from "../../services/user.service";
import {AlertService} from "../../services/alert.service";

import {Router} from "@angular/router";
import {ToastLevel} from "../../models/toast-level";
import {LoggerService} from "../../services/logger.service";
import {FormControl, FormGroup, Validators} from "@angular/forms";
import {RoutingUtils} from "../../utils/routing-utils";
import {StorageService} from "../../services/storage.service";
import {EventBusService} from "../../services/event-bus.service";
import {EventData} from "../../models/event-data";

@Component({
  selector: 'app-user-edition',
  templateUrl: './user-edition.component.html',
  styleUrls: ['./user-edition.component.css']
})
export class UserEditionComponent implements OnInit {
  public connectedUser: UserInfo;

  public editionFormGroup: FormGroup;

  public constructor(private userService: UserService,
                     private logger: LoggerService,
                     private alertService: AlertService,
                     private storageService: StorageService,
                     private eventBusService: EventBusService,
                     private router: Router) {
  }

  ngOnInit(): void {
    this.userService.getUserInfo().subscribe({
      next: res => {
        this.logger.log({status: res.status, data: res.data});
        if (res.status === 200) {
          this.connectedUser = res.data;
          this.editionFormGroup = new FormGroup({
            name: new FormControl(this.connectedUser.name, [Validators.required]),
            email: new FormControl(this.connectedUser.email, [Validators.email, Validators.required]),
            password: new FormControl('', [Validators.required, Validators.minLength(8)]),
          });

        }
      },
      error: error => {
        this.logger.error(error.message);
        this.logger.toast(ToastLevel.ERROR, error.error.error, 'getUserInfo() ERROR');
        this.alertService.error(error.error.error, false);
      },
      complete: () => this.logger.info('getUserInfo() DONE')
    })
  }

  public openPasswordEdition(): void {
    RoutingUtils.goToPasswordEditionPage(this.router);
  }

  public editUser(): void {
    let hasErrors = false;
    if (this.editionFormGroup.value.email === null || this.editionFormGroup.value.email === '') {
      this.logger.toast(ToastLevel.WARN, "Email cannot be empty", "Edit User Error");
      hasErrors = true;
    }
    if (this.editionFormGroup.value.name === null || this.editionFormGroup.value.name === '') {
      this.logger.toast(ToastLevel.WARN, "Name cannot be empty", "Edit User Error");
      hasErrors = true;
    }
    if (this.editionFormGroup.value.password === null || this.editionFormGroup.value.password === '') {
      this.logger.toast(ToastLevel.WARN, "Password cannot be empty", "Register User Error");
      hasErrors = true;
    }

    if (!hasErrors) {
      this.userService.editUser(<string>this.editionFormGroup.value.email, <string>this.editionFormGroup.value.name, <string>this.editionFormGroup.value.password)
        .subscribe({
          next: data => {
            this.storageService.saveUser(data.token, data.user.email);
            this.eventBusService.emit(new EventData('emailUpdate', data.user.email));
            this.alertService.success('Edition successful', true);
            this.goToUserInfoPage();
          },
          error: err => {
            this.alertService.error('ERROR: ' + err.error.error);
          },
          complete: () => this.logger.info('EditUser() DONE')
        })
    }
  }

  public goToUserInfoPage(): void {
    RoutingUtils.goToUserInformationPage(this.router);
  }
}
