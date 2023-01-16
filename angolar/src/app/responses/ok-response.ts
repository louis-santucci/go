import {Response} from "./response";

export interface OkResponse<T> extends Response {
  data: T;
}
