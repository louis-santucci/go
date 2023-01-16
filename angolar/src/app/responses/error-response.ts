import {Response} from './response';

export interface ErrorResponse extends Response {
  error: string;
}
