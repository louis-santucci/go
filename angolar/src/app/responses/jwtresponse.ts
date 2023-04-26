import {Response} from "./response";
import {UserInfo} from "../models/user-info";

export interface JWTResponse extends Response {
  token: string;
  user: UserInfo;
}
