package core

import (
	"github.com/innsanes/serv"
	"go.uber.org/zap"
	"time"
)

type View struct {
	*serv.Service
	ticker   *time.Ticker
	logger   *Logger
	ch       chan ViewSaveTask
	stop     chan struct{}
	tasks    map[ViewSaveId]int64
	saveFunc map[int]func(uint, int64)
}

type ViewSaveId struct {
	Type int
	Id   uint
}

type ViewSaveTask struct {
	ViewSaveId
	Count int64
}

func NewView() *View {
	return &View{
		logger: NewLog(),
		ch:     make(chan ViewSaveTask, 1000),
		stop:   make(chan struct{}, 1),
		tasks:  make(map[ViewSaveId]int64, 100),
	}
}

func (s *View) RegisterSaveFunc(kind int, saveFunc func(uint, int64)) {
	s.saveFunc[kind] = saveFunc
}

func (s *View) AddTask(task ViewSaveTask) {
	select {
	case _, hasClosed := <-s.stop:
		if hasClosed {
			return
		}
	default:
	}
	s.ch <- task
}

func (s *View) ForceStop() {
	s.stop <- struct{}{}
}

func (s *View) save() {
	for id, count := range s.tasks {
		f, ok := s.saveFunc[id.Type]
		if !ok {
			s.logger.Warn("no registered saveFunc", zap.Int("type", id.Type))
			continue
		}
		f(id.Id, count)
	}
}

func (s *View) Serve() (err error) {
	go func() {
		s.ticker = time.NewTicker(5 * time.Second)
		for {
			select {
			case t := <-s.ch:
				if count, ok := s.tasks[t.ViewSaveId]; ok {
					s.tasks[t.ViewSaveId] = max(count, t.Count)
				} else {
					s.tasks[t.ViewSaveId] = t.Count
				}
			case <-s.ticker.C:
				s.save()
			case <-s.stop:
				close(s.stop)
				close(s.ch)
				return
			}
		}
	}()
	return
}

func (s *View) AfterStop() {
	s.save()
}
