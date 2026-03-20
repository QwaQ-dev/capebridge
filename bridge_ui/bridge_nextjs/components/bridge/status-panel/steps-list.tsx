import { CheckCircle2, Circle, Loader2, XCircle } from "lucide-react"
import { cn } from "@/lib/utils"
import { STATUS_STEPS } from "./status-steps"
import type { BridgeStatus } from "@/hooks/use-bridge-store"

interface StepsListProps {
  status: BridgeStatus
  error: string | null
}

export function StepsList({ status, error }: StepsListProps) {
  const currentStepIndex = STATUS_STEPS.findIndex((s) => s.key === status)
  const isFailed = status === "failed"
  const isComplete = status === "relayed"

  return (
    <div className="space-y-3">
      {STATUS_STEPS.map((step, index) => {
        const isActive = step.key === status
        const isCompleted = currentStepIndex > index || isComplete
        const isPending = currentStepIndex < index && !isFailed

        return (
          <div
            key={step.key}
            className={cn(
              "flex items-center gap-4 rounded-xl border-2 p-4 transition-all",
              isActive && "border-primary bg-primary/5",
              isCompleted && "border-green-500/30 bg-green-500/5",
              isPending && "border-border bg-muted/20 opacity-50"
            )}
          >
            <div>
              {isCompleted ? (
                <CheckCircle2 className="h-6 w-6 text-green-500" />
              ) : isActive ? (
                isFailed ? (
                  <XCircle className="h-6 w-6 text-destructive" />
                ) : (
                  <Loader2 className="h-6 w-6 animate-spin text-primary" />
                )
              ) : (
                <Circle className="h-6 w-6 text-muted-foreground/30" />
              )}
            </div>
            <div className="flex-1">
              <p
                className={cn(
                  "text-xl",
                  isActive
                    ? "text-card-foreground"
                    : isCompleted
                    ? "text-green-600"
                    : "text-muted-foreground"
                )}
              >
                {step.label}
              </p>
              {isActive && !isFailed && (
                <p className="text-base text-muted-foreground">{step.description}</p>
              )}
              {isActive && isFailed && (
                <p className="text-base text-destructive">{error}</p>
              )}
            </div>
          </div>
        )
      })}
    </div>
  )
}