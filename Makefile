build:
	go build -o clockupdater .

run:
	go build -o clockupdater .
	./clockupdater

clean:
	go clean
	rm clockupdater

install:
	install clockupdater /usr/local/bin/clockupdater
	install -m 664 clockupdater.service /etc/systemd/system/clockupdater.service
	install -m 664 clockupdater.timer	/etc/systemd/system/clockupdater.timer
	systemctl enable clockupdater.timer
	systemctl enable clockupdater.service
	systemctl start clockupdater.timer
	systemctl daemon-reload
