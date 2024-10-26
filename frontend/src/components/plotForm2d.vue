<!-- eslint-disable vue/html-quotes -->
<template>
  <v-form ref="formRef" v-model="isFormValid">
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
            :value="Tabs.Expression"
          >
            Expressions
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
          <v-tabs-window-item :value="Tabs.Expression">
            <div v-for="(expr, index) in formState.plots" :key="index" class="mb-3">

              <v-text-field
                :id="'expression-' + index"
                v-model="expr.expression"
                density="compact"
                :label="`Expression ${index + 1}`"
                rounded="md"
                :rules="[expressionValidationRule]"
                single-line
                variant="outlined"
              >
                <template #prepend-inner>
                  <span class="whitespace-nowrap" v-html="`f<sub>${index + 1}</sub>(x) =`" />
                </template>

                <template v-if="formState.plots.length > 1" #append>
                  <v-btn
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
              <div
                :class="{ 'flex flex-row-reverse -mt-4': true, 'px-14': formState.plots.length > 1, 'px-2': !(formState.plots.length < 1) }"
              >

                <v-menu :close-on-content-click="false">
                  <template #activator="{ props }">
                    <v-sheet
                      v-bind="props"
                      class="rounded cursor-pointer"
                      :color="expr.color"
                      height="22px"
                      width="22px"
                    />
                  </template>

                  <v-color-picker v-model="expr.color" mode="hex" />
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
                            <div
                              class="w-full h-1"
                              :style="lineStyles.find(style => style.value === expr.line_style)?.style || 'border-bottom: 2px solid;'"
                            />
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
                        <template #activator="{ props }">
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
              :disabled="formState.plots.length >= MAX_EXPRS"
              prepend-icon="mdi-plus"
              variant="text"
              @click="addExpression"
            >
              Add Expression
            </v-btn>

            <div class="text-sm mb-5">
              <p>Examples of valid expressions:</p>
              <ul class="list-disc pl-5">
                <li>x^2 + 3*x - 5</li>
                <li>log10(x) + e^x - ln(x) + sin(deg(x))</li>
                <li>(x &lt; 0) * x^2 + (x>= 0) * (x + 1) </li>
              </ul>
            </div>

          </v-tabs-window-item>
          <v-tabs-window-item :value="Tabs.Config">
            <div class="rounded-lg">
              <v-row class="mb-0 pb-0" no-gutters>
                <v-col class="mb-1" cols="12">
                  <v-text-field
                    v-model="formState.title"
                    density="compact"
                    label="Plot Title"
                    rounded="md"
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
                    rounded="md"
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
                    rounded="md"
                    :rules="[validLabel]"
                    variant="outlined"
                  />

                </v-col>
                <v-col class="pr-4 mb-1" cols="6">
                  <v-number-input
                    v-model="formState.x_min"
                    dense
                    density="compact"
                    label="X-Axis Min"
                    rounded="md"
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
                    rounded="md"
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
                    rounded="md"
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
                    rounded="md"
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
                    rounded="md"
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
                    rounded="md"
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
                  <v-select
                    v-model="formState.grid"
                    class="pr-4"
                    dense
                    density="compact"
                    item-title="title"
                    item-value="value"
                    :items="Object.values(Grid)"
                    label="Grid"
                    rounded="md"
                    variant="outlined"
                  />
                </v-col>
                <v-col cols="6">
                  <v-number-input
                    v-model="lineWidth"
                    dense
                    density="compact"
                    label="Line Width"
                    :max="12"
                    :min="1"
                    rounded="md"
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
              <v-row class="mb-0 pb-0" no-gutters>

                <v-col>

                  <v-checkbox v-model="legend" class="-ml-2 -mt-2" hide-details label="Legend" />
                </v-col>
              </v-row>
              <v-row v-if="legend" class="mb-0 pb-0 mt-0" no-gutters>
                <v-col>

                  <v-select
                    v-if="legend"
                    v-model="formState.legend"
                    class="mr-4"
                    dense
                    density="compact"
                    item-title="title"
                    item-value="value"
                    :items="legendPositionItems"
                    label="Legend Position"
                    rounded="md"
                    variant="outlined"
                  />
                </v-col>
                <v-col v-for="(expr, index) in formState.plots" :key="index" cols="6">
                  <v-text-field
                    v-model="expr.legend"
                    :class="{ 'mr-4': index % 2 != 0 }"
                    clearable
                    dense
                    density="compact"
                    rounded="md"
                    variant="outlined"
                  >
                    <template #label>
                      <span v-html="`Legend For f<sub>${index + 1}</sub>(x)`" />
                    </template>
                  </v-text-field>

                </v-col></v-row>
            </div>

          </v-tabs-window-item>
        </v-tabs-window>

      </v-card-item>
    </v-card>

    <!-- Generate Plot Button -->
    <v-btn
      block
      class="mt-4 text-none text-lg font-medium"
      color="primary"
      :disabled="!isFormValid"
      :loading="loading"
      prepend-icon="mdi-play"
      rounded="md"
      @click="generateExpression"
    >
      Generate Plot

    </v-btn>
  </v-form>
</template>

