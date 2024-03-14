// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morph-l2/bindings/solc"
)

const L2SequencerStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L2/staking/L2Sequencer.sol:L2Sequencer\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/L2/staking/L2Sequencer.sol:L2Sequencer\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/L2/staking/L2Sequencer.sol:L2Sequencer\",\"label\":\"currentVersion\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_uint256\"},{\"astId\":1003,\"contract\":\"contracts/L2/staking/L2Sequencer.sol:L2Sequencer\",\"label\":\"currentVersionHeight\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_uint256\"},{\"astId\":1004,\"contract\":\"contracts/L2/staking/L2Sequencer.sol:L2Sequencer\",\"label\":\"preVersion\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_uint256\"},{\"astId\":1005,\"contract\":\"contracts/L2/staking/L2Sequencer.sol:L2Sequencer\",\"label\":\"preVersionHeight\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_uint256\"},{\"astId\":1006,\"contract\":\"contracts/L2/staking/L2Sequencer.sol:L2Sequencer\",\"label\":\"sequencerAddresses\",\"offset\":0,\"slot\":\"5\",\"type\":\"t_array(t_address)dyn_storage\"},{\"astId\":1007,\"contract\":\"contracts/L2/staking/L2Sequencer.sol:L2Sequencer\",\"label\":\"preSequencerAddresses\",\"offset\":0,\"slot\":\"6\",\"type\":\"t_array(t_address)dyn_storage\"},{\"astId\":1008,\"contract\":\"contracts/L2/staking/L2Sequencer.sol:L2Sequencer\",\"label\":\"sequencerInfos\",\"offset\":0,\"slot\":\"7\",\"type\":\"t_array(t_struct(SequencerInfo)1010_storage)dyn_storage\"},{\"astId\":1009,\"contract\":\"contracts/L2/staking/L2Sequencer.sol:L2Sequencer\",\"label\":\"preSequencerInfos\",\"offset\":0,\"slot\":\"8\",\"type\":\"t_array(t_struct(SequencerInfo)1010_storage)dyn_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_address)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"address[]\",\"numberOfBytes\":\"32\"},\"t_array(t_struct(SequencerInfo)1010_storage)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"struct Types.SequencerInfo[]\",\"numberOfBytes\":\"32\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_bytes_storage\":{\"encoding\":\"bytes\",\"label\":\"bytes\",\"numberOfBytes\":\"32\"},\"t_struct(SequencerInfo)1010_storage\":{\"encoding\":\"inplace\",\"label\":\"struct Types.SequencerInfo\",\"numberOfBytes\":\"96\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L2SequencerStorageLayout = new(solc.StorageLayout)

