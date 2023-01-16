import {Component, OnDestroy, OnInit, ViewChild} from '@angular/core';
import {animate, state, style, transition, trigger} from "@angular/animations";
import {Redirection} from "../../models/redirection";
import {MatTableDataSource} from "@angular/material/table";
import {Subscription} from "rxjs";
import {RedirectionService} from "../../services/redirection.service";
import {MatSort} from "@angular/material/sort";
import {LoggerService} from "../../services/logger.service";
import {DateUtils} from "../../utils/date-utils";

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
  displayColumns = ['shortcut', 'redirect_url', 'views', 'created_at'];
  displayColumnsExpanded = [...this.displayColumns, 'expand'];
  expandedRedirection?: Redirection;

  redirectionMapSubscription?: Subscription;

  public constructor(private redirectionService: RedirectionService, private logger: LoggerService) {
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
}
