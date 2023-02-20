import {Component} from '@angular/core';
import {Router} from "@angular/router";
import {RoutingUtils} from "../../utils/routing-utils";

@Component({
  selector: 'app-error',
  templateUrl: './error-not-found.component.html',
  styleUrls: ['./error-not-found.component.css']
})
export class ErrorNotFoundComponent {

  constructor(private router: Router) {
  }


  goToHomepage(): void {
    RoutingUtils.goToHomepage(this.router);
  }

}
