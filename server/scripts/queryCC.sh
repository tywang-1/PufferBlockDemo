#查询用脚本
FUNC=$1
OPNAME=$2
CHANNEL_NAME="mychannel"


function printHelp()
{
    echo "wrong Args"
}
if [ "${FUNC}" == "queryByOwner" ]; then
	docker exec cli /go/bin/peer -c  peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["queryByOwner",'\"${OPNAME}\"']}' 2>&1|grep "Query Result"
elif [ "${FUNC}" == "queryAllCarbonInfo" ]; then 

	docker exec cli /go/bin/peer -c  peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["queryAllCarbonInfo"]}' 2>&1|grep "Query Result"
else
	printHelp
fi