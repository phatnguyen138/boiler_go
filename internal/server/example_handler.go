package server

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	db "bolier_go/internal/db/sqlc"
)

type ExamName struct {
	Name string `json:"name"`
}

func (s *Server) getExample(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	selectedExample, err := s.query.GetExample(c.Request().Context(), int32(id))

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, selectedExample)
}

func (s *Server) listExample(c echo.Context) error {
	limit, err := strconv.Atoi(c.QueryParam("limit"))

	if err != nil || limit < 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(c.QueryParam("offset"))

	if err != nil || offset < 0 {
		offset = 0
	}

	listParam := db.ListExamplesParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	exs, err := s.query.ListExamples(c.Request().Context(), listParam)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, exs)
}

func (s *Server) addExample(c echo.Context) error {
	u := new(ExamName)

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	newExam, err := s.query.AddExample(c.Request().Context(), u.Name)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, newExam)
}
