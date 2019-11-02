import {messages} from "~/proto/web";

export default class Api {
  static getUrl(endpoint: string): string {
    return '/api' + endpoint;
  }

  static getAreas(): Promise<messages.AllAreaInfoResponse> {
    const u = this.getUrl('/areaInfo/get/all');
    return fetch(u, {
      method: 'GET',
    }).then(res => {
      return res.arrayBuffer();
    }).then((buf: ArrayBuffer) => {
      const _buf = new Uint8Array(buf);
      return messages.AllAreaInfoResponse.decode(_buf);
    })
  }
}
