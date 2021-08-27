const axios = require('axios');

// eslint-disable-next-line import/prefer-default-export
export const instance = axios.create({
  baseURL: 'http://localhost:8000/',
});
