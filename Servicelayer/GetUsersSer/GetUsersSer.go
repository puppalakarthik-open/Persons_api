package GetUsersSer

import (
	"example/persons/models"
	"example/persons/Pagination"
	"github.com/gin-gonic/gin"
	"example/persons/repository/GetUsersRepo"
)
func GetUsersSer(c *gin.Context)([]models.Person,error){
	personModels := []models.Person{}

	// var p models.Pagination
	// p.Page=c.DefaultQuery("page","1")
	// p.Limit=c.DefaultQuery("limit","3")
	// var perPage int
	// var page int
	// perPage, _ = strconv.Atoi(p.Limit)
	// page, _ = strconv.Atoi(p.Page)
	// var Totalrows int64
	// GetUsersRepo.CountRepo(&Totalrows)
	// totalPages:=math.Ceil(float64(Totalrows)/float64(perPage))
	// if page>int(totalPages){
	// 	c.JSON(400,gin.H{"Message":"Invalid page"})
	// 	return personModels,errors.New("error")
	// }
	// offset:=(page-1) * perPage
	perPage,offset,err:=Pagination.Pagination(c)
	if err!=nil{
		return personModels,err
	}
	err = GetUsersRepo.GetUsersRepo(perPage,offset,&personModels)
	if err != nil {
		c.JSON(400, gin.H{"Message": err})
		return personModels,err
	}
	return personModels,nil
}