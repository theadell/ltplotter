<template>
  <v-app>
    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script lang="ts" setup>
import { useDark } from "@vueuse/core"
import { useTheme } from "vuetify"

const theme = useTheme()

const isDark = useDark({
  selector: "html",
  attribute: "class",
  valueDark: "dark",
  valueLight: "light",
  onChanged: isDark => {
    theme.global.name.value = isDark ? "dark" : "light"
  },
})

onBeforeMount(() => {
  const savedTheme = localStorage.getItem("vueuse-color-scheme")
  if (savedTheme) {
    if (savedTheme === "dark") {
      isDark.value = true
      theme.global.name.value = "dark"
    } else if (savedTheme === "light") {
      isDark.value = false
      theme.global.name.value = "light"
    } else if (savedTheme === "auto") {
      const systemPrefersDark = window.matchMedia("(prefers-color-scheme: dark)").matches
      isDark.value = systemPrefersDark
      theme.global.name.value = systemPrefersDark ? "dark" : "light"
    }
  } else {
    const systemPrefersDark = window.matchMedia("(prefers-color-scheme: dark)").matches
    isDark.value = systemPrefersDark
    theme.global.name.value = systemPrefersDark ? "dark" : "light"
  }
})

</script>

<style>
html.dark .shiki,
html.dark .shiki span {
  color: var(--shiki-dark) !important;
  background-color: var(--shiki-dark-bg) !important;
  font-style: var(--shiki-dark-font-style) !important;
  font-weight: var(--shiki-dark-font-weight) !important;
  text-decoration: var(--shiki-dark-text-decoration) !important;
}
</style>
