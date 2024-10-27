import { HTTPError } from "ky"
import { DataPlotRequest } from "../models/plot"
import apiClient from "./client"
import { PlotError } from "../models/plotError"

const dataPlotApi = {
  plotExpression: async (
    requestBody: DataPlotRequest,
  ): Promise<Blob> => {
    try {
      return await apiClient
        .post("plot/data", { json: requestBody })
        .blob()
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
