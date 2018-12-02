package version

var Version string

const unsetVersion = "version-unset"

func init() {
	if Version == "" {
		Version = unsetVersion
	}
}
