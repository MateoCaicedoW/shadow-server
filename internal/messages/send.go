package messages

import (
	"fmt"
	"net/http"

	"github.com/shadow/backend/internal/json"
	"github.com/shadow/backend/internal/models"
	"github.com/shadow/backend/internal/websocket"
)

func Send(w http.ResponseWriter, r *http.Request) {
	var message models.MessageInfo

	if err := json.Decode(r, &message); err != nil {
		json.Response(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	fmt.Println("message", message)

	// messageService := r.Context().Value("messageService").(models.MessageService)
	// err := messageService.Create(&models.Message{
	// 	ElementID: message.ElementID,
	// 	Content:   message.Content,
	// 	SenderID:  message.UserID,
	// 	Kind:      message.Kind,
	// })

	// if err != nil {
	// 	json.Response(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	mesageByte, err := json.Marshal(message)
	if err != nil {
		json.Response(w, http.StatusInternalServerError, err.Error())
		return
	}

	websocket.L.Broadcast(mesageByte)
}
