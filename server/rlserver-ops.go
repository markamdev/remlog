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
		// TODO in future client will be identified by message content (not by IP:PORT string)
		clientAddress := client.String()

		id, ok := serverContext.clients[clientAddress]
		if !ok {
			// client not registered yet
			newID := rand.Uint32()
			serverContext.clients[clientAddress] = newID
			serverContext.clientsRev[newID] = clientAddress
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
	clientID, ok := serverContext.clientsRev[msg.Identifier]
	if !ok {
		// (maybe it will be good to inform sender that identifier is unknown?)
		return
	}
	logString := fmt.Sprintf("(%v) %s", clientID, string(msg.Content))
	// temporary solution - in final version there will be another log output
	fmt.Println(logString)
}
