import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {PropertiesService} from "./properties.service";
import {Observable} from "rxjs";
import {OkResponse} from "../responses/ok-response";
import {HistoryEntry} from "../models/history-entry";

const httpOptions = {
  headers: new HttpHeaders({'Content-Type': 'application/json'})
};

@Injectable({
  providedIn: 'root'
})
export class HistoryService {

  private readonly backendUrl?: string;
  private readonly historyUrl: string;

  constructor(private http: HttpClient,
              private propertiesService: PropertiesService) {
    this.backendUrl = this.propertiesService.backendUrl;
    this.historyUrl = this.backendUrl + '/api/history';
  }

  public getHistory(): Observable<OkResponse<HistoryEntry[]>> {
    return this.http.get<OkResponse<HistoryEntry[]>>(this.historyUrl, httpOptions);
  }

  public resetHistory(): Observable<OkResponse<string>> {
    const url = this.historyUrl + '/delete';
    return this.http.delete<OkResponse<string>>(url, httpOptions);
  }
}
