package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Search
// @Summary Search by username
// @Description search by username
// @ID search
// @Tags search
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id query int true "Key"
// @Success 200 {object} []SearchResponse
// @Failure 500 {object} Msg
// @Router /search [get]
func (ctrl *Controller) Search(c *gin.Context) {
	ID := c.MustGet("ID").(int)
	key := c.Query("k")
	res, err := ctrl.service.Profile.SearchName(c, ID, key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Msg{err.Error()})
	}

	var s interface{}
	err = json.Unmarshal([]byte(res), &s)
	jsonResponse(c, err,
		Response{http.StatusOK, s},
		ErrResponse{Code: http.StatusInternalServerError})
}
