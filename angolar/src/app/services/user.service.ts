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
import {UserInfo} from "../models/user-info";

const httpOptions = {
  headers: new HttpHeaders({'Content-Type': 'application/json'})
};

@Injectable({
  providedIn: 'root'
})
export class UserService {

  private readonly backendUrl?: string
  private readonly userUrl: string;

  constructor(private propertiesService: PropertiesService, logger: LoggerService, private http: HttpClient) {
    this.backendUrl = this.propertiesService.backendUrl;
    this.userUrl = this.backendUrl + '/api/user';
  }

  login(email: string, password: string): Observable<JWTResponse> {
    const loginInput: UserLoginInput = {
      email,
      password
    };
    return this.http.post<JWTResponse>(this.userUrl + '/login', loginInput, httpOptions);
  }

  register(email: string, name: string, password: string): Observable<OkResponse<User>> {
    const userInput: UserInput = {
      email,
      name,
      password
    };
    return this.http.post<OkResponse<User>>(this.userUrl + '/register', userInput, httpOptions);
  }

  getUserList(): Observable<OkResponse<UserInfo[]>> {
    return this.http.get<OkResponse<UserInfo[]>>(this.userUrl + '/list', httpOptions);
  }

  getUserInfo(): Observable<OkResponse<UserInfo>> {
    return this.http.get<OkResponse<UserInfo>>(this.userUrl + '/info', httpOptions);
  }

  deleteUser(): Observable<OkResponse<String>> {
    return this.http.delete<OkResponse<String>>(this.userUrl + '/delete', httpOptions);
  }

  editUser(email: string, name: string, password: string): Observable<JWTResponse> {
    const userInput: UserInput = {
      email,
      name,
      password
    };
    return this.http.post<JWTResponse>(this.userUrl + '/edit', userInput, httpOptions);
  }
}
