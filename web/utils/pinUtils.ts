import {structs} from "~/proto/web";

export default class PinUtils {
  static convertUiPin(pin: structs.ISpotInfo | undefined | null): any {
    if (pin === undefined || pin == null) {
      return
    }

    const location = pin.location;
    if (location === undefined || location == null) {
      return
    }
    const geoPos = {
      lat: location.latitude,
      lon: location.longitude
    };

    let domClass = 'disabled';

    const eventSpan = pin.eventSpan;
    if (eventSpan === undefined || eventSpan == null) {
      return
    }
    const start = this.convertITimeToDateTime(eventSpan.startingAt).getTime();
    const end = this.convertITimeToDateTime(eventSpan.endingAt).getTime();
    const now = new Date().getTime();
    
    if (now > start && now < end) {
      domClass = 'active';

      if (pin.hotScore > 40 && pin.hotScore < 60) {
        domClass = 'hot1'
      } else if (pin.hotScore < 90) {
        domClass = 'hot2'
      } else {
        domClass = 'hot3'
      }
    }

    const id = pin.spotId;
    if (id === undefined || id == null) {
      return
    }

    return {
      id: id,
      geoPos: geoPos,
      class: domClass,
    }
  }

  static convertITimeToDateTime(time: structs.ITime | undefined | null): Date {
    if (time === undefined || time == null) {
      return new Date();
    }

    return new Date(time.year + '-' + time.month + '-' + time.day + ' ' + time.hour + ':' + time.minute + ':' + time.second)
  }
}
