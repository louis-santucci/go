import {Injectable} from '@angular/core';
import {PropertiesService} from "./properties.service";
import {LoggerService} from "./logger.service";
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {Observable} from "rxjs";
import {UserLoginInput} from "../dtos/user/user.login.input";
import {UserInput} from "../dtos/user/user.input";
import {OkResponse} from "../responses/ok-response";
import {User} from "../models/user";
import {JWTResponse} from "../responses/jwtresponse";

const httpOptions = {
  headers: new HttpHeaders({'Content-Type': 'application/json'})
};

@Injectable({
  providedIn: 'root'
})
export class UserService {

  private readonly backendUrl?: string
  private readonly authUrl: string;

  constructor(private propertiesService: PropertiesService, logger: LoggerService, private http: HttpClient) {
    this.backendUrl = this.propertiesService.backendUrl;
    this.authUrl = this.backendUrl + '/api/user';
  }

  login(email: string, password: string): Observable<JWTResponse> {
    const loginInput: UserLoginInput = {
      email,
      password
    };
    return this.http.post<JWTResponse>(this.authUrl + '/login', loginInput, httpOptions);
  }

  register(email: string, name: string, password: string): Observable<OkResponse<User>> {
    const userInput: UserInput = {
      email,
      name,
      password
    };
    return this.http.post<OkResponse<User>>(this.authUrl + '/register', userInput, httpOptions);
  }
}
