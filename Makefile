update-poste-go:
	go get github.com/ondrejsika/poste-go@master

release:
	goreleaser --rm-dist
