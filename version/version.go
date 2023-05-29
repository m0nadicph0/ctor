package version

import (
	"fmt"
	"runtime/debug"
)

var version = ""

func GetVersion() string {
	if version != "" {
		return version
	}

	info, ok := debug.ReadBuildInfo()

	if !ok && info.Main.Version == "" {
		return "unknown"
	}

	if info.Main.Sum != "" {
		return fmt.Sprintf("%s (%s)", info.Main.Version, info.Main.Sum)
	}

	return info.Main.Version
}
