package main

import (
	"fmt"
	"math/rand"
	"time"
)

// basic interfaces
type Item struct {
	Title, Channel, Guid string
}

type Subscription interface {
	Updates() <-chan Item
	Close() error
}

// Fetcher
type Fetcher interface {
	Fetch() (items []Item, next time.Time, err error)
}

// Fetch returns a Fetcher for Items from domain.
func Fetch(domain string) Fetcher {
	return fakeFetch(domain)
}

func fakeFetch(domain string) Fetcher {
	return &fakeFetcher{channel: domain}
}

func (f *fakeFetcher) Fetch() (items []Item, next time.Time, err error) {
	now := time.Now()
	next = now.Add(time.Duration(rand.Intn(5)) * 500 * time.Millisecond)
	item := Item{
		Channel: f.channel,
		Title:   fmt.Sprintf("Item %d", len(f.items)),
	}
	item.Guid = item.Channel + "/" + item.Title
	f.items = append(f.items, item)
	return
}

type fakeFetcher struct {
	channel string
	items   []Item
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

type sub struct {
	fetcher    Fetcher
	updatesChn chan Item
	closed     bool
	err        error
}

func (s *sub) Updates() <-chan Item {

	return s.updatesChn
}

func (s *sub) Close() error {
	s.closed = true
	//close(s.updatesChn)

	return s.err
}

func (s *sub) Loop() {
	for {
		if s.closed {
			close(s.updatesChn)
			return
		}

		items, next, err := s.fetcher.Fetch()
		if err != nil {
			s.err = err
			time.Sleep(10 * time.Second)
			continue
		}

		for _, item := range items {
			s.updatesChn <- item
		}

		if now := time.Now(); next.After(now) {
			time.Sleep(next.Sub(now))
		}

	}
}

func Subscribe(fetch Fetcher) Subscription {
	s := &sub{
		fetcher:    fetch,
		updatesChn: make(chan Item),
	}

	go s.Loop()
	return s
}

type merge struct {
	subs    []Subscription
	updates chan Item
	quit    chan struct{}
	errs    chan error
}

func Merge(subs ...Subscription) Subscription {

	for _, sub := range subs {
		go func(s Subscription) {
			for {
				var it Item
				select {
				case it = <-s.Updates():
				}

				//select {
				//  case
				//}
			}
		}(sub)
	}

	return nil
}

func main() {

	subscription := Subscribe(Fetch("blog.golang.org"))

	time.AfterFunc(3*time.Second, func() {
		println("closed: ", subscription.Close())
	})

	for it := range subscription.Updates() {
		println(it.Channel, it.Title)
	}

}
