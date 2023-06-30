// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Predeploys } from "../libraries/Predeploys.sol";
import { OptimismMintableERC20 } from "../universal/OptimismMintableERC20.sol";

/// @custom:legacy
/// @custom:proxied
/// @custom:predeploy 0x420000000000000000000000000000000000000A
/// @title LegacyOVMETH
/// @notice LegacyOVMETH is a legacy contract that held WETH balances before the Bedrock upgrade.
contract LegacyOVMETH is OptimismMintableERC20 {
    /// @notice Initializes the contract as an Optimism Mintable ERC20.
    constructor()
        OptimismMintableERC20(Predeploys.L2_STANDARD_BRIDGE, address(0), "Ether", "WETH")
    {}
}
