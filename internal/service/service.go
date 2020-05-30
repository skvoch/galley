package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/skvoch/galley/internal/model"
	"github.com/skvoch/galley/internal/repository"
	"net/http"
)

func New() (*Service, error) {
	logrus.Info("Service: new")

	pg, err := repository.New()

	if err != nil {
		return nil, err
	}

	logrus.Info("Service: created")
	return &Service{
		pg:     pg,
		router: gin.Default(),
	}, nil
}

type Service struct {
	pg     *repository.Repository
	router *gin.Engine
}

func (s *Service) Setup() {
	logrus.Info("Service: setup")

	s.router.Handle(http.MethodPost, "/scream", s.HandlerScream)

	s.router.Handle(http.MethodGet, "/users", s.HandlerUsers)
	s.router.Handle(http.MethodPost, "/users/register", s.HandleRegister)

	s.router.Handle(http.MethodGet, "/board", s.HandlerBoard)
	s.router.Handle(http.MethodPost, "/task/create", s.HandleTaskCreate)
	s.router.Handle(http.MethodPost, "/task/change", s.HandleTaskChange)

	s.router.Handle(http.MethodGet, "/clicks", s.HandleClicks)
	s.router.Handle(http.MethodPost, "/clicks/add", s.HandleAddClicks)
}

func (s *Service) Run() {
	logrus.Info("Service: run")

	logrus.Info(s.router.Handlers)
	s.router.Run()
}

func (s *Service) HandlerScream(ctx *gin.Context) {
	ctx.String(http.StatusOK, "scream - not implemented yet")
}

func (s *Service) HandlerUsers(ctx *gin.Context) {
	users, err := s.pg.Users().ReadAll()

	if err != nil {
		s.bindError(ctx, err)
		return
	} else {
		ctx.JSON(http.StatusOK, users)
	}
}

func (s *Service) HandleRegister(ctx *gin.Context) {
	user := model.User{}

	if err := ctx.BindJSON(&user); err != nil {
		s.bindError(ctx, err)
		return
	}

	exist, err := s.pg.Users().Exist(&user)

	if err != nil {
		s.bindError(ctx, err)

		return
	}

	if !exist {
		if err := s.pg.Users().Create(&user); err != nil {
			s.bindError(ctx, err)
			return
		} else {
			ctx.String(http.StatusOK, "User created")
		}
	}
}

func (s *Service) HandlerBoard(ctx *gin.Context) {
	tasks, err := s.pg.Board().ReadAll()

	if err != nil {
		s.bindError(ctx, err)
		return
	} else {
		ctx.JSON(http.StatusOK, tasks)
	}
}

func (s *Service) HandleTaskChange(ctx *gin.Context) {
	task := model.Task{}

	if err := ctx.BindJSON(&task); err != nil {
		s.bindError(ctx, err)
		return
	}

	if err := s.pg.Board().Change(&task); err != nil {
		s.bindError(ctx, err)
		return
	} else {
		ctx.String(http.StatusOK, "Changed")
	}
}

func (s *Service) HandleTaskCreate(ctx *gin.Context) {
	task := model.Task{}

	if err := ctx.BindJSON(&task); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	if err := s.pg.Board().Create(&task); err != nil {
		s.bindError(ctx, err)
		return
	} else {
		ctx.String(http.StatusOK, "Created")
	}
}

func (s *Service) HandleClicks(ctx *gin.Context) {
	ctx.String(http.StatusOK, "change task - not implemented yet")
}

func (s *Service) HandleAddClicks(ctx *gin.Context) {
	ctx.String(http.StatusOK, "change task - not implemented yet")
}

func (s *Service) bindError(ctx *gin.Context, err error) {
	ctx.String(http.StatusInternalServerError, err.Error())
	ctx.Abort()
}
