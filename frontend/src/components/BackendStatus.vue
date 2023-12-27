<template>
  Backend status: 
  <template v-if="props.status === true">
    Up
  </template>
  <template v-else-if="props.status === false">
    Down
  </template>
  <template v-else>
    Unknown
  </template>
</template>

<script setup lang="ts">

const props = defineProps<{
  status: boolean|undefined;
}>();

</script>
