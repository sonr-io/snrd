import { AxiosInstance } from "axios";
export declare function getBase64FromBytes(bytes: number[] | Uint8Array): string;
export declare function getBase64WebEncodingFromBytes(bytes: number[] | Uint8Array): string;
export declare function getBytesFromBase64(str: string): Uint8Array;
export declare function arrayBufferEncode(value: ArrayBuffer): string;
export declare function arrayBufferDecode(value: string): ArrayBuffer;
export declare function baseUrl(): "http://0.0.0.0:8080" | "https://highway.build";
export declare function getAxios(authenticated: boolean): AxiosInstance;
export declare function publicKeyCredentialAttestionToJson(credential: PublicKeyCredential): string;
export declare function publicKeyCredentialAssertionToJson(credential: PublicKeyCredential): string;
