module github.com/maldan/gam-app-artbook

go 1.18

replace github.com/maldan/go-rapi => ../../go_lib/rapi

replace github.com/maldan/go-cmhp => ../../go_lib/cmhp

require (
	github.com/maldan/go-cmhp v0.0.23
	github.com/maldan/go-rapi v0.0.6
	github.com/maldan/go-restserver v1.2.8
)

require (
	github.com/aws/aws-sdk-go v1.40.6 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
)