<script setup lang="ts">
import { AxisLines, Grid, LegendPosition, PlotRequest } from "@/lib/models/plot"
import { validateExpression } from "@/lib/utils/expressionValidator"
import { mergeProps, ref } from "vue"
import { VForm } from "vuetify/components"

interface Props {
  loading: boolean
}

const MAX_EXPRS = 6

const enum Tabs {
  Expression = "expression",
  Config = "config",
}

const { loading = false } = defineProps<Props>()
const isFormValid = ref(true)
const tab = ref(Tabs.Expression)
const emit = defineEmits<{(e: "submit", formData: PlotRequest): void }>()
const colorPalette = [
  "#1f77b4",
  "#2ca02c",
  "#d62728",
  "#9467bd",
  "#8c564b",
  "#e377c2",
]

const lineStyles = [
  { value: "solid", label: "Solid", style: "border-bottom: 2px solid;" },
  { value: "dashed", label: "Dashed", style: "border-bottom: 2px dashed;" },
  { value: "dotted", label: "Dotted", style: "border-bottom: 2px dotted;" },
  { value: "dashdotdotted", label: "Dash-Dot-Dotted", style: "border-bottom: 2px dashed; border-bottom-style: dashdotdotted;" },
]

let colorIndex = 1
const legend = ref(false)
const formState = ref<PlotRequest>({
  title: "",
  x_label: "",
  y_label: "",
  x_min: undefined,
  x_max: undefined,
  y_min: undefined,
  y_max: undefined,
  grid: Grid.None,
  axis_lines: AxisLines.Box,
  border: 5,
  plots: [{ expression: "", color: colorPalette[0], line_style: "solid", line_width: "1pt" }],
})

const formRef = ref<VForm | null>(null)
const lineWidth = ref(1)

const axisLinesItemis = [
  { value: AxisLines.None, title: "none" },
  { value: AxisLines.Box, title: "box" },
  { value: AxisLines.Center, title: "center" },
  { value: AxisLines.Left, title: "left" },
  { value: AxisLines.Middle, title: "middle" },
  { value: AxisLines.Right, title: "right" },
]

const legendPositionItems = [
  { value: LegendPosition.NorthEast, title: "Up Right" },
  { value: LegendPosition.NorthWest, title: "Up Left" },
  { value: LegendPosition.SouthEast, title: "Bottom Right" },
  { value: LegendPosition.SouthhWest, title: "Bottom Left" },
  { value: LegendPosition.OuterNorthEast, title: "Outer up right" },
]

const expressionValidationRule = (expression: string): true | string => {
  if (expression.trim() === "") {
    return "Expression cannot be empty"
  }
  const result = validateExpression(expression)
  return result.ok ? true : result.error
}

const validNum = (value: string | undefined) => {
  if (value === undefined) {
    return true
  }
  return !isNaN(Number(value)) || "Must be a number"
}

const validTitle = (value: string) => {
  if (value === "") {
    return true
  }
  if (value.length > 128) {
    return "Label is too long"
  }

  return true
}

const validLabel = (value: string) => {
  if (value.length > 32) {
    return "Label is too long"
  }
  return true
}

const validPadding = (value: number): boolean | string => {
  const num = Number(value)
  if (isNaN(num)) {
    return "Must be a number"
  }
  if (num < 0 || num > 50) {
    return "Padding must be between 0 and 50"
  }
  return true
}
const validLineWidth = (value: number): boolean | string => {
  const num = Number(value)
  if (isNaN(num)) {
    return "Must be a number"
  }
  if (num < 1 || num > 10) {
    return "Padding must be between 0 and 50"
  }
  return true
}

const addExpression = () => {
  const nextColor = colorPalette[colorIndex % colorPalette.length]
  formState.value.plots.push({ expression: "", color: nextColor, line_style: "solid", line_width: `${lineWidth.value}pt` })
  colorIndex++
  focusExpressionField()
}

const generateExpression = () => {
  if (isFormValid.value) {
    emit("submit", formState.value)
  }
}

const removeExpression = (index: number) => {
  if (formState.value.plots.length > 1) {
    formState.value.plots.splice(index, 1)
    colorIndex = formState.value.plots.length % colorPalette.length
  }
}

const focusExpressionField = () => {
  nextTick(() => {
    const lastIndex = formState.value.plots.length - 1
    const lastFieldId = "expression-" + lastIndex
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

watch(legend, enabled => {
  if (!enabled) {
    formState.value.legend = LegendPosition.None
    formState.value.plots.map(plot => plot.legend = "")
  } else {
    formState.value.legend = LegendPosition.NorthEast
    formState.value.plots.map((plot, index) => plot.legend = `f${index}(x)`)
  }
})
</script>

<style scoped>
.line-style-preview {
  height: 2px;
  width: 100%;
  margin: 4px 0;
}

.selected-tab {
  background-color: rgba(var(--v-theme-background), 1);
  color: var(--v-theme-background);
  font-weight: 700;
}

/* Override CSS variables for the expansion panel */
.v-expansion-panel-title {
  padding: 0px 16px;
}

.v-expansion-panel-text__wrapper {
  padding: 0px;
  background-color: red;
}
</style>
