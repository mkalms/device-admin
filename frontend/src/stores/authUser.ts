import { defineStore } from "pinia";

export enum LoginState {
  Unknown,
  LoggedOut,
  LoggedIn,
}

export interface Credentials {
  username: string,
  password: string,
}

interface State {
  credentials: Credentials|null;
  loginState: LoginState;
}

export const useAuthUserStore = defineStore("authUser", {
  state: (): State => ({
    credentials: null,
    loginState: LoginState.LoggedOut, // TODO: this should be Unknown, and update to logged in/out when loaded
  }),

  actions: {
    setCredentials(credentials: Credentials) {
      this.credentials = credentials;
      this.loginState = LoginState.LoggedIn;
    },

    clearCredentials() {
      this.credentials = null;
      this.loginState = LoginState.LoggedOut;
    },
  },
});
