import { useState } from "react";
import { Input } from "./input";
import { Button } from "./button";
import { AvatarPicker } from "./avatarPicker";

type PlayerInfoProps = {
    name?: string,
    code?: string,
}

type PlayerJoinSetupProps = {
    joinOnly: true,
    handleCreateRoom?: undefined,
    handleJoinRoom: (playerName: string, roomCode: string) => void,
}

type PlayerCreateSetupProps = {
    joinOnly?: false,
    handleCreateRoom: (playerName: string) => void,
    handleJoinRoom: (playerName: string, roomCode: string) => void,
}

type PlayerSetupProps = PlayerInfoProps & (PlayerCreateSetupProps | PlayerJoinSetupProps);

export const PlayerSetup = ({ name, code, joinOnly, handleCreateRoom, handleJoinRoom }: PlayerSetupProps) => {
    const [playerName, setplayerName] = useState(name || "");
    const [roomCode, setRoomCode] = useState(code || "");
    const roomCodeValid = roomCode.length === 6;

    return (
        <div className="flex flex-col gap-8">
            <AvatarPicker />
            <div className="flex flex-col gap-4">
                <Input value= {playerName} onChange={(e) => setplayerName(e.target.value)} placeholder="Pseudo" />
                { !joinOnly && <Button className="w-full" disabled={!playerName} onClick={() => handleCreateRoom(playerName)}>Créer une partie</Button> }
                <div className="flex items-center justify-center gap-4">
                    <Input value = {roomCode} onChange={(e) => setRoomCode(e.target.value)} placeholder="Code de la partie" />
                    <Button variant='secondary' disabled={!roomCodeValid} onClick={() => handleJoinRoom(playerName, roomCode) }>Rejoindre</Button>
                </div>
            </div>
        </div>
    )
}