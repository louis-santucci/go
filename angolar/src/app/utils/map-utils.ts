import {Redirection} from "../models/redirection";

export class MapUtils {
  public static deleteById(map: Map<string, Redirection>, id: number): Map<string, Redirection> {
    for (let [key, value] of map) {
      if (value.id === id) {
        map.delete(key);
        return map;
      }
    }
    return map;
  }
}
