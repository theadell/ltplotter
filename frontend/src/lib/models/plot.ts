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
  None = 'none',
  Box = 'box',
  Left = 'left',
  Middle = 'middle',
  Center = 'center',
  Right = 'right',
  Empty = ''
}
export enum Grid {
  None = 'none',
  MAJOR = 'major',
  MINOR = 'minor',
  BOTH = 'both',
}

export enum LegendPosition {
  SouthhWest= 'south west',
  SouthEast= 'south east',
  NorthWest = 'north west',
  NorthEast = 'north east',
  OuterNorthEast = 'outer north east',
  None = '',
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

export interface PlotJobResponse {
    jobID: string;
}

