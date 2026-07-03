package article

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"sharing-vision-be/internal/common"
)

type ArticleHandler struct {
	service *ArticleService
}

func NewArticleHandler(service *ArticleService) *ArticleHandler {
	return &ArticleHandler{service: service}
}

func (h *ArticleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateArticleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.WriteError(w, err)
		return
	}
	defer r.Body.Close()

	if errs := common.Validate(&req); errs != nil {
		common.WriteValidateError(w, errs)
		return
	}

	article, err := h.service.Create(&req)
	if err != nil {
		common.WriteError(w, err)
		return
	}

	common.WriteJSON(w, &common.Response{
		Msg:    "success create article",
		Data:   article,
		Status: http.StatusCreated,
	})
}

func (h *ArticleHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	limitStr := r.PathValue("limit")
	offsetStr := r.PathValue("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	articles, total, err := h.service.FindAll(limit, offset)
	if err != nil {
		common.WriteError(w, err)
		return
	}

	if articles == nil {
		articles = []Article{}
	}

	common.WriteJSON(w, &common.Response{
		Msg: "success get articles",
		Data: map[string]any{
			"articles": articles,
			"total":    total,
		},
	})
}

func (h *ArticleHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		common.WriteJSON(w, &common.Response{
			Msg:    "invalid article id",
			Status: http.StatusBadRequest,
		})
		return
	}

	article, err := h.service.FindByID(id)
	if err != nil {
		common.WriteJSON(w, &common.Response{
			Msg:    err.Error(),
			Status: http.StatusNotFound,
		})
		return
	}

	common.WriteJSON(w, &common.Response{
		Msg:  "success get article",
		Data: article,
	})
}

func (h *ArticleHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		common.WriteJSON(w, &common.Response{
			Msg:    "invalid article id",
			Status: http.StatusBadRequest,
		})
		return
	}

	var req UpdateArticleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.WriteError(w, err)
		return
	}
	defer r.Body.Close()

	if errs := common.Validate(&req); errs != nil {
		common.WriteValidateError(w, errs)
		return
	}

	article, err := h.service.Update(id, &req)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			common.WriteJSON(w, &common.Response{
				Msg:    err.Error(),
				Status: http.StatusNotFound,
			})
			return
		}
		common.WriteError(w, err)
		return
	}

	common.WriteJSON(w, &common.Response{
		Msg:  "success update article",
		Data: article,
	})
}

func (h *ArticleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		common.WriteJSON(w, &common.Response{
			Msg:    "invalid article id",
			Status: http.StatusBadRequest,
		})
		return
	}

	if err := h.service.Delete(id); err != nil {
		if strings.Contains(err.Error(), "not found") {
			common.WriteJSON(w, &common.Response{
				Msg:    err.Error(),
				Status: http.StatusNotFound,
			})
			return
		}
		common.WriteError(w, err)
		return
	}

	common.WriteJSON(w, &common.Response{
		Msg: "success delete article",
	})
}
