package main

// This will be filled in by the compiler
var (
	GoVersion string
	GitCommit string
	Version   string
	BuildTime string
)

// A pre-release marker for the version. If this is "" (empty string)
// then it means that it is a final release. Otherwise, this is a pre-release
// such as "dev" (in development), "beta", "rc1", etc.
const VersionPrerelease = ""
