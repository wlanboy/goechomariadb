package application

import (
	"net/http"

	model "../model"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

//GetAll e.GET("/api/v1/event", GetAll)
func (goservice *GoService) GetAll(c echo.Context) (err error) {
	errdb, resp := model.GetAllEvents(goservice.DB)
	if errdb != "" {
		return handleModelError(c, errdb)
	}
	return c.JSON(http.StatusOK, resp)
}

//GetByID e.GET("/api/v1/event/:id", GetByID)
func (goservice *GoService) GetByID(c echo.Context) (err error) {
	// Event ID from path `event/:id`
	id := c.Param("id")
	uuid, uuiderr := uuid.FromString(id)

	if uuiderr != nil {
		return c.String(http.StatusInternalServerError, "cannot parse uuid")
	}

	c.Logger().Info(id)
	errdb, resp := model.GetEventByID(uuid.String(), goservice.DB)
	if errdb != "" {
		return handleModelError(c, errdb)
	}
	return c.JSON(http.StatusOK, resp)
}

//PostCreate e.POST("/api/v1/event", PostCreate)
func (goservice *GoService) PostCreate(c echo.Context) (err error) {
	var event *model.Event = &model.Event{}

	if err = c.Bind(event); err != nil {
		return err
	}
	errdb, resp := model.SaveEvent(*event, goservice.DB)
	if errdb != "" {
		return handleModelError(c, errdb)
	}
	return c.JSON(http.StatusCreated, resp)
}

func handleModelError(c echo.Context, errdb string) (err error) {
	c.Logger().Error("Model error")
	c.Logger().Error(errdb)
	return c.String(http.StatusInternalServerError, errdb)
}
