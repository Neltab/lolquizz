import { useEffect, useState } from "react"
import { cn } from "@/lib/utils"

interface CountdownProps extends React.ComponentProps<"div"> {
  duration: number
  onComplete?: () => void
  running?: boolean
}

export const Countdown = ({
  duration,
  onComplete,
  running = true,
  className,
  ...props
}: CountdownProps) => {
  const [remaining, setRemaining] = useState(duration)

  useEffect(() => {
    setRemaining(duration)
  }, [duration])

  useEffect(() => {
    if (!running || remaining <= 0) return

    const interval = setInterval(() => {
      setRemaining((prev) => {
        if (prev <= 1) {
          clearInterval(interval)
          onComplete?.()
          return 0
        }
        return prev - 1
      })
    }, 1000)

    return () => clearInterval(interval)
  }, [running, remaining <= 0, onComplete])

  const minutes = Math.floor(remaining / 60)
  const seconds = remaining % 60
  const progress = duration > 0 ? remaining / duration : 0

  return (
    <div
      data-slot="countdown"
      className={cn(
        "relative flex items-center justify-center",
        className
      )}
      {...props}
    >
      <svg viewBox="0 0 100 100" className="h-full w-full -rotate-90">
        <circle
          cx="50"
          cy="50"
          r="45"
          fill="none"
          stroke="currentColor"
          strokeWidth="4"
          className="text-muted-foreground/20"
        />
        <circle
          cx="50"
          cy="50"
          r="45"
          fill="none"
          stroke="currentColor"
          strokeWidth="4"
          strokeDasharray={2 * Math.PI * 45}
          strokeDashoffset={2 * Math.PI * 45 * (1 - progress)}
          strokeLinecap="round"
          className="transition-[stroke-dashoffset] duration-1000 ease-linear"
        />
      </svg>
      <span className="absolute text-2xl font-bold tabular-nums">
        {minutes > 0
          ? `${minutes}:${seconds.toString().padStart(2, "0")}`
          : seconds}
      </span>
    </div>
  )
}
