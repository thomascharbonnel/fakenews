all: fakenews

fakenews: main.go error.go html_output.go headline.go
	go build -o fakenews main.go error.go html_output.go headline.go

clean: fakenews
	rm fakenews
