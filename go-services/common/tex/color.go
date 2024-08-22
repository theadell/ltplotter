package tex

import (
	"strconv"
	"strings"

	"github.com/lucasb-eyer/go-colorful"
)

func ColorToTikzColor(color string) string {
	color = strings.TrimSpace(color)

	if !strings.HasPrefix(color, "#") {
		if isValidNamedColor(color) {
			return color
		}
		return "black"
	}

	if strings.HasPrefix(color, "#") && len(color) == 7 {
		c, err := colorful.Hex(color)
		if err == nil {
			return convertColorfulToLaTeX(c)
		}
	}

	return "black"
}

func convertColorfulToLaTeX(c colorful.Color) string {
	r, g, b := c.RGB255()
	return "{rgb,255:red," + strconv.Itoa(int(r)) + ";green," + strconv.Itoa(int(g)) + ";blue," + strconv.Itoa(int(b)) + "}"
}

func isValidNamedColor(color string) bool {
	namedColors := map[string]bool{
		"red":          true,
		"green":        true,
		"blue":         true,
		"cyan":         true,
		"magenta":      true,
		"yellow":       true,
		"black":        true,
		"white":        true,
		"gray":         true,
		"darkgray":     true,
		"lightgray":    true,
		"brown":        true,
		"lime":         true,
		"olive":        true,
		"orange":       true,
		"pink":         true,
		"purple":       true,
		"teal":         true,
		"violet":       true,
		"indigo":       true,
		"lightblue":    true,
		"lightgreen":   true,
		"lightcyan":    true,
		"lightmagenta": true,
		"lightyellow":  true,
		"darkred":      true,
		"darkgreen":    true,
		"darkblue":     true,
		"darkcyan":     true,
		"darkmagenta":  true,
		"darkyellow":   true,
	}
	return namedColors[strings.ToLower(color)]
}
