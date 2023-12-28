<template>
  <v-dialog v-model="dialog" persistent max-width="600px" min-width="360px">
    <v-card class="px-4" min-height="280">
      <v-card-text>
        <v-form ref="loginForm" v-model="valid" lazy-validation>
          <v-row>
            <v-col cols="12">
              <v-text-field
                v-model="loginUsername"
                label="Username"
                :rules="[required]"
                required
              ></v-text-field>
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12">
              <v-text-field
                v-model="loginPassword"
                :append-icon="show1 ? 'mdi-eye-off' : 'mdi-eye'"
                :rules="[required]"
                :type="show1 ? 'text' : 'password'"
                name="input-10-1"
                label="Password"
                @click:append="show1 = !show1"
              ></v-text-field>
            </v-col>
          </v-row>
          <v-row>
            <v-col class="d-flex" cols="12" sm="6" xsm="12"> </v-col>
            <v-spacer></v-spacer>
            <v-col class="d-flex" cols="12" sm="3" xsm="12" align-end>
              <v-btn
                x-large
                block
                :disabled="!valid"
                :loading="loading"
                color="primary"
                @click="login"
              >
                Login
              </v-btn>
            </v-col>
          </v-row>
        </v-form>
      </v-card-text>
    </v-card>
  </v-dialog>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { api } from "../backendApi";
import { Credentials, useAuthUserStore } from "../stores/authUser";
import { LoginRequest } from "../generated/api";

const dialog = ref(true);
const valid = ref(false);
const loginUsername = ref("");
const loginPassword = ref("");
const show1 = ref(false);

const loading = ref(false);

function required(value: any): boolean | string {
  return !!value || "Field is required";
}

const authUserStore = useAuthUserStore();

async function login() {
  loading.value = true;
  try {
    const loginRequest = {
      username: loginUsername.value,
      password: loginPassword.value,
    } as LoginRequest;
    await api.login(loginRequest);

    loading.value = false;
    const credentials = {
      username: loginUsername.value,
      password: loginPassword.value,
    } as Credentials;
    authUserStore.setCredentials(credentials);
  } catch (error) {
    console.log("Login failed");
    loading.value = false;
    authUserStore.clearCredentials();
  }
}
</script>
