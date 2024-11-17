import { all, create, MathNode } from "mathjs"
import { Err, Ok, Result } from "../models/result"

const requiringArguments = [
  "abs", "acos", "add", "and", "array", "asin", "atan", "atan2", "bin", "ceil", "cos", "cosec",
  "cosh", "cot", "deg", "dim", "div", "divide", "equal", "floor", "frac", "gcd", "greater",
  "hex", "Hex", "int", "ifthenelse", "iseven", "isodd", "isprime", "less", "ln", "log10",
  "log2", "max", "min", "mod", "Mod", "multiply", "neg", "notequal", "notgreater", "notless",
  "oct", "or", "pow", "rad", "real", "round", "scalar", "sec", "sign", "sin", "sinh", "sqrt",
  "subtract", "tan", "tanh", "veclen",
]

const functionNameMappings: Record<string, string> = {
  log: "log10 or log2",
  exp: "e^x",
}

const config = {}
const math = create(all, config)

const additionalFunctions = {
  deg: (x: number) => x * (180 / Math.PI),
  cosec: (x: number) => 1 / Math.sin(x),
  ln: (x: number) => Math.log(x),
  // Additional pgfplots functions
  depth: (x: number) => x, // Placeholder: Depth is usually a property of an object; define it as needed
  dim: (x: number, y: number) => Math.max(0, x - y), // Commonly used for dimensional analysis
  greater: (x: number, y: number) => (x > y ? 1 : 0),
  height: (x: number) => x, // Placeholder: Height is usually a property of an object; define it as needed
  int: (x: number) => Math.floor(x), // Integer part of the number
  ifthenelse: (condition: boolean, trueVal: number, falseVal: number) => (condition ? trueVal : falseVal),
  iseven: (x: number) => (x % 2 === 0 ? 1 : 0),
  isodd: (x: number) => (x % 2 !== 0 ? 1 : 0),
  isprime: (x: number) => {
    if (x < 2) return 0
    for (let i = 2; i <= Math.sqrt(x); i++) {
      if (x % i === 0) return 0
    }
    return 1
  },
  less: (x: number, y: number) => (x < y ? 1 : 0),
  notequal: (x: number, y: number) => (x !== y ? 1 : 0),
  notgreater: (x: number, y: number) => (x <= y ? 1 : 0),
  notless: (x: number, y: number) => (x >= y ? 1 : 0),
  rad: (x: number) => x * (Math.PI / 180),
  rnd: () => Math.random(),
  veclen: (x: number, y: number) => Math.sqrt(x * x + y * y), // Vector length for 2D vectors
  width: (x: number) => x, // Placeholder: Width is usually a property of an object; define it as needed
}

math.import(additionalFunctions)

math.import({
  e: Math.E,
  "!": math.factorial,
}, { override: true })

export function validateExpression (expression: string): Result<true, string> {
  try {
    const node: MathNode = math.parse(expression)

    node.traverse((node: MathNode, path: string, parent: MathNode) => {
      if (node.type === "FunctionNode") {
        const functionNode = node as any
        const functionName = functionNode.fn.name as keyof typeof math | keyof typeof additionalFunctions

        // Ensure that the function has the required arguments
        if (functionNode.args.length === 0) {
          throw new Error(`Function '${functionName}' requires an argument enclosed in parentheses.`)
        }

        if (requiringArguments.includes(functionName) && functionNode.args.length === 0) {
          throw new Error(`Function '${functionName}' requires an argument. Did you mean ${functionName}(x)'`)
        }

        // Check for renamed functions and provide suggestions
        if (functionNameMappings[functionName]) {
          throw new Error(`Function '${functionName}' is named differently in pgfplots. Did you mean ${functionNameMappings[functionName]}?`)
        }
      }

      if (node.type === "SymbolNode") {
        const symbolNode = node as any
        const name = symbolNode.name as keyof typeof math | keyof typeof additionalFunctions

        if (requiringArguments.includes(name)) {
          if (!(parent && parent.type === "FunctionNode" && (parent as any).fn.name === name)) {
            throw new Error(`Function '${name}' requires an argument. Did you mean '${name}(x)'?`)
          }
        }

        if (functionNameMappings[symbolNode.name]) {
          throw new Error(`Function '${symbolNode.name}' is not valid. Did you mean ${functionNameMappings[symbolNode.name]}?`)
        }
        if (!(name in math) && !(name in additionalFunctions) && symbolNode.name !== "x") {
          throw new Error(`Unknown variable or function: ${symbolNode.name}`)
        }
      }
    })

    // If no errors occur, return Ok
    return Ok(true)
  } catch (error) {
    if (error instanceof Error && error.message.startsWith("Function")) {
      return Err(error.message) // Return specific error message with suggestions
    } else if (error instanceof Error && error.message.startsWith("Unknown variable or function")) {
      return Err(error.message) // Return specific error message
    } else if (error instanceof SyntaxError) {
      return Err("Syntax error in expression.")
    } else {
      return Err("Invalid mathematical expression.")
    }
  }
}

