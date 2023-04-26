import {Router} from "@angular/router";

export class RoutingUtils {
  private static NOT_FOUND_URL = '/error/notFound';
  private static UNAUTHORIZED_URL = '/error/unauthorized';
  private static HISTORY_URL = '/history';
  private static LOGIN_URL = '/login';
  private static REGISTER_URL = '/register';
  private static USER_INFO_URL = '/user/info';
  private static USER_EDITION_URL = '/user/edit';
  private static USER_PASSWORD_EDITION_URL = "/user/edit/password";
  private static ROOT = '/';


  private static EMPTY = '';

  public static goToHomepage(router: Router): void {
    router.navigateByUrl(RoutingUtils.ROOT);
  }

  public static goToLoginPage(router: Router, returnUrl: string): void {
    if (returnUrl !== RoutingUtils.EMPTY) {
      router.navigate([RoutingUtils.LOGIN_URL], {
        queryParams: {returnUrl: returnUrl}
      });
    } else {
      router.navigateByUrl(RoutingUtils.LOGIN_URL);
    }
  }

  public static goToRegisterPage(router: Router): void {
    router.navigateByUrl(RoutingUtils.REGISTER_URL);
  }

  public static goToUserEditionPage(router: Router): void {
    router.navigateByUrl(RoutingUtils.USER_EDITION_URL);
  }

  public static goToPasswordEditionPage(router: Router): void {
    router.navigateByUrl(RoutingUtils.USER_PASSWORD_EDITION_URL);
  }

  public static goToNotFoundPage(router: Router): void {
    router.navigateByUrl(RoutingUtils.NOT_FOUND_URL);
  }

  public static goToUnauthorizedPage(router: Router): void {
    router.navigateByUrl(RoutingUtils.UNAUTHORIZED_URL);
  }

  public static goToHistoryPage(router: Router): void {
    router.navigateByUrl(RoutingUtils.HISTORY_URL);
  }

  public static goToUserInformationPage(router: Router): void {
    router.navigateByUrl(RoutingUtils.USER_INFO_URL);
  }
}
