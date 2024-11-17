import { describe, expect, it } from "vitest"
import { validateExpression } from "./expressionValidator"
import { Result } from "../models/result"

// The documentation about valid expressions of pgfplots can be found in the TikZ/PGF manual.
// https://ftp.rrzn.uni-hannover.de/pub/mirror/tex-archive/graphics/pgf/base/doc/pgfmanual.pdf
// 94.2 Syntax for Mathematical Expressions: Operators - Page 1031
// 94.3 Syntax for Mathematical Expressions: Functions - Page 1033

describe("validateExpression", () => {
  const testCases = [
    // Simple valid expressions
    { expression: "x + 1", expected: true },
    { expression: "x - 1", expected: true },
    { expression: "x * 2", expected: true },
    { expression: "x / 2", expected: true },
    { expression: "x ^ 2", expected: true },

    // Simple invalid expressions
    { expression: "x +", expected: "Syntax error in expression." },
    { expression: "x / 0", expected: true },

    // Polynomials
    { expression: "x^3 + 2*x^2 - 5*x + 7", expected: true },
    { expression: "3*x^4 - 4*x^3 + 2*x^2 - x + 5", expected: true },
    { expression: "x^5 + 3*x^4 - x + 1", expected: true },
    { expression: "2*x^3 + 3*x - 5", expected: true },
    { expression: "x^6 - 4*x^2 + 2*x + 9", expected: true },

    // Polynomial inequalities
    { expression: "x^2 + x - 6 > 0", expected: true },
    { expression: "x^3 - x < 0", expected: true },
    { expression: "x^4 - 16 >= 0", expected: true },
    { expression: "(x^2 - 4)*(x + 2) <= 0", expected: true },

    // trigonometric expressions
    { expression: "sin(deg(x))", expected: true },
    { expression: "cos(rad(x))", expected: true },
    { expression: "tan(x)", expected: true },

    // Invalid function calls
    { expression: "unknownFunc(x)", expected: "Unknown variable or function: unknownFunc" },

    // Valid expressions with conditionals
    { expression: "(x < 0) * x^2 + (x >= 0) * (x + 1)", expected: true },
    { expression: "ifthenelse(x < 5, x^2, sqrt(x))", expected: true },

    //  complex expressions
    { expression: "log10(x) + e^x - ln(x)", expected: true },
    { expression: "sec(x) * cosec(x) + cot(x)", expected: true },

    //  nested expressions
    { expression: "sin(cos(deg(x))) + tan(ln(x + 1))", expected: true },
    { expression: "max(sin(x), cos(x)) + min(tan(x), x)", expected: true },

    // Invalid complex expressions
    { expression: "max(sin(x), unknownFunc(x))", expected: "Unknown variable or function: unknownFunc" },
    { expression: "min(cos(x), 2 +)", expected: "Syntax error in expression." },

    // vector length and modulus
    { expression: "veclen(x, x + 2)", expected: true },
    { expression: "mod(x, 3)", expected: true },

    // special functions
    { expression: "iseven(x) * x + isodd(x)", expected: true },
    { expression: "isprime(x) + factorial(x)", expected: true },

    // arguments
    { expression: "log10(x)", expected: true },
    { expression: "sin(x)", expected: true },
    { expression: "sqrt(4)", expected: true },

    // Invalid function calls
    { expression: "log10", expected: "Function 'log10' requires an argument. Did you mean 'log10(x)'?" },
    { expression: "sin", expected: "Function 'sin' requires an argument. Did you mean 'sin(x)'?" },
    { expression: "sqrt", expected: "Function 'sqrt' requires an argument. Did you mean 'sqrt(x)'?" },
  ]

  testCases.forEach(({ expression, expected }) => {
    it(`should ${expected === true ? "pass" : `fail with message "${expected}"`} for expression "${expression}"`, () => {
      const result: Result<true, string> = validateExpression(expression)

      if (expected === true) {
        expect(result.ok).toBe(true)
      } else {
        expect(result.ok).toBe(false)
        if (!result.ok) {
          expect(result.error).toBe(expected)
        }
      }
    })
  })
})
