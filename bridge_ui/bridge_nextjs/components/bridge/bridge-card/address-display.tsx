import { BaseIcon, StacksIcon } from "./icons"

type Chain = "Base" | "Stacks"

interface AddressDisplayProps {
  label: string
  address: string
  chain: Chain
  variant?: "muted" | "bordered"
}

export function AddressDisplay({
  label,
  address,
  chain,
  variant = "muted",
}: AddressDisplayProps) {
  const ChainIcon = chain === "Base" ? BaseIcon : StacksIcon

  const containerClass =
    variant === "muted"
      ? "rounded-lg bg-muted/50 px-4 py-3"
      : "rounded-lg border border-border bg-card px-4 py-3"

  return (
    <div className={containerClass}>
      <p className="mb-1 text-xs text-muted-foreground">{label}</p>
      <div className="flex items-center gap-2">
        <ChainIcon className="h-4 w-4 shrink-0" />
        {address ? (
          <span className="break-all font-mono text-sm">{address}</span>
        ) : (
          <span className="text-sm text-muted-foreground">
            Connect {chain} wallet in header
          </span>
        )}
      </div>
    </div>
  )
}