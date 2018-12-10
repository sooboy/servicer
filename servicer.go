package servicer

import (
	"log"
	"time"
)

// Function 函数别名
type Function = func()

// Servicer 运行容器
type Servicer interface {
	Run()
	ID() string
}

// Service 基础函数
type Service struct {
	name string
	fn   Function
}

// Run implement Servicer
func (s *Service) Run() {
	s.fn()
}

// ID implement Servicer
func (s *Service) ID() string {
	return s.name
}

// LogtimeService 基础输出
type LogtimeService struct {
	service Servicer
}

// Run implement Servicer
func (l *LogtimeService) Run() {
	now := time.Now()
	l.service.Run()
	log.Println(l.service.ID(), "Runing Time is :", time.Now().Sub(now).String())
}

// ID implement Servicer
func (l *LogtimeService) ID() string {
	return l.service.ID() + "logTime "
}
