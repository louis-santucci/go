import {Injectable} from "@angular/core";
import {HttpClient} from "@angular/common/http";
import {Config} from "./models/config";

@Injectable()
export class AppConfig {
  private config: Config | null;

  constructor(private http: HttpClient) {
    this.config = null;
  }

  public loadConfig() {
    return this.http.get<Config>('./assets/properties.json')
      .subscribe(res => {
        this.config = res;
      })
  }

  public getConfig() {
    return this.config;
  }
}
