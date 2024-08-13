<template>
  <v-form ref="formRef" v-model="isFormValid">

    <div>
      <h3 class="text-xl font-medium text-gray-700 mb-4">Math Expressions</h3>
      <div v-for="(expr, index) in formState.plots" :key="index" class="mb-1">
        <v-text-field
          :id="'expression-' + index"
          v-model="expr.expression"
          density="compact"
          :label="`Expression ${index + 1}`"
          :rules="[expressionValidationRule]"
          single-line
          variant="outlined"
        >
          <template #prepend-inner>
            <span class="text-gray-500 whitespace-nowrap">y =</span>
          </template>

          <template #append>
            <v-btn
              v-if="formState.plots.length > 1"
              class="-ml-2"
              color="red-darken-3"
              icon
              size="small"
              variant="text"
              @click="removeExpression(index)"
            >
              <v-icon>mdi-delete-outline</v-icon>
            </v-btn>
          </template>

        </v-text-field>
      </div>

      <v-btn
        block
        class="mb-5 text-none text-subtitle-2"
        color="blue-grey-darken-1"
        prepend-icon="mdi-plus"
        variant="text"
        @click="addExpression"
      >
        Add Expression
      </v-btn>

      <!-- Guide for Valid Expressions -->
      <div class="text-gray-600 text-sm mb-5">
        <p><strong>Examples of valid expressions:</strong></p>
        <ul class="list-disc pl-5">
          <li>x^2 + 3*x - 5</li>
          <li>log10(x) + e^x - ln(x) + sin(deg(x))</li>
          <li>(x &lt; 0) * x^2 + (x>= 0) * (x + 1) </li>
        </ul>
      </div>
    </div>

    <!-- Plot Settings -->
    <div>
      <v-expansion-panels>
        <v-expansion-panel title="Plot Configurations">
          <v-expansion-panel-text>
            <div class="rounded-lg">
              <v-row class="mb-0 pb-0" no-gutters>
                <v-col class="mb-1" cols="12">
                  <v-text-field
                    v-model="formState.title"
                    density="compact"
                    label="Plot Title"
                    :rules="[validTitle]"
                    variant="outlined"
                  />
                </v-col>
                <v-col class="mb-1" cols="6">
                  <v-text-field
                    v-model="formState.x_label"
                    class="pr-4"
                    density="compact"
                    label="X-Axis Label"
                    :rules="[validLabel]"
                    variant="outlined"
                  />
                </v-col>
                <v-col cols="6">
                  <v-text-field
                    v-model="formState.y_label"
                    dense
                    density="compact"
                    label="Y-Axis Label"
                    :rules="[validLabel]"
                    variant="outlined"
                  />

                </v-col>
                <v-col
                  class="pr-4 mb-1"
                  cols="6"
                >
                  <v-number-input
                    v-model="formState.x_min"
                    dense
                    density="compact"
                    label="X-Axis Min"
                    :rules="[validNum]"
                    variant="outlined"
                  >
                    <template #increment />
                    <template #decrement />

                  </v-number-input>
                </v-col>
                <v-col class="mb-1" cols="6">
                  <v-number-input
                    v-model="formState.x_max"
                    dense
                    density="compact"
                    label="X-Axis Max"
                    :rules="[validNum]"
                    variant="outlined"
                  >
                    <template #increment />
                    <template #decrement />

                  </v-number-input>
                </v-col>
                <v-col class="pr-4 mb-1" cols="6">
                  <v-number-input
                    v-model="formState.y_min"
                    dense
                    density="compact"
                    label="Y-Axis Min"
                    :rules="[validNum]"
                    variant="outlined"
                  >
                    <template #increment />
                    <template #decrement />
                  </v-number-input>
                </v-col>
                <v-col cols="6">
                  <v-number-input
                    v-model="formState.y_max"
                    dense
                    density="compact"
                    label="Y-Axis Max"
                    :rules="[validNum]"
                    variant="outlined"
                  >
                    <template #increment />
                    <template #decrement />
                  </v-number-input>
                </v-col>
                <v-col class="pr-4" cols="6">
                  <v-select
                    v-model="formState.axis_lines"
                    dense
                    density="compact"
                    item-title="title"
                    item-value="value"
                    :items="axisLinesItemis"
                    label="Axis Lines"
                    variant="outlined"
                  />
                </v-col>
                <v-col cols="6">
                  <v-number-input
                    v-model="formState.border"
                    dense
                    density="compact"
                    label="Padding Around"
                    :max="50"
                    :min="0"
                    :rules="[validPadding]"
                    variant="outlined"
                  >
                    <template #increment />
                    <template #decrement />
                    <template #append-inner>
                      <span class="text-gray-500 whitespace-nowrap mr-4">mm</span>
                    </template>
                  </v-number-input>
                </v-col>
                <v-col cols="6">
                  <v-checkbox v-model="formState.grid" class="mt-0" density="compact" label="Show Grid" />
                </v-col>

              </v-row>
            </div>
          </v-expansion-panel-text>
        </v-expansion-panel>
      </v-expansion-panels>
    </div>

    <v-alert v-if="error" class="my-4 w-full" closable type="error">
      Something went wrong
    </v-alert>

    <!-- Generate Plot Button -->
    <v-btn
      block
      class="mt-4 py-3 text-lg font-medium text-white"
      color="#7986CB"
      :disabled="!isFormValid"
      :loading="loading"
      prepend-icon="mdi-auto-fix"
      variant="elevated"
      @click="generateExpression"
    >
      Generate Plot
    </v-btn>
  </v-form>
