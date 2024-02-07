// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {MockERC20} from "@rari-capital/solmate/src/test/utils/mocks/MockERC20.sol";
import {WETH} from "@rari-capital/solmate/src/tokens/WETH.sol";

import {L2GatewayBaseTest} from "./base/L2GatewayBase.t.sol";
import {Predeploys} from "../libraries/constants/Predeploys.sol";
import {AddressAliasHelper} from "../libraries/common/AddressAliasHelper.sol";
import {L2WETHGateway} from "../L2/gateways/L2WETHGateway.sol";
import {IL2ERC20Gateway} from "../L2/gateways/IL2ERC20Gateway.sol";
import {L2GatewayRouter} from "../L2/gateways/L2GatewayRouter.sol";
import {IL1ERC20Gateway} from "../L1/gateways/IL1ERC20Gateway.sol";
import {MockCrossDomainMessenger} from "../mock/MockCrossDomainMessenger.sol";
import {L2CrossDomainMessenger} from "../L2/L2CrossDomainMessenger.sol";

contract L2WETHGatewayTest is L2GatewayBaseTest {
    event WithdrawERC20(
        address indexed _l1Token,
        address indexed _l2Token,
        address indexed _from,
        address _to,
        uint256 _amount,
        bytes _data
    );
    event FinalizeDepositERC20(
        address indexed _l1Token,
        address indexed _l2Token,
        address indexed _from,
        address _to,
        uint256 _amount,
        bytes _data
    );

    WETH private l1weth;
    WETH private l2weth;

    L2WETHGateway private gateway;
    L2GatewayRouter private router;
    L2CrossDomainMessenger private l2Messenger;

    address private counterpartGateway;
    address private feeVault;
    address private l1Messenger;

    function setUp() public override {
        super.setUp();
        // Deploy tokens
        l1weth = new WETH();
        l2weth = new WETH();

        _deployWETH(address(l1weth), address(l2weth));
        gateway = l2WETHGateway;
        router = l2GatewayRouter;
        counterpartGateway = gateway.counterpart();
        l2Messenger = l2CrossDomainMessenger;
        feeVault = l2FeeVault;
        l1Messenger = address(NON_ZERO_ADDRESS);

        hevm.prank(multisig);
        router.setDefaultERC20Gateway(address(gateway));
        // Prepare token balances
        l2weth.deposit{value: address(this).balance / 2}();
        l2weth.approve(address(gateway), type(uint256).max);
    }

    function testDirectTransferETH(uint256 amount) public {
        amount = bound(amount, 0, address(this).balance);
        // solhint-disable-next-line avoid-low-level-calls
        (bool success, bytes memory result) = address(gateway).call{
            value: amount
        }("");
        assertBoolEq(success, false);
        assertEq(
            string(result),
            string(abi.encodeWithSignature("Error(string)", "only WETH"))
        );
    }

    function testWithdrawERC20(
        uint256 amount,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _withdrawERC20(false, amount, gasLimit, feePerGas);
    }

    function testWithdrawERC20WithRecipient(
        uint256 amount,
        address recipient,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _withdrawERC20WithRecipient(
            false,
            amount,
            recipient,
            gasLimit,
            feePerGas
        );
    }

    function testWithdrawERC20WithRecipientAndCalldata(
        uint256 amount,
        address recipient,
        bytes memory dataToCall,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _withdrawERC20WithRecipientAndCalldata(
            false,
            amount,
            recipient,
            dataToCall,
            gasLimit,
            feePerGas
        );
    }

    function testRouterDepositERC20(
        uint256 amount,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _withdrawERC20(true, amount, gasLimit, feePerGas);
    }

    function testRouterDepositERC20WithRecipient(
        uint256 amount,
        address recipient,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _withdrawERC20WithRecipient(
            true,
            amount,
            recipient,
            gasLimit,
            feePerGas
        );
    }

    function testRouterDepositERC20WithRecipientAndCalldata(
        uint256 amount,
        address recipient,
        bytes memory dataToCall,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _withdrawERC20WithRecipientAndCalldata(
            true,
            amount,
            recipient,
            dataToCall,
            gasLimit,
            feePerGas
        );
    }

    function testFinalizeDepositERC20FailedMocking(
        address sender,
        address recipient,
        uint256 amount,
        bytes memory dataToCall
    ) public {
        amount = bound(amount, 1, 100000);

        // revert when caller is not messenger
        hevm.expectRevert("only messenger can call");
        gateway.finalizeDepositERC20(
            address(l1weth),
            address(l2weth),
            sender,
            recipient,
            amount,
            dataToCall
        );

        MockCrossDomainMessenger mockMessenger = new MockCrossDomainMessenger();
        hevm.store(
            address(gateway),
            bytes32(eth_erc20_messenger_slot),
            bytes32(abi.encode(address(mockMessenger)))
        );

        // only call by counterpart
        hevm.expectRevert("only call by counterpart");
        mockMessenger.callTarget(
            address(gateway),
            abi.encodeWithSelector(
                gateway.finalizeDepositERC20.selector,
                address(l1weth),
                address(l2weth),
                sender,
                recipient,
                amount,
                dataToCall
            )
        );

        mockMessenger.setXDomainMessageSender(address(counterpartGateway));

        // l1 token not WETH
        hevm.expectRevert("l1 token not WETH");
        mockMessenger.callTarget(
            address(gateway),
            abi.encodeWithSelector(
                gateway.finalizeDepositERC20.selector,
                address(l2weth),
                address(l2weth),
                sender,
                recipient,
                amount,
                dataToCall
            )
        );

        // l2 token not WETH
        hevm.expectRevert("l2 token not WETH");
        mockMessenger.callTarget(
            address(gateway),
            abi.encodeWithSelector(
                gateway.finalizeDepositERC20.selector,
                address(l1weth),
                address(l1weth),
                sender,
                recipient,
                amount,
                dataToCall
            )
        );

        // msg.value mismatch
        hevm.expectRevert("msg.value mismatch");
        mockMessenger.callTarget(
            address(gateway),
            abi.encodeWithSelector(
                gateway.finalizeDepositERC20.selector,
                address(l1weth),
                address(l2weth),
                sender,
                recipient,
                amount,
                dataToCall
            )
        );
    }

    function testFinalizeDepositERC20Failed(
        address sender,
        address recipient,
        uint256 amount,
        bytes memory dataToCall
    ) public {
        // blacklist some addresses
        hevm.assume(recipient != address(0));
        hevm.assume(recipient != address(gateway));

        amount = bound(amount, 1, l2weth.balanceOf(address(this)));

        // send some WETH to L2CrossDomainMessenger
        gateway.withdrawERC20(address(l2weth), amount, 21000);

        // do finalize withdraw eth
        bytes memory message = abi.encodeWithSelector(
            IL2ERC20Gateway.finalizeDepositERC20.selector,
            address(l1weth),
            address(l2weth),
            sender,
            recipient,
            amount,
            dataToCall
        );
        bytes memory xDomainCalldata = abi.encodeWithSignature(
            "relayMessage(address,address,uint256,uint256,bytes)",
            address(uint160(address(counterpartGateway)) + 1),
            address(gateway),
            amount,
            0,
            message
        );

        // counterpart is not L1WETHGateway
        // emit FailedRelayedMessage from L2CrossDomainMessenger
        hevm.expectEmit(true, false, false, true);
        emit FailedRelayedMessage(keccak256(xDomainCalldata));

        uint256 messengerBalance = address(l2Messenger).balance;
        uint256 recipientBalance = l2weth.balanceOf(recipient);
        assertBoolEq(
            false,
            l2Messenger.isL1MessageExecuted(keccak256(xDomainCalldata))
        );
        hevm.startPrank(
            AddressAliasHelper.applyL1ToL2Alias(address(l1Messenger))
        );
        l2Messenger.relayMessage(
            address(uint160(address(counterpartGateway)) + 1),
            address(gateway),
            amount,
            0,
            message
        );
        hevm.stopPrank();
        assertEq(messengerBalance, address(l2Messenger).balance);
        assertEq(recipientBalance, l2weth.balanceOf(recipient));
        assertBoolEq(
            false,
            l2Messenger.isL1MessageExecuted(keccak256(xDomainCalldata))
        );
    }

    function testFinalizeDepositERC20(
        address sender,
        uint256 amount,
        bytes memory dataToCall
    ) public {
        address recipient = address(2048);

        amount = bound(amount, 1, l2weth.balanceOf(address(this)));

        // send some WETH to L1CrossDomainMessenger
        gateway.withdrawERC20(address(l2weth), amount, 21000);

        // do finalize withdraw weth
        bytes memory message = abi.encodeWithSelector(
            IL2ERC20Gateway.finalizeDepositERC20.selector,
            address(l1weth),
            address(l2weth),
            sender,
            address(recipient),
            amount,
            dataToCall
        );
        bytes memory xDomainCalldata = abi.encodeWithSignature(
            "relayMessage(address,address,uint256,uint256,bytes)",
            address(counterpartGateway),
            address(gateway),
            amount,
            0,
            message
        );

        // emit FinalizeDepositERC20 from L2WETHGateway
        {
            hevm.expectEmit(true, true, true, true);
            emit FinalizeDepositERC20(
                address(l1weth),
                address(l2weth),
                sender,
                address(recipient),
                amount,
                dataToCall
            );
        }

        // emit RelayedMessage from L2CrossDomainMessenger
        {
            hevm.expectEmit(true, false, false, true);
            emit RelayedMessage(keccak256(xDomainCalldata));
        }

        uint256 messengerBalance = address(l2Messenger).balance;
        uint256 recipientBalance = l2weth.balanceOf(address(recipient));
        assertBoolEq(
            false,
            l2Messenger.isL1MessageExecuted(keccak256(xDomainCalldata))
        );
        hevm.startPrank(
            AddressAliasHelper.applyL1ToL2Alias(address(l1Messenger))
        );
        l2Messenger.relayMessage(
            address(counterpartGateway),
            address(gateway),
            amount,
            0,
            message
        );
        hevm.stopPrank();
        assertEq(messengerBalance - amount, address(l2Messenger).balance);
        assertEq(
            recipientBalance + amount,
            l2weth.balanceOf(address(recipient))
        );
        assertBoolEq(
            true,
            l2Messenger.isL1MessageExecuted(keccak256(xDomainCalldata))
        );
    }

    function _withdrawERC20(
        bool useRouter,
        uint256 amount,
        uint256 gasLimit,
        uint256 feePerGas
    ) private {
        amount = bound(amount, 0, l2weth.balanceOf(address(this)));
        gasLimit = bound(gasLimit, 21000, 1000000);
        feePerGas = 0;

        setL1BaseFee(feePerGas);

        uint256 feeToPay = feePerGas * gasLimit;
        bytes memory message = abi.encodeWithSelector(
            IL1ERC20Gateway.finalizeWithdrawERC20.selector,
            address(l1weth),
            address(l2weth),
            address(this),
            address(this),
            amount,
            new bytes(0)
        );
        bytes memory xDomainCalldata = abi.encodeWithSignature(
            "relayMessage(address,address,uint256,uint256,bytes)",
            address(gateway),
            address(counterpartGateway),
            amount,
            0,
            message
        );

        if (amount == 0) {
            hevm.expectRevert("withdraw zero amount");
            if (useRouter) {
                router.withdrawERC20{value: feeToPay}(
                    address(l2weth),
                    amount,
                    gasLimit
                );
            } else {
                gateway.withdrawERC20{value: feeToPay}(
                    address(l2weth),
                    amount,
                    gasLimit
                );
            }
        } else {
            // token is not l2WETH
            hevm.expectRevert("only WETH is allowed");
            gateway.withdrawERC20(address(l1weth), amount, gasLimit);

            // emit AppendMessage from L2MessageQueue
            {
                hevm.expectEmit(false, false, false, true);
                emit AppendMessage(0, keccak256(xDomainCalldata));
            }

            // emit SentMessage from L2CrossDomainMessenger
            {
                hevm.expectEmit(true, true, false, true);
                emit SentMessage(
                    address(gateway),
                    address(counterpartGateway),
                    amount,
                    0,
                    gasLimit,
                    message
                );
            }

            // emit WithdrawERC20 from L2WETHGateway
            hevm.expectEmit(true, true, true, true);
            emit WithdrawERC20(
                address(l1weth),
                address(l2weth),
                address(this),
                address(this),
                amount,
                new bytes(0)
            );

            uint256 messengerBalance = address(l2Messenger).balance;
            uint256 feeVaultBalance = address(feeVault).balance;
            assertEq(
                l2Messenger.messageSendTimestamp(keccak256(xDomainCalldata)),
                0
            );
            if (useRouter) {
                router.withdrawERC20{value: feeToPay}(
                    address(l2weth),
                    amount,
                    gasLimit
                );
            } else {
                gateway.withdrawERC20{value: feeToPay}(
                    address(l2weth),
                    amount,
                    gasLimit
                );
            }
            assertEq(amount + messengerBalance, address(l2Messenger).balance);
            assertEq(feeToPay + feeVaultBalance, address(feeVault).balance);
            assertGt(
                l2Messenger.messageSendTimestamp(keccak256(xDomainCalldata)),
                0
            );
        }
    }

    function _withdrawERC20WithRecipient(
        bool useRouter,
        uint256 amount,
        address recipient,
        uint256 gasLimit,
        uint256 feePerGas
    ) private {
        amount = bound(amount, 0, l2weth.balanceOf(address(this)));
        gasLimit = bound(gasLimit, 21000, 1000000);
        feePerGas = 0;

        setL1BaseFee(feePerGas);

        uint256 feeToPay = feePerGas * gasLimit;
        bytes memory message = abi.encodeWithSelector(
            IL1ERC20Gateway.finalizeWithdrawERC20.selector,
            address(l1weth),
            address(l2weth),
            address(this),
            recipient,
            amount,
            new bytes(0)
        );
        bytes memory xDomainCalldata = abi.encodeWithSignature(
            "relayMessage(address,address,uint256,uint256,bytes)",
            address(gateway),
            address(counterpartGateway),
            amount,
            0,
            message
        );

        if (amount == 0) {
            hevm.expectRevert("withdraw zero amount");
            if (useRouter) {
                router.withdrawERC20{value: feeToPay}(
                    address(l2weth),
                    recipient,
                    amount,
                    gasLimit
                );
            } else {
                gateway.withdrawERC20{value: feeToPay}(
                    address(l2weth),
                    recipient,
                    amount,
                    gasLimit
                );
            }
        } else {
            // token is not l1WETH
            hevm.expectRevert("only WETH is allowed");
            gateway.withdrawERC20(address(l1weth), recipient, amount, gasLimit);

            // emit AppendMessage from L2MessageQueue
            {
                hevm.expectEmit(false, false, false, true);
                emit AppendMessage(0, keccak256(xDomainCalldata));
            }

            // emit SentMessage from L2CrossDomainMessenger
            {
                hevm.expectEmit(true, true, false, true);
                emit SentMessage(
                    address(gateway),
                    address(counterpartGateway),
                    amount,
                    0,
                    gasLimit,
                    message
                );
            }

            // emit WithdrawERC20 from L2WETHGateway
            hevm.expectEmit(true, true, true, true);
            emit WithdrawERC20(
                address(l1weth),
                address(l2weth),
                address(this),
                recipient,
                amount,
                new bytes(0)
            );

            uint256 messengerBalance = address(l2Messenger).balance;
            uint256 feeVaultBalance = address(feeVault).balance;
            assertEq(
                l2Messenger.messageSendTimestamp(keccak256(xDomainCalldata)),
                0
            );
            if (useRouter) {
                router.withdrawERC20{value: feeToPay}(
                    address(l2weth),
                    recipient,
                    amount,
                    gasLimit
                );
            } else {
                gateway.withdrawERC20{value: feeToPay}(
                    address(l2weth),
                    recipient,
                    amount,
                    gasLimit
                );
            }
            assertEq(amount + messengerBalance, address(l2Messenger).balance);
            assertEq(feeToPay + feeVaultBalance, address(feeVault).balance);
            assertGt(
                l2Messenger.messageSendTimestamp(keccak256(xDomainCalldata)),
                0
            );
        }
    }

    function _withdrawERC20WithRecipientAndCalldata(
        bool useRouter,
        uint256 amount,
        address recipient,
        bytes memory dataToCall,
        uint256 gasLimit,
        uint256 feePerGas
    ) private {
        amount = bound(amount, 0, l2weth.balanceOf(address(this)));
        gasLimit = bound(gasLimit, 21000, 1000000);
        feePerGas = 0;

        setL1BaseFee(feePerGas);

        uint256 feeToPay = feePerGas * gasLimit;
        bytes memory message = abi.encodeWithSelector(
            IL1ERC20Gateway.finalizeWithdrawERC20.selector,
            address(l1weth),
            address(l2weth),
            address(this),
            recipient,
            amount,
            dataToCall
        );
        bytes memory xDomainCalldata = abi.encodeWithSignature(
            "relayMessage(address,address,uint256,uint256,bytes)",
            address(gateway),
            address(counterpartGateway),
            amount,
            0,
            message
        );

        if (amount == 0) {
            hevm.expectRevert("withdraw zero amount");
            if (useRouter) {
                router.withdrawERC20AndCall{value: feeToPay}(
                    address(l2weth),
                    recipient,
                    amount,
                    dataToCall,
                    gasLimit
                );
            } else {
                gateway.withdrawERC20AndCall{value: feeToPay}(
                    address(l2weth),
                    recipient,
                    amount,
                    dataToCall,
                    gasLimit
                );
            }
        } else {
            // token is not l1WETH
            hevm.expectRevert("only WETH is allowed");
            gateway.withdrawERC20AndCall(
                address(l1weth),
                recipient,
                amount,
                dataToCall,
                gasLimit
            );

            // emit AppendMessage from L2MessageQueue
            {
                hevm.expectEmit(false, false, false, true);
                emit AppendMessage(0, keccak256(xDomainCalldata));
            }

            // emit SentMessage from L2CrossDomainMessenger
            {
                hevm.expectEmit(true, true, false, true);
                emit SentMessage(
                    address(gateway),
                    address(counterpartGateway),
                    amount,
                    0,
                    gasLimit,
                    message
                );
            }

            // emit WithdrawERC20 from L2WETHGateway
            hevm.expectEmit(true, true, true, true);
            emit WithdrawERC20(
                address(l1weth),
                address(l2weth),
                address(this),
                recipient,
                amount,
                dataToCall
            );

            uint256 messengerBalance = address(l2Messenger).balance;
            uint256 feeVaultBalance = address(feeVault).balance;
            assertEq(
                l2Messenger.messageSendTimestamp(keccak256(xDomainCalldata)),
                0
            );
            if (useRouter) {
                router.withdrawERC20AndCall{value: feeToPay}(
                    address(l2weth),
                    recipient,
                    amount,
                    dataToCall,
                    gasLimit
                );
            } else {
                gateway.withdrawERC20AndCall{value: feeToPay}(
                    address(l2weth),
                    recipient,
                    amount,
                    dataToCall,
                    gasLimit
                );
            }
            assertEq(amount + messengerBalance, address(l2Messenger).balance);
            assertEq(feeToPay + feeVaultBalance, address(feeVault).balance);
            assertGt(
                l2Messenger.messageSendTimestamp(keccak256(xDomainCalldata)),
                0
            );
        }
    }
}
