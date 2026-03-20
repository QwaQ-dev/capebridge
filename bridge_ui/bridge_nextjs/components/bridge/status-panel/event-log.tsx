import { formatEventMessage } from "@/lib/utils"
import type { BridgeEvent } from "@/hooks/use-bridge-store"

interface EventLogProps {
  events: BridgeEvent[]
}

export function EventLog({ events }: EventLogProps) {
  if (events.length === 0) return null

  return (
    <div className="mt-6 rounded-xl border-2 border-border bg-muted/20 p-4">
      <p className="mb-3 text-base text-muted-foreground">Live Updates</p>
      <div className="max-h-32 space-y-2 overflow-y-auto">
        {events.map((event, i) => (
          <div key={i} className="flex items-start gap-2 text-base">
            <span className="text-muted-foreground">
              {new Date(event.timestamp).toLocaleTimeString()}
            </span>
            <span className="text-card-foreground">
              {formatEventMessage(event)}
            </span>
          </div>
        ))}
      </div>
    </div>
  )
}