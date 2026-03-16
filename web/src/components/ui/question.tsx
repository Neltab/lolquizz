import { Countdown } from "./countdown"
import { Input } from "./input"
import { useState } from "react"

type QuestionProps = {
    text: string,
    duration: number,
    sendAnswer: (answer: string) => void,
}

export const Question = ({ text, duration, sendAnswer }: QuestionProps) => {
    const [answer, setAnswer] = useState("")

    return (
        <div className="w-full h-full flex flex-col items-center justify-center">
                <Countdown duration={duration} />
                <p>{text}</p>
                <Input onBlur={() => sendAnswer(answer)} value={answer} onChange={(e) => setAnswer(e.target.value)} />
        </div>
    )
}