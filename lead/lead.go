package lead

import (
	"github.com/KING-SAMM/go-fiber-crm-basic/database"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/gofiber/fiber"
)

type Lead struct {
	gorm.Model
	Name	string
	Company	string
	Email	string
	Phone	int
}

func GetLeads(c *fiber.Ctx) {

}

func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lead lead 
	db.Find(&lead, id)
	c.JSON(lead)
}

func NewLead(c *fiber.Ctx) {

}

func DeleteLead(c *fiber.Ctx) {

}