import { createFileRoute, useRouterState } from '@tanstack/react-router'
import { useCallback, useEffect, useRef, useState } from 'react'
import { GameSocket } from '@/lib/gameSocket'
import { Message } from '../../lib/gameSocket'
import { Card, CardContent } from '../../components/ui/card'
import { RoomSettings } from '../../components/ui/roomSettings'
import PlayerList from '../../components/ui/playerList'
import { Button } from '../../components/ui/button'
import { Question } from '../../components/ui/question'


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
    const [hasStarted, setHasStarted] = useState(false)
    const [question, setQuestion] = useState<any>(null)

    useEffect(() => {
        setIsHost(gameState.isHost)
    }, [gameState.isHost])

    useEffect(() => {
        if (!gameState.token) return

        const socket = new GameSocket(gameState.token)
        socketRef.current = socket

        socket.onOpen(() => {
            setConnected(true)
            socket.send('join_room', { room_code: code, player_name: gameState.playerName })
        })

        socket.onClose(() => {
            setConnected(false)
        })

        socket.on('player_joined', (data: Message) => {
            setPlayers(data.state.players)
        })

        socket.on('player_left', (data: Message) => {
            setPlayers(data.state.players)
        })

        socket.on('question_started', (data: Message) => {
            console.log('question_started', data)
            setHasStarted(true)
            setQuestion(data.payload)
        })

        socket.connect()

        return () => {
            socket.disconnect()
        }
    }, [gameState.token])

    const startGame = useCallback(() => {
        if (!gameState.token) {
            console.log(gameState)
            return
        }

        socketRef.current?.send('start_game', { room_code: code })
       
    }, [gameState.token])

    const sendAnswer = useCallback((answer: string) => {
        if (!gameState.token) return
        socketRef.current?.send('submit_answer', { room_code: code, answer })
    }, [gameState.token])

    if (hasStarted) {
        return <Question text={question.text} duration={question.duration} sendAnswer={sendAnswer} />
    }

    return (
        <div className="w-full h-full flex flex-row gap-8 items-center justify-center">
            <Card>
                <CardContent>
                    <RoomSettings />
                </CardContent>
            </Card>

            <Card>
                <CardContent>
                    <PlayerList players={players}/>
                    <Button onClick={startGame}>Start Game</Button>
                </CardContent>
            </Card>
        </div>
    )
}
