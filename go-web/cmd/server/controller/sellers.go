package controller

import (
	"io/ioutil"
	seller "mercado-frescos-time-7/go-web/internal/Seller"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Sellers struct {
	service seller.Service
}

type request struct {
	Cid  string  `json:"cid" binding:"required"`
	CompanyName string  `json:"company_name" binding:"required"`
	Address string  `json:"address" binding:"required"`
	Telephone  string `json:"telephone" binding:"required"`
}

func (c *Sellers) SellersStore() gin.HandlerFunc  {
	return func(ctx *gin.Context) {
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(422, gin.H{
				"error": err.Error(),
			})
			return
		}
		p, err := c.service.Store(req.Cid, req.CompanyName, req.Address, req.Telephone)
		if err != nil {
			ctx.JSON(409, gin.H{"error": err.Error(),})
			return
		}
		ctx.JSON(201, p)
	}
}


func (c *Sellers) SellersGetAll() gin.HandlerFunc  {
	return func(ctx *gin.Context) {
	p, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Sem resultados",})
			return
	}	
	ctx.JSON(200, p)
	}
}

func (c *Sellers) SellersGetId() gin.HandlerFunc  {
	return func(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(404, gin.H{
			"error": "Parametro de busca invalido"})
		return	
	}
	p, err := c.service.GetId(int(id))
	if err != nil {
		ctx.JSON(404, gin.H{
			"error": err.Error()})
		return	
	}
	ctx.JSON(200, p)
	}
}

func (c *Sellers) SellersUpdate() gin.HandlerFunc  {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": "Parametro de busca invalido"})
			return	
		}
		body := ctx.Request.Body
		defer body.Close()
	
		data, err := ioutil.ReadAll(body)
		if err != nil {
			ctx.JSON(404, err)
			return
		}
		p, err := c.service.Update(data, id)
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error()})
			return	
		}
		ctx.JSON(200, p)	
	}
}

func (c *Sellers) SellersDelete() gin.HandlerFunc  {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": "Parametro de busca invalido"})
			return	
		}
			err = c.service.Delete(id)
			if err != nil {
				ctx.JSON(404, gin.H{
					"error": err.Error()})
				return	
			}
			ctx.JSON(204, gin.H{
				"sucess": "Vendedor deletado"})
	}
}

func NewSellers(p seller.Service) *Sellers {
	return &Sellers{
		service: p,
	}
}