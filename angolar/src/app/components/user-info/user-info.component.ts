import {Component, OnInit} from '@angular/core';
import {UserService} from "../../services/user.service";
import {UserInfo} from "../../models/user-info";
import {LoggerService} from "../../services/logger.service";
import {ToastLevel} from "../../models/toast-level";
import {AlertService} from "../../services/alert.service";
import {DateUtils} from "../../utils/date-utils";

@Component({
  selector: 'app-user-info',
  templateUrl: './user-info.component.html',
  styleUrls: ['./user-info.component.css']
})
export class UserInfoComponent implements OnInit{
  public connectedUser?: UserInfo;

  public constructor(private userService: UserService,
                     private logger: LoggerService,
                     private alertService: AlertService) {
  }

  ngOnInit() {
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

  public openEdition() {

  }

  public deleteUser() {

  }

  public getCleanDate(date: string | undefined): string {
    if (date === undefined) {
      return '';
    }
    return DateUtils.CleanDate(new Date(date));
  }
}
