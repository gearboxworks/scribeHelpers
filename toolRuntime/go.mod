module github.com/gearboxworks/scribeHelpers/toolRuntime

go 1.14

replace github.com/gearboxworks/scribeHelpers/ux => ../ux

require (
	github.com/blang/semver v3.5.1+incompatible
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0
	github.com/gearboxworks/scribeHelpers/ux v0.0.0-20200604000029-dbb313f0fedc
)
