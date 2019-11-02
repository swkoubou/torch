import {messages} from "~/proto/web";

export default class Api {
  private static getUrl(endpoint: string): string {
    return '/api' + endpoint;
  }

  static addSpot(form: FormData): Promise<any> {
    const u = this.getUrl('/spotInfo/create');
    return fetch(u, {
      method: "POST",
      body: form
    }).then((res: any) => {
      return res.text();
    });
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

  static getPins(): Promise<messages.AllSpotsResponse> {
    const u = this.getUrl('/spotInfo/get/all');
    return fetch(u, {
      method: 'GET',
    }).then(res => {
      return res.arrayBuffer();
    }).then((buf: ArrayBuffer) => {
      const _buf = new Uint8Array(buf);
      return messages.AllSpotsResponse.decode(_buf);
    })
  }

  static getPinInfo(id: number): Promise<messages.SpotInfoResponse> {
    const req = new messages.SpotInfoRequest();
    req.spotId = id;
    const bin = messages.SpotInfoRequest.encode(req).finish();

    const u = this.getUrl('/spotInfo/get');
    return fetch(u, {
      method: 'POST',
      body: bin,
    }).then(res => {
      return res.arrayBuffer();
    }).then((buf: ArrayBuffer) => {
      const _buf = new Uint8Array(buf);
      return messages.SpotInfoResponse.decode(_buf);
    })
  }
}
