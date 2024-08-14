package tex

import (
	"testing"
)

func TestTexEscape(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    `Special characters ^ and ~ should be escaped.`,
			expected: `Special characters \^{}  and \textasciitilde{} should be escaped.`,
		},
		{
			input:    `No special characters here.`,
			expected: `No special characters here.`,
		},
		{
			input:    `Multiple $$ and && characters.`,
			expected: `Multiple \$\$ and \&\& characters.`,
		},
		{
			input:    `Comparison between A & B`,
			expected: `Comparison between A \& B`,
		},
		{
			input:    `Number of occurrences of # in text`,
			expected: `Number of occurrences of \# in text`,
		},
		{
			input:    `X-axis: Time (s), Y-axis: Voltage (V)`,
			expected: `X-axis: Time (s), Y-axis: Voltage (V)`,
		},
		{
			input:    `Growth rate (r%) vs Population size (N)`,
			expected: `Growth rate (r\%) vs Population size (N)`,
		},
		{
			input:    `Evolutionary Pressure: \lambda > 0`,
			expected: `Evolutionary Pressure: \textbackslash{}lambda \textgreater{} 0`,
		},
		{
			input:    `Temperature ~ Pressure`,
			expected: `Temperature \textasciitilde{} Pressure`,
		},
		{
			input:    `Error rate > 5%`,
			expected: `Error rate \textgreater{} 5\%`,
		},
		{
			input:    `Confidence interval: 95%`,
			expected: `Confidence interval: 95\%`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := Escape(tt.input)
			if result != tt.expected {
				t.Errorf("TexEscape(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestTexEscapeCommands(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    `\textbf{Injected Title}`,
			expected: `\textbackslash{}textbf\{Injected Title\}`,
		},
		{
			input:    `XLabel: \includegraphics[width=\textwidth]{evil.jpg}`,
			expected: `XLabel: \textbackslash{}includegraphics[width=\textbackslash{}textwidth]\{evil.jpg\}`,
		},
		{
			input:    `YLabel: \input{malicious.tex}`,
			expected: `YLabel: \textbackslash{}input\{malicious.tex\}`,
		},
		{
			input:    `Title: }{\end{axis}\begin{document}Some text`,
			expected: `Title: \}\{\textbackslash{}end\{axis\}\textbackslash{}begin\{document\}Some text`,
		},
		{
			input:    `Legend: \write18{rm -rf /}`,
			expected: `Legend: \textbackslash{}write18\{rm -rf /\}`,
		},
		{
			input:    `Title: \expandafter\def\csname mycommand\endcsname{}`,
			expected: `Title: \textbackslash{}expandafter\textbackslash{}def\textbackslash{}csname mycommand\textbackslash{}endcsname\{\}`,
		},
		{
			input:    `XLabel: $100 & 50% > 40% \texttt{exploit}`,
			expected: `XLabel: \$100 \& 50\% \textgreater{} 40\% \textbackslash{}texttt\{exploit\}`,
		},
		{
			input:    `YLabel: \begin{tikzpicture}\node{Hacked};\end{tikzpicture}`,
			expected: `YLabel: \textbackslash{}begin\{tikzpicture\}\textbackslash{}node\{Hacked\};\textbackslash{}end\{tikzpicture\}`,
		},
		{
			input:    `LineStyle: ;\immediate\write18{rm -rf /}`,
			expected: `LineStyle: ;\textbackslash{}immediate\textbackslash{}write18\{rm -rf /\}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := Escape(tt.input)
			if result != tt.expected {
				t.Errorf("TexEscape(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
