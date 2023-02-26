import {Component, OnDestroy, OnInit, ViewChild} from '@angular/core';
import {animate, state, style, transition, trigger} from "@angular/animations";
import {Redirection} from "../../models/redirection";
import {MatTableDataSource} from "@angular/material/table";
import {Subscription} from "rxjs";
import {RedirectionService} from "../../services/redirection.service";
import {MatSort} from "@angular/material/sort";
import {LoggerService} from "../../services/logger.service";
import {DateUtils} from "../../utils/date-utils";
import {Router} from "@angular/router";
import {UserService} from "../../services/user.service";
import {UserInfo} from "../../models/user-info";
import {ToastLevel} from "../../models/toast-level";
import {StorageService} from "../../services/storage.service";
import {TooltipUtils} from "../../utils/tooltip-utils";

@Component({
  selector: 'app-redirection-table',
  templateUrl: './redirection-table.component.html',
  styleUrls: ['./redirection-table.component.css'],
  animations: [
    trigger('detailExpand', [
      state('collapsed', style({height: '0px', minHeight: '0'})),
      state('expanded', style({height: '*'})),
      transition('expanded <=> collapsed', animate('225ms cubic-bezier(0.4, 0.0, 0.2, 1)')),
    ]),
  ]
})
export class RedirectionTableComponent implements OnInit, OnDestroy {
  dataSource = new MatTableDataSource<Redirection>();
  @ViewChild(MatSort, {static: true}) sort: MatSort | null = null;
  displayColumns = ['shortcut', 'redirect_url', 'views', 'created_at', 'created_by', 'edit', 'delete'];
  displayColumnsExpanded = [...this.displayColumns, 'expand'];
  expandedRedirection?: Redirection;
  isSearchBarHidden = true;

  redirectionMapSubscription?: Subscription;
  userMap: Map<number, UserInfo> = new Map();
  userInfo?: UserInfo;

  public constructor(private redirectionService: RedirectionService,
                     private logger: LoggerService,
                     private router: Router,
                     private userService: UserService,
                     private storageService: StorageService) {
  }

  public applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.dataSource.filter = filterValue.trim().toLowerCase();
  }

  public ngOnInit() {
    this.redirectionMapSubscription = this.redirectionService.getRedirectionMapObservable()
      .subscribe(map => {
        if (map !== undefined) {
          this.dataSource.data = [...map.values()];
          this.dataSource.sort = this.sort;
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
    })
  }

  public getCleanDate(date: string): string {
    return DateUtils.CleanDate(new Date(date));
  }

  public ngOnDestroy() {
    this.redirectionMapSubscription?.unsubscribe();
  }

  public deleteRedirection(id: number): void {
    this.logger.log("Deleting redirection #" + id);
    this.redirectionService.deleteRedirection(id);
  }

  public editRedirection(id: number): void {
    this.logger.log("Editing redirection #" + id);
    const url = '/redirection/edit/' + id;
    this.router.navigateByUrl(url)
      .then(window.location.reload);
  }

  public displayCreatorEmail(id: number): string {
    const user = this.userMap.get(id);
    if (user !== undefined) {
      return user.email;
    }
    return 'Deleted User';
  }

  public isUserLoggedIn(): boolean {
    return this.storageService.isLoggedIn();
  }

  public displayTooltip(id: number, message: string): string {
    if (!this.isUserLoggedIn() || this.userInfo === undefined) {
      return TooltipUtils.LOGIN_ERROR;
    }
    if (this.userInfo.id !== id) {
      return TooltipUtils.OWNERSHIP_ERROR;
    }
    return message;
  }

  public get tooltipUtils() {
    return TooltipUtils;
  }

  public displaySearchBar(): void {
    this.isSearchBarHidden = !this.isSearchBarHidden;
    if (this.isSearchBarHidden) {
      this.dataSource.filter = '';
    }
  }

  public redirect(url: string, id: number): void {
    this.redirectionService.incrementRedirectionView(id).subscribe({
      next: res => {
        this.logger.log({status: res.status, data: res.data});
        window.location.href = url;
      },
      error: error => {
        this.logger.error(error)
      }
    });
  }
}
