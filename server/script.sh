
function generateArtifacts(){
	cd ../blockchain/network
	source generateArtifacts.sh >gen.txt
}

function test(){
	cd ../blockchain/network
	ls >log.txt
}

test