</template>

<script setup lang="ts">
import { AxisLines, PlotRequest } from '@/lib/models/plot'
import { validateExpression } from '@/lib/utils/expressionValidator'

import { ref } from 'vue'
import { VForm } from 'vuetify/components'

interface Props {
  loading: boolean
  error: boolean
}

const { loading = false, error = false } = defineProps<Props>()
const isFormValid = ref(true)
const emit = defineEmits<{(e: 'submit', formData: PlotRequest): void }>()

const formState = ref<PlotRequest>({
  title: '',
  x_label: '',
  y_label: '',
  x_min: undefined,
  x_max: undefined,
  y_min: undefined,
  y_max: undefined,
  grid: false,
  axis_lines: AxisLines.BOX,
  border: 5,
  plots: [{ expression: '' }],
})

const formRef = ref<VForm | null>(null)

const axisLinesItemis = [
  { value: AxisLines.BOX, title: 'box' },
  { value: AxisLines.CENTER, title: 'center' },
  { value: AxisLines.LEFT, title: 'left' },
  { value: AxisLines.MIDDLE, title: 'middle' },
  { value: AxisLines.RIGHT, title: 'right' },
]

const expressionValidationRule = (expression: string): true | string => {
  if (expression.trim() === '') {
    return 'Expression cannot be empty'
  }
  const result = validateExpression(expression)
  return result.ok ? true : result.error
}

const validNum = (value: string | undefined) => {
  if (value === undefined) {
    return true
  }
  return !isNaN(Number(value)) || 'Must be a number'
}

const validTitle = (value: string) => {
  if (value === '') {
    return true
  }
  if (value.length > 128) {
    return 'Label is too long'
  }

  return true
}

const validLabel = (value: string) => {
  if (value.length > 32) {
    return 'Label is too long'
  }
  return true
}

const validPadding = (value: number): boolean | string => {
  const num = Number(value)
  if (isNaN(num)) {
    return 'Must be a number'
  }
  if (num < 0 || num > 50) {
    return 'Padding must be between 0 and 50'
  }
  return true
}

const addExpression = () => {
  formState.value.plots.push({ expression: '' })
  focusExpressionField()
}

const generateExpression = () => {
  if (isFormValid.value) {
    emit('submit', formState.value)
  }
}

const removeExpression = (index : number) => {
  if (formState.value.plots.length > 1) {
    formState.value.plots.splice(index, 1)
  }
}

const focusExpressionField = () => {
  nextTick(() => {
    const lastIndex = formState.value.plots.length - 1
    const lastFieldId = 'expression-' + lastIndex
    const lastField = document.getElementById(lastFieldId)
    if (lastField) {
      lastField.focus()
    }
  })
}

</script>
