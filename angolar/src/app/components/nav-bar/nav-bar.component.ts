import {Component, OnDestroy, OnInit} from '@angular/core';
import {Subscription} from "rxjs";
import {StorageService} from "../../services/storage.service";
import {UserService} from "../../services/user.service";
import {EventBusService} from "../../services/event-bus.service";
import {Router} from "@angular/router";
import {RoutingUtils} from "../../utils/routing-utils";

@Component({
  selector: 'app-nav-bar',
  templateUrl: './nav-bar.component.html',
  styleUrls: ['./nav-bar.component.css']
})
export class NavBarComponent implements OnInit, OnDestroy {
  eventBusSubscription?: Subscription;
  isLoggedIn?: boolean;
  loggedInEmail?: string;

  constructor(private storageService: StorageService, private userService: UserService, private eventBusService: EventBusService, private router: Router) {
    this.isLoggedIn = false;
  }

  ngOnInit(): void {
    this.eventBusSubscription = this.eventBusService.on('logout', () => {
      this.logout();
    })
    this.isLoggedIn = this.storageService.isLoggedIn();
    if (this.isLoggedIn) {
      this.loggedInEmail = this.storageService.getUserEmail();
    }
  }

  ngOnDestroy(): void {
    this.eventBusSubscription?.unsubscribe();
  }

  login(): void {
    this.router.navigateByUrl('/login');
  }

  register(): void {
    this.router.navigateByUrl('/register');
  }

  logout(): void {
    this.storageService.clean();
    RoutingUtils.goToHomepage(this.router);
    window.location.reload();
  }

  isOnHomepage(): boolean {
    return this.router.url == '/';
  }

  getUserInfo(): void {
    this.router.navigateByUrl('/user/info');
  }

  goToHomepage(): void {
    RoutingUtils.goToHomepage(this.router);
  }
}
