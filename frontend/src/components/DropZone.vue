<template>
  <div
    ref="dropZoneRef"
    :class="[
      'dropzone',
      isOverDropZone ? 'active-dropzone' : '',
    ]"
    @click="openFileDialog"
  >
    <span>Drag or Drop a .raw File</span>
    <span>OR</span>
    <v-btn
      class="text-none"
      prepend-icon="mdi-upload"
      variant="text"
      color="indigo-darken-3"
      border
      @click.stop="openFileDialog"
    >
      Select File
    </v-btn>
    <input type="file" accept=".raw, .bin" ref="fileInputRef" @change="handleFileChange" hidden />
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useDropZone } from '@vueuse/core'

const filesData = ref<{ name: string, size: number, type: string, lastModified: number }[]>([])

function onDrop(files: File[] | null) {
  filesData.value = []
  if (files) {
    filesData.value = files.slice(0, 1).map(file => ({
      name: file.name,
      size: file.size,
      type: file.type,
      lastModified: file.lastModified,
    }))
  }
}

const dropZoneRef = ref<HTMLElement | null>(null)
const fileInputRef = ref<HTMLInputElement | null>(null)

const { isOverDropZone } = useDropZone(dropZoneRef, {
  onDrop,
  dataTypes: ['.raw', '.RAW', '.bin']
})

const openFileDialog = () => {
  fileInputRef.value?.click()
}

function handleFileChange(event: Event) {
  const input = event.target as HTMLInputElement
  if (input.files && input.files.length > 0) {
    onDrop(Array.from(input.files))
  }
}
</script>

<style scoped>
.dropzone {
  @apply w-96 h-48 flex flex-col justify-center items-center gap-4 border-2 border-dashed border-gray-400 bg-white transition-all duration-300;
}

.dropzone.active-dropzone {
  @apply border-blue-500 bg-blue-50;
}

.dropzone label {
  @apply px-4 py-2 text-white bg-blue-500 transition-all duration-300 cursor-pointer;
}

.dropzone.active-dropzone label {
  @apply bg-blue-600;
}
</style>
