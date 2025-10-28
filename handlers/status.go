package handlers

import (
	
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func GetStatus (c  *gin.Context, db *sqlx.DB)