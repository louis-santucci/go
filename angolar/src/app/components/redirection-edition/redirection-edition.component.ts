import {Component, OnDestroy, OnInit} from '@angular/core';
import {FormControl, FormGroup, Validators} from "@angular/forms";
import {LoggerService} from "../../services/logger.service";
import {RedirectionService} from "../../services/redirection.service";
import {ActivatedRoute, Router} from "@angular/router";
import {ToastLevel} from "../../models/toast-level";
import {Subscription} from "rxjs";
import {Redirection} from "../../models/redirection";
import {RegexUtils} from "../../utils/regex-utils";
import {RoutingUtils} from "../../utils/routing-utils";
import {UserService} from "../../services/user.service";
import {AlertService} from "../../services/alert.service";

@Component({
  selector: 'app-redirection-edition',
  templateUrl: './redirection-edition.component.html',
  styleUrls: ['./redirection-edition.component.css']
})
export class RedirectionEditionComponent implements OnInit, OnDestroy {

  public id?: any;
  private routeSubscription?: Subscription;
  public redirectionSubscription?: Subscription;
  public redirection?: Redirection;

  public redirectionEditionFormGroup = new FormGroup({
    shortcut: new FormControl('', [Validators.required]),
    redirectUrl: new FormControl('', [Validators.required, Validators.pattern(RegexUtils.URL_REGEX)])
  })

  public constructor(private logger: LoggerService,
                     private redirectionService: RedirectionService,
                     private router: Router,
                     private route: ActivatedRoute,
                     private userService: UserService,
                     private alertService: AlertService) {
  }

  ngOnInit(): void {
    this.routeSubscription = this.route.params.subscribe(params => {
      this.id = params['id'];
      this.redirectionService.getRedirection(this.id).subscribe({
        next: res => {
          this.logger.log({status: res.status, data: res.data});
          if (res.status === 200) {
            this.redirection = res.data;
            this.userService.getUserInfo().subscribe({
              next: userRes => {
                this.logger.log({status: userRes.status, data: userRes.data});
                if (userRes.status === 200) {
                  const userInfo = userRes.data;
                  if (userInfo.id !== this.redirection?.creator_id) {
                    this.alertService.error('You are not the owner of this redirection. You can\'t edit it.', true);
                    RoutingUtils.goToUnauthorizedPage(this.router);
                  }
                }
              },
              error: error => {
                this.logger.error(error);
                this.logger.toast(ToastLevel.ERROR, error.error.error, 'getRedirection(' + this.id + ') ERROR');
              }
            });
            this.redirectionEditionFormGroup.setValue({
              shortcut: <string>this.redirection?.shortcut,
              redirectUrl: <string>this.redirection?.redirect_url
            })
          }
        },
        error: error => {
          this.logger.error(error);
          this.logger.toast(ToastLevel.ERROR, error.error.error, 'getRedirection(' + this.id + ') ERROR');
          if (error.status === 404) {
            RoutingUtils.goToNotFoundPage(this.router);
          }
        },
        complete: () => this.logger.info('getRedirection(' + this.id + ') DONE')
      });
    })
  }

  ngOnDestroy() {
    this.routeSubscription?.unsubscribe();
  }

  public editRedirection() {
    let hasErrors = false;
    if (this.redirectionEditionFormGroup.value.shortcut === null || this.redirectionEditionFormGroup.value.shortcut === '') {
      this.logger.toast(ToastLevel.WARN, 'Shortcut cannot be empty', 'Create Redirection Error');
      hasErrors = true;
    }
    if (this.redirectionEditionFormGroup.value.redirectUrl === null || this.redirectionEditionFormGroup.value.redirectUrl === '') {
      this.logger.toast(ToastLevel.WARN, 'Redirection URL cannot be empty', 'Create Redirection Error');
      hasErrors = true;
    }

    if (!hasErrors) {
      if (this.redirection) {
        this.redirectionService.editRedirection(this.redirection?.id, <string>this.redirectionEditionFormGroup.value.shortcut, <string>this.redirectionEditionFormGroup.value.redirectUrl);
        RoutingUtils.goToHomepage(this.router);
      }
    }
  }

  public resetRedirectionViews(): void {
    if (this.redirection !== undefined) {
      this.redirectionService.resetRedirectionView(this.redirection?.id)
        .subscribe({
          next: res => {
            this.logger.log({status: res.status, data: res.data});
            if (this.redirection !== undefined) {
              this.redirection.views = 0;
            }
          },
          error: error => {
            this.logger.error(error);
            this.logger.toast(ToastLevel.ERROR, error.error.error, 'getRedirection(' + this.id + ') ERROR');
            this.alertService.error(error.error.error, false);
          },
          complete: () => this.logger.info('resetRedirectionViews(' + this.redirection?.id + ') DONE')
        });
    }
  }
}
