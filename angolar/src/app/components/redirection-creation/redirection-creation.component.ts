import {Component, OnInit} from '@angular/core';
import {FormControl, Validators} from "@angular/forms";
import {RedirectionService} from "../../services/redirection.service";
import {LoggerService} from "../../services/logger.service";
import {ToastLevel} from "../../models/toast-level";
import {Router} from "@angular/router";

@Component({
  selector: 'app-redirection-creation',
  templateUrl: './redirection-creation.component.html',
  styleUrls: ['./redirection-creation.component.css']
})
export class RedirectionCreationComponent implements OnInit {

  private static URL_REGEX = /^(?:http(s)?:\/\/)?[\w.-]+(?:\.[\w\.-]+)+[\w\-\._~:/?#[\]@!\$&'\(\)\*\+,;=.]+$/;

  public shortcut = new FormControl('', [Validators.required]);
  public redirectUrl = new FormControl('', [Validators.required, Validators.pattern(RedirectionCreationComponent.URL_REGEX)]);

  public constructor(private logger: LoggerService,
                     private redirectionService: RedirectionService,
                     private router: Router) {
  }

  ngOnInit(): void {

  }

  public createRedirection() {
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
      this.redirectionService.createRedirection(<string>this.shortcut.value, <string>this.redirectUrl.value);
      this.router.navigateByUrl('/');
    }
  }
}
