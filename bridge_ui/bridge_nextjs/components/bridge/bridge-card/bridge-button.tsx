import { Button } from "@/components/ui/button"

interface BridgeButtonProps {
  bothConnected: boolean
  insufficientBalance: boolean
  canBridge: boolean
  sourceToken: string
  destChain: string
  onClick: () => void
}

export function BridgeButton({
  bothConnected,
  insufficientBalance,
  canBridge,
  sourceToken,
  destChain,
  onClick,
}: BridgeButtonProps) {
  const label = !bothConnected
    ? "Connect both wallets to bridge"
    : insufficientBalance
    ? `Insufficient ${sourceToken}`
    : `Bridge ${sourceToken} → ${destChain}`

  return (
    <Button
      onClick={onClick}
      disabled={!canBridge}
      className="w-full rounded-full py-6 text-base font-medium"
      size="lg"
    >
      {label}
    </Button>
  )
}