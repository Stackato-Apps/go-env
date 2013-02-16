all:	bin/env
	foreman start

bin/env:
	GOBIN=bin go install


