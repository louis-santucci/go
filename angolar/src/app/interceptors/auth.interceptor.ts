import {Injectable} from '@angular/core';
import {HttpErrorResponse, HttpEvent, HttpHandler, HttpInterceptor, HttpRequest} from '@angular/common/http';
import {catchError, Observable, throwError} from 'rxjs';
import {StorageService} from "../services/storage.service";
import {LoggerService} from "../services/logger.service";
import {UserService} from "../services/user.service";
import {PropertiesService} from "../services/properties.service";
import {EventBusService} from "../services/event-bus.service";
import {EventData} from "../models/event-data";

@Injectable()
export class AuthInterceptor implements HttpInterceptor {

  private static BAD_JWT_TOKEN_ERROR = 'bad JWT token';
  private static AUTHORIZATION = 'Authorization';
  private static BEARER = 'Bearer';

  constructor(private storageService: StorageService,
              private logger: LoggerService,
              private userService: UserService,
              private propertiesService: PropertiesService,
              private eventBusService: EventBusService) {
  }

  private createAuthHeader(token: string): string {
    return AuthInterceptor.BEARER + ' ' + token;
  }

  intercept(request: HttpRequest<unknown>, next: HttpHandler): Observable<HttpEvent<unknown>> {
    var headers = request.headers;
    headers = headers
      .set('Content-Type', 'application/json')

    const userToken = this.storageService.getUserToken();
    if (userToken !== undefined) {
      headers = headers.set(AuthInterceptor.AUTHORIZATION, this.createAuthHeader(userToken));
    }

    const finalRequest = request.clone({
      headers: headers
    });

    return next.handle(finalRequest)
      .pipe(catchError((error) => {
        // Handle automatic logout when JWT expired
        if (error instanceof HttpErrorResponse && error.status === 401 && error.error.error === AuthInterceptor.BAD_JWT_TOKEN_ERROR) {
          if (this.storageService.isLoggedIn()) {
            this.eventBusService.emit(new EventData('logout', null));
          }
          return next.handle(request);
        }
        // Else throw error normally
        return throwError(() => error);
    }));
  }
}
