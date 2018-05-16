function test() {
	cd ../blockchain/network
	ls -a >log.txt
}

function generateArtifacts() {
	cd ../blockchain/network
	source generateArtifacts.sh >gen.txt
}

function networkUp() {

}
