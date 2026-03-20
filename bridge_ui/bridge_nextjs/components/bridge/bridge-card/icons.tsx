export function BaseIcon({ className }: { className?: string }) {
  return (
    <svg viewBox="0 0 111 111" fill="none" className={className}>
      <path d="M54.921 110.034C85.359 110.034 110.034 85.402 110.034 55.017C110.034 24.6319 85.359 0 54.921 0C26.0432 0 2.35281 22.1714 0 50.3923H72.8467V59.6416H0C2.35281 87.8625 26.0432 110.034 54.921 110.034Z" fill="#0052FF" />
    </svg>
  )
}

export function StacksIcon({ className }: { className?: string }) {
  return (
    <svg viewBox="0 0 24 24" fill="none" className={className}>
      <path d="M12 2L2 7L12 12L22 7L12 2Z" fill="#5546FF" />
      <path d="M2 17L12 22L22 17" stroke="#5546FF" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" />
      <path d="M2 12L12 17L22 12" stroke="#5546FF" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" />
    </svg>
  )
}

export function USDCIcon({ className }: { className?: string }) {
  return (
    <svg viewBox="0 0 32 32" fill="none" className={className}>
      <circle cx="16" cy="16" r="16" fill="#2775CA" />
      <path d="M20.5 18.5C20.5 16.5 19 15.5 16 15C14 14.5 13.5 14 13.5 13C13.5 12 14.5 11.5 16 11.5C17.5 11.5 18.5 12 19 13L21 12C20 10.5 18.5 9.5 17 9.5V7H15V9.5C13 10 11.5 11.5 11.5 13.5C11.5 15.5 13 16.5 16 17C18 17.5 18.5 18 18.5 19C18.5 20 17.5 20.5 16 20.5C14.5 20.5 13 20 12.5 18.5L10.5 19.5C11.5 21 13 22 15 22.5V25H17V22.5C19.5 22 21 20.5 20.5 18.5Z" fill="white" />
    </svg>
  )
}