/**
 * .eslint.js
 *
 * ESLint configuration file.
 */

module.exports = {
  root: true,
  env: {
    node: true,
  },
  extends: [
    "vuetify",
    "@vue/eslint-config-typescript",
    "./.eslintrc-auto-import.json",
  ],
  rules: {
    "vue/multi-word-component-names": "off",
    "vue/script-indent": ["warn", 2, {
      baseIndent: 0,
      switchCase: 0,
      ignores: [],
    }],
    "no-multiple-empty-lines": ["warn", { max: 1 }],
    quotes: [2, "double", { avoidEscape: true }],
  },
}
