"use client"

import { useCallback, useState, useEffect } from "react"
import { useAccount, useWriteContract, useReadContract, useChainId } from "wagmi"
import { parseUnits } from "viem"
import { BRIDGE_CONTRACT, USDC_CONTRACT, USDC_ABI, BRIDGE_ABI, BACKEND_API } from "@/lib/wagmi"
import { subscribeToBridgeEvents } from "@/lib/bridge-sse"
import { useBridgeStore } from "@/hooks/use-bridge-store"

export function useBridge() {
  const { address } = useAccount()
  const chainId = useChainId()
  const state = useBridgeStore() 

  const [balance, setBalance] = useState("0.00")
  const [balanceWei, setBalanceWei] = useState(BigInt(0))

  const { data: balanceRaw, refetch: refetchBalance } = useReadContract({
    address: USDC_CONTRACT as `0x${string}`,
    abi: USDC_ABI,
    functionName: "balanceOf",
    args: address ? [address] : undefined,
    query: { enabled: !!address, refetchInterval: 10_000 },
  })

  const { data: allowanceRaw, refetch: refetchAllowance } = useReadContract({
    address: USDC_CONTRACT as `0x${string}`,
    abi: USDC_ABI,
    functionName: "allowance",
    args: address ? [address, BRIDGE_CONTRACT as `0x${string}`] : undefined,
    query: { enabled: !!address, refetchInterval: 10_000 },
  })

  useEffect(() => {
    if (!balanceRaw) return
    const raw = balanceRaw as bigint
    setBalanceWei(raw)
    setBalance((Number(raw) / 1_000_000).toFixed(2))
  }, [balanceRaw])

  const { writeContractAsync: approve } = useWriteContract()
  const { writeContractAsync: deposit } = useWriteContract()

  const bridge = useCallback(
    async (amount: string, receiver: string) => {
      if (!address) return

      state.setError(null)
      state.setEvents([])

      const amountWei = parseUnits(amount, 6)

      try {
        const currentAllowance = (allowanceRaw as bigint | undefined) ?? BigInt(0)
        if (currentAllowance < amountWei) {
          state.setStatus("approving")
          await approve({
            address: USDC_CONTRACT as `0x${string}`,
            abi: USDC_ABI,
            functionName: "approve",
            args: [BRIDGE_CONTRACT as `0x${string}`, amountWei],
          })

          await new Promise<void>((resolve) => {
            const check = setInterval(async () => {
              const result = await refetchAllowance()
              if (result.data && (result.data as bigint) >= amountWei) {
                clearInterval(check)
                resolve()
              }
            }, 2000)
          })
        }

        state.setStatus("depositing")
        const depositTx = await deposit({
          address: BRIDGE_CONTRACT as `0x${string}`,
          abi: BRIDGE_ABI,
          functionName: "deposit",
          args: [USDC_CONTRACT as `0x${string}`, amountWei, receiver],
        })

        state.setCurrentTxHash(depositTx)

        await fetch(`${BACKEND_API}/api/bridge/deposit`, {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({
            tx_hash: depositTx,
            sender: address,
            receiver,
            amount: amountWei.toString(),
            chain_id: chainId,
            direction: "base-to-stacks",
          }),
        })

        subscribeToBridgeEvents(depositTx, state.setStatus, state.setEvents, state.setError)
        refetchBalance()
      } catch (err) {
        state.setStatus("failed")
        state.setError(err instanceof Error ? err.message : "Transaction failed")
      }
    },
    [address, allowanceRaw, approve, deposit, refetchBalance, chainId, state]
  )

  return { balance, balanceWei, bridge }
}