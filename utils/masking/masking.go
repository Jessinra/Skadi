package masking

import (
	"fmt"
	"strings"
)

const mask = "*"

// simple constants to avoid noisy magic number linter.
const (
	TwoThird = 2.0 / 3
	Half     = 0.5
	AThird   = 1.0 / 3
	AQuarter = 0.25
)

// Email mask email address (ignore domain), showing the first and last 20% chars
// e.g. myAmazingEmail@gmail.com -> my*********ail@gmail.com.
func Email(email string) string {
	const maskPower = 0.6

	if email == "" {
		return ""
	}

	ss := strings.Split(email, "@")
	addr := Center(ss[0], maskPower)

	if len(ss) == 1 {
		return addr
	}

	domain := ss[1]
	return fmt.Sprintf("%s@%s", addr, domain)
}

// Phone mask phone number, showing the first and last 3 chars
// e.g. +6287800990099 -> +62********099.
func Phone(phone string) string {
	const start, end = 3, 3

	if len(phone) < start+end {
		return phone
	}

	return Mask(phone, start, len(phone)-end)
}

// Name mask space-delimited name, showing the first 40% of each word.
// e.g. Richard James Wayne -> Ric**** Ja*** Wa***.
func Name(name string) string {
	const maskPower = 0.6

	if name == "" {
		return ""
	}

	ss := strings.Split(name, " ")
	for i, s := range ss {
		ss[i] = Right(s, maskPower)
	}

	return strings.Join(ss, " ")
}

// Left mask a string partially starting from index 0.
// e.g. mask 80% abcde -> ****e.
func Left(s string, maskPercent float64) string {
	maskPercent = sanitizePercent(maskPercent)

	if res, ok := isEdgeCases(s, maskPercent); ok {
		return res
	}

	end := int(float64(len(s)) * maskPercent)
	return Mask(s, 0, end)
}

// Right mask a string partially starting from the last index.
// e.g. mask 80% abcde -> a****.
func Right(s string, maskPercent float64) string {
	maskPercent = sanitizePercent(maskPercent)

	if res, ok := isEdgeCases(s, maskPercent); ok {
		return res
	}

	start := int(float64(len(s)) * (1 - maskPercent))
	return Mask(s, start, len(s))
}

// Center mask a string partially starting from the middle, to the left and right (both mask half of the percentage).
// e.g. mask 60% abcde -> a***e.
func Center(s string, maskPercent float64) string {
	maskPercent = sanitizePercent(maskPercent)

	if res, ok := isEdgeCases(s, maskPercent); ok {
		return res
	}

	start := int(float64(len(s)) * (1 - maskPercent) / 2) // nolint
	end := int(float64(len(s)) * (1 + maskPercent) / 2)   // nolint
	return Mask(s, start, end)
}

// Mask mask a string from [start, end)
// e.g. abcde [1, 3) => a**de.
func Mask(str string, start, end int) string {
	// sanitize input
	l := len(str)
	if l == 0 {
		return ""
	}
	if start < 0 {
		start = 0
	}
	if start > l {
		start = l
	}
	if end < 0 {
		end = 0
	}
	if end > l {
		end = l
	}
	if start > end {
		start, end = end, start
	}

	r := []rune(str)
	masked := fmt.Sprintf("%s%s%s", string(r[:start]), strings.Repeat(mask, (end-start)), string(r[end:]))
	return masked
}

// isEdgeCases handle edge cases for Left(), Right(), Center().
func isEdgeCases(s string, maskPercent float64) (string, bool) {
	if len(s) == 1 {
		switch {
		case maskPercent >= 0.5: // nolint
			return mask, true
		default:
			return s, true
		}
	}

	return s, false
}

func sanitizePercent(n float64) float64 {
	if n > 1 {
		return 1
	}
	if n < 0 {
		return 0
	}

	return n
}
