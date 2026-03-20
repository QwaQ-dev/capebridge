import type { BridgeStatus, BridgeEvent } from "@/hooks/use-bridge-store"
import { BACKEND_API } from "@/lib/wagmi"

type StatusSetter = (s: BridgeStatus) => void
type EventsSetter = (fn: (prev: BridgeEvent[]) => BridgeEvent[]) => void
type ErrorSetter = (msg: string) => void

export function subscribeToBridgeEvents(
  txHash: string,
  setStatus: StatusSetter,
  setEvents: EventsSetter,
  setError: ErrorSetter
): () => void {
  const eventSource = new EventSource(`${BACKEND_API}/bridge/events/${txHash}`)

  eventSource.onmessage = (event) => {
  console.log("SSE EVENT:", event.data)

  const raw = event.data?.trim()
  if (!raw || raw === ":" || !raw.startsWith("{")) return

  try {
    const data = JSON.parse(raw)

    setEvents((prev) => [...prev, { ...data, timestamp: Date.now() }])

    switch (data.status) {
      case "detected":
        setStatus("detecting")
        break
      case "confirmed":
        setStatus("confirmed")
        break
      case "relaying":
        setStatus("relaying")
        break
      case "relayed":
        setStatus("relayed")
        eventSource.close()
        break
      case "failed":
        setStatus("failed")
        setError(data.message || "Bridge failed")
        eventSource.close()
        break
    }
  } catch {
    console.warn("SSE: skipped non-JSON message:", event.data)
  }
}

  eventSource.onerror = () => eventSource.close()

  return () => eventSource.close()
}