import { Component } from '@angular/core';
import {Router} from "@angular/router";
import {RoutingUtils} from "../../utils/routing-utils";

@Component({
  selector: 'app-error-unauthorized',
  templateUrl: './error-unauthorized.component.html',
  styleUrls: ['./error-unauthorized.component.css']
})
export class ErrorUnauthorizedComponent {
  constructor(private router: Router) {
  }

  goToHomepage(): void {
    RoutingUtils.goToHomepage(this.router);
  }
}
