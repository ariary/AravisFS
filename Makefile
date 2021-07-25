build_adret:
	@echo "build in ${PWD}";go build -o adret cmd/adret/main.go

build_ubac:
	@echo "build in ${PWD}";go build -o ubac cmd/ubac/main.go

debug_adret:
	@dlv debug github.com/ariary/AravisFS/cmd/adret

debug_ubac:
	@dlv debug github.com/ariary/AravisFS/cmd/ubac
