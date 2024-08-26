import { createHighlighterCore } from 'shiki/core'
import getWasm from 'shiki/wasm'

let highlighter: any = null

export const createLaTeXHighlighter = async () => {
  if (!highlighter) {
    highlighter = await createHighlighterCore({
      themes: [import('shiki/themes/snazzy-light.mjs'),
        import('shiki/themes/catppuccin-frappe.mjs'),
      ],
      langs: [() => import('shiki/langs/latex.mjs')],
      loadWasm: getWasm,
    })
  }

  return highlighter
}

export const highlightLaTeX = async (code: string) => {
  const highlighter = await createLaTeXHighlighter()
  const html = highlighter.codeToHtml(code,
    {
      lang: 'latex',
      themes: {
        light: 'snazzy-light',
        dark: 'catppuccin-frappe',
      },
    })
  return { html }
}
