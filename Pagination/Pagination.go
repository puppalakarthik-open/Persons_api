package Pagination
import(
	"example/persons/models"
	"errors"
	"strconv"
	"github.com/gin-gonic/gin"
	"math"
	"example/persons/repository/GetUsersRepo"
)
func Pagination(c *gin.Context)(int,int,error){
	var p models.Pagination
	p.Page=c.DefaultQuery("page","1")
	p.Limit=c.DefaultQuery("limit","3")
	var perPage int
	var page int
	perPage, _ = strconv.Atoi(p.Limit)
	page, _ = strconv.Atoi(p.Page)
	var Totalrows int64
	GetUsersRepo.CountRepo(&Totalrows)
	totalPages:=math.Ceil(float64(Totalrows)/float64(perPage))
	if page>int(totalPages){
		c.JSON(400,gin.H{"Message":"Invalid page"})
		return 0,0,errors.New("error")
	}
	offset:=(page-1) * perPage
	return perPage,offset,nil
}