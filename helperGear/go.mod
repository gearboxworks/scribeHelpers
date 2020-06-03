module github.com/newclarity/scribeHelpers/helperGear

go 1.14

replace github.com/newclarity/scribeHelpers/ux => ../ux

replace github.com/newclarity/scribeHelpers/helperDocker => ../helperDocker

require (
    github.com/newclarity/scribeHelpers/helperDocker v0.0.0
    github.com/newclarity/scribeHelpers/ux v0.0.0-00010101000000-000000000000
)
