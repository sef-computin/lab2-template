package handlers

import (
	"github.com/gin-gonic/gin"
)

type ServicesStruct struct {
	TicketServiceAddress string
	FlightServiceAddress string
	BonusServiceAddress  string
}

type GatewayService struct {
	Config ServicesStruct
}

func NewGatewayService(config *ServicesStruct) *GatewayService {
	return &GatewayService{Config: *config}
}

func (gs *GatewayService) GetAllFlights(c *gin.Context) {

}

func (gs *GatewayService) GetUserInfo(c *gin.Context) {

}

func (gs *GatewayService) GetUserTickets(c *gin.Context) {

}

func (gs *GatewayService) GetUserTicket(c *gin.Context) {

}

func (gs *GatewayService) BuyTicket(c *gin.Context) {

}

func (gs *GatewayService) CancelTicket(c *gin.Context) {

}

func (gs *GatewayService) GetPrivilege(c *gin.Context) {

}
