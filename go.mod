module github.com/maldan/gam-app-english

go 1.18

// replace github.com/maldan/go-restserver => ../../../go_lib/restserver
// replace github.com/maldan/go-rapi => ../../../go_lib/rapi
replace github.com/maldan/go-cmhp => ../../go_lib/cmhp

require (
	github.com/maldan/go-cmhp v0.0.20
	github.com/maldan/go-rapi v0.0.6
)

require (
	github.com/PuerkitoBio/goquery v1.8.0 // indirect
	github.com/andybalholm/cascadia v1.3.1 // indirect
	golang.org/x/net v0.0.0-20210916014120-12bc252f5db8 // indirect
)
