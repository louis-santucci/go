import { Injectable } from '@angular/core';

@Injectable({
    providedIn: 'root'
})
export class PropertiesService {

    private _backendUrl: string = "https://angolar";
    private _angolarSecretKey: string = 'lkMzojzKlshozdgZeidjbfbrShgdisgFHHzysiztDsyzhghejfvrjgvbchxgsywgfquzysoedirfyruoYghshxUvsh';

    constructor() { }


    get backendUrl(): string {
        return this._backendUrl;
    }


    get angolarSecretKey(): string {
        return this._angolarSecretKey;
    }
}
