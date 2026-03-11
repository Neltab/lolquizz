import  { Player, PlayerProps } from "./player";

export type PlayerListProps = {
    players: PlayerProps[]
}

export default function PlayerList({ players }: PlayerListProps) {
    return (
        <div className="flex flex-col gap-2">
            {players.map((player) => (
                <Player key={player.name} {...player} />
            ))}
        </div>
    )
}