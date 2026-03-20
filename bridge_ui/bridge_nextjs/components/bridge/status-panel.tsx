"use client"

import { ArrowLeft } from "lucide-react"
import { StepsList } from "./status-panel/steps-list"
import { EventLog } from "./status-panel/event-log"
import { CompletionBanner } from "./status-panel/completion-banner"
import type { BridgeStatus, BridgeEvent } from "@/hooks/use-bridge-store"
import { Direction } from "./types"

interface StatusPanelProps {
  status: BridgeStatus
  events: BridgeEvent[]
  error: string | null
  onReset: () => void
  amount: string
  recipient: string
  direction: Direction
}

export function BridgeStatusPanel({
  status,
  events,
  error,
  onReset,
  amount,
  recipient,
  direction
}: StatusPanelProps) {
  return (
    <div className="mx-auto max-w-lg">
      <div className="rounded-2xl border-2 border-border bg-card p-6">
        <div className="mb-6 flex items-center justify-between">
          <button
            onClick={onReset}
            className="flex items-center gap-2 text-lg text-muted-foreground hover:text-card-foreground"
          >
            <ArrowLeft className="h-5 w-5" />
            Back
          </button>
          <span className="text-lg text-card-foreground">{amount} USDC</span>
        </div>

        <StepsList status={status} error={error} />
        <EventLog events={events} />
        <CompletionBanner
          status={status}
          events={events}
          recipient={recipient}
          onReset={onReset}
          direction={direction}
        />
      </div>
    </div>
  )
}