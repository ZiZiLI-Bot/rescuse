package controller

import (
	"rescues/model"
	"rescues/service"

	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type QuizzController interface {
	GetById(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type quizzController struct {
	quizzService service.QuizzService
}

// Get quizz by id godoc
// @tags quizz-manager-apis
// @Summary Get quizz by id
// @Description input: quizz's id => output: struct quizz
// @Accept json
// @Produce json
// @Param id path integer true "quizz's id"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /quizz/{id} [get]
func (c *quizzController) GetById(w http.ResponseWriter, r *http.Request) {
	var res *model.Response

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		res = &model.Response{
			Data:    nil,
			Message: "get quizz failed: " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}

	tmp, err := c.quizzService.GetById(id)
	if err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    tmp,
			Message: "get quizz successfully.",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

// Get all quizzs godoc
// @tags quizz-manager-apis
// @Summary Get quizzs
// @Description output: struct quizzs
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /quizz/all [get]
func (c *quizzController) GetAll(w http.ResponseWriter, r *http.Request) {
	var res *model.Response

	tmp, err := c.quizzService.GetAll()
	if err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    tmp,
			Message: "get quizzs successfully.",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

// Create quizz godoc
// @tags quizz-manager-apis
// @Summary Create quizzs
// @Description input: quizz model.Quizz => output: status
// @Accept json
// @Produce json
// @param quizz body model.Quizz true "fill quizz"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /quizz/create [post]
func (c *quizzController) Create(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	var quizz model.Quizz

	err := json.NewDecoder(r.Body).Decode(&quizz)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		res = &model.Response{
			Data:    nil,
			Message: "create quizz failed: " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}
	tmp, err := c.quizzService.Create(&quizz)
	if err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    tmp,
			Message: "create quizzs successfully.",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

// Update quizz godoc
// @tags quizz-manager-apis
// @Summary Update quizzs
// @Description input: quizz model.Quizz => output: status
// @Accept json
// @Produce json
// @param quizz body model.Quizz true "change quizz"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /quizz/update [put]
func (c *quizzController) Update(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	var quizz model.Quizz

	err := json.NewDecoder(r.Body).Decode(&quizz)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		res = &model.Response{
			Data:    nil,
			Message: "update quizz failed: " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}

	tmp, err := c.quizzService.Update(quizz)
	if err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    tmp,
			Message: "update quizzs successfully.",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

// Delete quizz by id godoc
// @tags quizz-manager-apis
// @Summary Delete quizz by id
// @Description input: quizz's id => output: status
// @Accept json
// @Produce json
// @Param id path integer true "quizz's id"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /quizz/delete/{id} [delete]
func (c *quizzController) Delete(w http.ResponseWriter, r *http.Request) {
	var res *model.Response

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		res = &model.Response{
			Data:    nil,
			Message: "delete quizz failed: " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}

	if err := c.quizzService.Delete(id); err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    nil,
			Message: "delete quizz successfully.",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

func NewQuizzController() QuizzController {
	quizzService := service.NewQuizzService()
	return &quizzController{
		quizzService: quizzService,
	}
}
