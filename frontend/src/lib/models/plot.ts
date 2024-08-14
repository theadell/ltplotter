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
    BOX= 0,
    LEFT,
    RIGHT,
    MIDDLE,
    CENTER,
  }

export interface PlotRequest {
    title?: string
    x_label?: string
    y_label?: string
    x_min?: number
    x_max?: number
    y_min?: number
    y_max?: number
    grid?: boolean
    axis_lines?: AxisLines
    border?: number
    plots: PlotElement[]
  }

export interface PlotJobResponse {
    jobID: string;
}

