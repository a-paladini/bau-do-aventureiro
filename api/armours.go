package api

import (
	"database/sql"
	"net/http"

	db "github.com/a-paladini/bau-do-aventureiro/db/sqlc"
	"github.com/gin-gonic/gin"
)

type getArmourByIdRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}
type bodyArmourRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Slot        float64 `json:"slot"`
	Origin      string  `json:"origin"`
	CaBonus     int32   `json:"ca_bonus"`
	Penality    int32   `json:"penality"`
	Category    string  `json:"category"`
}

type listArmourRequest struct {
	Category string `form:"category"`
	PageID   int32  `form:"page_id" binding:"required,min=1"`
	PageSize int32  `form:"page_size" binding:"required,min=1,max=10"`
}

func (server *Server) createArmour(ctx *gin.Context) {
	var req bodyArmourRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateArmourParams{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Slot:        req.Slot,
		Origin:      req.Origin,
		CaBonus:     req.CaBonus,
		Penality:    req.Penality,
		Category:    req.Category,
	}

	armour, err := server.store.CreateArmour(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, armour)
}

func (server *Server) getArmour(ctx *gin.Context) {
	var req getArmourByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	armour, err := server.store.GetArmour(ctx, int32(req.ID))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, armour)
}

func (server *Server) listAllArmour(ctx *gin.Context) {
	var req listArmourRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListAllArmoursParams{
		Limit:  req.PageID,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	armours, err := server.store.ListAllArmours(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, armours)
}

func (server *Server) listArmours(ctx *gin.Context) {
	var req listArmourRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if req.Category != "" {
		arg := db.ListArmoursByCategoryParams{
			Category: req.Category,
			Limit:    req.PageSize,
			Offset:   (req.PageID - 1) * req.PageSize,
		}

		armours, err := server.store.ListArmoursByCategory(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, armours)
	} else {

		arg := db.ListAllWeaponsParams{
			Limit:  req.PageSize,
			Offset: (req.PageID - 1) * req.PageSize,
		}

		armours, err := server.store.ListAllWeapons(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, armours)
	}
}

func (server *Server) updateArmour(ctx *gin.Context) {
	var uriReq getArmourByIdRequest
	if err := ctx.ShouldBindUri(&uriReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var req bodyArmourRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateArmourParams{
		ID:          uriReq.ID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Slot:        req.Slot,
		Origin:      req.Origin,
		CaBonus:     req.CaBonus,
		Penality:    req.Penality,
		Category:    req.Category,
	}

	armour, err := server.store.UpdateArmour(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, armour)
}

func (server *Server) deleteArmour(ctx *gin.Context) {
	var req getArmourByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	armour, err := server.store.GetArmour(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = server.store.DeleteArmour(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, armour)
}
