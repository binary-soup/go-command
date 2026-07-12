package config

type Version struct {
	Version int `json:"version"`
}

func (v Version) IsUnsupported(currentVersion int) bool {
	return v.Version < 1 || v.Version > currentVersion
}

func (v Version) IsOutOfDate(currentVersion int) bool {
	return v.Version < currentVersion
}
