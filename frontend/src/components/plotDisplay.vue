<template>

  <div v-if="loading" class="w-full mb-6 mx-auto">
    <v-skeleton-loader class="mt-6 mb-6" loading-text="loading..." type="list-item, image" />
    <v-skeleton-loader loading-text="loading..." type="list-item, image" />
  </div>

  <div v-if="pdfUrl && !error" class="w-full mb-6">
    <h3 class="text-xl font-medium mb-2">PDF Preview</h3>
    <div class="overflow-hidden rounded-lg shadow-md border border-gray-200">
      <iframe class="w-full h-96" frameborder="0" :src="pdfUrl" />
    </div>
  </div>

  <!-- Empty State for PDF Preview -->
  <div v-if="!error && !loading && pdfUrl == null" class="w-full mb-6 flex items-center justify-center h-96">
    <v-empty-state
      text="Enter a math expression and click Generate Plot to see the result here."
      title="No plot generated yet"
    />
  </div>

  <div class="mt-8 w-100 px-10">
    <v-card v-if="error && errorMessage" color="red-lighten-2" variant="outlined">
      <v-card-title class="flex items-center bg-red-lighten-2 p-4">
        <v-icon class="mr-2" size="30">mdi-alert-circle-outline</v-icon>
        <span class="font-semibold text-lg">Error</span>
      </v-card-title>
      <v-divider />
      <v-card-text class="flex items-center p-4">
        <div class="text-md">
          {{ errorMessage }}
        </div>
      </v-card-text>
    </v-card>
  </div>
  <!-- LaTeX Source Code -->
  <div v-if="latexCode && !error" class="w-full mb-6">
    <h3 class="text-xl font-medium mb-2">LaTeX Source Code</h3>
    <div
      class="p-4 rounded-lg shadow-md border"
      style="max-height: 300px; overflow-y: auto;"
    >
      <pre class="whitespace-pre-wrap text-sm">{{ latexCode }}</pre>
    </div>
  </div>

</template>

<script setup lang="ts">

interface Props {
  pdfUrl: string | null
  latexCode: string | null
  loading: boolean
  error: boolean
  errorMessage: string | null
}

const {
  pdfUrl = null,
  latexCode = null,
  loading = false,
  error = false,
  errorMessage = null,
} = defineProps<Props>()

</script>
