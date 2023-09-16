package controller

import (
	"encoding/json"
	"net/http"

	"app/util"

	"github.com/gin-gonic/gin"
)

// GetFriendsDetail
// @Summary Get friend detail
// @Description get friend detail
// @ID get-friend-detail
// @Tags relationship
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Friend ID"
// @Success 200 {object} []FriendResponse
// @Failure 500 {object} Msg
// @Router /rel/friends/{id} [get]
func (ctrl *Controller) GetFriendsDetail(c *gin.Context) {
	id := util.ToInt(c.Param("id"))
	friends, err := ctrl.service.Relationship.FriendsDetail(c, id)
	var s interface{}
	json.Unmarshal([]byte(friends), &s)
	jsonResponse(c, err,
		Response{http.StatusOK, s},
		ErrResponse{Code: http.StatusInternalServerError})

}

// GetMutualFriends
// @Summary Get mutual friends count
// @Description get mutual friends
// @ID get-mutual-friend
// @Tags relationship
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Profile ID"
// @Success 200 {object} dataResponse
// @Failure 500 {object} Msg
// @Router /rel/mutual-friends/{id} [get]
func (ctrl *Controller) GetMutualFriends(c *gin.Context) {
	ID := c.MustGet("ID").(int)
	id := util.ToInt(c.Param("id"))
	mf, err := ctrl.service.Relationship.MutualFriends(c, ID, id)
	jsonResponse(c, err,
		Response{http.StatusOK, dataResponse{mf}},
		ErrResponse{Code: http.StatusInternalServerError})
}

// ChangeType
// @Summary Change relationship
// @Description get change relationship
// @ID change-relationship
// @Tags relationship
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Profile ID"
// @Param type path int true "Relationship type"
// @Success 200
// @Failure 500 {object} Msg
// @Router /rel/{id}/{type} [put]
func (ctrl *Controller) ChangeType(c *gin.Context) {
	ID := c.MustGet("ID").(int)
	id := util.ToInt(c.Param("id"))
	t := c.Param("type")
	err := ctrl.service.Relationship.ChangeType(c, ID, id, t)
	jsonResponse(c, err,
		Response{Code: http.StatusOK},
		ErrResponse{Code: http.StatusInternalServerError})
}

// GetMutualAndType
// @Summary Get mutual friends and type
// @Description get mutual friends and type
// @ID get-mutual-friends-and-type
// @Tags relationship
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Profile ID"
// @Success 200 {object} GetMutualAndTypeResponse
// @Failure 500 {object} Msg
// @Router /rel/mutual-type/{id} [get]
func (ctrl *Controller) GetMutualAndType(c *gin.Context) {
	ID := c.MustGet("ID").(int)
	id := util.ToInt(c.Param("id"))
	t := ctrl.service.Relationship.GetRelationshipWith(c, ID, id)
	m, err := ctrl.service.Relationship.MutualFriends(c, ID, id)

	jsonResponse(c, err,
		Response{http.StatusOK, GetMutualAndTypeResponse{t, m}},
		ErrResponse{Code: http.StatusInternalServerError})
}
