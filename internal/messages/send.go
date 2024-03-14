package messages

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid/v5"
	"github.com/shadow/backend/internal/json"
	"github.com/shadow/backend/internal/models"
	"github.com/shadow/backend/internal/multiroom"
)

func Send(w http.ResponseWriter, r *http.Request) {
	var message models.MessageInfo
	chatID := chi.URLParam(r, "id")

	if err := json.Decode(r, &message); err != nil {
		json.Response(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	messageService := r.Context().Value("messagesService").(models.MessageService)

	m := &models.Message{
		ElementID: uuid.FromStringOrNil(chatID),
		Content:   message.Content,
		SenderID:  message.UserID,
		Kind:      message.Kind,
	}
	err := messageService.Create(m)
	if err != nil {
		json.Response(w, http.StatusInternalServerError, err.Error())
		return
	}

	message.ID = m.ID

	client := multiroom.WebSocketServer.FindClientByID(message.UserID.String())
	room := multiroom.WebSocketServer.FindRoomByID(chatID)
	roomMessage := multiroom.Message{
		Action:  multiroom.SendMessageAction,
		Message: message.Content,
		Target:  room,
		Sender:  client,
	}

	client.HandleNewMessage(roomMessage.Encode())
	// websocket.Broadcast(messageByte)
}
