package api

import (
	"database/sql"
	"net/http"

	db "github.com/a-paladini/bau-do-aventureiro/db/sqlc"
	"github.com/gin-gonic/gin"
)

type bodyWeaponRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Slot        float64 `json:"slot"`
	Origin      string  `json:"origin"`
	Damage      string  `json:"damage"`
	Critical    string  `json:"critical"`
	Range       string  `json:"range"`
	TypeDamage  string  `json:"type_damage"`
	Property    string  `json:"property"`
	Proficiency string  `json:"proficiency"`
	Special     string  `json:"special"`
}

type listWeaponsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=10"`
}

type getWeaponByIdRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type listWeaponsByCategoryRequest struct {
	Category string `uri:"category" binding:"required"`
}

func (server *Server) createWeapon(ctx *gin.Context) {
	var req bodyWeaponRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateWeaponParams{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Slot:        req.Slot,
		Origin:      req.Origin,
		Damage:      req.Damage,
		Critical:    req.Critical,
		Range:       req.Range,
		TypeDamage:  req.TypeDamage,
		Property:    req.Property,
		Proficiency: req.Proficiency,
		Special:     sql.NullString{String: req.Special, Valid: true},
	}

	weapon, err := server.store.CreateWeapon(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, weapon)
}

func (server *Server) getWeapon(ctx *gin.Context) {
	var req getWeaponByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	weapon, err := server.store.GetWeapon(ctx, int32(req.ID))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, weapon)
}

func (server *Server) listAllWeapons(ctx *gin.Context) {
	var req listWeaponsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListAllWeaponsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	listWeapon, err := server.store.ListAllWeapons(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, listWeapon)
}

func (server *Server) listWeaponsByCategory(ctx *gin.Context) {
	var uriReq listWeaponsByCategoryRequest
	if err := ctx.ShouldBindUri(&uriReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var queryReq listWeaponsRequest
	if err := ctx.ShouldBindQuery(&queryReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListWeaponsByCategoryParams{
		TypeDamage: uriReq.Category,
		Limit:      queryReq.PageSize,
		Offset:     (queryReq.PageID - 1) * queryReq.PageSize,
	}

	listWeapon, err := server.store.ListWeaponsByCategory(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, listWeapon)
}

func (server *Server) deleteWeapon(ctx *gin.Context) {
	var req getWeaponByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	weapon, err := server.store.GetWeapon(ctx, int32(req.ID))
	if err != nil {
		if err != sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = server.store.DeleteWeapon(ctx, int32(req.ID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, weapon)
}

func (server *Server) updateWeapon(ctx *gin.Context) {
	var id getWeaponByIdRequest
	if err := ctx.BindUri(&id); err != nil {

	}

	var req bodyWeaponRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateWeaponParams{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Slot:        req.Slot,
		Origin:      req.Origin,
		Damage:      req.Damage,
		Critical:    req.Critical,
		Range:       req.Range,
		TypeDamage:  req.TypeDamage,
		Property:    req.Property,
		Proficiency: req.Proficiency,
		Special:     sql.NullString{String: req.Special, Valid: true},
	}

	weapon, err := server.store.UpdateWeapon(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, weapon)
}
