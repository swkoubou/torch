import * as geolib from 'geolib';
import {structs} from "~/proto/web";

export default class GeoUtils {
  public static start = {
    lat: 35.48832,
    lon: 139.34024,
  };
  public static end = {
    lat: 35.48491,
    lon: 139.34596,
  };

  static convertPos(lat: number, lon: number): any {
    let pointX = lon * 20037508.34 / 180.0;
    let pointY = Math.log(Math.tan((90.0 + lat) * Math.PI / 360.0)) / (Math.PI / 180.0);
    pointY = pointY * 20037508.34 / 180.0;

    return {
      x: pointX,
      y: pointY,
    }
  }

  static convertPosFromPx(x: number, y: number): any {
    let lon = x * 180 / 20037508.34;
    let lat = Math.atan(Math.exp(y * Math.PI / 20037508.34)) * 360 / Math.PI - 90;

    return {
      lat: lat,
      lon: lon,
    }
  }

  static containArea(areas: structs.AreaInfo[], lat: number, lon: number): structs.AreaInfo | undefined {
    let centerWithName = new Map<string, structs.AreaInfo>();
    const centerArray: Array<any> = [];

    areas.forEach((v: structs.AreaInfo) => {
      if (v.region === null || v.region === undefined) {
        return;
      }
      const leftTop = v.region.leftUp;
      const rightBottom = v.region.rightBottom;

      if (leftTop === null || rightBottom === null || typeof leftTop !== "object" || typeof rightBottom !== "object" ||
        typeof leftTop.latitude !== "number" || typeof leftTop.longitude !== "number"
        || typeof rightBottom.latitude !== "number" || typeof rightBottom.longitude !== "number") {
        return;
      }

      const center = geolib.getCenter([{
        latitude: leftTop.latitude,
        longitude: leftTop.longitude
      }, {
        latitude: rightBottom.latitude,
        longitude: rightBottom.longitude
      }]);
      if (!center) {
        return;
      }

      const key: string = center.latitude + '_' + center.longitude;
      centerWithName.set(key, v);
      centerArray.push(center);
    });

    const mostNear: any = geolib.findNearest({
      latitude: lat,
      longitude: lon
    }, centerArray);
    if (typeof mostNear === "undefined") {
      return;
    }

    const distance = geolib.getPreciseDistance({
      latitude: mostNear.latitude,
      longitude: mostNear.longitude,
    }, {
      latitude: lat,
      longitude: lon
    });
    const distanceM = geolib.convertDistance(distance, 'm');
    // 500M以上離れていればエリアにいないと判定
    if (distanceM > 500) {
      return;
    }

    const key = mostNear.latitude + '_' + mostNear.longitude;
    const areaInfo = centerWithName.get(key);
    if (typeof areaInfo !== "undefined") {
      return areaInfo;
    } else {
      return;
    }
  }
}


