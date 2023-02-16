import http from 'k6/http';
import {sleep} from 'k6';

// read ENDPOINT_URL from environment variables
const ENDPOINT_URL = __ENV.ENDPOINT_URL;

export default function () {
    if (!ENDPOINT_URL) {
        console.error("ENDPOINT_URL is not set");
        return;
    }
    let res = http.get(ENDPOINT_URL);
    if (res.status != 200) {
        console.error(`Request failed with status code ${res.status}\nResponse body: ${res.body}`);
    }
    sleep(0.1);
}
