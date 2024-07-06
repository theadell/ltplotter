export type ComplexNumber = {
    real: number;
    imag: number;
  };
  
  export enum DataSetType {
    Float64 = "float64",
    Complex = "complex"
  }
  
  interface DataSetBase {
    name: string;
    simType: string;
    steps: number;
    variables: string[];
    xAxis: string;
    type: DataSetType;
  }

  
  export interface Float64DataSet extends DataSetBase {
    type: DataSetType.Float64;
    data: Record<string, Record<string, number[]>>;
  }
  
  export interface ComplexDataSet extends DataSetBase {
    type: DataSetType.Complex;
    data: Record<string, Record<string, ComplexNumber[]>>;
  }
  
export type DataSet = Float64DataSet | ComplexDataSet;
  
export const isFloat64DataSet = (dataSet: DataSet): dataSet is Float64DataSet => dataSet.type === DataSetType.Float64;
export const isComplexDataSet = (dataSet: DataSet): dataSet is ComplexDataSet => dataSet.type === DataSetType.Complex;


export interface PlotState {
  x: string
  y: string[]
  title: string
  xAxisLabel: string
  yAxisLabel: string
  legendVisible: boolean,
  legendPos: string,
  customizeLegend: boolean
  legendNames: Array<string>
}
