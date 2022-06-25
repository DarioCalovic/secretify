package ticker

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DarioCalovic/secretify/pkg/api/secret"
	utilconfig "github.com/DarioCalovic/secretify/pkg/util/config"
)

type Ticker struct {
	ticker        *time.Ticker
	done          chan bool
	sigs          chan os.Signal
	cfg           *utilconfig.Configuration
	secretService secret.Service
}

func New(cfg *utilconfig.Configuration, secretService secret.Service) *Ticker {
	ticker := time.NewTicker(1 * time.Minute)
	done := make(chan bool)
	sigs := make(chan os.Signal, 1)
	t := &Ticker{
		ticker, done, sigs, cfg, secretService,
	}
	signal.Notify(t.sigs, syscall.SIGINT, syscall.SIGTERM)
	return t
}

func (t *Ticker) RunTask() {

	go func() {
		for {
			select {
			case <-t.done:
				return
			case <-t.ticker.C:
				err := t.secretService.DeleteExpired()
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}()

	<-t.sigs
	t.ticker.Stop()
	t.done <- true
}
