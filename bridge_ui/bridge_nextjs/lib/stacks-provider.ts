export type StacksProvider = {
  request: (method: string) => Promise<any>
}

export function getStacksProvider(): StacksProvider | null {
  if (typeof window === "undefined") return null
  return (
    (window as any).LeatherProvider ??
    (window as any).HiroWalletProvider ??
    null
  )
}