package filesystem

import (
	"os"
	"os/signal"
	"syscall"
)

type loop struct {
	Pre  chan struct{}
	Post chan struct{}
	Quit chan struct{}
	Comm chan string
	Sys  chan os.Signal
}

func newLoop() *loop {
	return &loop{
		make(chan struct{}, 0),
		make(chan struct{}, 0),
		make(chan struct{}, 0),
		make(chan string, 0),
		make(chan os.Signal, 0),
	}
}

func looping(f *FS) {
	go func() {
		f.Handle(f.Conn(), f.Pre, f.Post, f.Quit)
	}()

	signal.Notify(
		f.Sys,
		syscall.SIGINT,
		syscall.SIGKILL,
		//syscall.SIGHUP,
		syscall.SIGTERM,
		//syscall.SIGCHLD,
		//syscall.SIGPIPE,
	)

	go func() {
		for {
			select {
			case <-f.Pre:
				<-f.Post
			case sig := <-f.Sys:
				f.SignalHandler(sig)
			}
		}
	}()
}
