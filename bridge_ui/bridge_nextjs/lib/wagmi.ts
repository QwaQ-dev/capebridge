import { createConfig, http } from "wagmi"
import { base, baseSepolia } from "wagmi/chains"
import { injected } from "wagmi/connectors"

export const config = createConfig({
  chains: [base, baseSepolia],
  connectors: [injected()],
  transports: {
    [base.id]: http(),
    [baseSepolia.id]: http(),
  },
  ssr: true,
})

// Contract addresses
export const BRIDGE_CONTRACT = process.env.NEXT_PUBLIC_BRIDGE_CONTRACT 

export const USDC_CONTRACT = process.env.NEXT_PUBLIC_USDC_CONTRACT

export const BACKEND_API =  process.env.NEXT_PUBLIC_BRIDGE_API

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
  {
    name: "allowance",
    inputs: [
      { name: "owner", type: "address" },
      { name: "spender", type: "address" },
    ],
    outputs: [{ name: "", type: "uint256" }],
    stateMutability: "view",
    type: "function",
  },
] as const

export const BRIDGE_ABI = [
  {
    name: "deposit",
    inputs: [
      { name: "token", type: "address" },
      { name: "amount", type: "uint256" },
      { name: "receiver", type: "string" },
    ],
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
] as const
