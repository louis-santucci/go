import {Component, OnInit} from '@angular/core';
import {FormControl, FormGroup, Validators} from "@angular/forms";
import {RedirectionService} from "../../services/redirection.service";
import {LoggerService} from "../../services/logger.service";
import {ToastLevel} from "../../models/toast-level";
import {Router} from "@angular/router";
import {RegexUtils} from "../../utils/regex-utils";

@Component({
  selector: 'app-redirection-creation',
  templateUrl: './redirection-creation.component.html',
  styleUrls: ['./redirection-creation.component.css']
})
export class RedirectionCreationComponent implements OnInit {

  public redirectionEditionFormGroup = new FormGroup({
    shortcut: new FormControl('', [Validators.required]),
    redirectUrl: new FormControl('', [Validators.required, Validators.pattern(RegexUtils.URL_REGEX)])
  })

  public constructor(private logger: LoggerService,
                     private redirectionService: RedirectionService,
                     private router: Router) {
  }

  ngOnInit(): void {

  }

  public createRedirection() {
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
      this.redirectionService.createRedirection(
        <string>this.redirectionEditionFormGroup.value.shortcut,
        <string>this.redirectionEditionFormGroup.value.redirectUrl);
      this.router.navigateByUrl('/');
    }
  }
}
