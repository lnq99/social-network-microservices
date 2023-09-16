package controller

import (
	"app/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetPhoto
// @Summary Get photo
// @Description get photo
// @ID get-photo
// @Tags photo
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Photo ID"
// @Success 200 {object} model.Photo
// @Failure 500 {object} Msg
// @Router /photo/{id} [get]
func (ctrl *Controller) GetPhoto(c *gin.Context) {
	photo, err := ctrl.service.Photo.GetPhoto(c, util.ToInt(c.Param("id")))
	jsonResponse(c, err,
		Response{http.StatusCreated, photo},
		ErrResponse{Code: http.StatusInternalServerError})

}

// GetPhotoByUserId
// @Summary Get photo by user id
// @Description get photo by user id
// @ID get-photo-by-user-id
// @Tags photo
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Profile ID"
// @Success 200 {object} []model.Photo
// @Failure 500 {object} Msg
// @Router /photo/u/{id} [get]
func (ctrl *Controller) GetPhotoByUserId(c *gin.Context) {
	photo, err := ctrl.service.Photo.GetPhotoByUserId(c, util.ToInt(c.Param("id")))
	jsonResponse(c, err,
		Response{http.StatusCreated, photo},
		ErrResponse{Code: http.StatusInternalServerError})
}
