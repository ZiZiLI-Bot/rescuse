package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"rescues/model"
	"rescues/service"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type JudgeController interface {
	GetById(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetAdvices(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type judgeController struct {
	judgeService service.JudgeService
}

// Get judge by id godoc
// @tags judge-manager-apis
// @Summary Get judge by id
// @Description input: judge's id => output: struct judge
// @Accept json
// @Produce json
// @Param id path integer true "judge's id"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /judge/{id} [get]
func (c *judgeController) GetById(w http.ResponseWriter, r *http.Request) {
	var res *model.Response

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		res = &model.Response{
			Data:    nil,
			Message: "get judge failed: " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}

	tmp, err := c.judgeService.GetById(id)
	if err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    tmp,
			Message: "get judge successfully.",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

// Get all judges godoc
// @tags judge-manager-apis
// @Summary Get judges
// @Description output: struct judges
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /judge/all [get]
func (c *judgeController) GetAll(w http.ResponseWriter, r *http.Request) {
	var res *model.Response

	tmp, err := c.judgeService.GetAll()
	if err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    tmp,
			Message: "get judges successfully.",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

// Get judge-advice by score godoc
// @tags judge-manager-apis
// @Summary Get judge-advice by score
// @Description input: ScoreStress, ScoreDepress, ScoreAnxiety => output: struct judge-advice  
// @Description WARNING: ALl the score must be greater than 1 and less than 10
// @Accept json
// @Produce json
// @Param score_stress query integer true "score_Stress"
// @Param score_depress query integer true "score_depress"
// @Param score_anxiety query integer true "score_anxiety"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /judge/advice [get]
func (c *judgeController) GetAdvices(w http.ResponseWriter, r *http.Request) {
	var res *model.Response

	scoreStress, err := strconv.Atoi(r.URL.Query().Get("score_stress"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		res = &model.Response{
			Data:    nil,
			Message: "get scoreStress failed: " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}

	scoreDepress, err := strconv.Atoi(r.URL.Query().Get("score_depress"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		res = &model.Response{
			Data:    nil,
			Message: "get scoreDepress failed: " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}

	scoreAnxiety, err := strconv.Atoi(r.URL.Query().Get("score_anxiety"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		res = &model.Response{
			Data:    nil,
			Message: "get scoreAnxiety failed: " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}

	tmp, err := c.judgeService.GetAdvices(scoreStress, scoreDepress, scoreAnxiety)
	if err != nil {
		res = &model.Response{
			Data:    nil,
			Message: "get advice failed: " + err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    tmp,
			Message: "get advice successfully.",
			Success: true,
		}
	}

	render.JSON(w, r, res)
}

// Create judge godoc
// @tags judge-manager-apis
// @Summary Create judges
// @Description input: judge model.Judge => output: status
// @Accept json
// @Produce json
// @param judge body model.Judge true "fill judge"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /judge/create [post]
func (c *judgeController) Create(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	var judge model.Judge

	err := json.NewDecoder(r.Body).Decode(&judge)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		res = &model.Response{
			Data:    nil,
			Message: "create judge failed: " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}
	tmp, err := c.judgeService.Create(&judge)
	if err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    tmp,
			Message: "create judges successfully.",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

// Update judge godoc
// @tags judge-manager-apis
// @Summary Update judges
// @Description input: judge model.judge => output: status
// @Accept json
// @Produce json
// @param judge body model.Judge true "change judge"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /judge/update [put]
func (c *judgeController) Update(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	var judge model.Judge

	err := json.NewDecoder(r.Body).Decode(&judge)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		res = &model.Response{
			Data:    nil,
			Message: "update judge failed: " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}

	tmp, err := c.judgeService.Update(judge)
	if err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    tmp,
			Message: "update judges successfully.",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

// Delete judge by id godoc
// @tags judge-manager-apis
// @Summary Delete judge by id
// @Description input: judge's id => output: status
// @Accept json
// @Produce json
// @Param id path integer true "judge's id"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /judge/delete/{id} [delete]
func (c *judgeController) Delete(w http.ResponseWriter, r *http.Request) {
	var res *model.Response

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		res = &model.Response{
			Data:    nil,
			Message: "delete judge failed: " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}

	if err := c.judgeService.Delete(id); err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    nil,
			Message: "delete judge successfully.",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

func NewJudgeController() JudgeController {
	judgeService := service.NewJudgeService()
	return &judgeController{
		judgeService: judgeService,
	}
}
