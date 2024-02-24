package websocket

// Listener is used to listen for messages as a hub
type Listener struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

var L = &Listener{
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	clients:    make(map[*Client]bool),
}

func (l *Listener) Register(c *Client) {
	l.register <- c
}

func (l *Listener) Unregister(c *Client) {
	l.unregister <- c
}

func (l *Listener) Broadcast(m []byte) {
	l.broadcast <- m
}

// Run starts the listener
func (l *Listener) Run() {
	for {
		select {
		case client := <-l.register:
			l.clients[client] = true
		case client := <-l.unregister:
			if _, ok := l.clients[client]; ok {
				delete(l.clients, client)
				close(client.send)
			}
		case message := <-l.broadcast:
			for client := range l.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(l.clients, client)
				}
			}
		}
	}
}
