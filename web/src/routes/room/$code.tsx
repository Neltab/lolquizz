import { createFileRoute, useRouterState } from '@tanstack/react-router'
import { useCallback, useEffect, useMemo, useRef, useState } from 'react'
import { GameSocket } from '@/lib/gameSocket'
import { Message } from '../../lib/gameSocket'
import PlayerList from '../../components/ui/playerList'
import { Button } from '../../components/ui/button'


export const Route = createFileRoute('/room/$code')({
  component: RouteComponent,
})

function RouteComponent() {
    const { code } = Route.useParams()
    const { gameState }  = Route.useRouteContext()
    const socketRef = useRef<GameSocket | null>(null)
    const [connected, setConnected] = useState(false)
    const [players, setPlayers] = useState<any[]>([])
    const [isHost, setIsHost] = useState(false)

    useEffect(() => {
        setIsHost(gameState.isHost)
    }, [gameState.isHost])

    const joinRoom = useCallback(() => {
        console.log('connected', gameState)
        if (!gameState.token) {
            console.log(gameState)
            return
        }

        let cancelled = false

        const socket = new GameSocket(gameState.token)
        socketRef.current = socket

        socket.onOpen(() => {
            if (cancelled) return
            setConnected(true)
            socket.send('join_room', { room_code: code, player_name: gameState.playerName })
        })

        socket.onClose(() => {
            if (cancelled) return
            setConnected(false)
        })


        socket.on('player_joined', (data: Message) => {
            if (cancelled) return
            setPlayers(data.state.players)
            console.log('player_joined', data)
        })

        socket.on('player_left', (data: Message) => {
            if (cancelled) return
            setPlayers(data.state.players)
        })

        socket.on('round_started', (data: Message) => {
            if (cancelled) return
            console.log('round_started', data)
        })

        socket.connect()

        return () => {
            cancelled = true
            socket.disconnect()
        }
    }, [gameState.token])

    const startGame = useCallback(() => {
        if (!gameState.token) {
            console.log(gameState)
            return
        }

        console.log('start game')
        socketRef.current?.send('start_game', { room_code: code })
       
    }, [gameState.token])

    return (
        <div className='flex flex-col gap-4'>
            <h1>Room {code}</h1>
            <button onClick={joinRoom}>connect to ws</button>
            {isHost && <Button onClick={startGame}>Start Game</Button>}
            <PlayerList players={players} />
        </div>
    )
}
