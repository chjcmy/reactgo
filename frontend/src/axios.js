const axios = require("axios");

export const instance = axios.create({
    baseURL: 'http://choi1994.iptime.org:8000/'
});

