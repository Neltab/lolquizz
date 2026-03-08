import { createFileRoute } from '@tanstack/react-router'
import { useCallback, useEffect, useRef, useState } from 'react'
import { GameSocket } from '@/lib/gameSocket'

export const Route = createFileRoute('/room/$code')({
  component: RouteComponent,
})

function RouteComponent() {
    const { code } = Route.useParams()
    const { gameState } = Route.useRouteContext()
    const socketRef = useRef<GameSocket | null>(null)
    const [connected, setConnected] = useState(false)
    const [players, setPlayers] = useState<any[]>([])

    const joinRoom = useCallback(() => {
        console.log('connected')
        if (!gameState.token)
            return

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

        // Set up message handlers
        socket.on('room_state', (data) => {
            if (cancelled) return
            setPlayers(data.players)
            gameState.setIsHost(data.is_host)
        })

        socket.on('player_joined', (data) => {
            if (cancelled) return
            setPlayers(data.players)
            console.log('player_joined', data)
        })

        socket.on('player_left', (data) => {
            if (cancelled) return
            setPlayers(data.players)
            gameState.setIsHost(data.is_host)
        })

        socket.connect()

        return () => {
            cancelled = true
            socket.disconnect()
        }
    }, [])

    return (
        <div>
            <h1>Room {code}</h1>
            <button onClick={joinRoom}>connect to ws</button>
            {gameState.isHost && <button>Start Game</button>}
        </div>
    )
}
