package entity

const (
	BuildVersion = ""
	BuildAt      = ""
)

type Version struct {
	Build   string `json:"build"`
	BuildAt string `json:"build_at"`
}

func GetApplicationBuildInfo() (version *Version) {
	return &	Version{
		Build: BuildVersion,
		BuildAt: BuildAt,
	}
}