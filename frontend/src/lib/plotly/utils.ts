import Plotly, { ToImgopts } from "plotly.js-dist";
import { DataSet, PlotState } from "../models";

type PLU = Partial<Plotly.Layout>;

const initEmptyPlot = (
  chart: HTMLElement,
  title = "",
  xaxis = "",
  yaxis = ""
) => {
  Plotly.newPlot(chart, [], {
    title: title,
    xaxis: { title: xaxis },
    yaxis: { title: yaxis },
  });
};
const updateXAxis = (chart: HTMLElement, x: any[]) => {
  Plotly.restyle(chart, { x: [x] });
};

const updateTitle = (chart: HTMLElement, newTitle: string) => {
  const update: PLU = {
    title: {
      text: newTitle,
    },
  };
  Plotly.relayout(chart, update);
};

const updateXLabel = (chart: HTMLElement, newXLabel: string) => {
  const update: PLU = {
    xaxis: {
      title: {
        text: newXLabel,
      },
    },
  };
  Plotly.relayout(chart, update);
};

const updateYLabel = (chart: HTMLElement, newYLabel: string) => {
  const update: PLU = {
    yaxis: {
      title: {
        text: newYLabel,
      },
    },
  };
  Plotly.relayout(chart, update);
};

const updateXAxisData = (chart: HTMLElement, newXData: any[]) => {
  Plotly.restyle(chart, { x: [newXData] });
};

export enum ExportFormat {
  SVG = "svg",
  PNG = "png",
}

const redrawPlot = (
  chart: HTMLElement,
  dataSet: DataSet,
  state: PlotState,
  step = 0
) => {
  const x = dataSet.data[`${step}`][state.x];
  const traces = state.y.map((yaxis) => ({
    x: x,
    y: dataSet.data["0"][yaxis],
    name: yaxis,
  }));
  const layout: PLU = {
    title: state.title,
    xaxis: {
      title: state.xAxisLabel,
    },
    yaxis: {
      title: state.yAxisLabel,
    },
    showlegend: state.legendVisible,
  };
  Plotly.newPlot(chart, traces, layout);
};

const exportChart = (chart: HTMLElement, format: ExportFormat) => {
  const timestamp = Date.now();
  const filename = `chart_${timestamp}.${format}`;

  Plotly.toImage(chart, { format } as ToImgopts)
    .then((dataUrl: string) => {
      const a = document.createElement("a");
      a.href = dataUrl;
      a.download = filename;
      document.body.appendChild(a);
      a.click();
      document.body.removeChild(a);
    })
    .catch((e: Error) => {
      console.error(`Error exporting chart as ${format.toUpperCase()}:`, e);
    });
};

const updateLegends = (chart: HTMLElement, labels: string[]) => {
  const update = { name: labels };
  const traceIndices = labels.map((_, index) => index);
  // @ts-ignore
  Plotly.restyle(chart, update, traceIndices);
  Plotly.relayout(chart, { showlegend: true });
};
const toggleLengend = (chart: HTMLElement, visible: boolean) =>
  Plotly.relayout(chart, { showlegend: visible });

export const PlotUtils = {
  initEmptyPlot,
  updateXAxis,
  redrawPlot,
  toggleLengend,
  updateLegends,
  exportChart,
  updateTitle,
  updateXLabel,
  updateYLabel,
  updateXAxisData,
};
