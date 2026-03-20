"use client"

import { useCallback, useState, useEffect } from "react"
import { useAccount, useWriteContract, useReadContract, useChainId, usePublicClient } from "wagmi"
import { parseUnits } from "viem"
import { BRIDGE_CONTRACT, USDC_CONTRACT, USDC_ABI, BRIDGE_ABI, BACKEND_API } from "@/lib/wagmi"
import { subscribeToBridgeEvents } from "@/lib/bridge-sse"
import { useBridgeStore } from "@/hooks/use-bridge-store"

export function useBridge() {
  const { address } = useAccount()
  const chainId = useChainId()
  const publicClient = usePublicClient()

  const setStatus = useBridgeStore((s) => s.setStatus)
  const setError = useBridgeStore((s) => s.setError)
  const setEvents = useBridgeStore((s) => s.setEvents)
  const setCurrentTxHash = useBridgeStore((s) => s.setCurrentTxHash)

  const [balance, setBalance] = useState("0.00")
  const [balanceWei, setBalanceWei] = useState(BigInt(0))

  const { data: balanceRaw, refetch: refetchBalance } = useReadContract({
    address: USDC_CONTRACT as `0x${string}`,
    abi: USDC_ABI,
    functionName: "balanceOf",
    args: address ? [address] : undefined,
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
      if (!address || !publicClient) return

      setError(null)
      setEvents([])
      const amountWei = parseUnits(amount, 6)

      try {

        console.log(amountWei)
        setStatus("approving")
        const approveTxHash = await approve({
          address: USDC_CONTRACT as `0x${string}`,
          abi: USDC_ABI,
          functionName: "approve",
          args: [BRIDGE_CONTRACT as `0x${string}`, amountWei],
        })
        await publicClient.waitForTransactionReceipt({ hash: approveTxHash, confirmations: 1 })
        setStatus("depositing")
        const depositTxHash = await deposit({
          address: BRIDGE_CONTRACT as `0x${string}`,
          abi: BRIDGE_ABI,
          functionName: "deposit",
          args: [USDC_CONTRACT as `0x${string}`, amountWei, receiver],
        })
        await publicClient.waitForTransactionReceipt({ hash: depositTxHash, confirmations: 1 })

        setCurrentTxHash(depositTxHash)

        await fetch(`${BACKEND_API}/bridge/deposit`, {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({
            tx_hash: depositTxHash,
            sender: address,
            receiver,
            amount: amountWei.toString(),
            chain_id: chainId,
            direction: "base-to-stacks",
          }),
        })

        subscribeToBridgeEvents(depositTxHash, setStatus, setEvents, setError)
        refetchBalance()
      } catch (err) {
        setStatus("failed")
        setError(err instanceof Error ? err.message : "Transaction failed")
      }
    },
    [address, publicClient, approve, deposit, refetchBalance, chainId, setStatus, setError, setEvents, setCurrentTxHash]
  )

  return { balance, balanceWei, bridge }
}