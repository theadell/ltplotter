import { HTTPError } from "ky";
import { PlotJobResponse, PlotRequest } from "../models/plot";
import { escapePlotRequestStrings } from "../utils/plotExpressionHelpers";
import apiClient from "./client";
import { PlotError } from "../models/plotError";

const plotApi = {
  dispatchPlotJob: async (requestBody: PlotRequest): Promise<string> => {
    try {
      const r = escapePlotRequestStrings(requestBody);
      const response = await apiClient
        .post("plot/expr", { json: r })
        .json<PlotJobResponse>();
      return response.jobID;
    } catch (error: any) {
      throw new Error(error.message || "Error generating the plot");
    }
  },
  fetchPlotResult: async (
    jobId: string,
    format: "pdf" | "latex"
  ): Promise<Blob | string> => {
    try {
      const response = await apiClient.get(`plot/expr/${jobId}`, {
        searchParams: { format },
      });

      return format === "pdf" ? await response.blob() : await response.text();
    } catch (error: any) {
      throw new Error(
        error.message || `Failed to fetch the ${format.toUpperCase()} result`
      );
    }
  },
  plotExpression: async (
    requestBody: PlotRequest,
    waitTime: number = 1000
  ): Promise<{ pdfBlob: Blob; latexSource: string }> => {
    try {
      const r = escapePlotRequestStrings(requestBody);
      const jobResponse = await apiClient
        .post("plot/expr", { json: r })
        .json<PlotJobResponse>();
      const jobId = jobResponse.jobID;

      await new Promise((resolve) => setTimeout(resolve, waitTime));

      const [pdfBlob, latexSource] = await Promise.all([
        apiClient
          .get(`plot/expr/${jobId}`, { searchParams: { format: "pdf" } })
          .blob(),
        apiClient
          .get(`plot/expr/${jobId}`, { searchParams: { format: "latex" } })
          .text(),
      ]);

      return { pdfBlob, latexSource };
    } catch (error: any) {
      if (error instanceof HTTPError) {
        if (error.response.status === 400) {
          throw new PlotError(
            "Invalid request. Please check your input and try again.",
            400
          );
        } else if (error.response.status === 500) {
          throw new PlotError("Server error. Please try again later.", 500);
        } else {
          throw new PlotError(
            `Unexpected error: ${error.response.statusText}`,
            error.response.status
          );
        }
      } else {
        throw new PlotError("Failed to generate or fetch the plot results", 0);
      }
    }
  },
};

export default plotApi;
