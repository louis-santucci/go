export enum WeekDay {
  MON,
  TUE,
  WED,
  THU,
  FRI,
  SAT,
  SUN
}

export enum Month {
  JAN,
  FEB,
  MAR,
  APR,
  MAY,
  JUN,
  JUL,
  AUG,
  SEP,
  OCT,
  NOV,
  DEC
}

export class DateUtils {
  public static CleanDate(currentDate: Date): string {
    const day = WeekDay[currentDate.getDay()];
    const month = Month[currentDate.getMonth()];
    const date = currentDate.getDate();
    const year = currentDate.getFullYear();
    const hours = currentDate.getHours();
    const minutes = currentDate.getMinutes();
    const seconds = currentDate.getSeconds();
    let hoursStr: string;
    let minutesStr: string;
    let secondsStr: string;
    if (hours < 10) {
      hoursStr = '0' + hours;
    } else {
      hoursStr = hours.toString();
    }
    if (minutes < 10) {
      minutesStr = '0' + minutes;
    } else {
      minutesStr = minutes.toString();
    }
    if (seconds < 10) {
      secondsStr = '0' + seconds;
    } else {
      secondsStr = seconds.toString();
    }

    return day + ' ' + month + ' ' + date + ' ' + year + ' ' + hoursStr + ':' + minutesStr + ':' + secondsStr;
  }
}
