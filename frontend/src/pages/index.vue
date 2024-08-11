<template>
    <v-container>
        <h2 class="text-3xl mt-6 font-bold text-gray-800 mb-6 text-center">
            Plot Generator
        </h2>
        <v-row>

            <!-- Left Column: Inputs and Plot Configurations -->
            <v-col cols="12" md="6" class="bg-white rounded-lg p-6">
                <!-- Expression Fields Section -->
                <div>
                    <h3 class="text-xl font-medium text-gray-700 mb-4">Math Expressions</h3>
                    <div v-for="(expr, index) in expressions" :key="index" class="mb-1">
                        <v-text-field v-model="expr.expression" :label="`Expression ${index + 1}`" variant="outlined"
                            density="compact" single-line :rules="[validateExpression]">
                            <template #prepend-inner>
                                <span class="text-gray-500 whitespace-nowrap">y =</span>
                            </template>
                        </v-text-field>
                    </div>

                    <v-btn color="blue-grey-darken-1" prepend-icon="mdi-plus" variant="text" block
                        class="mb-5 text-none text-subtitle-2" @click="addExpression">
                        Add Expression
                    </v-btn>

                    <!-- Guide for Valid Expressions -->
                    <div class="text-gray-600 text-sm mb-5">
                        <p><strong>Examples of valid expressions:</strong></p>
                        <ul class="list-disc pl-5">
                            <li>x^2 + 3x - 5</li>
                            <li>sin(deg(x)) * cos(deg(x))</li>
                            <li>log10(x) + e^x - ln(x)</li>
                            <li>e^x - ln(x)</li>
                            <li>(x &lt; 0) * x^2 + (x>= 0) * (x + 1) </li>

                        </ul>
                    </div>
                </div>

                <!-- Advanced Plot Settings -->
                <div>
                    <v-expansion-panels>
                        <v-expansion-panel title="Plot Configurations">
                            <v-expansion-panel-text>
                                <div class="rounded-lg">
                                    <v-text-field density="compact" v-model="plotConfig.x_label" label="X-Axis Label"
                                        variant="outlined"></v-text-field>
                                    <v-text-field density="compact" v-model="plotConfig.y_label" label="Y-Axis Label"
                                        variant="outlined" dense></v-text-field>
                                    <v-row no-gutters class="mb-0 pb-0">
                                        <v-col cols="6" class="pr-4">
                                            <v-text-field density="compact" v-model="plotConfig.x_min"
                                                label="X-Axis Min" variant="outlined" dense
                                                :rules="[validNum]"></v-text-field>
                                        </v-col>
                                        <v-col cols="6">
                                            <v-text-field density="compact" v-model="plotConfig.x_max" :rules="validNum"
                                                label="X-Axis Max" variant="outlined" dense></v-text-field>
                                        </v-col>
                                        <v-col cols="6" class="pr-4">
                                            <v-text-field density="compact" v-model="plotConfig.y_min" :rules="validNum"
                                                label="Y-Axis Min" variant="outlined" dense></v-text-field>
                                        </v-col>
                                        <v-col cols="6">
                                            <v-text-field density="compact" v-model="plotConfig.y_max" :rules="validNum"
                                                label="Y-Axis Max" variant="outlined" dense></v-text-field>

                                        </v-col>
                                        <v-col cols="6">
                                            <v-checkbox density="compact" class="mt-0" v-model="plotConfig.grid"
                                                label="Show Grid"></v-checkbox>

                                        </v-col>
                                    </v-row>
                                </div>
                            </v-expansion-panel-text>
                        </v-expansion-panel>
                    </v-expansion-panels>

                </div>

                <v-alert v-if="errorMessage" type="error" closable v-model="errorMessage" class="my-4 w-full">
                    {{ errorMessage }}
                </v-alert>


                <!-- Generate Plot Button -->
                <v-btn :disabled="!isFormValid" color="#7986CB" prepend-icon="mdi-auto-fix" variant="elevated" block
                    :loading="loading" class="mt-4 py-3 text-lg font-medium text-white" @click="generateExpression">
                    Generate Plot
                </v-btn>
            </v-col>

            <!-- Right Column: Results Display -->
            <v-col cols="12" md="6" class="flex flex-col items-center">
                <!-- Error Message -->

                <!-- PDF Preview -->
                <div v-if="pdfUrl && !errorMessage" class="w-full mb-6">
                    <h3 class="text-xl font-medium text-gray-700 mb-2">PDF Preview</h3>
                    <div class="overflow-hidden rounded-lg shadow-md border border-gray-200">
                        <iframe :src="pdfUrl" class="w-full h-96" frameborder="0"></iframe>
                    </div>
                </div>

                <!-- Empty State for PDF Preview -->
                <div v-else class="w-full mb-6 flex items-center justify-center h-96">
                    <v-empty-state title="No plot generated yet"
                        text="Enter a math expression and click Generate Plot to see the result here."></v-empty-state>
                </div>

                <!-- LaTeX Source Code -->
                <div v-if="latexSource && !errorMessage" class="w-full mb-6">
                    <h3 class="text-xl font-medium text-gray-700 mb-2">LaTeX Source Code</h3>
                    <div class="p-4 bg-gray-100 rounded-lg shadow-md border border-gray-200"
                        style="max-height: 300px; overflow-y: auto;">
                        <pre class="whitespace-pre-wrap text-sm text-gray-800">{{ latexSource }}</pre>
                    </div>
                </div>

            </v-col>
        </v-row>
    </v-container>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';

