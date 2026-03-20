"use client"

import Link from "next/link"
import { Github, Play, ArrowDownUp, Loader2, LogOut } from "lucide-react"
import { BridgeCard } from "@/components/bridge/bridge-card"
import { useStacksWallet } from "@/contexts/stacks-wallet-context"
import { useAccount, useConnect, useDisconnect } from "wagmi"
import { injected } from "wagmi/connectors"

function StacksIcon({ className }: { className?: string }) {
  return (
    <svg viewBox="0 0 24 24" fill="none" className={className}>
      <path d="M12 2L2 7L12 12L22 7L12 2Z" fill="currentColor" />
      <path d="M2 17L12 22L22 17" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" />
      <path d="M2 12L12 17L22 12" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" />
    </svg>
  )
}

function EVMIcon({ className }: { className?: string }) {
  return (
    <svg viewBox="0 0 24 24" fill="none" className={className}>
      <path d="M12 2L4.5 12.5L12 16L19.5 12.5L12 2Z" fill="currentColor" opacity="0.9" />
      <path d="M4.5 13.5L12 22L19.5 13.5L12 17L4.5 13.5Z" fill="currentColor" />
    </svg>
  )
}

function EVMWalletButton() {
  const { address, isConnected } = useAccount()
  const { connect, isPending } = useConnect()
  const { disconnect } = useDisconnect()

  if (isConnected && address) {
    return (
      <div className="flex items-center gap-1 rounded-xl border border-border bg-card px-3 py-2">
        <span className="flex h-6 w-6 items-center justify-center rounded-full bg-[#627EEA] text-white">
          <EVMIcon className="h-3 w-3" />
        </span>
        <span className="hidden text-sm font-medium sm:inline">
          {address.slice(0, 5)}...{address.slice(-4)}
        </span>
        <button
          onClick={() => disconnect()}
          className="ml-1 text-muted-foreground transition-colors hover:text-foreground"
          title="Disconnect EVM wallet"
        >
          <LogOut className="h-3.5 w-3.5" />
        </button>
      </div>
    )
  }

  return (
    <button
      onClick={() => connect({ connector: injected() })}
      disabled={isPending}
      className="flex items-center gap-2 rounded-xl border border-border bg-card px-3 py-2 text-sm font-medium transition-colors hover:bg-muted disabled:opacity-60"
    >
      {isPending ? (
        <Loader2 className="h-4 w-4 animate-spin" />
      ) : (
        <span className="flex h-5 w-5 items-center justify-center rounded-full bg-[#627EEA] text-white">
          <EVMIcon className="h-3 w-3" />
        </span>
      )}
      <span className="hidden sm:inline">
        {isPending ? "Connecting..." : "EVM Wallet"}
      </span>
    </button>
  )
}

function LeatherWalletButton() {
  const stacks = useStacksWallet()

  if (stacks.isConnected && stacks.address) {
    return (
      <div className="flex items-center gap-1 rounded-xl border border-border bg-card px-3 py-2">
        <span className="flex h-6 w-6 items-center justify-center rounded-full bg-[#5546FF] text-white">
          <StacksIcon className="h-3 w-3" />
        </span>
        <span className="hidden text-sm font-medium sm:inline">
          {stacks.address.slice(0, 5)}...{stacks.address.slice(-4)}
        </span>
        <button
          onClick={stacks.disconnect}
          className="ml-1 text-muted-foreground transition-colors hover:text-foreground"
          title="Disconnect Stacks wallet"
        >
          <LogOut className="h-3.5 w-3.5" />
        </button>
      </div>
    )
  }

  return (
    <button
      onClick={stacks.connect}
      disabled={stacks.isConnecting}
      className="flex items-center gap-2 rounded-xl border border-border bg-card px-3 py-2 text-sm font-medium transition-colors hover:bg-muted disabled:opacity-60"
    >
      {stacks.isConnecting ? (
        <Loader2 className="h-4 w-4 animate-spin" />
      ) : (
        <span className="flex h-5 w-5 items-center justify-center rounded-full bg-[#5546FF] text-white">
          <StacksIcon className="h-3 w-3" />
        </span>
      )}
      <span className="hidden sm:inline">
        {stacks.isConnecting ? "Connecting..." : "Stacks Wallet"}
      </span>
    </button>
  )
}

export default function BridgePage() {
  return (
    <div className="min-h-screen bg-background">
      <header className="border-b border-border">
        <div className="mx-auto flex max-w-5xl items-center justify-between px-4 py-3">
          <div className="flex items-center gap-3">
            <span className="text-lg font-semibold">cape4labs</span>
            <span className="flex items-center gap-1.5 rounded-full bg-primary px-3 py-1 text-xs font-medium text-primary-foreground">
              <ArrowDownUp className="h-3 w-3" />
              Bridge
            </span>
          </div>

          <div className="flex items-center gap-2">
            <Link
              href="https://github.com/QwaQ-dev/capebridge"
              target="_blank"
              className="flex items-center gap-2 rounded-lg border border-border px-3 py-2 text-sm text-foreground transition-colors hover:bg-muted"
            >
              <Github className="h-4 w-4" />
              <span className="hidden sm:inline">GitHub</span>
            </Link>
            <Link
              href="https://youtu.be/VK5M9tOktmE"
              target="_blank"
              className="flex items-center gap-2 rounded-lg border border-border px-3 py-2 text-sm text-foreground transition-colors hover:bg-muted"
            >
              <Play className="h-4 w-4" />
              <span className="hidden sm:inline">Pitch</span>
            </Link>

            <LeatherWalletButton />
            <EVMWalletButton />
          </div>
        </div>
      </header>

      <main className="mx-auto max-w-xl px-4 py-12">
        <BridgeCard />
        <div className="mt-6 flex flex-wrap items-center justify-center gap-3">
          <div className="flex items-center gap-2 rounded-full border border-border bg-card px-4 py-2 text-sm">
            <span className="text-muted-foreground">Max amount</span>
            <span className="font-medium">1 USDC</span>
          </div>
          <div className="flex items-center gap-2 rounded-full border border-border bg-card px-4 py-2 text-sm">
            <span className="font-medium">ONLY FOR TEST</span>
          </div>
        </div>
      </main>

      <footer className="fixed bottom-0 left-0 right-0 border-t border-border bg-background py-4 text-center text-sm text-muted-foreground">
        This application is maintained by cape4labs and is powered by{" "}
        <Link href="https://stacks.co" target="_blank" className="text-foreground underline">
          Stacks
        </Link>
        .
      </footer>
    </div>
  )
}