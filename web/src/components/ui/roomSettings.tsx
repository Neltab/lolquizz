import { Slider } from "./slider"
import { Field, FieldContent, FieldLabel } from "./field"
import { Input } from "./input"
import { useState } from "react"
import { DifficultySlider } from "./difficultySlider"

export const RoomSettings = () => {
    const [maxPlayers, setMaxPlayers] = useState(10)
    const [maxQuestions, setMaxQuestions] = useState(10)
    const [repartitionQuestions, setRepartitionQuestions] = useState<[number, number]>([3, 6])

    return (
        <div className="w-full h-full flex flex-col items-center justify-center">
            <Field orientation="horizontal">
                <FieldContent>
                    <FieldLabel>Joueurs max</FieldLabel>
                </FieldContent>
                <Input />
            </Field>
            <Field orientation="horizontal">
                <FieldContent>
                    <FieldLabel>Nombre de questions</FieldLabel>
                </FieldContent>
                <Input />
            </Field>
            <Field orientation="horizontal">
                <FieldContent>
                    <FieldLabel>Répartition des questions</FieldLabel>
                </FieldContent>
                <DifficultySlider 
                    max={maxQuestions}
                    value={repartitionQuestions}
                    onValueChange={(value) => setRepartitionQuestions(value)}
                />
            </Field>
        </div>
    )
}