all:	bin/env
	@echo "Launching at http://localhost:5050/"
	foreman start -p 5050

bin/env:
	GOBIN=bin go install
