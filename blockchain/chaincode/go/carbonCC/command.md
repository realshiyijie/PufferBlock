# carbonCC链码操作指令示例

## 环境变量设置

### 在执行每条`peer chaincode`命令前需要指定

#### 1.选择 peer: org1/peer0 执行`peer chaincode`命令

```shell
CORE_PEER_LOCALMSPID="Org1MSP"
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
CORE_PEER_ADDRESS=peer0.org1.example.com:7051
```

#### 2.选择 peer: org1/peer1 执行`peer chaincode`命令

```shell
CORE_PEER_LOCALMSPID="Org1MSP"
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
CORE_PEER_ADDRESS=peer1.org1.example.com:7051
```

#### 3.选择 peer: org2/peer2 执行`peer chaincode`命令

```shell
CORE_PEER_LOCALMSPID="Org2MSP"
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
CORE_PEER_ADDRESS=peer0.org2.example.com:7051
```

#### 4.选择 peer: org2/peer3 执行`peer chaincode`命令

```shell
CORE_PEER_LOCALMSPID="Org2MSP"
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
CORE_PEER_ADDRESS=peer1.org2.example.com:7051
```

## 执行`peer chaincode`命令

### 初始化账户

#### 用户向认证中心注册，登记上链之后，为新用户初始化账户信息

```shell
```

### 更新账户信息

#### 运行过程中，可以对用户账户信息进行更新

```shell
```

### 查询账户信息

#### 查询所有账户信息

```shell
```

#### 查询指定用户账户信息

```shell
```

#### 查询指定类型账户信息

```shell
```

#### 查询指定额度账户信息

```shell
```

### 进行交易

#### 向指定对象账户进行交易

```shell
```