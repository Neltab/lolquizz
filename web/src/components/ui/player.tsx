export type PlayerProps = {
    name: string
    score: number
    isHost: boolean
}

export function Player({ name, score, isHost }: PlayerProps) {
    return (
        <div className="flex flex-col gap-2">
            <h1>{isHost ? 'Host' : 'Joueur'}</h1>
            <p>{name}</p>
            <p>{score}</p>
        </div>
    )
}