interface Expression {
    expression: string;
}

interface PlotConfig {
    x_label: string;
    y_label: string;
    x_min: string;
    x_max: string;
    y_min: string;
    y_max: string;
    grid: boolean;
}

const expressions = ref<Expression[]>([{ expression: '' }]);
const plotConfig = ref<PlotConfig>({
    x_label: '',
    y_label: '',
    x_min: '',
    x_max: '',
    y_min: '',
    y_max: '',
    grid: false,
});
const loading = ref(false);
const pdfUrl = ref<string | null>(null);
const latexSource = ref<string | null>(null);
const errorMessage = ref<string | null>(null);


const validNum = (v: any) => !v || !isNaN(Number(v)) || 'Must be a number';
const validateExpression = (v: any) => {
    if (!v) {
        return 'Expression cannot be empty';
    }
    if (!/x/.test(v)) {
        return 'Expression must contain at least one \'x\'';
    }
    return true;
};

const isFormValid = computed(() => {
    const areExpressionsValid = expressions.value.every(expr => expr.expression.trim() && /x/.test(expr.expression));
    const areMinMaxNumbers = (v: string | null) => !v || !isNaN(Number(v));
    return areExpressionsValid &&
        areMinMaxNumbers(plotConfig.value.x_min) &&
        areMinMaxNumbers(plotConfig.value.x_max) && areMinMaxNumbers(plotConfig.value.y_min) && areMinMaxNumbers(plotConfig.value.y_max);
});


const addExpression = () => {
    expressions.value.push({ expression: '' });
};

const generateExpression = async () => {
    if (!isFormValid.value) {
        errorMessage.value = "Please enter valid math expressions.";
        return;
    }

    loading.value = true;
    errorMessage.value = null;
    pdfUrl.value = null;
    latexSource.value = null;

    try {
        const payload = {
            ...plotConfig.value,
            plots: expressions.value.map((expression) => {
                const trimmedExpression = expression.expression.trim();
                const domain = plotConfig.value.x_min && plotConfig.value.x_max
                    ? `${plotConfig.value.x_min}:${plotConfig.value.x_max}`
                    : undefined;

                return {
                    expression: trimmedExpression,
                    ...(domain && { domain }),
                };
            }),
        };


        const res = await fetch('/api/v2/plot/expr', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(payload),
        });

        if (!res.ok) {
            throw new Error('Error generating the plot');
        }

        const data = await res.json();
        const jobId = data.jobID;

        await new Promise((resolve) => setTimeout(resolve, 1000));

        const pdfResponse = await fetch(`/api/v2/plot/expr/${jobId}?format=pdf`);
        if (!pdfResponse.ok) throw new Error('Failed to fetch PDF');
        const pdfBlob = await pdfResponse.blob();
        pdfUrl.value = URL.createObjectURL(pdfBlob);

        const latexResponse = await fetch(`/api/v2/plot/expr/${jobId}?format=latex`);
        if (!latexResponse.ok) throw new Error('Failed to fetch LaTeX');
        latexSource.value = await latexResponse.text();
    } catch (error: any) {
        errorMessage.value = error.message;
        pdfUrl.value = null;
        latexSource.value = null;
    } finally {
        loading.value = false;
    }
};
</script>
