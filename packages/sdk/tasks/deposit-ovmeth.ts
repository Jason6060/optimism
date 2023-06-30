import { promises as fs } from 'fs'

import { task, types } from 'hardhat/config'
import '@nomiclabs/hardhat-ethers'
import 'hardhat-deploy'
import { providers, utils } from 'ethers'
import { predeploys, sleep } from '@eth-optimism/core-utils'
import Artifact__L2ToL1MessagePasser from '@eth-optimism/contracts-bedrock/forge-artifacts/L2ToL1MessagePasser.sol/L2ToL1MessagePasser.json'
import Artifact__L2CrossDomainMessenger from '@eth-optimism/contracts-bedrock/forge-artifacts/L2CrossDomainMessenger.sol/L2CrossDomainMessenger.json'
import Artifact__L2StandardBridge from '@eth-optimism/contracts-bedrock/forge-artifacts/L2StandardBridge.sol/L2StandardBridge.json'
import Artifact__OptimismPortal from '@eth-optimism/contracts-bedrock/forge-artifacts/OptimismPortal.sol/OptimismPortal.json'
import Artifact__L1CrossDomainMessenger from '@eth-optimism/contracts-bedrock/forge-artifacts/L1CrossDomainMessenger.sol/L1CrossDomainMessenger.json'
import Artifact__L1StandardBridge from '@eth-optimism/contracts-bedrock/forge-artifacts/L1StandardBridge.sol/L1StandardBridge.json'
import Artifact__LegacyOVMETH from '@eth-optimism/contracts-bedrock/forge-artifacts/LegacyOVMETH.sol/LegacyOVMETH.json'

import {
  CrossChainMessenger,
  MessageStatus,
  CONTRACT_ADDRESSES,
  OEContractsLike,
  DEFAULT_L2_CONTRACT_ADDRESSES,
} from '../src'

const { formatEther } = utils

