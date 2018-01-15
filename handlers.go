package goclicker

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// e.GET("/:key", LinkHandler)
func LinkHandler(c echo.Context) error {
	key := c.Param("key")
	repo := NewLinkRepository()
	item := repo.FindByKey(key)
	if item == nil {
		return c.String(http.StatusNotFound, "key not found")
	}
	Logging(item.ID, c.Request())
	return c.Redirect(http.StatusFound, item.URL)
}

// e.POST("/api/links", SaveLinkHandler)
func SaveLinkHandler(c echo.Context) error {
	u := new(URLRequest)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "")
	}
	if len(u.LongURL) == 0 {
		return c.String(http.StatusBadRequest, "")
	}
	repo := NewLinkRepository()
	item := repo.FindByURL(u.LongURL)
	if item == nil {
		item := Link{URL: u.LongURL}
		repo.Save(item)
		found := repo.FindByURL(u.LongURL)
		found.CreateEndpoint(GetEndpoint())
		return c.JSON(http.StatusOK, found)
	}
	item.CreateEndpoint(GetEndpoint())
	return c.JSON(http.StatusOK, item)
}

// e.GET("/api/links", GetLinksHandler)
func GetLinksHandler(c echo.Context) error {
	repo := NewLinkRepository()
	items := repo.All()
	endpoint := GetEndpoint()
	resp := make([]Link, 0)
	for _, item := range items {
		item.CreateEndpoint(endpoint)
		resp = append(resp, item)
	}
	return c.JSON(http.StatusOK, resp)
}

// e.PATCH("/api/links/:id", UpdateLinkHandler)
func UpdateLinkHandler(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.String(http.StatusNotFound, "id not a number")
	}
	u := new(URLRequest)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "")
	}
	if len(u.LongURL) == 0 {
		return c.String(http.StatusBadRequest, "")
	}
	repo := NewLinkRepository()
	item := repo.Find(id)
	if item == nil {
		return c.String(http.StatusNotFound, "id not found")
	}
	item.URL = u.LongURL
	repo.Save(*item)
	return c.JSON(http.StatusOK, item)
}

// e.GET("/api/links/:key", GetLinkHandler)
func GetLinkHandler(c echo.Context) error {
	key := c.Param("key")
	repo := NewLinkRepository()
	item := repo.FindByKey(key)
	if item == nil {
		return c.String(http.StatusNotFound, "key not found")
	}
	resp := LinkResponse{}
	resp.ID = item.ID
	resp.Key = item.Key
	resp.URL = item.URL
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	resp.CreateEndpoint(GetEndpoint())

	logRepository := NewAccessLogRepository()
	logs := logRepository.FindByLinkID(item.ID)
	resp.Logs = logs

	return c.JSON(http.StatusOK, resp)
}
