export interface PlotElement {
    expression: string
    domain?: string
    samples?: number
    color?: string
    line_style?: string
    line_width?: string
    legend?: string
  }
export enum AxisLines {
  None = "none",
  Box = "box",
  Left = "left",
  Middle = "middle",
  Center = "center",
  Right = "right",
  Empty = ""
}
export enum Grid {
  None = "none",
  MAJOR = "major",
  MINOR = "minor",
  BOTH = "both",
}

export enum LegendPosition {
  SouthhWest= "south west",
  SouthEast= "south east",
  NorthWest = "north west",
  NorthEast = "north east",
  OuterNorthEast = "outer north east",
  None = "",
}

export interface PlotRequest {
    title?: string
    x_label?: string
    y_label?: string
    x_min?: number
    x_max?: number
    y_min?: number
    y_max?: number
    grid?: Grid
    axis_lines?: AxisLines
    border?: number
    legend?: LegendPosition
    plots: PlotElement[]
}

interface Labels {
  x: string;
  y: string;
}

interface DataSeries {
  values: number[];
}

interface Metadata {
  title: string;
  labels: Labels;
  legends: string[];
}

export interface DataPlotRequest {
  x: number[];
  y: DataSeries[];
  metadata: Metadata;
}

export interface PlotJobResponse {
    jobID: string;
}

export interface DataPlotResponse {
  pdf: string;
  latex: string;
}

export interface ManimRequest {
  pythonSource: string;
}

export interface ManimResponse {
  manimVideoUrl: string;
}
