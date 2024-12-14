gits_up:
	git status
	git add .
	git commit -m "Updates"
	git push

build:
	go build -o bin/goscope cmd/goscope/main.go

clean:
	rm -rf bin/