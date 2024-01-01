<template>
  <h2>Device Configuration</h2>
  <template v-if="deviceConfigState !== undefined">
    <DeviceConfig
      :device-config-state="deviceConfigState"
      @update="(newConfig) => update(newConfig)"
    />
  </template>
  <template v-else>
    <!-- Display progress bar until the device configuration has been loaded -->
    <v-progress-circular indeterminate color="primary"></v-progress-circular>
  </template>
</template>

<script setup lang="ts">
import { ref } from "vue";
import {
  DeviceConfigState,
  getDeviceConfig,
  setDeviceConfig,
} from "../services/deviceConfig";
import DeviceConfig from "../components/DeviceConfig.vue";

const deviceConfigState = ref(undefined as DeviceConfigState | undefined);

async function update(newConfig: DeviceConfigState) {
  deviceConfigState.value = newConfig;
  try {
    await setDeviceConfig(newConfig);
  } catch (error) {
    console.log("Error when calling setDeviceConfig: " + error);
  }
}

async function initialize() {
  try {
    let deviceConfig = await getDeviceConfig();
    deviceConfigState.value = deviceConfig;
  } catch (error) {
    console.log("Error when calling getDeviceConfig: " + error);
  }
}

initialize();
</script>
