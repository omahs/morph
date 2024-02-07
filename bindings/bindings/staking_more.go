// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morph-l2/bindings/solc"
)

const StakingStorageLayoutJSON = "{\"storage\":null,\"types\":{}}"

var StakingStorageLayout = new(solc.StorageLayout)

var StakingDeployedBin = "0x608060405260043610610157575f3560e01c8063b2145280116100bb578063dc6e13e111610071578063f83d08ba11610057578063f83d08ba146103bf578063fc15799d146103d4578063fd5e6dd1146103e7575f80fd5b8063dc6e13e114610371578063f2fde38b146103a0575f80fd5b8063cbd679cb116100a1578063cbd679cb1461032b578063d13f90b41461033f578063d702d8aa1461035e575f80fd5b8063b2145280146102f7578063c7cd469a1461030c575f80fd5b806396ca628111610110578063a16e0674116100f6578063a16e0674146102a2578063a4d66daf146102b5578063a7044836146102d8575f80fd5b806396ca6281146102455780639b19251a14610264575f80fd5b8063715018a611610140578063715018a6146101a65780637a9262a2146101ba5780638da5cb5b14610214575f80fd5b8063672729991461015b5780636ba7ccff14610171575b5f80fd5b348015610166575f80fd5b5061016f610406565b005b34801561017c575f80fd5b5061019061018b366004612155565b610543565b60405161019d91906121af565b60405180910390f35b3480156101b1575f80fd5b5061016f6105e9565b3480156101c5575f80fd5b506101f76101d43660046121e3565b606b6020525f908152604090208054600182015460029092015490919060ff1683565b60408051938452602084019290925215159082015260600161019d565b34801561021f575f80fd5b506033546001600160a01b03165b6040516001600160a01b03909116815260200161019d565b348015610250575f80fd5b5061016f61025f36600461220f565b6105fc565b34801561026f575f80fd5b5061029261027e3660046121e3565b60686020525f908152604090205460ff1681565b604051901515815260200161019d565b61016f6102b0366004612239565b610693565b3480156102c0575f80fd5b506102ca60665481565b60405190815260200161019d565b3480156102e3575f80fd5b5060655461022d906001600160a01b031681565b348015610302575f80fd5b506102ca606d5481565b348015610317575f80fd5b5061016f61032636600461229a565b61095b565b348015610336575f80fd5b506069546102ca565b34801561034a575f80fd5b5061016f610359366004612301565b610a34565b61016f61036c366004612239565b610d17565b34801561037c575f80fd5b5061039061038b3660046121e3565b611104565b60405161019d949392919061234a565b3480156103ab575f80fd5b5061016f6103ba3660046121e3565b6111bc565b3480156103ca575f80fd5b506102ca60675481565b61016f6103e23660046123af565b611249565b3480156103f2575f80fd5b5061022d610401366004612155565b6119a5565b335f908152606b602052604090206002015460ff1680156104345750335f908152606b602052604090205415155b80156104505750335f908152606b602052604090206001015443115b6104a15760405162461bcd60e51b815260206004820152601260248201527f696e76616c6964207769746864726177616c000000000000000000000000000060448201526064015b60405180910390fd5b335f818152606b602052604080822054905181156108fc0292818181858888f193505050501580156104d5573d5f803e3d5ffd5b50335f818152606b6020908152604091829020548251938452908301527fd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a910160405180910390a1335f908152606b602052604081208181556001810191909155600201805460ff19169055565b606e8181548110610552575f80fd5b905f5260205f20015f91509050805461056a90612473565b80601f016020809104026020016040519081016040528092919081815260200182805461059690612473565b80156105e15780601f106105b8576101008083540402835291602001916105e1565b820191905f5260205f20905b8154815290600101906020018083116105c457829003601f168201915b505050505081565b6105f16119cd565b6105fa5f611a27565b565b6106046119cd565b606d54821415801561061857506069548210155b801561062357505f82115b61066f5760405162461bcd60e51b815260206004820152601b60248201527f696e76616c6964206e65772073657175656e636572732073697a6500000000006044820152606401610498565b606954606d54101561068d57606d82905561068981611a90565b5050565b50606d55565b335f908152606b602052604090206002015460ff16156106f55760405162461bcd60e51b815260206004820152601060248201527f7374616b657220697320657869746564000000000000000000000000000000006044820152606401610498565b5f6106ff33611f31565b60408051606081018252335f908152606a6020908152929020600301548152606754929350919082019061073390436124f1565b815260016020918201819052335f818152606b845260409081902085518082558686015194820194909455948101516002909501805460ff1916951515959095179094558351908152918201527f6cca423c6ffc06e62a0acc433965e074b11c28479b0449250ce3ff65ac9e39fe910160405180910390a1805b6069546107bc9060019061250a565b8110156108535760696107d08260016124f1565b815481106107e0576107e061251d565b5f91825260209091200154606980546001600160a01b03909216918390811061080b5761080b61251d565b5f91825260209091200180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03929092169190911790556001016107ad565b5060698054806108655761086561254a565b5f828152602080822083015f1990810180547fffffffffffffffffffffffff00000000000000000000000000000000000000009081169091559301909355338152606a909252604082208054909116815560018101829055906108cb60028301826120cf565b505f60039190910181905560695490036109455760655f9054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b81526004015f604051808303815f87803b15801561092b575f80fd5b505af115801561093d573d5f803e3d5ffd5b505050505050565b606d548110156106895761068982611a90565b50565b6109636119cd565b5f5b838110156109c857600160685f8787858181106109845761098461251d565b905060200201602081019061099991906121e3565b6001600160a01b0316815260208101919091526040015f20805460ff1916911515919091179055600101610965565b505f5b81811015610a2d575f60685f8585858181106109e9576109e961251d565b90506020020160208101906109fe91906121e3565b6001600160a01b0316815260208101919091526040015f20805460ff19169115159190911790556001016109cb565b5050505050565b5f54610100900460ff1615808015610a5257505f54600160ff909116105b80610a6b5750303b158015610a6b57505f5460ff166001145b610add5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610498565b5f805460ff191660011790558015610b1b575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6001600160a01b038516610b715760405162461bcd60e51b815260206004820152601a60248201527f696e76616c69642073657175656e63657220636f6e74726163740000000000006044820152606401610498565b5f8411610be65760405162461bcd60e51b815260206004820152602260248201527f73657175656e6365727353697a65206d7573742067726561746572207468616e60448201527f20300000000000000000000000000000000000000000000000000000000000006064820152608401610498565b5f8311610c5b5760405162461bcd60e51b815260206004820152602160248201527f7374616b696e67206c696d6974206d7573742067726561746572207468616e2060448201527f30000000000000000000000000000000000000000000000000000000000000006064820152608401610498565b606580547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b038716179055606d84905560668390556067829055610ca5611fcb565b610cae86611a27565b801561093d575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050505050565b335f9081526068602052604090205460ff16610d755760405162461bcd60e51b815260206004820152601060248201527f6e6f7420696e2077686974656c697374000000000000000000000000000000006044820152606401610498565b5f805b606954811015610dc957336001600160a01b031660698281548110610d9f57610d9f61251d565b5f918252602090912001546001600160a01b031603610dc15760019150610dc9565b600101610d78565b5080610e175760405162461bcd60e51b815260206004820152601060248201527f7374616b6572206e6f74206578697374000000000000000000000000000000006044820152606401610498565b5f34118015610e435750606654335f908152606a6020526040902060030154610e419034906124f1565b115b610e8f5760405162461bcd60e51b815260206004820152601860248201527f7374616b696e672076616c7565206e6f7420656e6f75676800000000000000006044820152606401610498565b335f908152606a602052604081206003018054349290610eb09084906124f1565b9091555050335f818152606a6020908152604091829020600301548251938452908301527f9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d910160405180910390a15f610f0933611f31565b6069549091505f90610f1d9060019061250a565b90505b80156110c157606a5f6069610f3660018561250a565b81548110610f4657610f4661251d565b905f5260205f20015f9054906101000a90046001600160a01b03166001600160a01b03166001600160a01b031681526020019081526020015f2060030154606a5f60698481548110610f9a57610f9a61251d565b5f9182526020808320909101546001600160a01b0316835282019290925260400190206003015411156110af575f6069610fd560018461250a565b81548110610fe557610fe561251d565b5f91825260209091200154606980546001600160a01b03909216925090839081106110125761101261251d565b5f918252602090912001546001600160a01b0316606961103360018561250a565b815481106110435761104361251d565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555080606983815481106110825761108261251d565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b03160217905550505b806110b981612577565b915050610f20565b505f6110cc33611f31565b606c5490915060ff1680156110e35750606d548210155b80156110f05750606d5481105b156110fe576110fe84611a90565b50505050565b606a6020525f90815260409020805460018201546002830180546001600160a01b0390931693919261113590612473565b80601f016020809104026020016040519081016040528092919081815260200182805461116190612473565b80156111ac5780601f10611183576101008083540402835291602001916111ac565b820191905f5260205f20905b81548152906001019060200180831161118f57829003601f168201915b5050505050908060030154905084565b6111c46119cd565b6001600160a01b0381166112405760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610498565b61095881611a27565b335f9081526068602052604090205460ff166112a75760405162461bcd60e51b815260206004820152601060248201527f6e6f7420696e2077686974656c697374000000000000000000000000000000006044820152606401610498565b5f805b6069548110156112fb57336001600160a01b0316606982815481106112d1576112d161251d565b5f918252602090912001546001600160a01b0316036112f357600191506112fb565b6001016112aa565b50801561134a5760405162461bcd60e51b815260206004820152601260248201527f616c7265616479207265676973746572656400000000000000000000000000006044820152606401610498565b335f908152606b602052604090206002015460ff16156113ac5760405162461bcd60e51b815260206004820152601060248201527f7374616b657220697320657869746564000000000000000000000000000000006044820152606401610498565b5f606d54116114235760405162461bcd60e51b815260206004820152602260248201527f73657175656e6365727353697a65206d7573742067726561746572207468616e60448201527f20300000000000000000000000000000000000000000000000000000000000006064820152608401610498565b5f8490036114735760405162461bcd60e51b815260206004820152601960248201527f696e76616c69642074656e6465726d696e74207075626b6579000000000000006044820152606401610498565b8251610100146114c55760405162461bcd60e51b815260206004820152601260248201527f696e76616c696420626c73207075626b657900000000000000000000000000006044820152606401610498565b60665434116115165760405162461bcd60e51b815260206004820152601b60248201527f7374616b696e672076616c7565206973206e6f7420656e6f75676800000000006044820152606401610498565b5f5b60695481101561165d5784606a5f606984815481106115395761153961251d565b5f9182526020808320909101546001600160a01b03168352820192909252604001902060010154036115ad5760405162461bcd60e51b815260206004820152601860248201527f746d4b657920616c7265616479207265676973746572656400000000000000006044820152606401610498565b8380519060200120606a5f606984815481106115cb576115cb61251d565b5f9182526020808320909101546001600160a01b0316835282019290925260409081019091209051611600916002019061258c565b6040518091039020036116555760405162461bcd60e51b815260206004820152601960248201527f626c734b657920616c72656164792072656769737465726564000000000000006044820152606401610498565b600101611518565b50604080516080810182523380825260208083018881528385018881523460608601525f938452606a90925293909120825181547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03909116178155925160018401555190919060028201906116db9082612647565b5060609190910151600390910155606980546001810182555f919091527f7fb4302e8e91f9110a6554c2c0a24601252c2a42c2220ca988efcfe3999143080180547fffffffffffffffffffffffff000000000000000000000000000000000000000016339081179091556040517fb7f230b53b0f914ccf820ab0618ac8320e984eeb0fb6a740785cf7fdc3b5cee3916117799187908790349061234a565b60405180910390a16069545f906117929060019061250a565b90505b801561193f57606a5f60696117ab60018561250a565b815481106117bb576117bb61251d565b905f5260205f20015f9054906101000a90046001600160a01b03166001600160a01b03166001600160a01b031681526020019081526020015f2060030154606a5f6069848154811061180f5761180f61251d565b5f9182526020808320909101546001600160a01b031683528201929092526040019020600301541115611928575f606961184a60018461250a565b8154811061185a5761185a61251d565b5f91825260209091200154606980546001600160a01b03909216925090839081106118875761188761251d565b5f918252602090912001546001600160a01b031660696118a860018561250a565b815481106118b8576118b861251d565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555080606983815481106118f7576118f761251d565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b031602179055505061192d565b61193f565b8061193781612577565b915050611795565b606c5460ff161580156119555750606d54606954145b1561197657606c805460ff1916600117905561197083611a90565b506110fe565b606c5460ff1680156119975750606d546069541115806119975750606d5481105b15610a2d57610a2d83611a90565b606981815481106119b4575f80fd5b5f918252602090912001546001600160a01b0316905081565b6033546001600160a01b031633146105fa5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610498565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b611a9b606e5f612106565b606d54606954811115611aad57506069545b5f8167ffffffffffffffff811115611ac757611ac7612382565b604051908082528060200260200182016040528015611b1357816020015b60408051606080820183525f808352602083015291810191909152815260200190600190039081611ae55790505b5090505f5b82811015611d1957606e606a5f60698481548110611b3857611b3861251d565b5f9182526020808320909101546001600160a01b031683528281019390935260409091018120835460018101855593825291902090910190611b7d9060020182612703565b506040518060600160405280606a5f60698581548110611b9f57611b9f61251d565b5f9182526020808320909101546001600160a01b0390811684528382019490945260409092018120549092168352606980549390910192606a92919086908110611beb57611beb61251d565b905f5260205f20015f9054906101000a90046001600160a01b03166001600160a01b03166001600160a01b031681526020019081526020015f20600101548152602001606a5f60698581548110611c4457611c4461251d565b5f9182526020808320909101546001600160a01b0316835282019290925260400190206002018054611c7590612473565b80601f0160208091040260200160405190810160405280929190818152602001828054611ca190612473565b8015611cec5780601f10611cc357610100808354040283529160200191611cec565b820191905f5260205f20905b815481529060010190602001808311611ccf57829003601f168201915b5050505050815250828281518110611d0657611d0661251d565b6020908102919091010152600101611b18565b505f63ad01732f60e01b60655f9054906101000a90046001600160a01b03166001600160a01b03166373452a926040518163ffffffff1660e01b8152600401602060405180830381865afa158015611d73573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611d9791906127ce565b611da29060016124f1565b83604051602401611db49291906127e5565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909316929092179091526065549091506001600160a01b03166399a2807682606e87611e323390565b6040518563ffffffff1660e01b8152600401611e519493929190612953565b5f604051808303815f87803b158015611e68575f80fd5b505af1158015611e7a573d5f803e3d5ffd5b505050507ffccf12f2113375dab7436466a54baaa4c4e5eb61ea2b6964e6716dec4cb8c198606e60655f9054906101000a90046001600160a01b03166001600160a01b03166373452a926040518163ffffffff1660e01b8152600401602060405180830381865afa158015611ef1573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611f1591906127ce565b604051611f2392919061299e565b60405180910390a150505050565b5f805b606954811015611f8257826001600160a01b031660698281548110611f5b57611f5b61251d565b5f918252602090912001546001600160a01b031603611f7a5792915050565b600101611f34565b5060405162461bcd60e51b815260206004820152601060248201527f7374616b6572206e6f74206578697374000000000000000000000000000000006044820152606401610498565b5f54610100900460ff166120475760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610498565b6105fa5f54610100900460ff166120c65760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610498565b6105fa33611a27565b5080546120db90612473565b5f825580601f106120ea575050565b601f0160209004905f5260205f20908101906109589190612121565b5080545f8255905f5260205f20908101906109589190612139565b5b80821115612135575f8155600101612122565b5090565b80821115612135575f61214c82826120cf565b50600101612139565b5f60208284031215612165575f80fd5b5035919050565b5f81518084525f5b8181101561219057602081850181015186830182015201612174565b505f602082860101526020601f19601f83011685010191505092915050565b602081525f6121c1602083018461216c565b9392505050565b80356001600160a01b03811681146121de575f80fd5b919050565b5f602082840312156121f3575f80fd5b6121c1826121c8565b803563ffffffff811681146121de575f80fd5b5f8060408385031215612220575f80fd5b82359150612230602084016121fc565b90509250929050565b5f60208284031215612249575f80fd5b6121c1826121fc565b5f8083601f840112612262575f80fd5b50813567ffffffffffffffff811115612279575f80fd5b6020830191508360208260051b8501011115612293575f80fd5b9250929050565b5f805f80604085870312156122ad575f80fd5b843567ffffffffffffffff808211156122c4575f80fd5b6122d088838901612252565b909650945060208701359150808211156122e8575f80fd5b506122f587828801612252565b95989497509550505050565b5f805f805f60a08688031215612315575f80fd5b61231e866121c8565b945061232c602087016121c8565b94979496505050506040830135926060810135926080909101359150565b6001600160a01b0385168152836020820152608060408201525f612371608083018561216c565b905082606083015295945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f805f606084860312156123c1575f80fd5b83359250602084013567ffffffffffffffff808211156123df575f80fd5b818601915086601f8301126123f2575f80fd5b81358181111561240457612404612382565b604051601f8201601f19908116603f0116810190838211818310171561242c5761242c612382565b81604052828152896020848701011115612444575f80fd5b826020860160208301375f60208483010152809650505050505061246a604085016121fc565b90509250925092565b600181811c9082168061248757607f821691505b6020821081036124be577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b80820180821115612504576125046124c4565b92915050565b81810381811115612504576125046124c4565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b5f81612585576125856124c4565b505f190190565b5f80835461259981612473565b600182811680156125b157600181146125c6576125f2565b60ff19841687528215158302870194506125f2565b875f526020805f205f5b858110156125e95781548a8201529084019082016125d0565b50505082870194505b50929695505050505050565b601f82111561264257805f5260205f20601f840160051c810160208510156126235750805b601f840160051c820191505b81811015610a2d575f815560010161262f565b505050565b815167ffffffffffffffff81111561266157612661612382565b6126758161266f8454612473565b846125fe565b602080601f8311600181146126a8575f84156126915750858301515b5f19600386901b1c1916600185901b17855561093d565b5f85815260208120601f198616915b828110156126d6578886015182559484019460019091019084016126b7565b50858210156126f357878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b81810361270e575050565b6127188254612473565b67ffffffffffffffff81111561273057612730612382565b61273e8161266f8454612473565b5f601f82116001811461276f575f83156127585750848201545b5f19600385901b1c1916600184901b178455610a2d565b5f8581526020808220868352908220601f198616925b838110156127a55782860154825560019586019590910190602001612785565b50858310156126f3579301545f1960f8600387901b161c19169092555050600190811b01905550565b5f602082840312156127de575f80fd5b5051919050565b5f604080830185845260206040818601528186518084526060935060608701915060608160051b8801018389015f5b83811015612882578983037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0018552815180516001600160a01b03168452868101518785015288015188840188905261286f8885018261216c565b9587019593505090850190600101612814565b50909a9950505050505050505050565b5f828254808552602080860195506005818360051b850101865f52825f205f5b8581101561294557601f19878403018a525f82546128cf81612473565b808652600182811680156128ea57600181146129035761292e565b60ff1984168a890152898315158a1b890101945061292e565b865f52895f205f5b848110156129265781548a82018d0152908301908b0161290b565b89018b019550505b509c88019c929550505091909101906001016128b2565b509098975050505050505050565b608081525f612965608083018761216c565b82810360208401526129778187612892565b91505063ffffffff841660408301526001600160a01b038316606083015295945050505050565b604081525f6129b06040830185612892565b9050826020830152939250505056fea2646970667358221220837b4db63d36be87e4ad7d666aa3c0f85d692c45c23566353a3f668c071d9cdf64736f6c63430008180033"

func init() {
	if err := json.Unmarshal([]byte(StakingStorageLayoutJSON), StakingStorageLayout); err != nil {
		panic(err)
	}

	layouts["Staking"] = StakingStorageLayout
	deployedBytecodes["Staking"] = StakingDeployedBin
}
