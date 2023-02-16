import {Component, OnDestroy, OnInit} from '@angular/core';
import {FormControl, FormGroup, Validators} from "@angular/forms";
import {LoggerService} from "../../services/logger.service";
import {RedirectionService} from "../../services/redirection.service";
import {ActivatedRoute, Router} from "@angular/router";
import {ToastLevel} from "../../models/toast-level";
import {Subscription} from "rxjs";
import {Redirection} from "../../models/redirection";
import {RegexUtils} from "../../utils/regex-utils";

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
                     private route: ActivatedRoute) {
  }

  ngOnInit(): void {
    this.routeSubscription = this.route.params.subscribe(params => {
      this.id = params['id'];
      this.redirectionSubscription = this.redirectionService.getRedirectionObservable().subscribe(redirection => {
        this.redirection = redirection;
        this.redirectionEditionFormGroup.setValue({
          shortcut: <string>this.redirection?.shortcut,
          redirectUrl: <string>this.redirection?.redirect_url
        })
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
        this.router.navigateByUrl('/');
      }
    }
  }
}
