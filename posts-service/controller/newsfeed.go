package controller

import (
	"app/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Feed
// @Summary Feed
// @Description feed
// @ID feed
// @Tags feed
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param lim query int true "Limit"
// @Param off query int true "Offset"
// @Success 200 {object} []int64
// @Failure 500 {object} Msg
// @Router /feed [get]
func (ctrl *Controller) Feed(c *gin.Context) {
	//id := toInt(c.Param("id"))
	ID := c.MustGet("ID").(int)
	limit := util.ToInt(c.Query("lim"))
	offset := util.ToInt(c.Query("off"))

	var ids_arr []int32
	ids_arr, _ = parsePropTypes(c)
	ids_arr = append(ids_arr, int32(ID))

	fmt.Println(ids_arr)

	feed, err := ctrl.service.Feed.GetNewsfeed(c, ids_arr, int32(limit), int32(offset))
	log.Println(ID, limit, offset, feed, err)
	jsonResponse(c, err,
		Response{http.StatusOK, feed},
		ErrResponse{Code: http.StatusInternalServerError})
}

func parsePropTypes(c *gin.Context) ([]int32, bool) {
	var propertyTypes []int32
	var propertyTypesAsString string
	var propertyTypesAsArrayOfStrings []string
	propertyTypesAsString, success := c.GetQuery("ids")
	propertyTypesAsArrayOfStrings = strings.Split(propertyTypesAsString, ",")
	if success {
		for _, propertyTypeAsString := range propertyTypesAsArrayOfStrings {
			i, err := strconv.Atoi(propertyTypeAsString)
			if err != nil {
				//svc.ErrorWithJson(c, 400, errors.New("invalid property_types array"))
				return nil, true
			}
			propertyTypes = append(propertyTypes, int32(i))
		}
	}
	return propertyTypes, false
}
