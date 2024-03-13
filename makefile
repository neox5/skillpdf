run: build
	@bin/profileskills

build:
	@go build -o bin/profileskills cmd/profileskills/main.go

build-win:
	@GOOS=windows GOARCH=amd64 go build -ldflags "-s -w"  -o bin/profileskills.exe cmd/profileskills/main.go 

# Converts fonts for gofpdf from Google Fonts ttf files
fonts:
	@echo "Make font: Montserrat"
	@tools/makefont/makefont --embed --enc=tools/makefont/cp1252.map --dst=internal/fonts internal/fonts/Montserrat-Regular.ttf
	@echo "Make font: Ovo"
	@tools/makefont/makefont --embed --enc=tools/makefont/cp1252.map --dst=internal/fonts internal/fonts/Ovo-Regular.ttf