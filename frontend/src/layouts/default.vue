<template>
  <v-app>
    <v-app-bar
      class="border-b border-gray-200 dark:border-gray-700"
      density="comfortable"
      elevate-on-scroll
      flat
    >
      <v-app-bar-nav-icon class="mr-4 md:hidden" @click="drawer = !drawer">
        <v-icon>mdi-menu</v-icon>
      </v-app-bar-nav-icon>

      <v-app-bar-title class="text-lg font-medium">
        LTPplotter
      </v-app-bar-title>

      <!-- Flex container to align the links and the theme toggle button -->
      <div class="flex-grow" />

      <div class="hidden md:flex items-center space-x-2">
        <v-btn class="text-none" to="/">
          Symbolic
        </v-btn>
        <v-btn class="text-none" to="/numerical">
          Numerical
        </v-btn>
        <v-btn class="text-none" to="/ltspice">
          LTSpice
        </v-btn>
      </div>

      <!-- Divider -->
      <v-divider class="mx-10 my-2.5 hidden md:flex" thickness="2" vertical />

      <!-- Theme Toggle Button -->
      <v-btn icon @click="toggleTheme">
        <v-icon>{{ themeIcon }}</v-icon>
      </v-btn>
    </v-app-bar>    <!-- Navigation Drawer for Mobile -->
    <v-navigation-drawer
      v-model="drawer"
      app
      location="start"
      mobile
      temporary
      @click-outside="drawer = false"
    >
      <v-list class="mt-4" density="compact" nav>
        <v-list-item
          prepend-icon="mdi-function-variant"
          title="Symbolic Plotting"
          to="/"
          value="home"
          @click="drawer = false"
        />
        <v-list-item
          prepend-icon="mdi-scatter-plot"
          title="Numerical Plotting"
          to="/numerical"
          value="account"
          @click="drawer = false"
        />
        <v-list-item
          prepend-icon="mdi-lambda"
          title="LTSpice"
          to="/ltspice"
          value="users"
          @click="drawer = false"
        />
      </v-list>

    </v-navigation-drawer>
    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script lang="ts" setup>
import { useTheme } from 'vuetify'

const drawer = ref(false)

const theme = useTheme()
const themeIcon = ref('mdi-weather-sunny')

// Function to toggle the theme
function toggleTheme () {
  theme.global.name.value = theme.global.current.value.dark ? 'light' : 'dark'
  themeIcon.value = theme.global.current.value.dark ? 'mdi-weather-night' : 'mdi-weather-sunny'
}

</script>

<style scoped>
.v-divider {
  border-color: rgb(var(--v-theme-on-surface)) !important;
}
</style>

