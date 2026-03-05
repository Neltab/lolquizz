import { useState, useRef, useEffect, useCallback } from "react";

const MESSAGE_TEMPLATES = {
  join_room: { room_code: "ABCD", nickname: "Player1" },
  leave_room: {},
  update_settings: {
    question_count: 10,
    time_per_question: 30,
    difficulty: "medium",
    categories: [],
  },
  start_game: {},
  submit_answer: { answer: "" },
  judge_answer: { player_id: "", correct: true },
  next_round: {},
};

const TAG_COLORS = {
  sent: { bg: "#1a2f1a", border: "#2d5a2d", text: "#6fcf6f" },
  received: { bg: "#1a1a2f", border: "#2d2d5a", text: "#6f8fcf" },
  system: { bg: "#2f2a1a", border: "#5a4d2d", text: "#cfb86f" },
  error: { bg: "#2f1a1a", border: "#5a2d2d", text: "#cf6f6f" },
};

function formatTime(date) {
  return date.toLocaleTimeString("en-US", {
    hour12: false,
    hour: "2-digit",
    minute: "2-digit",
    second: "2-digit",
    fractionalSecondDigits: 3,
  });
}

function LogEntry({ entry }) {
  const colors = TAG_COLORS[entry.tag];
  const [expanded, setExpanded] = useState(false);

  let parsed = null;
  try {
    parsed = JSON.parse(entry.data);
  } catch {
    parsed = null;
  }

  const preview = parsed
    ? parsed.type || Object.keys(parsed)[0] || "..."
    : entry.data?.slice(0, 60);

  return (
    <div
      onClick={() => setExpanded(!expanded)}
      style={{
        padding: "8px 12px",
        borderLeft: `3px solid ${colors.border}`,
        background: expanded ? colors.bg : "transparent",
        cursor: "pointer",
        fontFamily: "'IBM Plex Mono', 'SF Mono', 'Fira Code', monospace",
        fontSize: "12px",
        lineHeight: "1.5",
        transition: "background 0.15s ease",
      }}
    >
      <div style={{ display: "flex", alignItems: "center", gap: "10px" }}>
        <span style={{ color: "#555", flexShrink: 0, fontSize: "11px" }}>
          {formatTime(entry.time)}
        </span>
        <span
          style={{
            color: colors.text,
            fontWeight: 600,
            fontSize: "10px",
            textTransform: "uppercase",
            letterSpacing: "0.08em",
            background: colors.bg,
            padding: "1px 6px",
            borderRadius: "3px",
            flexShrink: 0,
          }}
        >
          {entry.tag === "sent" ? "→ SENT" : entry.tag === "received" ? "← RECV" : entry.tag === "error" ? "✕ ERR" : "● SYS"}
        </span>
        <span
          style={{
            color: "#aaa",
            overflow: "hidden",
            textOverflow: "ellipsis",
            whiteSpace: "nowrap",
            flex: 1,
          }}
        >
          {preview}
        </span>
        <span style={{ color: "#444", flexShrink: 0, fontSize: "11px" }}>
          {expanded ? "▾" : "▸"}
        </span>
      </div>

      {expanded && (
        <pre
          style={{
            marginTop: "8px",
            padding: "10px",
            background: "#0a0a0f",
            borderRadius: "4px",
            color: "#ccc",
            fontSize: "11px",
            whiteSpace: "pre-wrap",
            wordBreak: "break-all",
            overflowX: "auto",
            border: "1px solid #1a1a25",
          }}
        >
          {parsed ? JSON.stringify(parsed, null, 2) : entry.data}
        </pre>
      )}
    </div>
  );
}

function ConnectionBar({ status, url, setUrl, playerID, setPlayerID, onConnect, onDisconnect }) {
  const isConnected = status === "connected";
  const statusColor =
    status === "connected" ? "#4ade80" : status === "connecting" ? "#facc15" : "#666";

  return (
    <div
      style={{
        padding: "16px 20px",
        borderBottom: "1px solid #1a1a25",
        display: "flex",
        flexDirection: "column",
        gap: "10px",
        background: "#08080d",
      }}
    >
      <div style={{ display: "flex", alignItems: "center", gap: "12px" }}>
        <div
          style={{
            width: "8px",
            height: "8px",
            borderRadius: "50%",
            background: statusColor,
            boxShadow: isConnected ? `0 0 8px ${statusColor}` : "none",
            flexShrink: 0,
          }}
        />
        <input
          type="text"
          value={url}
          onChange={(e) => setUrl(e.target.value)}
          disabled={isConnected}
          placeholder="ws://localhost:8080/ws"
          style={{
            flex: 1,
            background: isConnected ? "#0a0a0f" : "#0f0f18",
            border: "1px solid #1a1a25",
            borderRadius: "6px",
            color: isConnected ? "#666" : "#ddd",
            padding: "8px 12px",
            fontFamily: "'IBM Plex Mono', monospace",
            fontSize: "13px",
            outline: "none",
          }}
        />
        <input
          type="text"
          value={playerID}
          onChange={(e) => setPlayerID(e.target.value)}
          disabled={isConnected}
          placeholder="player_id"
          style={{
            width: "140px",
            background: isConnected ? "#0a0a0f" : "#0f0f18",
            border: "1px solid #1a1a25",
            borderRadius: "6px",
            color: isConnected ? "#666" : "#ddd",
            padding: "8px 12px",
            fontFamily: "'IBM Plex Mono', monospace",
            fontSize: "13px",
            outline: "none",
          }}
        />
        <button
          onClick={isConnected ? onDisconnect : onConnect}
          style={{
            padding: "8px 20px",
            borderRadius: "6px",
            border: "none",
            background: isConnected ? "#2f1a1a" : "#1a2f1a",
            color: isConnected ? "#cf6f6f" : "#6fcf6f",
            fontFamily: "'IBM Plex Mono', monospace",
            fontWeight: 600,
            fontSize: "12px",
            cursor: "pointer",
            textTransform: "uppercase",
            letterSpacing: "0.05em",
            flexShrink: 0,
          }}
        >
          {isConnected ? "Disconnect" : "Connect"}
        </button>
      </div>
    </div>
  );
}

