import {Router} from "@angular/router";

export class RoutingUtils {
  private static NOT_FOUND_URL = '/error/notFound';
  private static UNAUTHORIZED_URL = '/error/unauthorized';
  private static HISTORY_URL = '/history';
  private static ROOT = '/';

  public static goToHomepage(router: Router): void {
    router.navigateByUrl(RoutingUtils.ROOT);
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
}
