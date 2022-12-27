import axios, { AxiosInstance } from "axios";


const isClient = () => typeof window !== 'undefined';

class ApiService {

    service: AxiosInstance
    constructor() {
        const service = axios.create();
        this.service = service;
    }
}

const singleton = new ApiService();
export default singleton;
