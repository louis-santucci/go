import {Injectable} from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {BehaviorSubject, Observable} from "rxjs";
import {Redirection} from "../models/redirection";
import {OkResponse} from "../responses/ok-response";
import {PropertiesService} from "./properties.service";
import {LoggerService} from "./logger.service";
import {ToastLevel} from "../models/toast-level";
import {RedirectionInput} from "../dtos/redirection/redirection.input";
import {AlertService} from "./alert.service";
import {MapUtils} from "../utils/map-utils";

const httpOptions = {
  headers: new HttpHeaders({'Content-Type': 'application/json'})
};

@Injectable({
  providedIn: 'root'
})
export class RedirectionService {

  private readonly backendUrl?: string
  private readonly redirectionUrl: string;

  // Observable source
  private redirectionMapSource = new BehaviorSubject<Map<string, Redirection> | undefined>(undefined)

  // Observable source
  private redirectionMapObservable: Observable<Map<string, Redirection> | undefined> = this.redirectionMapSource.asObservable();

  constructor(private http: HttpClient,
              private propertiesService: PropertiesService,
              private logger: LoggerService,
              private alertService: AlertService) {
    this.backendUrl = this.propertiesService.backendUrl;
    this.redirectionUrl = this.backendUrl + '/api/redirection';
  }

  public getRedirectionMapObservable(): Observable<Map<string, Redirection> | undefined> {
    return this.redirectionMapObservable;
  }

  public getRedirections(): void {
    try {
      this.http.get<OkResponse<Redirection[]>>(this.redirectionUrl, httpOptions)
        .subscribe({
          next: res => {
            this.logger.log({status: res.status, data: res.data});
            if (res.status === 200) {
              const data = res.data;
              const redirectionMap = new Map<string, Redirection>();
              data.forEach(redirection => {
                redirectionMap.set(redirection.shortcut, redirection);
              });
              this.redirectionMapSource.next(redirectionMap);
            }
          },
          error: error => {
            this.logger.error(error.message);
            this.logger.toast(ToastLevel.ERROR, error.error.error, 'getRedirections() ERROR');
          },
          complete: () => this.logger.info('getRedirections() DONE')
        })
    } catch (error) {
      this.logger.error(error);
    }
  }

  public getRedirection(id: number): Observable<OkResponse<Redirection>> {
    const url = this.redirectionUrl + '/' + id;
    return this.http.get<OkResponse<Redirection>>(url, httpOptions);
  }

  public editRedirection(id: number, shortcut: string, redirectUrl: string): void {
    try {
      const input: RedirectionInput = {
        shortcut: shortcut,
        redirect_url: redirectUrl
      };
      const url = this.redirectionUrl + '/' + id;
      this.http.post<OkResponse<Redirection>>(url, input, httpOptions)
        .subscribe({
          next: res => {
            this.logger.log({status: res.status, data: res.data});
            if (res.status === 200) {
              let currentMap = this.redirectionMapSource.getValue();
              if (currentMap !== undefined) {
                const newRedirection = res.data;
                currentMap = MapUtils.deleteById(currentMap, id);
                currentMap.set(newRedirection.shortcut, newRedirection);
                this.redirectionMapSource.next(currentMap);
              }
            }
          },
          error: error => {
            this.logger.error(error);
            this.logger.toast(ToastLevel.ERROR, error.error.error, 'editRedirection(' + id + ') ERROR')
          },
          complete: () => this.logger.info('editRedirection() DONE')
        })
    } catch (error) {
      this.logger.error(error);
    }
  }

  public createRedirection(shortcut: string, redirectUrl: string): void {
    try {
      const input: RedirectionInput = {
        shortcut: shortcut,
        redirect_url: redirectUrl
      };
      this.http.post<OkResponse<Redirection>>(this.redirectionUrl, input, httpOptions)
        .subscribe({
          next: res => {
            this.logger.log({status: res.status, data: res.data});
            if (res.status === 200) {
              const currentMap = this.redirectionMapSource.getValue();
              if (currentMap !== undefined) {
                const newRedirection = res.data;
                currentMap.set(newRedirection.shortcut, newRedirection);
                this.redirectionMapSource.next(currentMap);
                this.alertService.success('New Redirection for URL ' + newRedirection.redirect_url + ' created', false);
              }
            }
          },
          error: error => {
            this.logger.error(error);
            this.logger.toast(ToastLevel.ERROR, error.error.error, 'createRedirection() ERROR');
            this.alertService.error('ERROR: ' + error.error.error);
          },
          complete: () => this.logger.info('createRedirection() DONE')
        });
    } catch (error) {
      this.logger.error(error);
    }
  }

  public deleteRedirection(id: number) {
    try {
      const url = this.redirectionUrl + '/' + id;
      this.http.delete<OkResponse<string>>(url, httpOptions)
        .subscribe({
          next: res => {
            this.logger.log({status: res.status, data: res.data});
            if (res.status === 200) {
              const currentMap = this.redirectionMapSource.getValue();
              let deletedShortcut = '';
              if (currentMap !== undefined) {
                for (let entry of currentMap.values()) {
                  if (entry.id === id) {
                    deletedShortcut = entry.shortcut;
                    break;
                  }
                }
                if (deletedShortcut !== undefined) {
                  currentMap.delete(deletedShortcut);
                }
                this.redirectionMapSource.next(currentMap);
                this.alertService.success('Redirection #' + id + ' deleted', false);
                this.logger.toast(ToastLevel.SUCCESS, 'Redirection #' + id + ' deleted', 'Delete Redirection SUCCESS');
              }
            }
          },
          error: error => {
            this.logger.error(error);
            this.logger.toast(ToastLevel.ERROR, error.error.error, 'deleteRedirection(' + id + ') ERROR')
          },
          complete: () => this.logger.info('createRedirection() DONE')
        })
    } catch (error) {
      this.logger.error(error);
    }
  }

  public incrementRedirectionView(id: number): Observable<OkResponse<Redirection>> {
    const url = this.redirectionUrl + '/' + id;
    return this.http.put<OkResponse<Redirection>>(url, {}, httpOptions);
  }

  public resetRedirectionView(id: number): Observable<OkResponse<Redirection>> {
    const url = this.redirectionUrl + '/' + id;
    return this.http.patch<OkResponse<Redirection>>(url, {}, httpOptions);
  }
}
