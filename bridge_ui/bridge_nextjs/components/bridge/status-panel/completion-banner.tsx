import { Button } from "@/components/ui/button"
import { truncateAddress } from "@/lib/utils"
import type { BridgeStatus, BridgeEvent } from "@/hooks/use-bridge-store"
import { Direction } from "../types"

interface CompletionBannerProps {
  status: BridgeStatus
  events: BridgeEvent[]
  recipient: string
  onReset: () => void
  direction: Direction
}

export function CompletionBanner({
  status,
  events,
  recipient,
  onReset,
  direction
}: CompletionBannerProps) {
  const isFailed = status === "failed"
  const isComplete = status === "relayed"

  const latestRelayTx = events
    .filter((e) => e.relay_tx)
    .at(-1)?.relay_tx
    
  console.log("events:", events)
  console.log("latestRelayTx:", latestRelayTx)

  if (!isFailed && !isComplete) return null

  return (
    <div className="mt-6 space-y-4">
      {isComplete && latestRelayTx && (
        <div className="rounded-xl border-2 border-green-500/30 bg-green-500/10 p-4">
          <div className="flex items-center justify-between">
            <div>
              <p className="text-xl text-green-600">Bridge Complete!</p>
              <p className="text-base text-muted-foreground">
                Sent to {truncateAddress(recipient)}
              </p>
            </div>
          </div>
        </div>
      )}

      <Button
        onClick={onReset}
        className="w-full rounded-xl border-2 border-border bg-muted py-6 text-xl text-card-foreground hover:bg-muted/80"
        variant="ghost"
      >
        {isFailed ? "Try Again" : "Make Another Transfer"}
      </Button>
    </div>
  )
}