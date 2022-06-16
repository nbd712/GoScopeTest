package main

import (
	"log"
	"sync"
	"time"
)

var ()

func init() {

	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)

}

func main() {

	cache := NewWhatever()

	colors := []string{"RED", "GREEN", "YELLOW"}

	for i, color := range colors {

		log.Printf("Color: %s ColorByIndex: %s", color, colors[i])

		var mainWg sync.WaitGroup

		mainWg.Add(1)
		go func(wg *sync.WaitGroup, color string, indexColor string) {
			defer wg.Done()

			var fwg sync.WaitGroup

			cache.ForEach(color, func(key string, value bool) {

				log.Printf("Color: %s ColorByIndex: %s Key: %s, Value: %t", color, colors[i], key, value)

				fwg.Add(1)

				go func() {
					defer fwg.Done()

					time.Sleep(time.Second * 1)

					log.Printf("Color: %s ColorByIndex: %s Key: %s, Value: %t", color, colors[i], key, value)

				}()

			})

			fwg.Wait()

			log.Printf("Color: %s ColorByIndex: %s Complete", color, indexColor)

		}(&mainWg, color, colors[i])

		mainWg.Wait()
	}

	log.Println("We done!")

}

type Whatever struct {
	Data map[string]bool
	mute sync.Mutex
}

func NewWhatever() *Whatever {
	return &Whatever{
		Data: map[string]bool{"1": true, "2": true, "3": true, "4": true},
		mute: sync.Mutex{},
	}
}

func (w *Whatever) ForEach(color string, cb func(key string, value bool)) {
	w.mute.Lock()
	defer w.mute.Unlock()

	for key, value := range w.Data {

		cb(key, value)

	}

}
