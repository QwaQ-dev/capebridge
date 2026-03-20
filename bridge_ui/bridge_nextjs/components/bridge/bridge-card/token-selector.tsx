import { ArrowLeftRightIcon } from "lucide-react"
import { BaseIcon, StacksIcon, USDCIcon } from "./icons"
import type { Direction } from "../types"

interface TokenSelectorProps {
  direction: Direction
  onToggle: () => void
}

export function TokenSelector({ direction, onToggle }: TokenSelectorProps) {
  const isBaseToStacks = direction === "base-to-stacks"

  const sourceToken = isBaseToStacks ? "USDC" : "USDCx"
  const destToken = isBaseToStacks ? "USDCx" : "USDC"
  const sourceChain = isBaseToStacks ? "Base" : "Stacks"
  const destChain = isBaseToStacks ? "Stacks" : "Base"

  const SourceChainIcon = isBaseToStacks ? BaseIcon : StacksIcon
  const DestChainIcon = isBaseToStacks ? StacksIcon : BaseIcon

  return (
    <div className="mb-2 flex items-start justify-between">
      <div className="flex-1">
        <p className="mb-2 text-xs text-muted-foreground">From</p>
        <div className="flex items-center gap-3">
          <div className="flex items-center gap-2">
            <USDCIcon className="h-6 w-6" />
            <span className="text-lg font-semibold">{sourceToken}</span>
          </div>
          <div className="flex items-center gap-1.5 rounded-full bg-muted px-2.5 py-1">
            <SourceChainIcon className="h-4 w-4" />
            <span className="text-xs font-medium">{sourceChain}</span>
          </div>
        </div>
      </div>

      <button
        onClick={onToggle}
        className="mx-4 mt-6 rounded-full border border-border p-2 transition-colors hover:bg-muted"
      >
        <ArrowLeftRightIcon className="h-4 w-4" />
      </button>

      <div className="flex-1 text-right">
        <p className="mb-2 text-xs text-muted-foreground">To</p>
        <div className="flex items-center justify-end gap-3">
          <div className="flex items-center gap-2">
            <USDCIcon className="h-6 w-6" />
            <span className="text-lg font-semibold">{destToken}</span>
          </div>
          <div className="flex items-center gap-1.5 rounded-full bg-muted px-2.5 py-1">
            <DestChainIcon className="h-4 w-4" />
            <span className="text-xs font-medium">{destChain}</span>
          </div>
        </div>
      </div>
    </div>
  )
}