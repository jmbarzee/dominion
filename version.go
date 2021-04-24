package dominion

import "github.com/blang/semver"

const versionString = "0.1.0"

func init() {
	var err error
	version, err = semver.Parse(versionString)
	if err != nil {
		panic(err)
	}
}

var version semver.Version

func Version() semver.Version {
	return version
}
