const STACKS_NETWORK = process.env.NEXT_PUBLIC_STACKS_NETWORK || "testnet"
const SUSDCX_CONTRACT = process.env.NEXT_PUBLIC_SUSDCX_CONTRACT || ""
const SUSDCX_DECIMALS = 6

const STACKS_API =
  STACKS_NETWORK === "mainnet"
    ? "https://api.mainnet.hiro.so"
    : "https://api.testnet.hiro.so"

export async function fetchStacksUSDCBalance(address: string): Promise<string> {
  if (!SUSDCX_CONTRACT || !address) return "0.00"

  try {
    const res = await fetch(`${STACKS_API}/extended/v1/address/${address}/balances`)

    if (!res.ok) throw new Error(`HTTP ${res.status}`)

    const data = await res.json()

    const ftKey = Object.keys(data.fungible_tokens ?? {}).find(
      (k) => k.toLowerCase().includes(SUSDCX_CONTRACT.toLowerCase())
    )

    if (!ftKey) return "0.00"

    const raw = data.fungible_tokens[ftKey].balance as string
    return (parseInt(raw) / 10 ** SUSDCX_DECIMALS).toFixed(2)
  } catch (err) {
    console.error("fetchStacksUSDCBalance error:", err)
    return "0.00"
  }
}