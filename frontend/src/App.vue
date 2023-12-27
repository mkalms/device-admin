<template>
  <v-app>
    <v-app-bar color="primary">
      <v-app-bar-title>STB admin panel</v-app-bar-title>
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
  </v-app>
</template>

<script setup lang="ts">

</script>