function MessageComposer({ onSend, connected }) {
  const [selectedType, setSelectedType] = useState("join_room");
  const [payload, setPayload] = useState(
    JSON.stringify(MESSAGE_TEMPLATES["join_room"], null, 2)
  );
  const [rawMode, setRawMode] = useState(false);
  const [rawMessage, setRawMessage] = useState('{"type":"ping"}');

  const handleTypeChange = (type) => {
    setSelectedType(type);
    setPayload(JSON.stringify(MESSAGE_TEMPLATES[type], null, 2));
  };

  const handleSend = () => {
    if (!connected) return;
    if (rawMode) {
      onSend(rawMessage);
    } else {
      try {
        const parsed = JSON.parse(payload);
        onSend(JSON.stringify({ type: selectedType, payload: parsed }));
      } catch {
        onSend(JSON.stringify({ type: selectedType, payload: {} }));
      }
    }
  };

  const handleKeyDown = (e) => {
    if (e.key === "Enter" && (e.metaKey || e.ctrlKey)) {
      handleSend();
    }
  };

  return (
    <div
      style={{
        borderTop: "1px solid #1a1a25",
        background: "#08080d",
        padding: "12px 16px",
        display: "flex",
        flexDirection: "column",
        gap: "10px",
      }}
    >
      <div style={{ display: "flex", alignItems: "center", gap: "8px", flexWrap: "wrap" }}>
        <button
          onClick={() => setRawMode(!rawMode)}
          style={{
            padding: "4px 10px",
            borderRadius: "4px",
            border: "1px solid #1a1a25",
            background: rawMode ? "#1a1a2f" : "transparent",
            color: rawMode ? "#6f8fcf" : "#666",
            fontFamily: "'IBM Plex Mono', monospace",
            fontSize: "10px",
            cursor: "pointer",
            textTransform: "uppercase",
            letterSpacing: "0.08em",
          }}
        >
          {rawMode ? "Raw JSON" : "Template"}
        </button>

        {!rawMode &&
          Object.keys(MESSAGE_TEMPLATES).map((type) => (
            <button
              key={type}
              onClick={() => handleTypeChange(type)}
              style={{
                padding: "4px 10px",
                borderRadius: "4px",
                border: `1px solid ${selectedType === type ? "#2d5a2d" : "#1a1a25"}`,
                background: selectedType === type ? "#1a2f1a" : "transparent",
                color: selectedType === type ? "#6fcf6f" : "#888",
                fontFamily: "'IBM Plex Mono', monospace",
                fontSize: "10px",
                cursor: "pointer",
              }}
            >
              {type}
            </button>
          ))}
      </div>

      <div style={{ display: "flex", gap: "10px", alignItems: "flex-end" }}>
        <textarea
          value={rawMode ? rawMessage : payload}
          onChange={(e) =>
            rawMode ? setRawMessage(e.target.value) : setPayload(e.target.value)
          }
          onKeyDown={handleKeyDown}
          rows={rawMode ? 2 : 4}
          placeholder={rawMode ? '{"type":"ping"}' : "payload JSON"}
          style={{
            flex: 1,
            background: "#0a0a0f",
            border: "1px solid #1a1a25",
            borderRadius: "6px",
            color: "#ddd",
            padding: "10px 12px",
            fontFamily: "'IBM Plex Mono', monospace",
            fontSize: "12px",
            resize: "vertical",
            outline: "none",
            lineHeight: "1.5",
          }}
        />
        <button
          onClick={handleSend}
          disabled={!connected}
          style={{
            padding: "10px 24px",
            borderRadius: "6px",
            border: "none",
            background: connected ? "#1a2f1a" : "#111",
            color: connected ? "#6fcf6f" : "#444",
            fontFamily: "'IBM Plex Mono', monospace",
            fontWeight: 700,
            fontSize: "12px",
            cursor: connected ? "pointer" : "not-allowed",
            textTransform: "uppercase",
            letterSpacing: "0.05em",
            flexShrink: 0,
            alignSelf: "stretch",
            minHeight: "60px",
          }}
        >
          Send
          <div style={{ fontSize: "9px", fontWeight: 400, marginTop: "2px", opacity: 0.6 }}>
            ⌘ Enter
          </div>
        </button>
      </div>
    </div>
  );
}

