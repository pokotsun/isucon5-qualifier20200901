branch=master
build: 
	(git reset --hard HEAD && git fetch && git checkout ${branch})
	(cd /home/isucon/webapp/go && go build -o app)
	sudo systemctl restart nginx
	sudo systemctl restart isuxi.go 
analyze:
	sudo cp /dev/null /var/log/nginx/access.log
	(cd /home/isucon && ./bench.sh)
	sudo alp --file=/var/log/nginx/access.log ltsv -r --sort sum

.PHONY: build
