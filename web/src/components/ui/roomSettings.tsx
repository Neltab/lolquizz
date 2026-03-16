import { Slider } from "./slider"
import { Field, FieldContent, FieldLabel } from "./field"
import { Input } from "./input"

export const RoomSettings = () => {
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
                    <FieldLabel>Questions faciles</FieldLabel>
                </FieldContent>
                <Slider />
            </Field>
            <Field orientation="horizontal">
                <FieldContent>
                    <FieldLabel>Questions moyennes</FieldLabel>
                </FieldContent>
                <Slider />
            </Field>
            <Field orientation="horizontal">
                <FieldContent>
                    <FieldLabel>Questions difficiles</FieldLabel>
                </FieldContent>
                <Slider />
            </Field>
        </div>
    )
}