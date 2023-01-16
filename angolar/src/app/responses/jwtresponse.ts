import {Response} from "./response";

export interface JWTResponse extends Response {
  token: string;
}
