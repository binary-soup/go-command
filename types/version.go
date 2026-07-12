package types

// Version represents a version code.
type Version int

// Check if the version is considered unsupported
// (ie. less than the min version or greater than the current version).
func (v Version) IsUnsupported(minVersion, currentVersion int) bool {
	return v < Version(minVersion) || v > Version(currentVersion)
}

// Check if the version is considered out-of-date (ie. less than the current version).
// Note: does NOT check for a minimum supported version (see IsUnsupported).
func (v Version) IsOutOfDate(currentVersion int) bool {
	return v < Version(currentVersion)
}
