import { HTTPError } from "ky"
import { DataPlotRequest, PlotJobResponse } from "../models/plot"
import apiClient from "./client"
import { PlotError } from "../models/plotError"

const dataPlotApi = {
  dispatchPlotJob: async (requestBody: DataPlotRequest): Promise<string> => {
    try {
      const response = await apiClient
        .post("plot/data", { json: r })
        .json<PlotJobResponse>()
      return response.jobID
    } catch (error: any) {
      throw new Error(error.message || "Error generating the plot")
    }
  },
  fetchPlotResult: async (
    jobId: string,
    format: "pdf" | "latex"
  ): Promise<Blob | string> => {
    try {
      const response = await apiClient.get(`plot/data/${jobId}`, {
        searchParams: { format },
      })

      return format === "pdf" ? await response.blob() : await response.text()
    } catch (error: any) {
      throw new Error(
        error.message || `Failed to fetch the ${format.toUpperCase()} result`
      )
    }
  },
  plotExpression: async (
    requestBody: DataPlotRequest,
  ): Promise<Blob> => {
    try {
      const jobResponse = await apiClient
        .post("plot/data", { json: requestBody })

      return jobResponse.blob()
    } catch (error: any) {
      if (error instanceof HTTPError) {
        if (error.response.status === 400) {
          throw new PlotError(
            "Invalid request. Please check your input and try again.",
            400
          )
        } else if (error.response.status === 500) {
          throw new PlotError("Server error. Please try again later.", 500)
        } else {
          throw new PlotError(
            `Unexpected error: ${error.response.statusText}`,
            error.response.status
          )
        }
      } else {
        throw new PlotError("Failed to generate or fetch the plot results", 0)
      }
    }
  },
}

export default dataPlotApi
