package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server : %v", err)
	}
}

func run(ctx context.Context) error {
	s := &http.Server{
		Addr: ":18080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello %s", r.URL.Path[1:])
		}),
	}
	//sync.WaitGroupの上位互換、ついでにラップしたctxを返す
	eg, ctx := errgroup.WithContext(ctx)

	//別のGo routineでHTTPサーバを建てる
	eg.Go(func() error {
		if err := s.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
			return err
		}
		return nil
	})
	//チャンネルからの通知を待機する
	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown %+v", err)
	}
	return eg.Wait()
}
