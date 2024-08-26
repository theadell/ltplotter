
<template>
  <div class="max-w-screen-xl mx-auto pa-4">
    <h2 class="text-2xl mt-6 font-bold mb-6">
      Symbolic Graphing
    </h2>
    <v-row>
      <v-col class="rounded-lg p-6" cols="12" md="6">
        <PlotForm2d :error="status.error" :loading="status.loading" @submit="onFormSubmit" />
      </v-col>

      <v-col class="flex flex-col items-center" cols="12" md="6">
        <plotDisplay
          :error="status.error"
          :error-message="status.errorMessage"
          :latex-code="latexCode"
          :loading="status.loading"
          :pdf-url="pdfURL"
        />
      </v-col>
    </v-row>
  </div>
</template>

<script setup lang="ts">
import { PlotRequest } from '@/lib/models/plot'
import api from '@/lib/api/plotApi'
import { setPlotRequestDomainFromXBounds } from '@/lib/utils/plotExpressionHelpers'
import { PlotError } from '@/lib/models/plotError'

interface Status {
  loading: boolean
  error: boolean
  errorMessage: string | null
}

const status = ref<Status>({ loading: false, error: false, errorMessage: null })
const pdfURL = ref<string | null>(null)
const latexCode = ref<string | null>(null)

const resetState = () => {
  status.value.loading = true
  status.value.error = false
  status.value.errorMessage = null
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
    if (error instanceof PlotError) {
      status.value.errorMessage = error.message
    } else {
      console.error('An unexpected error occurred:', error)
      status.value.errorMessage = 'An unexpected error occurred'
    }
  } finally {
    status.value.loading = false
  }
}

</script>
