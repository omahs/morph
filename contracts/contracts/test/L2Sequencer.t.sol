// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import "forge-std/console2.sol";
import {Predeploys} from "../libraries/constants/Predeploys.sol";
import {L2StakingBaseTest} from "./base/L2StakingBase.t.sol";
import {Types} from "../libraries/common/Types.sol";

contract L2SequencerTest is L2StakingBaseTest {
    function testUpdateSequencers() external {
        address[] memory newSequencers = new address[](SEQUENCER_SIZE);
        beginSeq = 100;
        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            address user = address(uint160(beginSeq + i));
            Types.StakerInfo memory stakerInfo = ffi.generateStakingInfo(user);
            sequencerBLSKeys.push(stakerInfo.blsKey);

            newSequencers[i] = stakerInfo.addr;
        }
        hevm.prank(address(Predeploys.L2_STAKING));
        // updateSequencers
        l2Sequencer.updateSequencerSet(newSequencers);

        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            assertEq(
                l2Sequencer.getCurrentSeqeuncerSet()[i],
                sequencerAddrs[i]
            );
        }
        hevm.roll(2);
        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            assertEq(
                l2Sequencer.getCurrentSeqeuncerSet()[i],
                sequencerAddrs[i]
            );
        }

        hevm.roll(3);
        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            assertEq(l2Sequencer.getCurrentSeqeuncerSet()[i], newSequencers[i]);
        }
    }
}
