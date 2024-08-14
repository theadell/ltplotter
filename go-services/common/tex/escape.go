package tex

import "regexp"

var escapePattern = regexp.MustCompile(`([\\&%$#_{}~^<>])`)

func Escape(text string) string {
	conv := map[string]string{
		"\\": `\textbackslash{}`,
		"&":  `\&`,
		"%":  `\%`,
		"$":  `\$`,
		"#":  `\#`,
		"_":  `\_`,
		"{":  `\{`,
		"}":  `\}`,
		"~":  `\textasciitilde{}`,
		"^":  `\^{} `,
		"<":  `\textless{}`,
		">":  `\textgreater{}`,
	}

	escaped := escapePattern.ReplaceAllStringFunc(text, func(match string) string {
		return conv[match]
	})

	return escaped
}
