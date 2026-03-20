"use client"

import { useState, useEffect, useCallback } from "react"
import { getStacksProvider } from "@/lib/stacks-provider"

export interface StacksWalletState {
  address: string | null
  isConnected: boolean
  isConnecting: boolean
}

const STORAGE_KEY = "stacks_address"

export function useStacksWallet() {
  const [state, setState] = useState<StacksWalletState>({
    address: null,
    isConnected: false,
    isConnecting: false,
  })

  useEffect(() => {
    const saved = localStorage.getItem(STORAGE_KEY)
    if (saved) setState({ address: saved, isConnected: true, isConnecting: false })
  }, [])

  const connect = useCallback(async () => {
    setState((s) => ({ ...s, isConnecting: true }))

    try {
      const provider = getStacksProvider()

      if (!provider) {
        window.open("https://leather.io/install-extension", "_blank")
        setState((s) => ({ ...s, isConnecting: false }))
        return
      }

      const response = await provider.request("getAddresses")

      const accounts = response?.result?.addresses as Array<{
        type: string
        symbol: string
        address: string
      }> | undefined

      const stxAccount = accounts?.find(
        (a) => a.symbol === "STX" || a.type === "stx"
      )

      if (!stxAccount?.address) throw new Error("No Stacks address returned")

      setState({ address: stxAccount.address, isConnected: true, isConnecting: false })
      localStorage.setItem(STORAGE_KEY, stxAccount.address)
    } catch (err) {
      console.error("Stacks connect error:", err)
      setState((s) => ({ ...s, isConnecting: false }))
    }
  }, [])

  const disconnect = useCallback(() => {
    localStorage.removeItem(STORAGE_KEY)
    setState({ address: null, isConnected: false, isConnecting: false })
  }, [])

  return { ...state, connect, disconnect }
}