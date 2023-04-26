import {Component, OnInit} from '@angular/core';
import {UserInfo} from "../../models/user-info";
import {UserService} from "../../services/user.service";
import {AlertService} from "../../services/alert.service";

import {Router} from "@angular/router";
import {ToastLevel} from "../../models/toast-level";
import {LoggerService} from "../../services/logger.service";

@Component({
  selector: 'app-user-edition',
  templateUrl: './user-edition.component.html',
  styleUrls: ['./user-edition.component.css']
})
export class UserEditionComponent implements OnInit {
  public connectedUser?: UserInfo;

  public constructor(private userService: UserService,
                     private logger: LoggerService,
                     private alertService: AlertService,
                     private router: Router) {
  }

  ngOnInit(): void {
    this.userService.getUserInfo().subscribe({
      next: res => {
        this.logger.log({status: res.status, data: res.data});
        if (res.status === 200) {
          this.connectedUser = res.data;
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


}
