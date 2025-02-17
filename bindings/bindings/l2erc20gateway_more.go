// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"morph-l2/bindings/solc"
)

const L2ERC20GatewayStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/l2/gateways/L2ERC20Gateway.sol:L2ERC20Gateway\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/l2/gateways/L2ERC20Gateway.sol:L2ERC20Gateway\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/l2/gateways/L2ERC20Gateway.sol:L2ERC20Gateway\",\"label\":\"_status\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_uint256\"},{\"astId\":1003,\"contract\":\"contracts/l2/gateways/L2ERC20Gateway.sol:L2ERC20Gateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_array(t_uint256)1013_storage\"},{\"astId\":1004,\"contract\":\"contracts/l2/gateways/L2ERC20Gateway.sol:L2ERC20Gateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_array(t_uint256)1014_storage\"},{\"astId\":1005,\"contract\":\"contracts/l2/gateways/L2ERC20Gateway.sol:L2ERC20Gateway\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_address\"},{\"astId\":1006,\"contract\":\"contracts/l2/gateways/L2ERC20Gateway.sol:L2ERC20Gateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_array(t_uint256)1013_storage\"},{\"astId\":1007,\"contract\":\"contracts/l2/gateways/L2ERC20Gateway.sol:L2ERC20Gateway\",\"label\":\"counterpart\",\"offset\":0,\"slot\":\"151\",\"type\":\"t_address\"},{\"astId\":1008,\"contract\":\"contracts/l2/gateways/L2ERC20Gateway.sol:L2ERC20Gateway\",\"label\":\"router\",\"offset\":0,\"slot\":\"152\",\"type\":\"t_address\"},{\"astId\":1009,\"contract\":\"contracts/l2/gateways/L2ERC20Gateway.sol:L2ERC20Gateway\",\"label\":\"messenger\",\"offset\":0,\"slot\":\"153\",\"type\":\"t_address\"},{\"astId\":1010,\"contract\":\"contracts/l2/gateways/L2ERC20Gateway.sol:L2ERC20Gateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"154\",\"type\":\"t_array(t_uint256)1012_storage\"},{\"astId\":1011,\"contract\":\"contracts/l2/gateways/L2ERC20Gateway.sol:L2ERC20Gateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"200\",\"type\":\"t_array(t_uint256)1014_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)1012_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[46]\",\"numberOfBytes\":\"1472\"},\"t_array(t_uint256)1013_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)1014_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L2ERC20GatewayStorageLayout = new(solc.StorageLayout)

var L2ERC20GatewayDeployedBin = "0x"

func init() {
	if err := json.Unmarshal([]byte(L2ERC20GatewayStorageLayoutJSON), L2ERC20GatewayStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2ERC20Gateway"] = L2ERC20GatewayStorageLayout
	deployedBytecodes["L2ERC20Gateway"] = L2ERC20GatewayDeployedBin
}
