#set -e
function test() {
	docker ps > log.txt
}

test