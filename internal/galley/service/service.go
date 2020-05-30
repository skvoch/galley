package service

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/skvoch/galley/internal/galley/model"
	"github.com/skvoch/galley/internal/galley/repository"
	"net/http"
	"strconv"
	"sync/atomic"
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
		pushMessages: []string{
			"Греби сильнее, качественнее, быстрее!",
			"Дедлайн не за горами, пиши быстрее, ты же не хочешь всех подвести..",
			"Не переживай, ты не раб, точно-точно. Мы же команда!",
			"Уже готово? Нет?",
			"А сейчас?",
			"Ты должен писать свой код быстрее! БЫСТРЕЕ!!",
			"Зачем тебе тесты? документация? не смеши, пиши быстрее",
		},
	}, nil
}

type Service struct {
	pg     *repository.Repository
	router *gin.Engine

	pushIndex    int64
	pushMessages []string
}

func (s *Service) Setup() {
	logrus.Info("Service: setup")
	s.router.Use(cors.Default())

	s.router.Handle(http.MethodPost, "/scream", s.HandlerScream)

	s.router.Handle(http.MethodGet, "/users", s.HandlerUsers)
	s.router.Handle(http.MethodPost, "/users/register", s.HandleRegister)

	s.router.Handle(http.MethodGet, "/board", s.HandlerBoard)
	s.router.Handle(http.MethodPost, "/task/create", s.HandleTaskCreate)
	s.router.Handle(http.MethodPost, "/task/change", s.HandleTaskChange)

	s.router.Handle(http.MethodGet, "/clicks/:count", s.HandleClicks)
	s.router.Handle(http.MethodPost, "/clicks/add", s.HandleAddClicks)

	s.router.Handle(http.MethodPost, "/push/send", s.handlePushSend)
	s.router.Handle(http.MethodGet, "/push/get", s.handlePushGet)

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

	ctx.String(http.StatusOK, "Registered")
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
	count := ctx.Param("count")

	count_int, err := strconv.Atoi(count)

	if err != nil {
		s.bindError(ctx, err)
		return
	}

	stats, err := s.pg.Clicks().ReadLast(count_int)

	if err != nil {
		s.bindError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, stats)
}

func (s *Service) HandleAddClicks(ctx *gin.Context) {
	stats := model.ClickStats{}

	if err := ctx.BindJSON(&stats); err != nil {
		s.bindError(ctx, err)
	}
	ctx.String(http.StatusOK, "Clicks added")

	s.pg.Clicks().Add(&stats)
	logrus.Info("Hash: ", stats.Hash, " count: ", stats.Count)
}

func (s *Service) bindError(ctx *gin.Context, err error) {
	ctx.String(http.StatusInternalServerError, err.Error())
	ctx.Abort()
}

func (s *Service) handlePushSend(ctx *gin.Context) {
	atomic.AddInt64(&s.pushIndex, 1)

	ctx.String(http.StatusOK, "OK")
}

func (s *Service) handlePushGet(ctx *gin.Context) {
	push := model.Push{
		Index: atomic.LoadInt64(&s.pushIndex),
	}

	l := int64(len(s.pushMessages))
	push.Message = s.pushMessages[push.Index%l]

	ctx.JSON(http.StatusOK, push)
}
