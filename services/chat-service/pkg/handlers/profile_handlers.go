package handlers

import (
	"encoding/json"
	"github.com/satoshi1975/smartChat/services/chat-service/internal/models"
	"github.com/satoshi1975/smartChat/services/chat-service/internal/services"
	"github.com/satoshi1975/smartChat/services/chat-service/internal/utils"
	"net/http"

	"strconv"

	"github.com/julienschmidt/httprouter"
)

// ProfileHandler handles profile-related requests.
type ProfileHandler struct {
	service *services.ProfileService
}

func NewProfileHandler(service *services.ProfileService) *ProfileHandler {
	return &ProfileHandler{service: service}
}

// CreateProfile godoc
// @Summary Create a new profile
// @Description Create a new profile for the authenticated user
// @Tags profiles
// @Accept  json
// @Produce  json
// @Param profile body models.Profile true "Profile"
// @Success 201 {object} models.Profile
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /profiles [post]
func (h *ProfileHandler) CreateProfile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var profile models.Profile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("userID").(int)
	profile.UserID = userID

	if err := h.service.CreateProfile(r.Context(), &profile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, profile)
}

// GetProfile godoc
// @Summary Get a profile by ID
// @Description Get a profile by its ID
// @Tags profiles
// @Produce  json
// @Param id path int true "Profile ID"
// @Success 200 {object} models.Profile
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /profiles/{id} [get]
func (h *ProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid profile ID", http.StatusBadRequest)
		return
	}

	profile, err := h.service.GetProfileByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if profile == nil {
		http.Error(w, "Profile not found", http.StatusNotFound)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, profile)
}

// UpdateProfile godoc
// @Summary Update a profile
// @Description Update a profile by its ID
// @Tags profiles
// @Accept  json
// @Produce  json
// @Param id path int true "Profile ID"
// @Param profile body models.Profile true "Profile"
// @Success 200 {object} models.Profile
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /profiles/{id} [put]
func (h *ProfileHandler) UpdateProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid profile ID", http.StatusBadRequest)
		return
	}

	var profile models.Profile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	profile.ID = id

	if err := h.service.UpdateProfile(r.Context(), &profile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, profile)
}

// DeleteProfile godoc
// @Summary Delete a profile
// @Description Delete a profile by its ID
// @Tags profiles
// @Param id path int true "Profile ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /profiles/{id} [delete]
func (h *ProfileHandler) DeleteProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid profile ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteProfile(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// AddFriend godoc
// @Summary Add a friend
// @Description Add a friend to the user's friend list
// @Tags profiles
// @Accept json
// @Param id path int true "Profile ID"
// @Param friend body models.Profile true "Friend ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /profiles/{id}/friends [post]
func (h *ProfileHandler) AddFriend(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profileID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid profile ID", http.StatusBadRequest)
		return
	}

	var friend struct {
		FriendID int `json:"friend_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&friend); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.AddFriend(r.Context(), profileID, friend.FriendID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// BlockUser godoc
// @Summary Block a user
// @Description Block a user
// @Tags profiles
// @Accept json
// @Param id path int true "Profile ID"
// @Param block body models.Profile true "Blocked ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /profiles/{id}/block [post]
func (h *ProfileHandler) BlockUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profileID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid profile ID", http.StatusBadRequest)
		return
	}

	var block struct {
		BlockedID int `json:"blocked_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&block); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.BlockUser(r.Context(), profileID, block.BlockedID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}
