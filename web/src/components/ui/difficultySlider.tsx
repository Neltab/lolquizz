
import * as SliderPrimitive from "@radix-ui/react-slider";
import { cn } from "@/lib/utils";

function DifficultySlider({
  value,
  onValueChange,
  min = 0,
  max = 100,
  step = 1,
  className,
}: {
  value: [number, number];
  onValueChange: (value: [number, number]) => void;
  min?: number;
  max?: number;
  step?: number;
  className?: string;
}) {
  const range = max - min;
  const v0Pct = ((value[0] - min) / range) * 100;
  const v1Pct = ((value[1] - min) / range) * 100;

  return (
    <SliderPrimitive.Root
      className={cn("relative flex w-full touch-none select-none items-center", className)}
      value={value}
      onValueChange={(v) => onValueChange(v as [number, number])}
      min={min}
      max={max}
      step={step}
    >
      <SliderPrimitive.Track className="relative h-2 w-full grow overflow-hidden rounded-full bg-easy">
        <div
          className="absolute h-full rounded-l-full"
          style={{
            left: 0,
            width: `${v0Pct}%`,
          }}
        />
        <div
          className="absolute h-full bg-medium"
          style={{
            left: `${v0Pct}%`,
            width: `${v1Pct - v0Pct}%`,
          }}
        />
        <div
          className="absolute h-full rounded-r-full bg-hard"
          style={{
            left: `${v1Pct}%`,
            right: 0,
          }}
        />
        <SliderPrimitive.Range className="absolute h-full opacity-0" />
      </SliderPrimitive.Track>

      <SliderPrimitive.Thumb
        className={cn(
          "block h-5 w-5 rounded-full border-2 bg-background ring-offset-background",
          "transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2",
          "disabled:pointer-events-none disabled:opacity-50 cursor-grab active:cursor-grabbing"
        )}
        aria-label="Easy / Medium boundary"
      />

      <SliderPrimitive.Thumb
        className={cn(
          "block h-5 w-5 rounded-full border-2 bg-background ring-offset-background",
          "transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2",
          "disabled:pointer-events-none disabled:opacity-50 cursor-grab active:cursor-grabbing"
        )}
        aria-label="Medium / Hard boundary"
      />
    </SliderPrimitive.Root>
  );
}

export { DifficultySlider }