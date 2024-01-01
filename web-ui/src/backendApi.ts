import axios, { InternalAxiosRequestConfig } from "axios";

import {
  Configuration,
  ConfigurationParameters,
} from "./generated/configuration";
import { DefaultApiFactory } from "./generated/api";
import { backendAPIEndpoint } from "./appConfig";
import { useAuthUserStore } from "./stores/authUser";

const axiosInstance = axios.create({
  baseURL: backendAPIEndpoint,
  timeout: 5000,
});

// Use interceptor to inject the token to requests
axiosInstance.interceptors.request.use(
  async (config: InternalAxiosRequestConfig) => {
    const authUserStore = useAuthUserStore();

    if (authUserStore.credentials) {
      const authString = `${authUserStore.credentials.username}:${authUserStore.credentials.password}`;
      const authBase64 = btoa(
        String.fromCodePoint(...new TextEncoder().encode(authString)),
      );
      config.headers["Authorization"] = `Basic ${authBase64}`;
    }
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
