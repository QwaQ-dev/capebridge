import { clsx, type ClassValue } from 'clsx'
import { twMerge } from 'tailwind-merge'
import type { BridgeEvent } from "@/hooks/use-bridge-store"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export function formatEventMessage(event: BridgeEvent): string {
  switch (event.status) {
    case "detected":
      return `Transaction found on Stacks`
    case "voted":
      return `Consensus in progress (votes collected)`
    case "relaying":
      return "Executing transfer on Base..."
    case "relayed":
      return `Success! Bridged to Base`
    case "failed":
      return `Error: ${event.message || "Liquidity issue or revert"}`
    default:
      return `Status: ${event.status}`
  }
}

export function truncateAddress(addr: string): string {
  if (!addr) return ""
  return `${addr.slice(0, 8)}...${addr.slice(-6)}`
}