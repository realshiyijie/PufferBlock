#ledger.sh及network.sh设置环境变量用

#设置全局变量
PEER=$1

#设置环境变量
setGlobals () {

	if [ $PEER -eq 0 -o $PEER -eq 1 ] ; then
		CORE_PEER_LOCALMSPID="Org1MSP"
		CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
		CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
		if [ $PEER -eq 0 ]; then
			CORE_PEER_ADDRESS=peer0.org1.example.com:7051
			echo "PEER${PEER} is ready"
		else
			CORE_PEER_ADDRESS=peer1.org1.example.com:7051
			echo "PEER${PEER} is ready"
		fi
	else
		CORE_PEER_LOCALMSPID="Org2MSP"
		CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
		CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
		if [ $PEER -eq 2 ]; then
			CORE_PEER_ADDRESS=peer0.org2.example.com:7051
			echo "PEER${PEER} is ready"
		else
			CORE_PEER_ADDRESS=peer1.org2.example.com:7051
			echo "PEER${PEER} is ready"
		fi
	fi

	env |grep CORE
}

#帮助
printHelp() {

	echo "setGlobals.sh-check yr chosen peer"
}

if [ $PEER -ge 0 -o $PEER le 3 ]; then
	setGlobals
else
	printHelp
fi