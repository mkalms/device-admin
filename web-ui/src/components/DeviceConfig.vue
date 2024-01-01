<template>
  <v-text-field
    label="IP Address"
    :model-value="clonedDeviceConfigState.ipAddress"
    @update:model-value="
      (newValue) => (clonedDeviceConfigState.ipAddress = newValue)
    "
  ></v-text-field>
  <v-text-field
    label="Netmask"
    :model-value="clonedDeviceConfigState.netMask"
    @update:model-value="
      (newValue) => (clonedDeviceConfigState.netMask = newValue)
    "
  ></v-text-field>
  <v-text-field
    label="Bit rate"
    type="number"
    :model-value="clonedDeviceConfigState.bitRate"
    @update:model-value="
      (newValue) =>
        (clonedDeviceConfigState.bitRate = newValue as unknown as number)
    "
  ></v-text-field>
  <v-select
    label="Duplex"
    :model-value="clonedDeviceConfigState.duplex"
    :items="duplexItems"
    @update:model-value="
      (newValue) => (clonedDeviceConfigState.duplex = newValue)
    "
  ></v-select>
  <v-btn color="primary" @click="$emit('update', clonedDeviceConfigState)"
    >Update</v-btn
  >
  <v-btn @click="cancel()">Cancel</v-btn>
</template>

<script setup lang="ts">
import { ref, toRaw } from "vue";

import { DeviceConfigState } from "../services/deviceConfig";
import { GetDeviceConfigResponseDuplexEnum } from "../generated/api";

const duplexItems = ref([
  {
    title: "Half",
    value: GetDeviceConfigResponseDuplexEnum.Half,
  },
  {
    title: "Full",
    value: GetDeviceConfigResponseDuplexEnum.Full,
  },
]);

const props = defineProps<{
  deviceConfigState: DeviceConfigState;
}>();

defineEmits<{
  (e: "update", value: DeviceConfigState): void;
}>();

const clonedDeviceConfigState = ref({} as DeviceConfigState);

function copyReferenceToClone() {
  clonedDeviceConfigState.value = structuredClone(
    toRaw(props.deviceConfigState),
  );
}

function cancel() {
  copyReferenceToClone();
}

function initialize() {
  copyReferenceToClone();
}

initialize();
</script>
