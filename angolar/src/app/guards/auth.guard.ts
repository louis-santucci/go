import { Injectable } from '@angular/core';
import {ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot, UrlTree} from '@angular/router';
import { Observable } from 'rxjs';
import {StorageService} from "../services/storage.service";
import {AlertService} from "../services/alert.service";
import {RoutingUtils} from "../utils/routing-utils";

@Injectable({
  providedIn: 'root'
})
export class AuthGuard implements CanActivate {

  constructor(private router: Router,
              private storageService: StorageService,
              private alertService: AlertService) {
  }

  canActivate(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot): Observable<boolean | UrlTree> | Promise<boolean | UrlTree> | boolean | UrlTree {
    if (this.storageService.getUserToken() !== undefined) {
      // Logged in so returns true
      return true;
    }

    // Not logged so redirect to login page
    this.alertService.error('ERROR: You need to be logged in to access this page.', true);
    RoutingUtils.goToLoginPage(this.router, state.url);
    return false;
  }

}