task('deposit-ovmeth', 'Deposits ether onto L2 ovm_eth(WETH).')
  .addParam(
    'l2ProviderUrl',
    'L2 provider URL.',
    'http://localhost:9545',
    types.string
  )
  .addOptionalParam('to', 'Recipient of the ether', '', types.string)
  .addOptionalParam(
    'amount',
    'Amount of ether to send (in ETH)',
    '',
    types.string
  )
  .addOptionalParam(
    'withdraw',
    'Follow up with a withdrawal',
    true,
    types.boolean
  )
  .addOptionalParam(
    'l1ContractsJsonPath',
    'Path to a JSON with L1 contract addresses in it',
    '',
    types.string
  )
  .addOptionalParam('withdrawAmount', 'Amount to withdraw', '', types.string)
  .addParam(
    'opNodeProviderUrl',
    'op-node provider URL',
    'http://localhost:7545',
    types.string
  )
  .setAction(async (args, hre) => {
    const signers = await hre.ethers.getSigners()
    if (signers.length === 0) {
      throw new Error('No configured signers')
    }
    // Use the first configured signer for simplicity
    const signer = signers[0]
    const address = await signer.getAddress()
    console.log(`Using signer ${address}`)

    // Ensure that the signer has a balance before trying to
    // do anything
    const balance = await signer.getBalance()
    if (balance.eq(0)) {
      throw new Error('Signer has no balance')
    }

    const l2Provider = new providers.StaticJsonRpcProvider(args.l2ProviderUrl)

    // send to self if not specified
    const to = args.to ? args.to : address
    const amount = args.amount
      ? utils.parseEther(args.amount)
      : utils.parseEther('1')
    const withdrawAmount = args.withdrawAmount
      ? utils.parseEther(args.withdrawAmount)
      : amount.div(2)

    const l2Signer = new hre.ethers.Wallet(
      hre.network.config.accounts[0],
      l2Provider
    )

    const l2ChainId = await l2Signer.getChainId()
    let contractAddrs = CONTRACT_ADDRESSES[l2ChainId]
    if (args.l1ContractsJsonPath) {
      const data = await fs.readFile(args.l1ContractsJsonPath)
      contractAddrs = {
        l1: JSON.parse(data.toString()),
        l2: DEFAULT_L2_CONTRACT_ADDRESSES,
      } as OEContractsLike
    }

    const OptimismPortal = new hre.ethers.Contract(
      contractAddrs.l1.OptimismPortal,
      Artifact__OptimismPortal.abi,
      signer
    )

    const L1CrossDomainMessenger = new hre.ethers.Contract(
      contractAddrs.l1.L1CrossDomainMessenger,
      Artifact__L1CrossDomainMessenger.abi,
      signer
    )

    const L1StandardBridge = new hre.ethers.Contract(
      contractAddrs.l1.L1StandardBridge,
      Artifact__L1StandardBridge.abi,
      signer
    )

    const L2ToL1MessagePasser = new hre.ethers.Contract(
      predeploys.L2ToL1MessagePasser,
      Artifact__L2ToL1MessagePasser.abi
    )

    const L2CrossDomainMessenger = new hre.ethers.Contract(
      predeploys.L2CrossDomainMessenger,
      Artifact__L2CrossDomainMessenger.abi
    )

    const L2StandardBridge = new hre.ethers.Contract(
      predeploys.L2StandardBridge,
      Artifact__L2StandardBridge.abi
    )

    const messenger = new CrossChainMessenger({
      l1SignerOrProvider: signer,
      l2SignerOrProvider: l2Signer,
      l1ChainId: await signer.getChainId(),
      l2ChainId,
      bedrock: true,
      contracts: contractAddrs,
    })

    const l1BridgeBalanceBefore = await signer.provider.getBalance(
      L1StandardBridge.address
    )

    console.log('Depositing ETH to L2 ovm_eth')
    const depositTx = await messenger.depositETH(amount, { recipient: to })
    await depositTx.wait()
    console.log(`ETH deposited - ${depositTx.hash}`)

    // Deposit might get reorged, wait 30s and also log for reorgs.
    let prevBlockHash: string = ''
    for (let i = 0; i < 30; i++) {
      const messageReceipt = await messenger.waitForMessageReceipt(depositTx)
      if (messageReceipt.receiptStatus !== 1) {
        console.log(`Deposit failed, retrying...`)
      }

      if (
        prevBlockHash !== '' &&
        messageReceipt.transactionReceipt.blockHash !== prevBlockHash
      ) {
        console.log(
          `Block hash changed from ${prevBlockHash} to ${messageReceipt.transactionReceipt.blockHash}`
        )

        // Wait for stability, we want at least 30 seconds after any reorg
        i = 0
      }

      prevBlockHash = messageReceipt.transactionReceipt.blockHash
      await sleep(1000)
    }

    const l1BridgeBalanceAfter = await signer.provider.getBalance(
      L1StandardBridge.address
    )

    console.log(
      `L1StandardBridge balance before: ${formatEther(l1BridgeBalanceBefore)}`
    )

    console.log(
      `L1StandardBridge balance after: ${formatEther(l1BridgeBalanceAfter)}`
    )

    if (!l1BridgeBalanceBefore.add(amount).eq(l1BridgeBalanceAfter)) {
      throw new Error(`L1StandardBridge balance mismatch`)
    }

    const LegacyOVMETH = new hre.ethers.Contract(
      predeploys.LegacyOVMETH,
      Artifact__LegacyOVMETH.abi,
      l2Signer
    )

    const l2Balance = await LegacyOVMETH.balanceOf(
      typeof to === 'string' ? to : to.address
    )
    console.log(
      `L2 balance of deposit recipient: ${utils.formatEther(
        l2Balance.toString()
      )}`
    )

    if (l2Balance.lt(utils.parseEther('1'))) {
      throw new Error('bad deposit')
    }
    console.log(`Deposit success`)

    console.log('Withdrawing OVM_ETH')
    // const preBalance = await LegacyOVMETH.balanceOf(signer.address)
    const withdraw = await messenger.withdrawETH(withdrawAmount)
    console.log(`Transaction hash: ${withdraw.hash}`)
    const withdrawalReceipt = await withdraw.wait()
    console.log(
      `OVM_ETH withdrawn on L2 - included in block ${withdrawalReceipt.blockNumber}`
    )
    for (const log of withdrawalReceipt.logs) {
      switch (log.address) {
        case L2ToL1MessagePasser.address: {
          const parsed = L2ToL1MessagePasser.interface.parseLog(log)
          console.log(`Log ${parsed.name} from ${log.address}`)
          console.log(parsed.args)
          console.log()
          break
        }
        case L2StandardBridge.address: {
          const parsed = L2StandardBridge.interface.parseLog(log)
          console.log(`Log ${parsed.name} from ${log.address}`)
          console.log(parsed.args)
          console.log()
          break
        }
        case L2CrossDomainMessenger.address: {
          const parsed = L2CrossDomainMessenger.interface.parseLog(log)
          console.log(`Log ${parsed.name} from ${log.address}`)
          console.log(parsed.args)
          console.log()
          break
        }
        default: {
          console.log(`Unknown log from ${log.address} - ${log.topics[0]}`)
        }
      }
    }

    setInterval(async () => {
      const currentStatus = await messenger.getMessageStatus(withdraw)
      console.log(`Message status: ${MessageStatus[currentStatus]}`)
    }, 3000)

    const now = Math.floor(Date.now() / 1000)

    console.log('Waiting for message to be able to be proved')
    await messenger.waitForMessageStatus(withdraw, MessageStatus.READY_TO_PROVE)

    console.log('Proving eth withdrawal...')
    const prove = await messenger.proveMessage(withdraw)
    const proveReceipt = await prove.wait()
    console.log(proveReceipt)
    if (proveReceipt.status !== 1) {
      throw new Error('Prove withdrawal transaction reverted')
    }
    console.log('Successfully proved withdrawal')

    console.log('Waiting for message to be able to be relayed')
    await messenger.waitForMessageStatus(
      withdraw,
      MessageStatus.READY_FOR_RELAY
    )

    console.log('Finalizing eth withdrawal...')
    // TODO: Update SDK to properly estimate gas
    const finalize = await messenger.finalizeMessage(withdraw, {
      overrides: { gasLimit: 500_000 },
    })
    console.log(`Transaction hash: ${finalize.hash}`)
    const finalizeReceipt = await finalize.wait()
    console.log('finalizeReceipt:', finalizeReceipt)
    console.log(`Took ${Math.floor(Date.now() / 1000) - now} seconds`)

    for (const log of finalizeReceipt.logs) {
      switch (log.address) {
        case OptimismPortal.address: {
          const parsed = OptimismPortal.interface.parseLog(log)
          console.log(`Log ${parsed.name} from OptimismPortal (${log.address})`)
          console.log(parsed.args)
          console.log()
          break
        }
        case L1CrossDomainMessenger.address: {
          const parsed = L1CrossDomainMessenger.interface.parseLog(log)
          console.log(
            `Log ${parsed.name} from L1CrossDomainMessenger (${log.address})`
          )
          console.log(parsed.args)
          console.log()
          break
        }
        case L1StandardBridge.address: {
          const parsed = L1StandardBridge.interface.parseLog(log)
          console.log(
            `Log ${parsed.name} from L1StandardBridge (${log.address})`
          )
          console.log(parsed.args)
          console.log()
          break
        }
        default:
          console.log(
            `Unknown log emitted from ${log.address} - ${log.topics[0]}`
          )
      }
    }

    const l1BridgeBalanceFinally = await signer.provider.getBalance(
      L1StandardBridge.address
    )

    if (!l1BridgeBalanceFinally.add(withdrawAmount).eq(l1BridgeBalanceAfter)) {
      throw new Error('L1StandardBridge balance mismatch')
    }
    console.log('Withdrawal success')
  })
