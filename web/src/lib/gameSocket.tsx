export type MessageHandler = (payload: any) => void

export class GameSocket {
  private ws: WebSocket | null = null
  private handlers = new Map<string, Set<MessageHandler>>()
  private openHandler: (() => void) | null = null
  private closeHandler: (() => void) | null = null
  private token: string

  constructor(token: string) {
    this.token = token
  }

  connect() {
    const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:'
    this.ws = new WebSocket(
      `${protocol}/${window.location.host}/ws?token=${this.token}`
    )

    this.ws.onopen = () => this.openHandler?.()

    this.ws.onmessage = (event) => {
      const msg = JSON.parse(event.data)
      console.log('ws message', msg)
      const handlers = this.handlers.get(msg.type)
      if (handlers) {
        handlers.forEach(h => h(msg.payload))
      }
    }

    this.ws.onclose = () => this.closeHandler?.()
  }

  send(type: string, payload: any) {
    if (this.ws?.readyState === WebSocket.OPEN) {
      this.ws.send(JSON.stringify({ type, payload }))
    }
  }

  on(type: string, handler: MessageHandler): () => void {
    if (!this.handlers.has(type)) {
      this.handlers.set(type, new Set())
    }
    this.handlers.get(type)!.add(handler)
    return () => this.handlers.get(type)?.delete(handler)
  }

  onOpen(handler: () => void) { this.openHandler = handler }
  onClose(handler: () => void) { this.closeHandler = handler }

  disconnect() {
    if (this.ws) {
        this.ws.onopen = null
        this.ws.onmessage = null
        this.ws.onclose = null
        this.ws.onerror = null
        this.ws.close()
        this.ws = null
    }
  }
}