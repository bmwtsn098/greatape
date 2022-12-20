package handlers

import (
	"net/http"

	. "github.com/xeronith/diamante/contracts/network/http"
	pipeline "github.com/xeronith/diamante/network/http"
	. "rail.town/infrastructure/components/api/protobuf"
	. "rail.town/infrastructure/components/contracts"
)

type getOutboxHandler struct {
}

func GetOutboxHandler() IHttpHandler {
	return &getOutboxHandler{}
}

func (handler *getOutboxHandler) Method() string {
	return http.MethodGet
}

func (handler *getOutboxHandler) Path() string {
	return "/u/:username/outbox"
}

func (handler *getOutboxHandler) HandlerFunc() HttpHandlerFunc {
	return func(x IServerDispatcher) error {
		request := &GetOutboxRequest{}
		result := &GetOutboxResult{}

		onRequestUnmarshalled := func(request *GetOutboxRequest) {
			request.Username = x.Param("username")
		}

		return pipeline.Handle(x,
			"get_outbox",
			GET_OUTBOX_REQUEST,
			GET_OUTBOX_RESULT,
			request, result,
			onRequestUnmarshalled,
			false,
		)
	}
}
