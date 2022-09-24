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

type QuestionController interface {
	GetById(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	FilterByGroup(w http.ResponseWriter, r *http.Request)
}

type questionController struct {
	questionService service.QuestionService
}

// Get question by id godoc
// @tags question-manager-apis
// @Summary Get question by id
// @Description input: question's id => output: struct question
// @Accept json
// @Produce json
// @Param id path integer true "question's id"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /question/{id} [get]
func (c *questionController) GetById(w http.ResponseWriter, r *http.Request) {
	var res *model.Response

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		res = &model.Response{
			Data:    nil,
			Message: "get question failed: " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}

	tmp, err := c.questionService.GetById(id)
	if err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    tmp,
			Message: "get question successfully.",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

// Get all questions godoc
// @tags question-manager-apis
// @Summary Get questions
// @Description output: struct questions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /question/all [get]
func (c *questionController) GetAll(w http.ResponseWriter, r *http.Request) {
	var res *model.Response

	tmp, err := c.questionService.GetAll()
	if err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    tmp,
			Message: "get questions successfully.",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

// Create question godoc
// @tags question-manager-apis
// @Summary Create questions
// @Description input: question model.Question => output: status
// @Accept json
// @Produce json
// @param question body model.Question true "fill question"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /question/create [post]
func (c *questionController) Create(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	var question model.Question

	err := json.NewDecoder(r.Body).Decode(&question)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		res = &model.Response{
			Data:    nil,
			Message: "create question failed: " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}
	tmp, err := c.questionService.Create(&question)
	if err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    tmp,
			Message: "create questions successfully.",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

// Update question godoc
// @tags question-manager-apis
// @Summary Update questions
// @Description input: question model.Question => output: status
// @Accept json
// @Produce json
// @param question body model.Question true "change question"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /question/update [put]
func (c *questionController) Update(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	var question model.Question

	err := json.NewDecoder(r.Body).Decode(&question)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		res = &model.Response{
			Data:    nil,
			Message: "update question failed: " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}

	tmp, err := c.questionService.Update(question)
	if err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    tmp,
			Message: "update questions successfully.",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

// Delete question by id godoc
// @tags question-manager-apis
// @Summary Delete question by id
// @Description input: question's id => output: status
// @Accept json
// @Produce json
// @Param id path integer true "question's id"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /question/delete/{id} [delete]
func (c *questionController) Delete(w http.ResponseWriter, r *http.Request) {
	var res *model.Response

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		res = &model.Response{
			Data:    nil,
			Message: "delete question failed: " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}

	if err := c.questionService.Delete(id); err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    nil,
			Message: "delete question successfully.",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

// Filter question by group godoc
// @tags question-manager-apis
// @Summary filter question by group_id
// @Description input: question's group_id => output: group questions
// @Accept json
// @Produce json
// @Param group_id path integer true "group_id"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /question/filter/{group_id} [get]
func (c *questionController) FilterByGroup(w http.ResponseWriter, r *http.Request) {
	var res *model.Response

	idStr := chi.URLParam(r, "group_id")
	group_id, err := strconv.Atoi(idStr)

	questions, err := c.questionService.FilterByGroup(group_id)
	if err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    questions,
			Message: "filter question successfully.",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

func NewQuestionController() QuestionController {
	questionService := service.NewquestionService()
	return &questionController{
		questionService: questionService,
	}
}
