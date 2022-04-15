package rnd

import "strings"

// ValidID checks if the string is a valid unique ID.
func ValidID(s string, prefix byte) bool {
	// Regular UUID.
	if ValidUUID(s) {
		return true
	}

	// Not a known GenerateUID format.
	if len(s) != 16 {
		return false
	}

	return EntityUID(s, prefix)
}

// ValidIDs checks if a slice of strings contains ValidIDs only.
func ValidIDs(s []string, prefix byte) bool {
	if len(s) < 1 {
		return false
	}

	for _, id := range s {
		if !ValidID(id, prefix) {
			return false
		}
	}

	return true
}

// EntityUID returns true if string is a unique id as generated by PhotoPrism.
func EntityUID(s string, prefix byte) bool {
	if len(s) != 16 {
		return false
	}

	if !IsAlnum(s) {
		return false
	}

	return prefix == 0 || s[0] == prefix
}

// ValidUUID tests if the string looks like a standard UUID.
func ValidUUID(s string) bool {
	return len(s) == 36 && IsHex(s)
}

// SanitizeUUID normalizes UUIDs found in XMP or Exif metadata.
func SanitizeUUID(s string) string {
	if s == "" {
		return ""
	}

	s = strings.Replace(strings.TrimSpace(s), "\"", "", -1)

	if start := strings.LastIndex(s, ":"); start != -1 {
		s = s[start+1:]
	}

	if !ValidUUID(s) {
		return ""
	}

	return strings.ToLower(s)
}

// IsAlnum returns true if the string only contains alphanumeric ascii chars without whitespace.
func IsAlnum(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s {
		if (r < 48 || r > 57) && (r < 97 || r > 122) {
			return false
		}
	}

	return true
}

// IsHex returns true if the string only contains hex numbers, dashes and letters without whitespace.
func IsHex(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s {
		if (r < 48 || r > 57) && (r < 97 || r > 102) && (r < 65 || r > 70) && r != 45 {
			return false
		}
	}

	return true
}
