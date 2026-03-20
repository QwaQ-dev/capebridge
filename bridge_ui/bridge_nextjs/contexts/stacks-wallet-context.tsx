"use client"

import { createContext, useContext, useState, useEffect, type ReactNode } from "react"
import { useStacksWallet as useStacksWalletHook, type StacksWalletState } from "@/hooks/use-stacks-wallet"
import { fetchStacksUSDCBalance } from "@/lib/stacks-balance"

interface StacksWalletContextType extends StacksWalletState {
  usdcxBalance: string
  connect: () => Promise<void>
  disconnect: () => void
}

const StacksWalletContext = createContext<StacksWalletContextType>(
  {} as StacksWalletContextType
)

export function StacksWalletProvider({ children }: { children: ReactNode }) {
  const wallet = useStacksWalletHook()
  const [usdcxBalance, setUsdcxBalance] = useState("0.00")

  useEffect(() => {
    if (!wallet.address) {
      setUsdcxBalance("0.00")
      return
    }

    fetchStacksUSDCBalance(wallet.address).then(setUsdcxBalance)

    const interval = setInterval(
      () => fetchStacksUSDCBalance(wallet.address!).then(setUsdcxBalance),
      30_000
    )

    return () => clearInterval(interval)
  }, [wallet.address])

  return (
    <StacksWalletContext.Provider value={{ ...wallet, usdcxBalance }}>
      {children}
    </StacksWalletContext.Provider>
  )
}

export const useStacksWallet = () => useContext(StacksWalletContext)