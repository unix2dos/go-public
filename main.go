package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	TimeTemplate = "15:04:05.999999999"
)

type Service interface {
	GetName() string
	Serve(ctx context.Context)
	Shutdown() error
}

type BusinessService struct {
}

func (b *BusinessService) GetName() string {
	return "BusinessService"
}

func (b *BusinessService) Serve(ctx context.Context) {
	for {
		fmt.Printf("BusinessService serve run at %s\n", time.Now().Format(TimeTemplate))
		select {
		case <-ctx.Done():
			fmt.Printf("111111111111 %s\n", time.Now().Format(TimeTemplate))
			return
		default:
		}
		time.Sleep(time.Second)
	}
	return
}

func (b *BusinessService) Shutdown() error {
	fmt.Printf("BusinessService shutdown begin... at %s\n", time.Now().Format(TimeTemplate))
	defer func() {
		fmt.Printf("BusinessService shutdown end... at %s\n", time.Now().Format(TimeTemplate))
	}()
	return nil
}

type LogService struct {
	buffer []string
}

func (l *LogService) Serve(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("2222222222 %s\n", time.Now().Format(TimeTemplate))
			return
		default:
			// 0.5s append one log
			time.Sleep(500 * time.Millisecond)
			l.buffer = append(l.buffer, fmt.Sprintf("Time: %d", time.Now().Unix()))
		}
	}
}

func (b *LogService) GetName() string {
	return "LogService"
}

func (l *LogService) Shutdown() (err error) {
	fmt.Printf("LogService shutdown begin... at %s\n", time.Now().Format(TimeTemplate))
	defer fmt.Printf("LogService shutdown end... at %s\n", time.Now().Format(TimeTemplate))
	if len(l.buffer) == 0 {
		return
	}
	fmt.Printf("cache [%d] wait to send \n", len(l.buffer))
	for _, log := range l.buffer {
		fmt.Printf("send Log [%s]\n", log)
	}
	return
}

type ServiceGroup struct {
	ctx      context.Context
	cancel   func()
	services []Service //service list
}

func NewServiceGroup(ctx context.Context) *ServiceGroup {
	g := ServiceGroup{}
	g.ctx, g.cancel = context.WithCancel(ctx)
	return &g
}

func (s *ServiceGroup) Add(service Service) {
	s.services = append(s.services, service)
}

func (s *ServiceGroup) run(service Service) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
			fmt.Printf("receive panic msg: %s\n", err.Error())
		}
	}()
	//with cancel ctx to child context
	service.Serve(s.ctx)
	return
}

func (s *ServiceGroup) watchDog() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case signalData := <-signalChan:
			switch signalData {
			case syscall.SIGINT:
				fmt.Println("receive signal sigint")
			case syscall.SIGTERM:
				fmt.Println("receive signal sigerm")
			default:
				fmt.Println("receive singal unknown")
			}
			// do cancel notify all services cancel
			s.cancel()
			goto CLOSE
		case <-s.ctx.Done():
			goto CLOSE
		}
	}
CLOSE:
	for _, service := range s.services {
		if err := service.Shutdown(); err != nil {
			fmt.Printf("shutdown failed err: %s", err)
		}
	}
}

func (s *ServiceGroup) ServeAll() {
	var wg sync.WaitGroup
	for idx := range s.services {
		service := s.services[idx]
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := s.run(service); err != nil {
				fmt.Printf("receive service [%s] has error: 【%s】, do cancel\n", service.GetName(), err.Error())
				s.cancel()
			}
		}()
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		s.watchDog()
	}()
	wg.Wait()
}

func main() {
	rand.Seed(time.Now().Unix())
	ctx := context.Background()

	g := NewServiceGroup(ctx)
	g.Add(&LogService{})
	g.Add(&BusinessService{})
	g.ServeAll()
}
