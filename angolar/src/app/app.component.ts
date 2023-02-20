import {Component, OnInit} from '@angular/core';
import {Subscription} from "rxjs";
import {StorageService} from "./services/storage.service";
import {UserService} from "./services/user.service";
import {EventBusService} from "./services/event-bus.service";
import {LoggerService} from "./services/logger.service";
import {Router} from "@angular/router";
import {RoutingUtils} from "./utils/routing-utils";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  title = 'angolar';

  eventBusSub?: Subscription;

  constructor(private eventBusService: EventBusService,
              private logger: LoggerService,
              private userService: UserService,
              private storageService: StorageService,
              private router: Router) {
  }

  ngOnInit(): void {
    this.eventBusSub = this.eventBusService.on('logout', () => {
      this.logout();
    });
  }

  logout(): void {
    this.storageService.clean();
    RoutingUtils.goToHomepage(this.router);
    window.location.reload();
  }

}
