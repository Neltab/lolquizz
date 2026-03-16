export type PlayerProps = {
    name: string
    avatar: string
    isHost: boolean
}

export function Player({ name, avatar, isHost }: PlayerProps) {
    return (
        <div className="flex flex-col">
            <img src={avatar} alt="avatar" />
            <h1>{isHost ? 'Host' : 'Joueur'}</h1>
            <p>{name}</p>
        </div>
    )
}
