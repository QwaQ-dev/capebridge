import { createPublicClient, http, formatUnits } from "viem"
import { base, baseSepolia } from "wagmi/chains"
import { USDC_CONTRACT, USDC_ABI } from "@/lib/wagmi"

const CHAIN = baseSepolia

const publicClient = createPublicClient({
  chain: CHAIN,
  transport: http(),
})

export async function fetchEvmUSDCBalance(address: string): Promise<string> {
  if (!address || !USDC_CONTRACT) return "0.00"

  try {
    const raw = await publicClient.readContract({
      address: USDC_CONTRACT as `0x${string}`,
      abi: USDC_ABI,
      functionName: "balanceOf",
      args: [address as `0x${string}`],
    })

    return formatUnits(raw as bigint, 6)
  } catch (err) {
    console.error("fetchEvmUSDCBalance error:", err)
    return "0.00"
  }
} 