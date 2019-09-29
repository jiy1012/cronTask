fmt:
	gofmt -l -w -s ./

build: fmt
	go build -o cronTask main.go

run:
	./cronTask -conf conf/cron.yaml >> ./run.log 2>&1 &

stop:
	killall -9 "cronTask"

rerun: build stop run

