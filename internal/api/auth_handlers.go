package api

import (
	"main/internal/jsonutils"
	"net/http"
)

func (api *Api) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (api *Api) handleLogout(w http.ResponseWriter, r *http.Request) {
	err := api.Sessions.RenewToken(r.Context())
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": "unexpected internal server error",
		})
		return
	}

	api.Sessions.Remove(r.Context(), "AuthenticatedUserId")
	jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{
		"message": "logged out succesfully",
	})
}