var L2SequencerDeployedBin = "0x608060405234801561000f575f80fd5b5060043610610149575f3560e01c8063aeaf9f41116100c7578063d1c55fe31161007d578063dd967ee911610063578063dd967ee91461030e578063e597c19e14610321578063f81e02a714610341575f80fd5b8063d1c55fe3146102db578063d958646714610305575f80fd5b8063be6c5d68116100ad578063be6c5d681461029f578063c9406b1a146102bf578063cfd1eff3146102d2575f80fd5b8063aeaf9f411461026a578063b95cdb781461027d575f80fd5b80635942e7c71161011c578063927ede2d11610102578063927ede2d146102275780639d888e861461024e578063ad01732f14610257575f80fd5b80635942e7c7146101ff5780637ad9e3ac14610214575f80fd5b8063342b63451461014d5780633cb747bf1461017a5780634a3c980c146101c15780634bbf5252146101d8575b5f80fd5b61016061015b366004611291565b610368565b604080519283526020830191909152015b60405180910390f35b7f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610171565b6101ca60045481565b604051908152602001610171565b61019c7f000000000000000000000000000000000000000000000000000000000000000081565b61021261020d3660046114c4565b61049d565b005b6101606102223660046114fe565b61074b565b61019c7f000000000000000000000000000000000000000000000000000000000000000081565b6101ca60015481565b61021261026536600461151e565b610771565b61019c610278366004611562565b610b77565b61029061028b366004611562565b610bac565b604051610171939291906115da565b6102b26102ad3660046114fe565b610c7e565b6040516101719190611617565b6102906102cd366004611562565b610eaf565b6101ca60035481565b6102ee6102e9366004611291565b610ebe565b604080519215158352602083019190915201610171565b6101ca60025481565b61019c61031c366004611562565b610fa2565b61033461032f3660046114fe565b610fb1565b60405161017191906116ca565b61019c7f000000000000000000000000000000000000000000000000000000000000000081565b5f808315610437575f5b6006548110156103cf576006818154811061038f5761038f611723565b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff908116908516036103c7576003549092509050610496565b600101610372565b506040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f73657175656e636572206e6f742065786973740000000000000000000000000060448201526064015b60405180910390fd5b5f5b6005548110156103cf576005818154811061045657610456611723565b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff9081169085160361048e576001549092509050610496565b600101610439565b9250929050565b5f54610100900460ff16158080156104bb57505f54600160ff909116105b806104d45750303b1580156104d457505f5460ff166001145b610560576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161042e565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156105bc575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b5f5b82518110156106e35760058382815181106105db576105db611723565b6020908102919091018101515182546001810184555f938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909216919091179055825160079084908390811061065257610652611723565b6020908102919091018101518254600180820185555f94855293839020825160039092020180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9092169190911781559181015192820192909255604082015160028201906106d890826117ed565b5050506001016105be565b508015610747575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a15b5050565b5f8082156107625750506006546003549092909150565b50506005546001549092909150565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614801561088d57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610851573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108759190611909565b73ffffffffffffffffffffffffffffffffffffffff16145b610919576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603f60248201527f53657175656e6365723a2066756e6374696f6e2063616e206f6e6c792062652060448201527f63616c6c65642066726f6d20746865206f746865722073657175656e63657200606482015260840161042e565b6040517f16e2994a00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906316e2994a9061098c90600590600401611977565b5f604051808303815f87803b1580156109a3575f80fd5b505af11580156109b5573d5f803e3d5ffd5b5050600154600355506109cb905060085f611091565b6109d660065f6110b2565b600780546109e6916008916110cd565b50600580546109f79160069161117d565b506002546004556001829055610a0e60075f611091565b610a1960055f6110b2565b436002555f5b8151811015610b44576005828281518110610a3c57610a3c611723565b6020908102919091018101515182546001810184555f938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9092169190911790558151600790839083908110610ab357610ab3611723565b6020908102919091018101518254600180820185555f94855293839020825160039092020180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909216919091178155918101519282019290925560408201516002820190610b3990826117ed565b505050600101610a1f565b507f71e1b9989bdd3dbcfe04813f0785646335737b50dd32355cc19eeb58d618279660058360405161073e929190611989565b60058181548110610b86575f80fd5b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b60088181548110610bbb575f80fd5b5f91825260209091206003909102018054600182015460028301805473ffffffffffffffffffffffffffffffffffffffff9093169450909291610bfd90611750565b80601f0160208091040260200160405190810160405280929190818152602001828054610c2990611750565b8015610c745780601f10610c4b57610100808354040283529160200191610c74565b820191905f5260205f20905b815481529060010190602001808311610c5757829003601f168201915b5050505050905083565b60608115610da0576008805480602002602001604051908101604052809291908181526020015f905b82821015610d95575f8481526020908190206040805160608101825260038602909201805473ffffffffffffffffffffffffffffffffffffffff16835260018101549383019390935260028301805492939291840191610d0690611750565b80601f0160208091040260200160405190810160405280929190818152602001828054610d3290611750565b8015610d7d5780601f10610d5457610100808354040283529160200191610d7d565b820191905f5260205f20905b815481529060010190602001808311610d6057829003601f168201915b50505050508152505081526020019060010190610ca7565b505050509050919050565b6007805480602002602001604051908101604052809291908181526020015f905b82821015610d95575f8481526020908190206040805160608101825260038602909201805473ffffffffffffffffffffffffffffffffffffffff16835260018101549383019390935260028301805492939291840191610e2090611750565b80601f0160208091040260200160405190810160405280929190818152602001828054610e4c90611750565b8015610e975780601f10610e6e57610100808354040283529160200191610e97565b820191905f5260205f20905b815481529060010190602001808311610e7a57829003601f168201915b50505050508152505081526020019060010190610dc1565b60078181548110610bbb575f80fd5b5f808315610f33575f5b600654811015610f265760068181548110610ee557610ee5611723565b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff90811690851603610f1e57505060035460019150610496565b600101610ec8565b50506003545f9150610496565b5f5b600554811015610f935760058181548110610f5257610f52611723565b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff90811690851603610f8b57600180549250925050610496565b600101610f35565b50506001545f91509250929050565b60068181548110610b86575f80fd5b6060811561102657600680548060200260200160405190810160405280929190818152602001828054801561101a57602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610fef575b50505050509050919050565b600580548060200260200160405190810160405280929190818152602001828054801561101a57602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610fef5750505050509050919050565b5080545f8255600302905f5260205f20908101906110af91906111c5565b50565b5080545f8255905f5260205f20908101906110af9190611211565b828054828255905f5260205f2090600302810192821561116d575f5260205f209160030282015b8281111561116d57825482547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90911617825560018084015490830155828260028082019061115b908401826119aa565b505050916003019190600301906110f4565b506111799291506111c5565b5090565b828054828255905f5260205f209081019282156111b9575f5260205f209182015b828111156111b957825482559160010191906001019061119e565b50611179929150611211565b808211156111795780547fffffffffffffffffffffffff00000000000000000000000000000000000000001681555f600182018190556112086002830182611225565b506003016111c5565b5b80821115611179575f8155600101611212565b50805461123190611750565b5f825580601f10611240575050565b601f0160209004905f5260205f20908101906110af9190611211565b8035801515811461126b575f80fd5b919050565b73ffffffffffffffffffffffffffffffffffffffff811681146110af575f80fd5b5f80604083850312156112a2575f80fd5b6112ab8361125c565b915060208301356112bb81611270565b809150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6040516060810167ffffffffffffffff81118282101715611316576113166112c6565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611363576113636112c6565b604052919050565b5f601f83601f84011261137c575f80fd5b8235602067ffffffffffffffff80831115611399576113996112c6565b8260051b6113a883820161131c565b93845286810183019383810190898611156113c1575f80fd5b84890192505b858310156114b7578235848111156113dd575f80fd5b890160607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828d038101821315611412575f80fd5b61141a6112f3565b8884013561142781611270565b81526040848101358a830152928401359288841115611444575f80fd5b83850194508e603f860112611457575f80fd5b8985013593508884111561146d5761146d6112c6565b61147c8a848e8701160161131c565b92508383528e81858701011115611491575f80fd5b838186018b8501375f9383018a01939093529182015283525091840191908401906113c7565b9998505050505050505050565b5f602082840312156114d4575f80fd5b813567ffffffffffffffff8111156114ea575f80fd5b6114f68482850161136b565b949350505050565b5f6020828403121561150e575f80fd5b6115178261125c565b9392505050565b5f806040838503121561152f575f80fd5b82359150602083013567ffffffffffffffff81111561154c575f80fd5b6115588582860161136b565b9150509250929050565b5f60208284031215611572575f80fd5b5035919050565b5f81518084525f5b8181101561159d57602081850181015186830182015201611581565b505f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b73ffffffffffffffffffffffffffffffffffffffff84168152826020820152606060408201525f61160e6060830184611579565b95945050505050565b5f60208083018184528085518083526040925060408601915060408160051b8701018488015f5b838110156116bc578883037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc00185528151805173ffffffffffffffffffffffffffffffffffffffff168452878101518885015286015160608785018190526116a881860183611579565b96890196945050509086019060010161163e565b509098975050505050505050565b602080825282518282018190525f9190848201906040850190845b8181101561171757835173ffffffffffffffffffffffffffffffffffffffff16835292840192918401916001016116e5565b50909695505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b600181811c9082168061176457607f821691505b60208210810361179b577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b601f8211156117e857805f5260205f20601f840160051c810160208510156117c65750805b601f840160051c820191505b818110156117e5575f81556001016117d2565b50505b505050565b815167ffffffffffffffff811115611807576118076112c6565b61181b816118158454611750565b846117a1565b602080601f83116001811461186d575f84156118375750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555611901565b5f858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b828110156118b95788860151825594840194600190910190840161189a565b50858210156118f557878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b505060018460011b0185555b505050505050565b5f60208284031215611919575f80fd5b815161151781611270565b5f815480845260208085019450835f5260205f205f5b8381101561196c57815473ffffffffffffffffffffffffffffffffffffffff168752958201956001918201910161193a565b509495945050505050565b602081525f6115176020830184611924565b604081525f61199b6040830185611924565b90508260208301529392505050565b8181036119b5575050565b6119bf8254611750565b67ffffffffffffffff8111156119d7576119d76112c6565b6119e5816118158454611750565b5f601f821160018114611a35575f83156119ff5750848201545b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b1784556117e5565b5f85815260208082208683529082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616925b83811015611a895782860154825560019586019590910190602001611a69565b5085831015611ac557818501547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b0190555056fea164736f6c6343000818000a"

func init() {
	if err := json.Unmarshal([]byte(L2SequencerStorageLayoutJSON), L2SequencerStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2Sequencer"] = L2SequencerStorageLayout
	deployedBytecodes["L2Sequencer"] = L2SequencerDeployedBin
}
