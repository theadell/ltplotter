
<template>
  <v-container>
    <h2 class="text-3xl mt-6 font-bold text-gray-800 mb-6 text-center">
      Plot Generator
    </h2>
    <v-row>
      <v-col class="bg-white rounded-lg p-6" cols="12" md="6">
        <PlotForm2d :error="status.error" :loading="status.loading" @submit="onFormSubmit" />
      </v-col>

      <v-col class="flex flex-col items-center" cols="12" md="6">
        <plotDisplay :error="status.error" :latex-code="latexCode" :loading="status.loading" :pdf-url="pdfURL" />
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup lang="ts">
import { PlotRequest } from '@/lib/models/plot'
import api from '@/lib/api/plotApi'
import { setPlotRequestDomainFromXBounds } from '@/lib/utils/plotExpressionHelpers'

interface Status {
  loading: boolean
  error: boolean
}

const status = ref<Status>({ loading: false, error: false })
const pdfURL = ref<string | null>(null)
const latexCode = ref<string | null>(null)

const resetState = () => {
  status.value.loading = true
  status.value.error = false
  pdfURL.value = null
  latexCode.value = null
}
const onFormSubmit = async (plotRequest : PlotRequest) => {
  const payload = setPlotRequestDomainFromXBounds(plotRequest)
  resetState()
  try {
    const { pdfBlob, latexSource } = await api.plotExpression(payload)
    pdfURL.value = URL.createObjectURL(pdfBlob)
    latexCode.value = latexSource
  } catch (error) {
    status.value.error = true
    pdfURL.value = null
    latexCode.value = null
  } finally {
    status.value.loading = false
  }
}

</script>
