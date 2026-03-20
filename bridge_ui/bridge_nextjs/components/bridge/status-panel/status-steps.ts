export const STATUS_STEPS = [
  { key: "approving",  label: "Approving USDC",       description: "Waiting for approval transaction..." },
  { key: "depositing", label: "Depositing",            description: "Confirm deposit in your wallet..." },
  { key: "detecting",  label: "Transaction Detected",  description: "Bridge is watching your transaction..." },
  { key: "confirmed",  label: "Confirmed",             description: "Waiting for quorum confirmation..." },
  { key: "relaying",   label: "Relaying",              description: "Sending to destination network..." },
  { key: "relayed",    label: "Complete!",             description: "Your tokens have been bridged successfully" },
] as const