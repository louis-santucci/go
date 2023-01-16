import {Injectable} from '@angular/core';
import {ToastLevel} from "../models/toast-level";
import {ToastrService} from "ngx-toastr";

@Injectable({
  providedIn: 'root'
})
export class LoggerService {

  constructor(private toastr: ToastrService) {
  }

  log(message: any): void {
    console.log(message);
  }

  error(message: any): void {
    console.error(message);
  }

  info(message: any): void {
    console.info(message);
  }

  toast(level: ToastLevel, message: any, title: any) {
    switch (level) {
      case ToastLevel.SUCCESS:
        this.toastr.success(JSON.stringify(message), JSON.stringify(title));
        break;
      case ToastLevel.ERROR:
        this.toastr.error(JSON.stringify(message), JSON.stringify(title));
        break;
      case ToastLevel.INFO:
        this.toastr.info(JSON.stringify(message), JSON.stringify(title));
        break;
      case ToastLevel.WARN:
        this.toastr.warning(JSON.stringify(message), JSON.stringify(title));
        break;
    }
  }
}
