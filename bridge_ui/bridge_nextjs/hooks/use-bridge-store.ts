import { create } from 'zustand'

export type BridgeStatus =
  | "idle" | "approving" | "depositing" | "detecting"
  | "confirmed" | "relaying" | "relayed" | "failed"

export interface BridgeEvent {
  status: string
  block?: number
  relay_tx?: string
  message?: string
  timestamp: number
}

type StateSetter<T> = T | ((prev: T) => T)

interface BridgeState {
  status: BridgeStatus
  events: BridgeEvent[]
  currentTxHash: string | null
  error: string | null
  setStatus: (val: StateSetter<BridgeStatus>) => void
  setEvents: (val: StateSetter<BridgeEvent[]>) => void
  setCurrentTxHash: (hash: string | null) => void
  setError: (error: string | null) => void
  reset: () => void
}

export const useBridgeStore = create<BridgeState>((set) => ({
  status: "idle",
  events: [],
  currentTxHash: null,
  error: null,
  setStatus: (val) => 
    set((state) => ({ 
      status: typeof val === 'function' ? val(state.status) : val 
    })),
  setEvents: (val) => 
    set((state) => ({ 
      events: typeof val === 'function' ? val(state.events) : val 
    })),
  setCurrentTxHash: (hash) => set({ currentTxHash: hash }),
  setError: (error) => set({ error }),
  reset: () => set({ status: "idle", events: [], error: null, currentTxHash: null }),
}))