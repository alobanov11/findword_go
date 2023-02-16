package app

import (
	"bufio"
	"findword/internal/app/parser"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type App struct {
	word string
	tops int
}

func New(w string, t int) (*App, error) {
	a := &App{
		word: w,
		tops: t,
	}
	return a, nil
}

func (a *App) Run() {
	start := time.Now()
	defer fmt.Printf("Elapsed: %s\n", time.Since(start))

	fmt.Printf("Word '%s' | Tops %d\n", a.word, a.tops)

	wg := &sync.WaitGroup{}
	sem := make(chan struct{}, a.tops)
	total := int64(0)
	reader := bufio.NewReader(os.Stdin)

	for {
		text, err := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if err == io.EOF {
			break
		}

		wg.Add(1)
		go a.find(text, &total, wg, sem)
	}

	wg.Wait()
	fmt.Printf("%d total\n", total)
}

func (a *App) find(text string, total *int64, wg *sync.WaitGroup, sem chan struct{}) {
	defer func() {
		<-sem
		wg.Done()
	}()

	sem <- struct{}{}

	parser := parser.New(a.word, text)
	count, err := parser.Parse()

	if err != nil {
		fmt.Printf("error parse %s with %s\n", text, err)
		return
	}

	atomic.AddInt64(total, int64(count))
	fmt.Printf("%d for %s\n", count, text)
}
