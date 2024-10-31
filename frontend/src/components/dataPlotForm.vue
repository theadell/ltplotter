<template>
  <v-card
    border="sm"
    class="mb-4 pa-2"
    density="compact"
    :elevation="0.5"
    min-height="380px"
    rounded="lg"
  >
    <v-card-item>
      <v-tabs
        v-model="tab"
        align-tabs="center"
        bg-color="surface-variant"
        class="rounded-lg pa-1 mb-4"
        density="compact"
        grow
        height="44px"
      >
        <v-tab
          active-color="background"
          class="rounded text-none"
          density="compact"
          height="34px"
          rounded="lg"
          selected-class="selected-tab"
          :value="Tabs.Values"
        >
          Values
        </v-tab>
        <v-tab
          active-color="background"
          class="rounded text-none"
          density="compact"
          height="34px"
          rounded="lg"
          selected-class="selected-tab"
          :value="Tabs.Config"
        >
          Configuration
        </v-tab>
      </v-tabs>

      <v-tabs-window v-model="tab" class="px-0 py-2">
        <v-tabs-window-item :value="Tabs.Values">
          <v-text-field
            v-model="xValueString"
            density="compact"
            label="x values"
            persistent-hint
            single-line
            variant="outlined"
          />
          <v-divider class="mb-5" :thickness="3" />
          <v-text-field
            v-model="yValueString"
            density="compact"
            label="y values"
            persistent-hint
            single-line
            variant="outlined"
          />
          <v-btn
            block
            class="mb-5 text-subtitle-2"
            color="blue-grey-lighten-1"
            prepend-icon="mdi-plus"
            variant="text"
          >
            Add value series
          </v-btn>
        </v-tabs-window-item>
        <v-tabs-window-item :value="Tabs.Config">
          <v-text-field
            v-model="title"
            density="compact"
            label="title"
            persistent-hint
            single-line
            variant="outlined"
          />
          <v-text-field
            v-model="xLabel"
            density="compact"
            label="x label"
            persistent-hint
            single-line
            variant="outlined"
          />
          <v-text-field
            v-model="yLabel"
            density="compact"
            label="y label"
            persistent-hint
            single-line
            variant="outlined"
          />
          <v-text-field
            v-model="legends"
            density="compact"
            label="legend1, legend2, ..."
            persistent-hint
            single-line
            variant="outlined"
          />
        </v-tabs-window-item>
      </v-tabs-window>
    </v-card-item>

  </v-card>
  <v-btn
    block
    class="mt-4 text-none text-lg font-medium"
    color="primary"
    :loading="loading"
    prepend-icon="mdi-play"
    rounded="md"
    @click="generatePlot"
  >
    Generate Plot
  </v-btn>
</template>

<script setup lang="ts">

import { ref } from "vue"
import { DataPlotRequest } from "@/lib/models/plot"

const enum Tabs {
  Values = "values",
  Config = "config",
}

const tab = ref(Tabs.Values)

const xValueString = ref<string>("")
const yValueString = ref<string>("")
const loading = ref<boolean>(false)
const emit = defineEmits<{(e: "submit", formData: DataPlotRequest): void }>()
const title = ref<string>("title")
const xLabel = ref<string>("x label")
const yLabel = ref<string>("y label")
const legends = ref<string>("legend1, legend2, ...")

const parseToFloatArray = (str: string): number[] => {
  return str.split(",").map(s => s.trim()).filter(s => s.length > 0).map(Number)
}

const generatePlot = () => {
  loading.value = true

  const xValues = parseToFloatArray(xValueString.value)
  const yValues = parseToFloatArray(yValueString.value)

  console.log("X Values:", xValues)
  console.log("Y Values:", yValues)

  if (xValues.some(isNaN) || yValues.some(isNaN)) {
    console.error("One or more values could not be converted to a number.")
  } else {
    const funnyDataPlotRequest: DataPlotRequest = {
      x: xValues,
      y: [
        { values: yValues },
      ],
      metadata: {
        title: title.value,
        labels: {
          x: xLabel.value,
          y: yLabel.value,
        },
        legends: legends.value.split(",").map(s => s.trim()).filter(s => s.length > 0),
      },
    }
    emit("submit", funnyDataPlotRequest)
  }

  loading.value = false
}
</script>

<style scoped>
.selected-tab {
  background-color: rgba(var(--v-theme-background), 1);
  color: var(--v-theme-background);
  font-weight: 700;
}
</style>
