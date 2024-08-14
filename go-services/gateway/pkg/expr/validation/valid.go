package validation

import (
	"errors"
	"fmt"
	"ltplotter/gen/pb"
	"regexp"
	"strconv"
)

var (
	ErrExpressionEmpty         = errors.New("expression cannot be empty")
	ErrInvalidExpressionFormat = errors.New("expression contains invalid characters")
	ErrUnsafeLaTeXCommand      = errors.New("input contains unsafe LaTeX command")
	ErrInvalidNumericalField   = errors.New("field must be a valid number")
	ErrInvalidDomainFormat     = errors.New("domain must be in the format 'start:end', with both as valid numbers")
	ErrPlotExpressionNil       = errors.New("plot expression cannot be nil")
	ErrRequestNil              = errors.New("request cannot be nil")
)

var blackList = []string{
	`\input`, `\include`, `\write`, `\openin`, `\read`, `\newread`, `\immediate`, `\catcode`,
	`\lstinputlisting`, `\verbatiminput`, `\usepackage`, `\href`, `\url`, `\textbackslash`,
	"rm", "sudo", "chmod", "chown", "dd", "mkfs", "shutdown", "reboot", "kill", "killall",
	"mv", "cp", "wget", "curl", "scp", "apt-get", "yum", "brew", "pip",
	"service", "systemctl", "mount", "umount", "ifconfig", "ip", "iptables",
	"ssh", "telnet", "echo", "cat", "eval", "sh", "bash", "lsof", "ps",
}

func checkUnsafeCommands(value string) error {
	for _, cmd := range blackList {
		// Create a regex to match the whole word
		pattern := fmt.Sprintf(`\b%s\b`, regexp.QuoteMeta(cmd))
		matched, err := regexp.MatchString(pattern, value)
		if err != nil {
			return err
		}
		if matched {
			return fmt.Errorf("%w: %s", ErrUnsafeLaTeXCommand, cmd)
		}
	}
	return nil
}

// validateNumericalField ensures that a string field contains a valid number if not empty.
func validateNumericalField(fieldName, value string) error {
	if value == "" {
		return nil // Field is optional
	}
	if _, err := strconv.ParseFloat(value, 64); err != nil {
		return fmt.Errorf("%w: %s, error: %v", ErrInvalidNumericalField, fieldName, err)
	}
	return nil
}

func validateExpression(expression string) error {
	if expression == "" {
		return ErrExpressionEmpty
	}

	allowedPattern := `^[a-zA-Z0-9\+\-\*/\^\(\)\s=.,<>!&|%]*$`
	matched, err := regexp.MatchString(allowedPattern, expression)
	if err != nil {
		return err
	}
	if !matched {
		return fmt.Errorf("%w: %s", ErrInvalidExpressionFormat, expression)
	}

	// Check for unsafe LaTeX commands
	return checkUnsafeCommands(expression)
}

func validateTextualField(value string) error {
	if value == "" {
		return nil // Field is optional
	}
	return checkUnsafeCommands(value)
}

// validateDomain checks if a domain is a valid range (e.g., "0:10").
func validateDomain(domain string) error {
	if domain == "" {
		return nil // Domain is optional
	}

	domainPattern := `^\-?\d+(\.\d+)?\:\-?\d+(\.\d+)?$`
	matched, err := regexp.MatchString(domainPattern, domain)
	if err != nil {
		return err
	}
	if !matched {
		return ErrInvalidDomainFormat
	}
	return nil
}

func validatePlotExpression(plot *pb.PlotExpression) error {
	if plot == nil {
		return ErrPlotExpressionNil
	}

	// Validate the expression
	if err := validateExpression(plot.Expression); err != nil {
		return err
	}

	// Validate domain
	if err := validateDomain(plot.Domain); err != nil {
		return err
	}

	// Validate other optional fields (if provided)
	if err := validateNumericalField("Samples", fmt.Sprintf("%d", plot.Samples)); err != nil {
		return err
	}

	if err := validateTextualField(plot.Color); err != nil {
		return err
	}

	if err := validateTextualField(plot.LineStyle); err != nil {
		return err
	}

	if err := validateTextualField(plot.LineWidth); err != nil {
		return err
	}

	if err := validateTextualField(plot.Legend); err != nil {
		return err
	}

	return nil
}

func ValidateRequest(req *pb.ExprPlotRequest) error {
	if req == nil {
		return ErrRequestNil
	}

	for _, plot := range req.GetPlots() {
		if err := validatePlotExpression(plot); err != nil {
			return err
		}
	}

	if err := validateNumericalField("XMin", req.GetXMin()); err != nil {
		return err
	}
	if err := validateNumericalField("XMax", req.GetXMax()); err != nil {
		return err
	}
	if err := validateNumericalField("YMin", req.GetYMin()); err != nil {
		return err
	}
	if err := validateNumericalField("YMax", req.GetYMax()); err != nil {
		return err
	}

	if err := validateTextualField(req.GetXLabel()); err != nil {
		return err
	}
	if err := validateTextualField(req.GetYLabel()); err != nil {
		return err
	}

	return nil
}
