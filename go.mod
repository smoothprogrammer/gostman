module github.com/smoothprogrammer/gostman

go 1.17

require (
	github.com/sirupsen/logrus v1.8.1
	github.com/smoothprogrammer/testr v1.0.0
	gopkg.in/yaml.v2 v2.4.0
)

require golang.org/x/sys v0.0.0-20191026070338-33540a1f6037 // indirect

// Old github handle.
retract (
	v0.1.0
	v0.1.1
)
