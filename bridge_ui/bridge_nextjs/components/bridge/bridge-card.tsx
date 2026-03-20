"use client"

import { useState } from "react"
import { useAccount } from "wagmi"
import { ArrowDown } from "lucide-react"
import { useBridge } from "@/hooks/use-base-bridge"
import { useStacksWallet } from "@/contexts/stacks-wallet-context"
import { useStacksBridge } from "@/hooks/use-stacks-bridge"
import { useBridgeStore } from "@/hooks/use-bridge-store" 
import { BridgeStatusPanel } from "./status-panel"
import { TokenSelector } from "./bridge-card/token-selector"
import { AddressDisplay } from "./bridge-card/address-display"
import { AmountInput } from "./bridge-card/amount-input"
import { BridgeButton } from "./bridge-card/bridge-button"
import type { Direction } from "./types"

export function BridgeCard() {
  const { address: evmAddress, isConnected: evmConnected } = useAccount()
  const stacks = useStacksWallet()

  const { status, events, error, reset } = useBridgeStore()

  const evmBridge = useBridge() 
  const stacksBridge = useStacksBridge()

  const [direction, setDirection] = useState<Direction>("base-to-stacks")
  const [amount, setAmount] = useState("")

  const MAX_AMOUNT = 1

  const isBaseToStacks = direction === "base-to-stacks"
  const { balance: usdcBalance, balanceWei } = evmBridge

  const sourceToken = isBaseToStacks ? "USDC" : "USDCx"
  const destChain = isBaseToStacks ? "Stacks" : "Base"
  const sourceChain = isBaseToStacks ? "Base" : "Stacks"

  const sourceBalance = isBaseToStacks ? usdcBalance : stacks.usdcxBalance
  const sourceAddress = isBaseToStacks ? evmAddress ?? "" : stacks.address ?? ""
  const recipientAddress = isBaseToStacks ? stacks.address ?? "" : evmAddress ?? ""

  const bothConnected = evmConnected && stacks.isConnected

  const amountFloat = Math.min(parseFloat(amount) || 0, MAX_AMOUNT)
  const amountWei = BigInt(Math.floor(amountFloat * 1_000_000))

  const insufficientBalance =
    !!amount &&
    amountFloat > 0 &&
    (isBaseToStacks
      ? amountWei > balanceWei
      : amountFloat > parseFloat(stacks.usdcxBalance))

  const handleAmountChange = (value: string) => {
    if (value === "") {
      setAmount("")
      return
    }

    const num = parseFloat(value)

    if (isNaN(num)) return

    if (num > MAX_AMOUNT) {
      setAmount(MAX_AMOUNT.toString())
      return
    }

    if (num < 0) {
      setAmount("0")
      return
    }

    setAmount(value)
  }

  const handleMax = () => {
    const balanceFloat = parseFloat(sourceBalance) || 0
    const capped = Math.min(balanceFloat, MAX_AMOUNT)
    setAmount(capped.toString())
  }

  const handleBridge = () => {
    if (!amount || !recipientAddress) return

    const num = parseFloat(amount)

    if (isNaN(num) || num <= 0 || num > MAX_AMOUNT) {
      console.error("Invalid amount")
      return
    }

    if (isBaseToStacks) {
      evmBridge.bridge(amount, recipientAddress)
    } else {
      const senderAddress = stacks.address ?? ""
      if (senderAddress) {
        stacksBridge.bridge(amount, recipientAddress, senderAddress)
      }
    }
  }

  const handleToggleDirection = () => {
    setDirection((d) => (d === "base-to-stacks" ? "stacks-to-base" : "base-to-stacks"))
    setAmount("")
    reset() 
  }

  if (status !== "idle") {
    return (
      <BridgeStatusPanel
        status={status}
        events={events}
        error={error}
        onReset={reset}
        amount={amount}
        recipient={recipientAddress}
        direction={direction}
      />
    )
  }

  return (
    <div className="rounded-xl border border-border bg-card">
      <div className="p-6">
        <TokenSelector direction={direction} onToggle={handleToggleDirection} />
        <Divider />
        <AddressDisplay
          label="Source address"
          address={sourceAddress}
          chain={sourceChain as "Base" | "Stacks"}
          variant="muted"
        />
        <AmountInput
          amount={amount}
          balance={sourceBalance}
          token={sourceToken}
          insufficientBalance={insufficientBalance}
          onChange={handleAmountChange}
          onMax={handleMax}
        />
        <Divider />
        <AddressDisplay
          label="Recipient address"
          address={recipientAddress}
          chain={destChain as "Base" | "Stacks"}
          variant="bordered"
        />
        <div className="mt-6">
          <BridgeButton
            bothConnected={bothConnected}
            insufficientBalance={insufficientBalance}
            canBridge={!insufficientBalance && !!amount && bothConnected}
            sourceToken={sourceToken}
            destChain={destChain}
            onClick={handleBridge}
          />
        </div>
      </div>
    </div>
  )
}

function Divider() {
  return (
    <div className="relative my-4 flex items-center justify-center">
      <div className="absolute left-0 right-0 border-t border-border" />
      <div className="relative rounded-full border border-border bg-card p-1.5">
        <ArrowDown className="h-3 w-3 text-muted-foreground" />
      </div>
    </div>
  )
}