import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { USDCIcon } from "./icons"

interface AmountInputProps {
  amount: string
  balance: string
  token: string
  insufficientBalance: boolean
  onChange: (value: string) => void
  onMax: () => void
}



export function AmountInput({
  amount,
  balance,
  token,
  insufficientBalance,
  onChange,
  onMax,
}: AmountInputProps) {
  return (
    <div className="rounded-lg border border-border bg-card px-4 py-3">
      <div className="flex items-center justify-between">
        <p className="text-xs text-muted-foreground">Amount</p>
        <p className={`text-xs ${insufficientBalance ? "font-medium text-destructive" : "text-muted-foreground"}`}>
          Available: {balance} {token}
        </p>
      </div>
      <div className="mt-1 flex items-center justify-between">
        <div className="flex items-center gap-2">
          <USDCIcon className="h-5 w-5" />
          <Input
            type="number"
            placeholder="0.00"
            value={amount}
            onChange={(e) => onChange(e.target.value)}
            className="h-auto w-32 border-0 bg-transparent p-0 text-xl font-semibold shadow-none focus-visible:ring-0"
          />
          <span className="text-muted-foreground">{token}</span>
        </div>
        <Button
          variant="outline"
          size="sm"
          onClick={onMax}
          className="h-7 px-2 text-xs"
        >
          Max
        </Button>
      </div>
      {insufficientBalance && (
        <p className="mt-1 text-xs text-destructive">Insufficient {token} balance</p>
      )}
    </div>
  )
}