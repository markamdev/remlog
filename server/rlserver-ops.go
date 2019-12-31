package server

import (
	"fmt"
	"math/rand"
	"net"

	rlp "github.com/markamdev/remlog/protocol"
)

func registerClient(client net.Addr, request rlp.Message) {
	if request.Identifier != rlp.IdentifierGlobal && request.Identifier != rlp.IdentifierDebug {
		// invalid identifier set in request - only debug and global allowed
		fmt.Println("Invalid identifier set in registration request")
		return
	}
	response := rlp.Message{}
	response.Type = rlp.Confirm

	if request.Identifier == rlp.IdentifierDebug {
		// if debug identifier used then respond with same one
		response.Identifier = request.Identifier
	} else {
		clientName := string(request.Content)

		id, ok := serverContext.clients[clientName]
		if !ok {
			// client not registered yet
			newID := rand.Uint32() // TODO consider using some hash instead of random uint
			serverContext.clients[clientName] = newID
			serverContext.clientsRev[newID] = clientName
			response.Identifier = newID
		} else {
			// client already registered
			response.Identifier = id
		}
	}

	data, err := rlp.MessageToData(&response)
	if err != nil {
		fmt.Println("Failed to convert registration response to bytes")
		return
	}

	n, err := serverContext.listener.WriteTo(data, client)
	if n != len(data) || err != nil {
		fmt.Println("Failed to send registration response")
	}

	fmt.Println("Client registration request processed succesfully")
}

func saveLogContent(client net.Addr, msg rlp.Message) {
	var clientName string
	if msg.Identifier == rlp.IdentifierDebug {
		clientName = "DEBUG"
	} else {
		name, ok := serverContext.clientsRev[msg.Identifier]
		if !ok {
			// (maybe it will be good to inform sender that identifier is unknown?)
			return
		}
		clientName = name
	}
	logString := fmt.Sprintf("(%s) %s", clientName, string(msg.Content))
	// temporary solution - in final version there will be another log output
	fmt.Println(logString)
}
