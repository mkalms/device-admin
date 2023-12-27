import axios, { InternalAxiosRequestConfig } from "axios";

import {
  Configuration,
  ConfigurationParameters,
} from "./generated/configuration";
import { DefaultApiFactory } from "./generated/api";
import { backendAPIEndpoint } from "./appConfig";

const axiosInstance = axios.create({
  baseURL: backendAPIEndpoint,
  timeout: 5000,
});

// Use interceptor to inject the token to requests
axiosInstance.interceptors.request.use(
  async (config: InternalAxiosRequestConfig) => {

    // Add hardcoded auth header for test:1234
    // TODO: replace this with actual tracking of logged-in identity
    config.headers["Authorization"] = "Basic dGVzdDoxMjM0";
    return config;
  },
);

const apiConfigurationParameters = {
  basePath: backendAPIEndpoint,
} as ConfigurationParameters;

const apiConfiguration = new Configuration(apiConfigurationParameters);

export const api = DefaultApiFactory(
  apiConfiguration,
  undefined,
  axiosInstance,
);
