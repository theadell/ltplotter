<template>
  <div class="fill-height pa-4">
    <v-row align="stretch" class="fill-height" justify="center">
      <!-- Editor Column -->
      <v-col class="d-flex flex-column" cols="12" md="6">
        <v-card class="flex-grow-1 d-flex flex-column" variant="flat">
          <v-card-title class="font-weight-bold">Editor</v-card-title>
          <v-divider />
          <v-card-text class="my-0 py-0">
            <v-toolbar density="compact">
              <!-- Slide Group for Tabs -->
              <v-slide-group v-model="activeTabName" class="editor-tabs ml-2" show-arrows>
                <v-slide-group-item
                  v-for="(tab) in tabs"
                  :key="tab.name"
                  v-slot="{ isSelected }"
                  :value="tab.name"
                >
                  <v-btn
                    class="text-none pa-2 ma-1"
                    :class="{ 'active-file-tab': isSelected }"
                    size="small"
                    variant="text"
                    @click="() => switchTab(tab.name)"
                  >
                    <span v-if="isSelected" class="v-btn__overlay" />
                    {{ tab.name }}
                  </v-btn>
                </v-slide-group-item>
              </v-slide-group>

              <v-spacer />

              <!-- Add New Tab Button -->
              <v-btn icon @click="addTab">
                <v-icon>mdi-plus</v-icon>
              </v-btn>
            </v-toolbar>

          </v-card-text>
          <v-card-text class="p-0 flex-grow-1 border-secondary mb-0 pb-1 pt-0 mt-0">
            <div id="editor" class="editor border rounded" />
          </v-card-text>
          <v-card-text class="pt-0 mt-0">
            <v-sheet class="px-2 py-1 d-flex justify-end rounded elevation-1 border" color="background">
              <v-btn class="text-none font-weight-bold" color="green-lighten-1" variant="outlined" @click="runCode">
                <v-icon size="x-large">mdi-play</v-icon>
                Run
              </v-btn>
            </v-sheet>
          </v-card-text>
        </v-card>
      </v-col>

      <!-- Results Column -->
      <v-col class="d-flex flex-column" cols="12" md="6">
        <v-card class="flex-grow-1 d-flex flex-column " variant="flat">
          <v-card-title class="font-weight-bold">Output</v-card-title>
          <v-divider />
          <v-card-text class="flex-grow-1 d-flex align-center justify-center">
            <div v-if="result?.success" class="video-container">
              <video autoplay controls style="width: 100%; height: auto;">
                <source :src="manimVideoUrl" type="video/mp4">
                Your browser does not support the video tag.
              </video>
            </div>

            <v-alert v-else-if="result !== null" class="mt-3" type="error">
              {{ result.output }}
            </v-alert>

            <div v-else class="text-center text-muted">
              Run your code to see the output here.
            </div>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </div>
</template>

<script setup>
import Manim from "@/pages/manim.vue"
import { onMounted, ref } from "vue"
import * as monaco from "monaco-editor"
import { createMonacoHlighter, MONACO_DEFAULT_THEME, setmonacoTheme } from "@/lib/utils/shiki"
import { useTheme } from "vuetify"
import { shikiToMonaco } from "@shikijs/monaco"
import manimApi from "@/lib/api/manimApi"

const theme = useTheme()
const tabs = [
  {
    name: "main.py",
    model: null,
    viewState: null,
  },
]

const activeTabName = ref("main.py")
const result = ref(null)
const manimVideoUrl = ref("")

let editor = null

onMounted(async () => {
  console.debug("Create shikijs highlighter for monaco")
  const highlighter = await createMonacoHlighter()

  // Register the languageIds first. Only registered languages will be highlighted.
  monaco.languages.register({ id: "python" })

  // Register the themes from Shiki, and provide syntax highlighting for Monaco.
  shikiToMonaco(highlighter, monaco)

  console.debug("Initializing Monaco Editor...")
  const editorContainer = document.getElementById("editor")

  if (!editorContainer) {
    console.error("Editor container not found!")
    return
  }

  // Create the initial tab's model
  const initialTab = tabs[0]
  initialTab.model = monaco.editor.createModel(
    "# Write your Python code here\nprint('Hello, World!')",
    "python"
  )

  // Initialize the Monaco Editor with no model initially
  editor = monaco.editor.create(editorContainer, {
    model: null,
    theme: MONACO_DEFAULT_THEME,
    automaticLayout: true,
    fontSize: 16,
  })

  editor.setModel(initialTab.model)
  console.debug("Monaco Editor initialized.")
})

const addTab = () => {
  const newTabName = `file${tabs.length + 1}.py`
  const newModel = monaco.editor.createModel("", "python")

  tabs.push({
    name: newTabName,
    model: newModel,
    viewState: null,
  })

  switchTab(newTabName)
}

const switchTab = tabName => {
  if (!editor) return

  const currentTab = tabs.find(tab => tab.name === activeTabName.value)
  const newTab = tabs.find(tab => tab.name === tabName)

  if (!newTab) {
    console.error("Tab not found:", tabName)
    return
  }

  if (currentTab) {
    currentTab.viewState = editor.saveViewState()
  }

  activeTabName.value = tabName

  editor.setModel(newTab.model)

  if (newTab.viewState) {
    editor.restoreViewState(newTab.viewState)
  } else {
    editor.setPosition({
      lineNumber: 1,
      column: 1,
    })
  }

  editor.focus()
}

const runCode = async () => {
  if (!editor) return
  console.log("Run code triggered")
  console.log("Code content:", editor.getValue())
  try {
    const manimVideoResponse = await manimApi.createVideo({ pythonSource: editor.getValue() })
    manimVideoUrl.value = manimVideoResponse.manimVideoUrl
    result.value = { success: true }
  } catch (e) {
    console.log("Error while getting manim video response", e)
  }
}

watch(() => theme.current.value, v => {
  setmonacoTheme(monaco, v.dark)
})
</script>

<style scoped>
.editor {
  height: 100%;
}

.video-container {
  width: 100%;
}

.v-card-title {
  font-weight: bold;
}

.active-file-tab {
  background-color: rgba(0, 0, 0, 0.08);
  border-radius: 3px;
  position: relative;
  overflow: hidden;
}

.active-file-tab .v-btn__overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.08);
  pointer-events: none;
}

</style>
