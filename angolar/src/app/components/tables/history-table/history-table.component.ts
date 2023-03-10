import {Component, OnDestroy, OnInit, ViewChild} from '@angular/core';
import {MatTableDataSource} from "@angular/material/table";
import {HistoryEntry} from "../../../models/history-entry";
import {Subscription} from "rxjs";
import {UserInfo} from "../../../models/user-info";
import {RedirectionService} from "../../../services/redirection.service";
import {LoggerService} from "../../../services/logger.service";
import {Router} from "@angular/router";
import {UserService} from "../../../services/user.service";
import {StorageService} from "../../../services/storage.service";
import {MatPaginator} from "@angular/material/paginator";
import {HistoryService} from "../../../services/history.service";
import {ToastLevel} from "../../../models/toast-level";
import {DateUtils} from "../../../utils/date-utils";
import {Redirection} from "../../../models/redirection";

@Component({
  selector: 'app-history-table',
  templateUrl: './history-table.component.html',
  styleUrls: ['./history-table.component.css']
})
export class HistoryTableComponent implements OnInit, OnDestroy {
  dataSource = new MatTableDataSource<HistoryEntry>();
  @ViewChild(MatPaginator) paginator: MatPaginator | null = null;
  displayColumns = ['redirection', 'visited_at', 'user_id'];

  redirectionMapSubscription?: Subscription;
  redirections?: Map<number, Redirection>;

  history?: HistoryEntry[];
  userMap: Map<number, UserInfo> = new Map();
  userInfo?: UserInfo;

  public constructor(private redirectionService: RedirectionService,
                     private logger: LoggerService,
                     private router: Router,
                     private userService: UserService,
                     private storageService: StorageService,
                     private historyService: HistoryService) {
  }

  ngOnDestroy(): void {
    this.redirectionMapSubscription?.unsubscribe();
  }

  ngOnInit(): void {
    this.dataSource.paginator = this.paginator;
    this.redirectionMapSubscription = this.redirectionService.getRedirectionMapObservable()
      .subscribe(map => {
        if (map !== undefined) {
          this.redirections = map;
        }
      });
    this.userService.getUserList().subscribe({
      next: res => {
        this.logger.log({status: res.status, data: res.data});
        if (res.status === 200) {
          const userList = res.data;
          userList.forEach(user => {
            this.userMap.set(user.id, user);
          })
        }
      },
      error: error => {
        this.logger.error(error.message);
        this.logger.toast(ToastLevel.ERROR, error.error.error, 'getUserList() ERROR');
      },
      complete: () => this.logger.info('getUserList() DONE')
    });
    this.userService.getUserInfo().subscribe({
      next: res => {
        this.logger.log({status: res.status, data: res.data});
        if (res.status === 200) {
          this.userInfo = res.data;
        }
      },
      error: () => {
        this.logger.info('User not connected');
      },
      complete: () => this.logger.info('getUserInfo() DONE')
    });
    this.historyService.getHistory()
      .subscribe({
        next: res => {
          this.logger.log({status: res.status, data: res.data});
          if (res.status === 200) {
            this.dataSource.data = res.data;
          }
        },
        error: error => {
          this.logger.error(error.message);
          this.logger.toast(ToastLevel.ERROR, error.error.error, 'getHistory() ERROR');
        },
        complete: () => this.logger.info('getHistory() DONE')
      });
    this.redirectionService.getRedirections();
  }

  public getCleanDate(date: string): string {
    return DateUtils.CleanDate(new Date(date));
  }

  public displayCreatorEmail(id: number): string {
    const user = this.userMap.get(id);
    if (user !== undefined) {
      return user.email;
    }
    return '';
  }

  public displayRedirectionUrl(id: number): string {
    const redirection = this.redirections?.get(id);
    if (redirection !== undefined) {
      return redirection.redirect_url;
    }
    return '';
  }


  public isUserLoggedIn(): boolean {
    return this.storageService.isLoggedIn();
  }
}
