<template>
  <div class="mt-20">
    <v-empty-state :image="notFoundImgSrc">
      <template #title>
        <div class="mt-8">
          Whoops, 404
        </div>
      </template>

      <template #text>
        <div class="text-subtitle-2">
          The page you were looking for does not exist.
        </div>
      </template>

      <template #actions>
        <v-btn color="primary" to="/">Go Back to Home</v-btn>
      </template>
    </v-empty-state>
  </div>
</template>
<route lang="yaml">
    meta:
      layout: 404
</route>

<script setup lang="ts">
import notFoundImgSrc from '@/assets/undraw_page_not_found.svg'

</script>
