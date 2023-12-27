<template>
  <v-app>
    <!-- Display main UI when user is logged in -->

    <template v-if="isLoggedIn()">
      <v-app-bar color="primary">
        <v-app-bar-title>STB admin panel</v-app-bar-title>

        <v-spacer></v-spacer>

        <v-btn @click="logout">Logout</v-btn>
      </v-app-bar>

      <v-navigation-drawer permanent expand-on-hover rail>
        <v-list density="compact" nav>
          <v-list-item
            prepend-icon="mdi-speedometer"
            title="Dashboard"
            :to="{ name: 'dashboard' }"
            link
          />
          <v-list-item
            prepend-icon="mdi-cog"
            title="Device configuration"
            :to="{ name: 'deviceConfig' }"
            link
          />
        </v-list>
      </v-navigation-drawer>

      <v-main>
        <v-container fluid>
          <router-view />
        </v-container>
      </v-main>
    </template>

    <!-- Display a spinner when the app doesn't yet know the user's login status -->

    <template v-else-if="isLoginStateUnknown()">
      <v-container class="fill-height" fluid>
        <v-row justify="center">
          <v-progress-circular
            indeterminate
            color="primary"
          ></v-progress-circular>
        </v-row>
      </v-container>
    </template>

    <!-- Display a login prompt when user is logged out -->

    <template v-else>
      <v-container class="fill-height" fluid>
        <v-row justify="center">
          <LoginView/>
        </v-row>
      </v-container>
    </template>
  </v-app>
</template>

<script setup lang="ts">

import { useAuthUserStore, LoginState } from "./stores/authUser";

import LoginView from "./views/LoginView.vue";

const authUserStore = useAuthUserStore();

function logout(): void {
  authUserStore.clearCredentials();
}

function isLoggedIn(): boolean {
  return authUserStore.credentials != undefined;
}

function isLoginStateUnknown(): boolean {
  return authUserStore.loginState == LoginState.Unknown;
}
</script>
