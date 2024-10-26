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
import { defineComponent, ref } from "vue"
import {PlotRequest} from "@/lib/models/plot";

// Define props and emits can be used here if you had any, but they're not in your original script
// const props = defineProps<{ /* your props here */ }>()
// const emit = defineEmits<{(e: 'submit', formData: DataPlotRequest): void}>()

// Reactive references for the input strings
const xValueString = ref<string>("")
const yValueString = ref<string>("")
const loading = ref<boolean>(false)
const emit = defineEmits<{(e: "submit", formData: PlotRequest): void }>()

// Convert string to array of floats
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
        title: "The Rise and Fall of Sitcom Hilarity",
        labels: {
          x: "Years of Laugh Tracks",
          y: "Giggle Quotient",
        },
        legends: ["LPM (Laughs Per Minute)", "PPE (Puns Per Episode)"],
      },
    }
    console.log(funnyDataPlotRequest)
    emit("submit", funnyDataPlotRequest)
  }

  loading.value = false
}
</script>

<style scoped>
</style>
