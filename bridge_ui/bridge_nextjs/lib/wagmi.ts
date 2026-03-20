import { createConfig, http } from "wagmi"
import { baseSepolia } from "wagmi/chains"
import { injected } from "wagmi/connectors"

export const config = createConfig({
  chains: [baseSepolia],
  connectors: [injected()],
  transports: {
    [baseSepolia.id]: http(),
  },
  ssr: true,
})

export const BRIDGE_CONTRACT = process.env.NEXT_PUBLIC_BASE_BRIDGE_CONTRACT
export const USDC_CONTRACT = process.env.NEXT_PUBLIC_USDC_CONTRACT
export const BACKEND_API = process.env.NEXT_PUBLIC_BRIDGE_API

export const USDC_ABI = [
  {
    name: "approve",
    inputs: [
      { name: "spender", type: "address" },
      { name: "value", type: "uint256" },
    ],
    outputs: [{ name: "", type: "bool" }],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    name: "balanceOf",
    inputs: [{ name: "account", type: "address" }],
    outputs: [{ name: "", type: "uint256" }],
    stateMutability: "view",
    type: "function",
  },
] as const

export const BRIDGE_ABI = [
  {
    name: "Deposit",
    inputs: [
      { name: "amount", type: "uint256" },
      { name: "receiver", type: "string" },
    ],
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
] as const