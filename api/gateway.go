package api

import (
	"net/http"

	"github.com/NebulousLabs/Sia/modules"

	"github.com/julienschmidt/httprouter"
)

type GatewayInfo struct {
	NetAddress modules.NetAddress `json:"netaddress"`
	Peers      []modules.Peer     `json:"peers"`
}

// gatewayHandler handles the API call asking for the gatway status.
func (srv *Server) gatewayHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	peers := srv.gateway.Peers()
	if peers == nil {
		peers = make([]modules.Peer, 0)
	}
	writeJSON(w, GatewayInfo{srv.gateway.Address(), peers})
}

// gatewayAddHandler handles the API call to add a peer to the gateway.
func (srv *Server) gatewayAddHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	addr := modules.NetAddress(ps.ByName("netaddress"))
	err := srv.gateway.Connect(addr)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeSuccess(w)
}

// gatewayRemoveHandler handles the API call to remove a peer from the gateway.
func (srv *Server) gatewayRemoveHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	addr := modules.NetAddress(ps.ByName("netaddress"))
	err := srv.gateway.Disconnect(addr)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeSuccess(w)
}
