import {Injectable} from '@angular/core';
import {HttpEvent, HttpHandler, HttpInterceptor, HttpRequest} from '@angular/common/http';
import {Observable} from 'rxjs';
import {StorageService} from "../services/storage.service";
import {LoggerService} from "../services/logger.service";
import {UserService} from "../services/user.service";
import {PropertiesService} from "../services/properties.service";

@Injectable()
export class AuthInterceptor implements HttpInterceptor {

  private static AUTHORIZATION = 'Authorization';
  private static BEARER = 'Bearer';
  private angolarSecretKey: string;

  constructor(private storageService: StorageService,
              private logger: LoggerService,
              private userService: UserService,
              private propertiesService: PropertiesService) {
    this.angolarSecretKey = propertiesService.angolarSecretKey;
  }

  private createAuthHeader(token: string): string {
    return AuthInterceptor.BEARER + ' ' + token;
  }

  intercept(request: HttpRequest<unknown>, next: HttpHandler): Observable<HttpEvent<unknown>> {
    var headers = request.headers;
    headers = headers
      .set('Content-Type', 'application/json')
      .set('ANGOLAR_SECRET', this.angolarSecretKey);

    const userToken = this.storageService.getUserToken();
    if (userToken !== undefined) {
      headers = headers.set(AuthInterceptor.AUTHORIZATION, this.createAuthHeader(userToken));
    }

    const finalRequest = request.clone({
      headers: headers
    });

    return next.handle(finalRequest);
  }
}
