<!-- eslint-disable vue/html-quotes -->
<template>
  <v-form ref="formRef" v-model="isFormValid">

    <div>
      <h3 class="text-xl font-medium mb-4">Math Expressions</h3>
      <div v-for="(expr, index) in formState.plots" :key="index" class="mb-3">

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
            <span class="whitespace-nowrap">y =</span>
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
        <div class="flex flex-row-reverse px-10 -mt-4">

          <v-menu
            :close-on-content-click="false"
          >
            <template #activator="{ props }">
              <v-sheet
                v-bind="props"
                class="rounded cursor-pointer"
                :color="expr.color"
                height="22px"
                width="22px"
              />
            </template>

            <v-color-picker
              v-model="expr.color"
              mode="hex"
            />
          </v-menu>

          <v-menu>
            <template #activator="{ props: menu }">
              <v-tooltip location="top">
                <template #activator="{ props: tooltip }">

                  <v-sheet
                    v-bind="mergeProps(menu, tooltip)"
                    class="mx-4 cursor-pointer "
                    height="24px"
                    variant="outlined"
                    width="30px"
                  >
                    <div class="w-full h-full flex items-center justify-center">
                      <div class="w-full h-1" :style="lineStyles.find(style => style.value === expr.line_style)?.style || 'border-bottom: 2px solid;'" />
                    </div>
                  </v-sheet>
                </template>
                <span>Set line style</span>
              </v-tooltip>

            </template>

            <v-list density="compact" width="100px">
              <v-list-item
                v-for="style in lineStyles"
                :key="style.value"
                :color="expr.color"
                @click="expr.line_style = style.value"
              >
                <v-tooltip location="right">
                  <template #activator="{props}">
                    <div v-bind="props" :style="style.style" />
                  </template>
                  <span> {{ style.label }}</span>
                </v-tooltip>
              </v-list-item>
            </v-list>
          </v-menu>

        </div>
      </div>

      <v-btn
        block
        class="mb-5 text-subtitle-2"
        color="blue-grey-lighten-1"
        prepend-icon="mdi-plus"
        variant="text"
        @click="addExpression"
      >
        Add Expression
      </v-btn>

      <!-- Guide for Valid Expressions -->
      <div class="text-sm mb-5">
        <p>Examples of valid expressions:</p>
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
                <v-col cols="6">
                  <v-number-input
                    v-model="lineWidth"
                    dense
                    density="compact"
                    label="Line Width"
                    :max="12"
                    :min="1"
                    :rules="[validLineWidth]"
                    variant="outlined"
                  >
                    <template #increment />
                    <template #decrement />
                    <template #append-inner>
                      <span class="text-gray-500 whitespace-nowrap mr-4">pt</span>
                    </template>
                  </v-number-input>
                </v-col>
              </v-row>
            </div>
          </v-expansion-panel-text>
        </v-expansion-panel>
      </v-expansion-panels>
    </div>

    <!-- Generate Plot Button -->
    <v-btn
      block
      class="mt-4 text-none text-lg font-medium"
      color="primary"
      :disabled="!isFormValid"
      :loading="loading"
      prepend-icon="mdi-auto-fix"
      @click="generateExpression"
    >
      Generate Plot
    </v-btn>
  </v-form>
</template>

<script setup lang="ts">
import { AxisLines, PlotRequest } from '@/lib/models/plot'
import { validateExpression } from '@/lib/utils/expressionValidator'

import { mergeProps, ref } from 'vue'
import { VForm } from 'vuetify/components'

interface Props {
  loading: boolean
}

const { loading = false } = defineProps<Props>()
const isFormValid = ref(true)
const emit = defineEmits<{(e: 'submit', formData: PlotRequest): void }>()
const colorPalette = [
  '#1f77b4',
  '#2ca02c',
  '#d62728',
  '#9467bd',
  '#8c564b',
  '#e377c2',
]

const lineStyles = [
  { value: 'solid', label: 'Solid', style: 'border-bottom: 2px solid;' },
  { value: 'dashed', label: 'Dashed', style: 'border-bottom: 2px dashed;' },
  { value: 'dotted', label: 'Dotted', style: 'border-bottom: 2px dotted;' },
  { value: 'dashdotdotted', label: 'Dash-Dot-Dotted', style: 'border-bottom: 2px dashed; border-bottom-style: dashdotdotted;' },
]

let colorIndex = 1

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
  plots: [{ expression: '', color: colorPalette[0], line_style: 'solid', line_width: '1pt' }],
})

const formRef = ref<VForm | null>(null)
const lineWidth = ref(1)

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
const validLineWidth = (value: number): boolean | string => {
  const num = Number(value)
  if (isNaN(num)) {
    return 'Must be a number'
  }
  if (num < 1 || num > 10) {
    return 'Padding must be between 0 and 50'
  }
  return true
}

const addExpression = () => {
  const nextColor = colorPalette[colorIndex % colorPalette.length]
  formState.value.plots.push({ expression: '', color: nextColor, line_style: 'solid', line_width: `${lineWidth.value}pt` })
  colorIndex++
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
    colorIndex = formState.value.plots.length % colorPalette.length
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
watch(lineWidth, (n, _) => {
  formState.value.plots.forEach(plot => {
    plot.line_width = `${n}pt`
  })
})
</script>

<style scoped>
.line-style-preview {
  height: 2px;
  width: 100%;
  margin: 4px 0;
}
</style>
