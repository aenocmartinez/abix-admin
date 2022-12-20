package controller

import (
	"abix360/src/usecase"
	"abix360/src/view/dto"
	formrequest "abix360/src/view/form-request"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateField(c *gin.Context) {
	var req formrequest.CreateFieldFormRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var subfields []dto.SubfieldDto
	if len(req.Subfield) > 0 {
		for _, field := range req.Subfield {
			subfields = append(subfields, dto.SubfieldDto{
				Id:    field.Id,
				Order: field.Order,
			})
		}
	}

	useCase := usecase.CreateFieldUseCase{}
	code, err := useCase.Execute(dto.FieldDto{
		Name:         req.Name,
		Description:  req.Description,
		Abbreviation: req.Abbreviation,
		Subfields:    subfields,
	})

	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(code, gin.H{"message": "el campo se ha creado exitosamente"})
}

func UpdateField(c *gin.Context) {
	var req formrequest.UpdateFieldFormRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var subfields []dto.SubfieldDto
	if len(req.Subfield) > 0 {
		for _, field := range req.Subfield {
			subfields = append(subfields, dto.SubfieldDto{
				Id:    field.Id,
				Order: field.Order,
			})
		}
	}

	useCase := usecase.UpdateFieldUseCase{}
	code, err := useCase.Execute(dto.FieldDto{
		Name:         req.Name,
		Description:  req.Description,
		Abbreviation: req.Abbreviation,
		Subfields:    subfields,
		Id:           req.Id,
	})

	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(code, gin.H{"message": "el campo se ha actualizado con éxito"})
}

func ViewField(c *gin.Context) {
	var strId string = c.Param("id")
	if len(strId) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parámetro no válido"})
		return
	}

	id, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parámetro no válido"})
		return
	}

	useCase := usecase.FindFieldUseCase{}
	field, err := useCase.Execute(int64(id))
	if err != nil {
		c.AbortWithStatusJSON(202, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"result": field})
}

func DeleteField(c *gin.Context) {
	var strId string = c.Param("id")
	if len(strId) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parámetro no válido"})
		return
	}

	id, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parámetro no válido"})
		return
	}

	useCase := usecase.DeleteFieldUseCase{}
	code, err := useCase.Execute(int64(id))
	if err != nil {
		c.AbortWithStatusJSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(code, gin.H{"result": "el campo se ha eliminado con éxito"})
}

func AllFields(c *gin.Context) {
	useCase := usecase.ListFieldsUseCase{}
	fields := useCase.Execute()
	c.JSON(200, gin.H{"result": fields})
}

func SearchFields(c *gin.Context) {
	name := c.Query("name")
	useCase := usecase.SearchFieldUseCase{}
	fields := useCase.Execute(name)
	c.JSON(200, gin.H{"result": fields})
}
