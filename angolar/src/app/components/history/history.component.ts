import { Component } from '@angular/core';
import {LoggerService} from "../../services/logger.service";
import {AlertService} from "../../services/alert.service";
import {HistoryService} from "../../services/history.service";
import {ToastLevel} from "../../models/toast-level";

@Component({
  selector: 'app-history',
  templateUrl: './history.component.html',
  styleUrls: ['./history.component.css']
})
export class HistoryComponent {

  constructor(private logger: LoggerService,
              private alertService: AlertService,
              private historyService: HistoryService) {
  }

  public clearHistory(): void {
    this.historyService.resetHistory()
      .subscribe({
        next: res => {
          this.logger.log({status: res.status, data: res.data});
          window.location.reload();
        },
        error: error => {
          this.logger.error(error);
          this.logger.toast(ToastLevel.ERROR, error.error.error, 'clearHistory() ERROR');
          this.alertService.error(error.error.error, false);
        },
        complete: () => this.logger.info('clearHistory() DONE')
      })
  }
}
