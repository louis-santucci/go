import {Injectable} from '@angular/core';
import {ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot, UrlTree} from '@angular/router';
import {Observable} from 'rxjs';
import {AlertService} from "../services/alert.service";

@Injectable({
  providedIn: 'root'
})
export class ParameterGuard implements CanActivate {

  constructor(private router: Router,
              private alertService: AlertService) {
  }

  canActivate(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot): Observable<boolean | UrlTree> | Promise<boolean | UrlTree> | boolean | UrlTree {
    const id = route.paramMap.get('id');
    if (!this.isNumber(id)) {
      this.alertService.error('ERROR: Given \'id\' is not a number', true);
      this.router.navigateByUrl('error/notFound');
      return false;
    }
    return true;
  }

  private isNumber(id: any) {
    return !isNaN(Number(id));
  }
}
