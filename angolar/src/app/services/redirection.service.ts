import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {BehaviorSubject, Observable} from "rxjs";
import {Redirection} from "../models/redirection";
import {OkResponse} from "../responses/ok-response";
import {PropertiesService} from "./properties.service";
import {LoggerService} from "./logger.service";
import {ToastLevel} from "../models/toast-level";
import {RedirectionInput} from "../dtos/redirection/redirection.input";
import {AlertService} from "./alert.service";
import {MapUtils} from "../utils/map-utils";

@Injectable({
  providedIn: 'root'
})
export class RedirectionService {

  private readonly backendUrl?: string
  private readonly redirectionUrl: string;

  // Observable source
  private redirectionSource = new BehaviorSubject<Redirection | undefined>(undefined);
  private redirectionMapSource = new BehaviorSubject<Map<string, Redirection> | undefined>(undefined)

  // Observable source
  private redirectionObservable: Observable<Redirection | undefined> = this.redirectionSource.asObservable();
  private redirectionMapObservable: Observable<Map<string, Redirection> | undefined> = this.redirectionMapSource.asObservable();

  constructor(private http: HttpClient,
              private propertiesService: PropertiesService,
              private logger: LoggerService,
              private alertService: AlertService) {
    this.backendUrl = this.propertiesService.backendUrl;
    this.redirectionUrl = this.backendUrl + '/api/redirection';
  }

  public getRedirectionObservable(): Observable<Redirection | undefined> {
    return this.redirectionObservable;
  }

  public getRedirectionMapObservable(): Observable<Map<string, Redirection> | undefined> {
    return this.redirectionMapObservable;
  }

  public getRedirections(): void {
    try {
      this.http.get<OkResponse<Redirection[]>>(this.redirectionUrl)
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
            this.logger.toast(ToastLevel.ERROR, error.message, 'getRedirections() ERROR');
          },
          complete: () => this.logger.info('getRedirections() DONE')
        })
    } catch (error) {
      this.logger.error(error);
    }
  }

  public getRedirection(id: number): void {
    try {
      const url = this.redirectionUrl + '/' + id;
      this.http.get<OkResponse<Redirection>>(url)
        .subscribe({
          next: res => {
            this.logger.log({status: res.status, data: res.data});
            if (res.status === 200) {
              this.redirectionSource.next(res.data);
            }
          },
          error: error => {
            this.logger.error(error);
            this.logger.toast(ToastLevel.ERROR, error.error.error, 'getRedirection(' + id + ') ERROR');
          },
          complete: () => this.logger.info('getRedirection(' + id + ') DONE')
        })
    } catch (error) {
      this.logger.error(error);
    }
  }

  public incrementRedirectionView(id: number): void {
    try {
      const url = this.redirectionUrl + '/' + id;
      this.http.put<OkResponse<Redirection>>(url, null)
        .subscribe({
          next: res => {
            this.logger.log({status: res.status, data: res.data});
            if (res.status === 200) {
              const currentMap = this.redirectionMapSource.getValue();
              if (currentMap !== undefined) {
                const newRedirection = res.data;
                currentMap.set(newRedirection.shortcut, newRedirection);
                this.redirectionMapSource.next(currentMap);
              }
            }
          },
          error: error => {
            this.logger.error(error);
            this.logger.toast(ToastLevel.ERROR, error, 'incrementRedirectionView(' + id + ') ERROR');
          },
          complete: () => this.logger.info('incrementRedirectionView(' + id + ') DONE')
        })
    } catch (error) {
      this.logger.error(error);
    }
  }

  public resetRedirectionView(id: number): void {
    try {
      const url = this.redirectionUrl + '/' + id;
      this.http.patch<OkResponse<Redirection>>(url, null)
        .subscribe({
          next: res => {
            this.logger.log({status: res.status, data: res.data});
            if (res.status === 200) {
              const currentMap = this.redirectionMapSource.getValue();
              if (currentMap !== undefined) {
                const newRedirection = res.data;
                currentMap.set(newRedirection.shortcut, newRedirection);
                this.redirectionMapSource.next(currentMap);
              }
            }
          },
          error: err => {
            this.logger.error(err);
            this.logger.toast(ToastLevel.ERROR, err, 'resetRedirectionView(' + id + ') ERROR');
          },
          complete: () => this.logger.info('resetRedirectionView(' + id + ')' + 'DONE')
        })
    } catch (error) {
      this.logger.error(error);
    }
  }

  public editRedirection(id: number, shortcut: string, redirectUrl: string): void {
    try {
      const input: RedirectionInput = {
        shortcut: shortcut,
        redirect_url: redirectUrl
      };
      const url = this.redirectionUrl + '/' + id;
      this.http.post<OkResponse<Redirection>>(url, input)
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
            this.logger.toast(ToastLevel.ERROR, error, 'editRedirection(' + id + ') ERROR')
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
      this.http.post<OkResponse<Redirection>>(this.redirectionUrl, input)
        .subscribe({
          next: res => {
            this.logger.log({status: res.status, data: res.data});
            if (res.status === 200) {
              const currentMap = this.redirectionMapSource.getValue();
              if (currentMap !== undefined) {
                const newRedirection = res.data;
                currentMap.set(newRedirection.shortcut, newRedirection);
                this.redirectionMapSource.next(currentMap);
                this.alertService.success('New Redirection for URL ' + newRedirection.redirect_url + ' created', true);
              }
            }
          },
          error: error => {
            this.logger.error(error);
            this.logger.toast(ToastLevel.ERROR, error, 'createRedirection() ERROR');
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
      this.http.delete<OkResponse<string>>(url)
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
              }
            }
          },
          error: error => {
            this.logger.error(error);
            this.logger.toast(ToastLevel.ERROR, error, 'deleteRedirection(' + id + ') ERROR')
          },
          complete: () => this.logger.info('createRedirection() DONE')
        })
    } catch (error) {
      this.logger.error(error);
    }
  }
}
