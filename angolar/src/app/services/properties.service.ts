import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Config} from "../models/config";

@Injectable({
  providedIn: 'root'
})
export class PropertiesService {

  private config?: Config;
  private loaded = false;

  constructor(private http: HttpClient) {
  }

  public loadConfig(): Promise<void> {
    return this.http.get<Config>('/assets/properties.json')
      .toPromise()
      .then(data => {
        this.config = data;
        this.loaded = true;
      })
  }

  public getConfig(): Config | undefined {
    return this.config;
  }
}
