<template>

  <v-card
    border="sm"
    class="pa-4"
    density="compact"
    :elevation="0.5"
    min-height="380px"
    rounded="lg"
    width="100%"
  >

    <div v-if="props.loading" class="w-full mb-6 mx-auto">
      <v-skeleton-loader class="mt-6 mb-6" loading-text="loading..." type="list-item, image" />
      <v-skeleton-loader loading-text="loading..." type="list-item, image" />
    </div>

    <div v-if="pdfUrl && !props.error" class="w-full mb-6">
      <div class="overflow-hidden rounded-lg shadow-md border border-gray-200">
        <iframe class="w-full h-96" frameborder="0" :src="pdfUrl" />
      </div>
    </div>

    <!-- Empty State for PDF Preview -->
    <div v-if="!props.error && !props.loading && pdfUrl == null" class="pa-4 h-full">
      <v-card
        v-if="!props.error && !props.loading && pdfUrl == null"
        class="d-flex align-center justify-center"
        color="surface-variant"
        height="100%"
        variant="flat"
        width="100%"
      >
        <v-card-text class="text-center">Plot output will be displayed here</v-card-text>
      </v-card>
    </div>
    <div v-if="props.error && props.errorMessage" class="mt-8 w-100 px-10">
      <v-card color="red-lighten-2" variant="outlined">
        <v-card-title class="flex items-center bg-red-lighten-2 p-4">
          <v-icon class="mr-2" size="30">mdi-alert-circle-outline</v-icon>
          <span class="font-semibold text-lg">Error</span>
        </v-card-title>
        <v-divider />
        <v-card-text class="flex items-center p-4">
          <div class="text-md">
            {{ props.errorMessage }}
          </div>
        </v-card-text>
      </v-card>
    </div>
    <!-- LaTeX Source Code -->
    <div v-if="props.latexCode && !props.error" class="w-full">
      <div class="shiki-container">
        <v-tooltip location="left" :text="downloaded ? 'Downloaded' : 'Download'">
          <template #activator="{ props: tprops }">
            <v-icon
              v-bind="tprops"
              class="download-button"
              :class="{'copied': downloaded}"
              icon="mdi-download-box-outline"
              size="large"
              @click="downloadCodeSnippet"
            />

          </template>
        </v-tooltip>
        <span class="lang-name text-caption">Latex</span>
        <v-tooltip v-model="tooltipOpen" location="left" :target="copyIcon!" text="Copied">
          <template #activator>
            <v-icon
              ref="copyIcon"
              class="copy-button"
              :class="{'copied': copied}"
              size="large"
              @click="copyToClipboard"
            >
              {{ copied ? 'mdi-clipboard-check-outline' : 'mdi-clipboard-outline' }}
            </v-icon>
          </template>
        </v-tooltip>

        <pre v-html="highlightedCode" />
      </div>
    </div>

  </v-card>
</template>

<script setup lang="ts">
import { highlightLaTeX } from "@/lib/utils/shiki"

interface Props {
  pdfUrl: string | null
  latexCode: string | null
  loading: boolean
  error: boolean
  errorMessage: string | null
}
const highlightedCode = ref("")
const copied = ref(false)
const downloaded = ref(false)
const tooltipOpen = ref(false)

const props = defineProps<Props>()
const copyIcon = ref(null)
const copyToClipboard = () => {
  if (props.latexCode) {
    navigator.clipboard
      .writeText(props.latexCode)
      .then(() => {
        copied.value = true // Set copied state to true
        tooltipOpen.value = true // Show tooltip
        setTimeout(() => {
          copied.value = false // Reset copied state after 2 seconds
          tooltipOpen.value = false // Hide tooltip
        }, 1000)
      })
      .catch(err => {
        console.error("Failed to copy text: ", err)
      })
  }
}
const downloadCodeSnippet = () => {
  const fileContent = props.latexCode || ""
  const ts = Date.now()
  const fileName = `plot-snippet-${ts}.tex`
  const blob = new Blob([fileContent], { type: "text/x-latex" })
  const a = document.createElement("a")
  a.href = URL.createObjectURL(blob)
  a.download = fileName
  downloaded.value = true
  a.click()
  URL.revokeObjectURL(a.href)
  a.remove()
  setTimeout(() => {
    downloaded.value = false
  }, 1000)
}

watch(() => props.latexCode, async () => {
  if (props.latexCode) {
    const { html } = await highlightLaTeX(props.latexCode)
    highlightedCode.value = html
  }
})

</script>

<style>
pre.shiki {
  font-size: 1rem;
    padding: 1.25rem;
    border-radius: 0.5rem;
    overflow: auto;
}
.shiki-container{
  position: relative;
}
.shiki-container:hover .copy-button  {
  opacity: 1; /* Fully opaque when parent is hovered */
  transform: translateY(0); /* Reset transform to original position */

}
.shiki-container:hover .download-button  {
  opacity: 1; /* Fully opaque when parent is hovered */
  transform: translateY(0); /* Reset transform to original position */

}
.copy-button {
  position: absolute;
  top: 0.5rem;
  right: 0.5rem;
  color: rgb(var(--v-theme-surface-variant));
  opacity: 0;
  transform: translateY(-10px); /* Optional: slide up from 10px */
  transition: opacity 0.3s ease, transform 0.3s ease; /* Smooth transition */
}
.download-button {
  position: absolute;
  top: 0.5rem;
  right: 2.5rem;
  color: rgb(var(--v-theme-surface-variant));
  opacity: 0;
  transform: translateY(-10px); /* Optional: slide up from 10px */
  transition: all 0.3s ease, transform 0.3s ease; /* Smooth transition */
}

.lang-name {
  position: absolute;
  top: 0.5rem;
  right: 0.5rem;
  background-color: rgb(var(--v-theme-badge-bg-color));
  color: rgb(var(--v-theme-surface-badge-text));
  transition: opacity 0.3s ease, transform 0.3s ease; /* Smooth transition */
  border-radius: 5px;
  padding: 1px 6px;
}
.shiki-container:hover .lang-name {
  opacity: 0; /* Fully opaque when parent is hovered */
  transform: translateY(0); /* Reset transform to original position */

}

.w-full.relative:hover .copy-button {
  display: inline-flex; /* Show button on hover */
}

.copy-button.copied {
  color: rgb(var(--v-theme-success))
}
.download-button.copied {
  color: rgb(var(--v-theme-success))
}
</style>
