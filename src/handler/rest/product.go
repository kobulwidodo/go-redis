package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get All Product
// @Description Get All Product
// @Tags Product
// @Produce json
// @Success 200 {object} entity.Response{}
// @Failure 400 {object} entity.Response{}
// @Failure 401 {object} entity.Response{}
// @Failure 404 {object} entity.Response{}
// @Failure 500 {object} entity.Response{}
// @Router /api/v1/product [GET]
func (r *rest) GetListProduct(ctx *gin.Context) {
	products, err := r.uc.Product.GetList(ctx.Copy().Request.Context())
	if err != nil {
		r.httpRespError(ctx, http.StatusInternalServerError, err)
		return
	}

	r.httpRespSuccess(ctx, http.StatusOK, "successfully get all product", products)
}
