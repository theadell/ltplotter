/* eslint-disable camelcase */
import { PlotRequest } from "../models/plot"
import lescape from "escape-latex"

export function setPlotRequestDomainFromXBounds (plotRequest: PlotRequest): PlotRequest {
  const { x_min, x_max, plots } = plotRequest

  if (x_min !== undefined && x_max !== undefined) {
    const newDomain = `${x_min}:${x_max}`
    plotRequest.plots = plots.map(plot => ({
      ...plot,
      domain: newDomain,
    }))
  }

  return plotRequest
}

export function escapePlotRequestStrings (request: PlotRequest): PlotRequest {
  const escapedRequest: PlotRequest = { ...request }

  if (escapedRequest.title) {
    escapedRequest.title = lescape(escapedRequest.title)
  }
  if (escapedRequest.x_label) {
    escapedRequest.x_label = lescape(escapedRequest.x_label)
  }
  if (escapedRequest.y_label) {
    escapedRequest.y_label = lescape(escapedRequest.y_label)
  }

  return escapedRequest
}
