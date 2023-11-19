package handlers

import (
	"lab2/src/flight-service/dbhandler"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FlightHandler struct {
	dbhandler dbhandler.DBHandler
}

func (hand *FlightHandler) GetAirportHandler(c *gin.Context) {

	airportID := c.Param("airportId")

	airport, err := hand.dbhandler.GetAirportByID(airportID)

	if err != nil {
		log.Printf("failed to get airport: %s", err)
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, airport)
}

func (hand *FlightHandler) GetAllFlightsHandler(c *gin.Context) {

	flights, err := hand.dbhandler.GetAllFlights()
	if err != nil {
		log.Printf("failed to get flghts: %s", err)
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, flights)
}

func (hand *FlightHandler) GetFlightHandler(c *gin.Context) {

	flightID := c.Param("flightNumber")

	flight, err := hand.dbhandler.GetFlightByNumber(flightID)

	if err != nil {
		log.Printf("failed to get flight: %s", err)
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, flight)
}
