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
    let lon = x *  180 / 20037508.34 ;
    let lat = Math.atan(Math.exp(y * Math.PI / 20037508.34)) * 360 / Math.PI - 90;

    return {
      lat: lat,
      lon: lon,
    }
  }

}


