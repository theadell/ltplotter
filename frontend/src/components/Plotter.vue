<template>
  <v-container>
    <v-row justify="center">
      <v-col cols="12" md="6">
        <v-file-input v-model="file" label="Upload .raw File" accept=".raw" :rules="[fileRules]" :max-size="20000000"
          single-line density="compact" :loading="loading" @update:modelValue="onUpload" :clearable="false" />
      </v-col>
    </v-row>

    <div v-show="dataSet?.data">
      <v-tabs v-model="tab" bg-color="transparent" color="basil" grow>
        <v-tab v-for="item in items" :key="item" :text="item" :value="item"></v-tab>
      </v-tabs>

      <v-tabs-window v-model="tab">
        <v-tabs-window-item value="Visualization">
          <v-container fluid class="fill-height">
            <v-row align="stretch">
              <!-- Controls Section -->
              <v-col cols="12" md="4">
                <v-card height="100%">
                  <v-card-text class="flex-grow-1">
                    <v-text-field variant="outlined" density="compact" label="Title"
                      @update:modelValue="PlotUtils.updateTitle(chart!, $event)" v-model="graph.title"></v-text-field>
                    <v-text-field variant="outlined" density="compact" label="X-Label"
                      @update:modelValue="PlotUtils.updateXLabel(chart!, $event)"
                      v-model="graph.xAxisLabel"></v-text-field>
                    <v-text-field variant="outlined" density="compact" label="Y-Label"
                      @update:modelValue="PlotUtils.updateYLabel(chart!, $event)"
                      v-model="graph.yAxisLabel"></v-text-field>
                    <v-select variant="outlined" density="compact" label="Select X-Axis Variables"
                      :items="dataSet?.variables" @update:modelValue="updateX" v-model="graph.x"
                      class="w-full"></v-select>
                    <v-select variant="outlined" density="compact" multiple label="Select Y-Axis Variables"
                      :items="dataSet?.variables" @update:modelValue="redraw" class="w-full"
                      v-model="graph.y"></v-select>

                    <v-row>
                      <v-col>
                        <v-checkbox v-model="graph.legendVisible" label="Legend" density="compact" hide-details
                          color="info" value="info"
                          @update:modelValue="PlotUtils.toggleLengend(chart!, $event)"></v-checkbox>
                      </v-col>
                      <v-col>
                        <v-btn-toggle v-if="graph.legendVisible" v-model="graph.legendPos" mandatory flat
                          class="text-none" density="compact">
                          <v-btn class="text-none pa-1" size="xs" value="default">
                            default
                          </v-btn>

                          <v-btn class="text-none pa-1" size="xs" value="center">
                            center
                          </v-btn>

                          <v-btn class="text-none pa-1" size="xs" value="right">
                            right
                          </v-btn>

                          <v-btn class="text-none pa-1" size="xs" value="left">
                            left
                          </v-btn>
                        </v-btn-toggle>

                      </v-col>
                    </v-row>

                    <v-row class="mt-0" v-if="graph.legendVisible">
                      <v-col class="mt-0 pt-0">
                        <v-checkbox density="compact" class="pl-0 ml-0" v-model="graph.customizeLegend"
                          label="Customize Legend Names"></v-checkbox>
                        <div class="ml-4" v-if="graph.customizeLegend">
                          <v-text-field @keyup.enter="PlotUtils.updateLegends(chart!, graph.legendNames)"
                            variant="outlined" density="compact" v-for="(_, index) in graph.y" :key="index"
                            v-model="graph.legendNames[index]" :label="'Legend for ' + graph.y[index]"></v-text-field>
                        </div>
                      </v-col>
                    </v-row>




                  </v-card-text>

                </v-card>
              </v-col>

              <!-- Chart Section -->
              <v-col cols="12" md="8">
                <v-card height="100%">
                  <v-card-text>
                    <div ref="chart"></div>
                  </v-card-text>
                  <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn @click="PlotUtils.exportChart(chart!, ExportFormat.SVG)" class="me-2 text-none"
                      prepend-icon="mdi-export-variant" variant="flat">
                      Download as SVG
                    </v-btn>
                    <v-btn @click="PlotUtils.exportChart(chart!, ExportFormat.PNG)" class="me-2 text-none"
                      prepend-icon="mdi-export-variant" variant="flat">
                      Download as PNG
                    </v-btn>
                  </v-card-actions>

                </v-card>
              </v-col>
            </v-row>
          </v-container> </v-tabs-window-item>
        <v-tabs-window-item value="Data & Export">
          <v-container>
            <v-data-table density="compact" :items="tableData"></v-data-table>
          </v-container>

        </v-tabs-window-item>

      </v-tabs-window>
    </div>
  </v-container>
</template>

<script setup lang="ts">
import { DataSet, isComplexDataSet, PlotState } from "@/lib/models";
import { Api } from "../lib/api";
import { ExportFormat, PlotUtils } from "@/lib/plotly/utils"
const file = ref<File | null>(null);
const tab = ref('Visualization')
const loading = ref(false);
const items = [
  'Visualization',
  'Data & Export',
]

const graph = ref<PlotState>({
  x: '',
  y: [],
  title: '',
  xAxisLabel: '',
  yAxisLabel: '',
  legendVisible: false,
  legendPos: 'default',
  customizeLegend: false,
  legendNames: []
})


const dataSet = ref<DataSet | undefined>(undefined);

const chart = ref<HTMLDivElement | null>(null);
const fileRules = (value: File | null) => {
  if (!value) return "File is required";
  if (value.size > 20000000) return "File size should be less than 20MB";
  return true;
};

const onUpload = async (file: File) => {
  if (file) {
    loading.value = true;
    try {
      const data = await Api.parseRawFile(file);
      dataSet.value = data;
      graph.value.title = file.name
      graph.value.xAxisLabel = 'x'
      graph.value.yAxisLabel = 'y'
      graph.value.x = isComplexDataSet(data) ? 'frequency' : 'time'
      PlotUtils.initEmptyPlot(chart.value!, graph.value.title, graph.value.xAxisLabel, graph.value.yAxisLabel)
    } catch (error) {
      loading.value = false;
      console.error(error);
    } finally {
      loading.value = false;
    }
  }
};

const tableData = computed(() => {
  if (dataSet.value) {
    const d = dataSet.value.data["0"]; // map[variableName][]values >> [{"variableName": value},...]
    const keys = Object.keys(d);
    const length = d[keys[0]].length;

    let result = [];

    for (let i = 0; i < length; i++) {
      let obj = {} as any;
      keys.forEach((key) => {
        obj[key] = d[key][i];
      });
      result.push(obj);
    }

    return result;
  }
});



const updateX = (newVal: string) => PlotUtils.updateXAxis(chart.value!, dataSet.value?.data["0"][newVal]!)
const redraw = (val: string[]) => {
  graph.value.y = val
  PlotUtils.redrawPlot(chart.value!, dataSet.value!, graph.value)
}
</script>
