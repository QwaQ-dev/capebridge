"use client"

import { useState, useCallback } from "react"
import { BACKEND_API } from "@/lib/wagmi"
import { subscribeToBridgeEvents } from "@/lib/bridge-sse"
import { fetchStacksUSDCBalance } from "@/lib/stacks-balance"
import { useBridgeStore } from "@/hooks/use-bridge-store"

const STACKS_NETWORK = process.env.NEXT_PUBLIC_STACKS_NETWORK || "testnet"
const STACKS_BRIDGE_CONTRACT = process.env.NEXT_PUBLIC_STACKS_BRIDGE_CONTRACT
const SUSDCX_DECIMALS = 6

function toMicroUnits(amount: string): number {
  return Math.round(parseFloat(amount) * 10 ** SUSDCX_DECIMALS)
}

export function useStacksBridge() {
  const state = useBridgeStore()
  const [stacksBalance, setStacksBalance] = useState("0.00")

  const refreshBalance = useCallback(async (address: string) => {
    const bal = await fetchStacksUSDCBalance(address)
    setStacksBalance(bal)
  }, [])

  const bridge = useCallback(
    async (amount: string, receiver: string, senderAddress: string) => {
      if (!senderAddress) { state.setError("Connect your Stacks wallet first"); return }
      if (!STACKS_BRIDGE_CONTRACT) { state.setError("USDCx contract address not configured"); return }

      state.setError(null)
      state.setEvents([])

      const microAmount = toMicroUnits(amount)

      try {
        const { openContractCall } = await import("@stacks/connect")
        const {
          uintCV,
          stringAsciiCV,
          contractPrincipalCV,
          PostConditionMode,
          Pc,
        } = await import("@stacks/transactions")

        const [contractAddress, contractName] = STACKS_BRIDGE_CONTRACT.split(".")

        state.setStatus("depositing")

        const postConditions = [
          Pc.principal(senderAddress)
            .willSendEq(microAmount)
            .ft("ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM.usdcx", "usdcx-token")
        ]

        await new Promise<void>((resolve, reject) => {
          openContractCall({
            network: STACKS_NETWORK === "mainnet" ? "mainnet" : "testnet",
            contractAddress,
            contractName,
            functionName: "deposit",
            functionArgs: [
              uintCV(microAmount),
              stringAsciiCV(receiver),
              contractPrincipalCV("ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM", "usdcx"),
            ],
            postConditions,
            postConditionMode: PostConditionMode.Deny,
            onFinish: async ({ txId }) => {
              try {
                state.setCurrentTxHash(txId)

                await fetch(`${BACKEND_API}/api/bridge/deposit`, {
                  method: "POST",
                  headers: { "Content-Type": "application/json" },
                  body: JSON.stringify({
                    tx_hash: txId,
                    sender: senderAddress,
                    receiver,
                    amount: microAmount.toString(),
                    chain_id: "stacks",
                    direction: "stacks-to-base",
                  }),
                }).catch(console.warn)

                subscribeToBridgeEvents(txId, state.setStatus, state.setEvents, state.setError)
                resolve()
              } catch (e) {
                reject(e)
              }
            },
            onCancel: () => {
              state.setStatus("idle")
              reject(new Error("User cancelled"))
            },
          })
        })
      } catch (err) {
        if (err instanceof Error && err.message === "User cancelled") return
        state.setStatus("failed")
        state.setError(err instanceof Error ? err.message : "Transaction failed")
      }
    },
    [state]
  )

  return { stacksBalance, refreshBalance, bridge }
}