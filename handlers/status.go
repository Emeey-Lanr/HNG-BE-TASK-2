package handlers

import (
	"be-task2/helpers"
	"be-task2/models"
	"be-task2/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

)

func GetStatus (c  *gin.Context, db *sqlx.DB) {

	status , err := repository.GetTotalCountryAndTimeStamp(db)

	if err != nil {
		helpers.ErrorResponse(http.StatusNotFound, models.ErrorResp{Error: "Unable to creat status summary", Details: err.Error() }, c)
	 return
	}

	helpers.SuccessResponse(http.StatusOK, status, c)
}