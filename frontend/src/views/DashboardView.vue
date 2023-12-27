<template>
  <h2>Dashboard</h2>
  <BackendStatus :status="backendStatus"/>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue';
import { api } from '../backendApi';
import { HttpStatusCode } from 'axios';

import BackendStatus from '../components/BackendStatus.vue';

const backendStatus = ref(undefined as boolean|undefined);

const BACKEND_STATUS_REFRESH_INTERVAL = 1*1000;

let backendStatusPollId: ReturnType<typeof setInterval>|undefined = undefined;

onMounted(() => {

  backendStatusPollId = setInterval(async () => {
    try {
      const result = await api.health();
      if (result.status == HttpStatusCode.Ok) {
        backendStatus.value = true;
      } else {
        backendStatus.value = false;
      }
    } catch (error) {
      backendStatus.value = false;
    }
  }, BACKEND_STATUS_REFRESH_INTERVAL);
});

onUnmounted(() => {
  clearInterval(backendStatusPollId);
  backendStatusPollId = undefined;
});
</script>
