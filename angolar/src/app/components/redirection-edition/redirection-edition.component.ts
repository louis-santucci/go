import {Component, OnDestroy, OnInit} from '@angular/core';
import {FormControl, Validators} from "@angular/forms";
import {LoggerService} from "../../services/logger.service";
import {RedirectionService} from "../../services/redirection.service";
import {ActivatedRoute, Router} from "@angular/router";
import {ToastLevel} from "../../models/toast-level";
import {Subscription} from "rxjs";
import {Redirection} from "../../models/redirection";

@Component({
  selector: 'app-redirection-edition',
  templateUrl: './redirection-edition.component.html',
  styleUrls: ['./redirection-edition.component.css']
})
export class RedirectionEditionComponent implements OnInit, OnDestroy {

  private static URL_REGEX = /^(?:http(s)?:\/\/)?[\w.-]+(?:\.[\w\.-]+)+[\w\-\._~:/?#[\]@!\$&'\(\)\*\+,;=.]+$/;

  public id?: any;
  private routeSubscription?: Subscription;
  public redirectionSubscription?: Subscription;
  public redirection?: Redirection;
  public shortcut = new FormControl('', [Validators.required]);
  public redirectUrl = new FormControl('', [Validators.required, Validators.pattern(RedirectionEditionComponent.URL_REGEX)]);

  public constructor(private logger: LoggerService,
                     private redirectionService: RedirectionService,
                     private router: Router,
                     private route: ActivatedRoute) {
  }

  ngOnInit(): void {
    this.routeSubscription = this.route.params.subscribe(params => {
      this.id = params['id'];
      this.redirectionSubscription = this.redirectionService.getRedirectionObservable().subscribe(redirection => {
        this.redirection = redirection;
        this.shortcut.setValue(<string>this.redirection?.shortcut);
        this.redirectUrl.setValue(<string>this.redirection?.redirect_url);
      });
      this.redirectionService.getRedirection(this.id);
    })
  }

  private static isNumber(id: any) {
    return typeof(id) === 'number';
  }

  ngOnDestroy() {
    this.routeSubscription?.unsubscribe();
  }

  public editRedirection() {
    let hasErrors = false;
    if (this.shortcut.value === null || this.shortcut.value === '') {
      this.logger.toast(ToastLevel.WARN, 'Shortcut cannot be empty', 'Create Redirection Error');
      hasErrors = true;
    }
    if (this.redirectUrl.value === null || this.redirectUrl.value === '') {
      this.logger.toast(ToastLevel.WARN, 'Redirection URL cannot be empty', 'Create Redirection Error');
      hasErrors = true;
    }

    if (!hasErrors) {
      if (this.redirection) {
        this.redirectionService.editRedirection(this.redirection?.id, <string>this.shortcut.value, <string>this.redirectUrl.value);
        this.router.navigateByUrl('/');
      }
    }
  }
}