export default function WSDebug() {
  const [url, setUrl] = useState("ws://localhost:8080/ws");
  const [playerID, setPlayerID] = useState(() => crypto.randomUUID().slice(0, 8));
  const [status, setStatus] = useState("disconnected");
  const [logs, setLogs] = useState([]);
  const wsRef = useRef(null);
  const logEndRef = useRef(null);

  const addLog = useCallback((tag, data) => {
    setLogs((prev) => [...prev, { tag, data, time: new Date(), id: Date.now() + Math.random() }]);
  }, []);

  useEffect(() => {
    logEndRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [logs]);

  const connect = () => {
    const fullUrl = `${url}?player_id=${playerID}`;
    setStatus("connecting");
    addLog("system", `Connecting to ${fullUrl}`);

    const ws = new WebSocket(fullUrl);

    ws.onopen = () => {
      setStatus("connected");
      addLog("system", "Connection established");
    };

    ws.onmessage = (event) => {
      addLog("received", event.data);
    };

    ws.onerror = () => {
      addLog("error", "WebSocket error");
    };

    ws.onclose = (event) => {
      setStatus("disconnected");
      addLog("system", `Connection closed (code: ${event.code}, reason: ${event.reason || "none"})`);
      wsRef.current = null;
    };

    wsRef.current = ws;
  };

  const disconnect = () => {
    wsRef.current?.close();
    wsRef.current = null;
  };

  const send = (data) => {
    if (wsRef.current?.readyState === WebSocket.OPEN) {
      wsRef.current.send(data);
      addLog("sent", data);
    }
  };

  const clearLogs = () => setLogs([]);

  return (
    <div
      style={{
        height: "100vh",
        display: "flex",
        flexDirection: "column",
        background: "#0c0c14",
        color: "#ddd",
        fontFamily: "'IBM Plex Mono', 'SF Mono', 'Fira Code', monospace",
      }}
    >
      <link
        href="https://fonts.googleapis.com/css2?family=IBM+Plex+Mono:wght@400;500;600;700&display=swap"
        rel="stylesheet"
      />

      {/* Header */}
      <div
        style={{
          padding: "12px 20px",
          borderBottom: "1px solid #1a1a25",
          display: "flex",
          alignItems: "center",
          justifyContent: "space-between",
          background: "#08080d",
        }}
      >
        <div style={{ display: "flex", alignItems: "center", gap: "10px" }}>
          <span style={{ fontSize: "16px" }}>⚡</span>
          <span
            style={{
              fontWeight: 700,
              fontSize: "14px",
              letterSpacing: "0.05em",
              color: "#eee",
            }}
          >
            WS DEBUG
          </span>
          <span style={{ fontSize: "11px", color: "#444" }}>quizgame</span>
        </div>
        <div style={{ display: "flex", gap: "8px", alignItems: "center" }}>
          <span style={{ fontSize: "11px", color: "#555" }}>
            {logs.length} messages
          </span>
          <button
            onClick={clearLogs}
            style={{
              padding: "4px 12px",
              borderRadius: "4px",
              border: "1px solid #1a1a25",
              background: "transparent",
              color: "#666",
              fontFamily: "'IBM Plex Mono', monospace",
              fontSize: "10px",
              cursor: "pointer",
              textTransform: "uppercase",
              letterSpacing: "0.08em",
            }}
          >
            Clear
          </button>
        </div>
      </div>

      <ConnectionBar
        status={status}
        url={url}
        setUrl={setUrl}
        playerID={playerID}
        setPlayerID={setPlayerID}
        onConnect={connect}
        onDisconnect={disconnect}
      />

      {/* Log area */}
      <div
        style={{
          flex: 1,
          overflowY: "auto",
          padding: "4px 0",
        }}
      >
        {logs.length === 0 && (
          <div
            style={{
              display: "flex",
              alignItems: "center",
              justifyContent: "center",
              height: "100%",
              color: "#333",
              fontSize: "13px",
              flexDirection: "column",
              gap: "8px",
            }}
          >
            <span style={{ fontSize: "28px", opacity: 0.3 }}>⚡</span>
            <span>Connect and send messages to see them here</span>
          </div>
        )}
        {logs.map((entry) => (
          <LogEntry key={entry.id} entry={entry} />
        ))}
        <div ref={logEndRef} />
      </div>

      <MessageComposer onSend={send} connected={status === "connected"} />
    </div>
  );
}