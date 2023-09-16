package controller

import (
	"app/service"
	"app/util"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetTreeComment
// @Summary Get comment tree
// @Description get comment tree
// @ID get-cmt-tree
// @Tags comment
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Post ID"
// @Success 200 {object} model.Comment
// @Failure 500 {object} Msg
// @Router /cmt/:id [get]
func (ctrl *Controller) GetTreeComment(c *gin.Context) {
	cmt, err := ctrl.service.Comment.GetTree(c, util.ToInt(c.Param("id")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, Msg{err.Error()})
	}

	var s interface{}
	err = json.Unmarshal([]byte(cmt), &s)
	jsonResponse(c, err,
		Response{http.StatusOK, s},
		ErrResponse{Code: http.StatusInternalServerError})
}

// PostComment
// @Summary Post a comment
// @Description post a comment
// @ID post-comment
// @Tags comment
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param  account body service.CommentBody true "Comment body"
// @Success 201
// @Failure 422,500 {object} Msg
// @Router /cmt [post]
func (ctrl *Controller) PostComment(c *gin.Context) {
	var cmtBody service.CommentBody
	ID := c.MustGet("ID").(int)
	if err := c.ShouldBindJSON(&cmtBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, Msg{"Invalid json provided"})
		return
	}
	err := ctrl.service.Comment.Add(c, ID, cmtBody)
	jsonResponse(c, err,
		Response{Code: http.StatusCreated},
		ErrResponse{Code: http.StatusInternalServerError})
}
