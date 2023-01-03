import axios, { AxiosInstance } from "axios";
import { GetConfig } from "../config";


const isClient = typeof window !== 'undefined';
const config = GetConfig();
class ApiService {

    service: AxiosInstance
    constructor() {
        const service = axios.create();
        this.service = service;
    }
    get(path: string) {
        let url = isClient ? config.API_EXTERNAL_URL : config.API_INTERNAL_URL;
        const resp = this.service.get(`${config.API_PROTOCOOL}://${url}:${config.API_PORT}/api/v1/${path}`, {
            headers: {
                "Accepts": "application/json",
            }
        }).then((resp) => {
            return resp
        });
        return resp;
    }
}

const singleton = new ApiService();
export default singleton;
