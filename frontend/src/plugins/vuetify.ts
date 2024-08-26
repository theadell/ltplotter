/**
 * plugins/vuetify.ts
 *
 * Framework documentation: https://vuetifyjs.com`
 */

// Styles
import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles'

// Composables
import { createVuetify, ThemeDefinition } from 'vuetify'
import { VNumberInput } from 'vuetify/labs/VNumberInput'

const lightTheme: ThemeDefinition = {
  dark: false,
  colors: {
    primary: '#415F91',
    'primary-darken-1': '#234373',
    secondary: '#565F71',
    'secondary-darken-1': '#3A4354',
    accent: '#705575',
    error: '#BA1A1A',
    info: '#5177B8',
    success: '#86ACF0',
    warning: '#FFB400',
    background: '#F9F9FF',
    surface: '#F9F9FF',
    'on-primary': '#FFFFFF',
    'on-secondary': '#FFFFFF',
    'on-accent': '#FFFFFF',
    'on-error': '#FFFFFF',
    'on-info': '#FFFFFF',
    'on-success': '#FFFFFF',
    'on-warning': '#FFFFFF',
    'on-background': '#191C20',
    'on-surface': '#191C20',
    'surface-variant': '#E0E2EC',
    'on-surface-variant': '#44474E',
    outline: '#74777F',
    'inverse-surface': '#2E3036',
    'inverse-on-surface': '#F0F0F7',
    'inverse-primary': '#AAC7FF',
    'badge-bg-color': '#8E8E93',
    'badge-text': '#111111',
  },

}

const darkTheme: ThemeDefinition = {
  dark: true,
  colors: {
    primary: '#AAC7FF',
    'primary-darken-1': '#7491C7',
    secondary: '#BEC6DC',
    'secondary-darken-1': '#8891A5',
    accent: '#DDBCE0',
    error: '#FFB4AB',
    info: '#6B91D3',
    success: '#AAC7FF',
    warning: '#FFB400',
    background: '#111318',
    surface: '#111318',
    'on-primary': '#0A305F',
    'on-secondary': '#283141',
    'on-accent': '#3F2844',
    'on-error': '#690005',
    'on-info': '#001B3E',
    'on-success': '#001B3E',
    'on-warning': '#111318',
    'on-background': '#E2E2E9',
    'on-surface': '#E2E2E9',
    'surface-variant': '#44474E',
    'on-surface-variant': '#C4C6D0',
    outline: '#8E9099',
    'inverse-surface': '#E2E2E9',
    'inverse-on-surface': '#2E3036',
    'inverse-primary': '#415F91',
    'badge-bg-color': '#8E8E93',
    'badge-text': '#FFFFFF',
  },

}

// https://vuetifyjs.com/en/introduction/why-vuetify/#feature-guides
export default createVuetify({
  components: {
    VNumberInput,
  },
  theme: {
    defaultTheme: 'light',
    themes: {
      light: lightTheme,
      dark: darkTheme,
    },
  },
})
