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
	return f.items, next, err
}

type fakeFetcher struct {
	channel string
	items   []Item
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

type sub struct {
	fetcher Fetcher
	updates chan Item
	closing chan chan error
}

func (s *sub) Updates() <-chan Item {

	return s.updates
}

func (s *sub) Close() error {
	errc := make(chan error)
	s.closing <- errc
	return <-errc
}

type fetchResult struct {
	fetched []Item
	next    time.Time
	err     error
}

func (s *sub) Loop() {

	var pending []Item
	var next time.Time
	var err error

	var fetchDone chan fetchResult

	for {
		var fetchDelay time.Duration

		if now := time.Now(); next.After(now) {
			fetchDelay = next.Sub(now)
		}

		startFetch := time.After(fetchDelay)

		select {
		case <-startFetch:
			fetchDone = make(chan fetchResult, 1)
			fetched, next, err := s.fetcher.Fetch()
			fetchDone <- fetchResult{fetched: fetched, next: next, err: err}

		case result := <-fetchDone:

			fetchDone = nil
			fetched := result.fetched
			next, err = result.next, result.err

			if err != nil {
				next = time.Now().Add(10 * time.Second)
				break
			}

			for _, item := range fetched {
				pending = append(pending, item)
			}
		case errc := <-s.closing:
			errc <- err
			close(s.updates)
			return
		}

	}
}

func Subscribe(fetch Fetcher) Subscription {
	s := &sub{
		fetcher: fetch,
		updates: make(chan Item),
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

//func Merge(subs ...Subscription) Subscription {
//
//	for _, sub := range subs {
//		go func(s Subscription) {
//			for {
//				var it Item
//				select {
//				case it = <-s.Updates():
//				}
//
//				//select {
//				//  case
//				//}
//			}
//		}(sub)
//	}
//
//	return nil
//}

func main() {

	subscription := Subscribe(Fetch("blog.golang.org"))

	time.AfterFunc(3*time.Second, func() {
		println("closed: ", subscription.Close())
	})

	for it := range subscription.Updates() {
		println(it.Channel, it.Title)
	}

}
