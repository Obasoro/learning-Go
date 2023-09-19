package main

import (
    "context"
    "fmt"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
)

func main() {
    userService := UserService{
        db:   db{},
        amqp: amqp{},
    }

    r := mux.NewRouter()
    r.HandleFunc("/user", func(rw http.ResponseWriter, req *http.Request) {
        name := "some name" // let's imagine we got it from the request
        if err := userService.RegisterUser(req.Context(), name); err != nil {
            rw.WriteHeader(http.StatusInternalServerError)
            return
        }

        rw.WriteHeader(http.StatusOK)
    }).Methods(http.MethodPost)

    srv := http.Server{}
    srv.Addr = ":8080"
    srv.Handler = r

    sigs := make(chan os.Signal, 1)
    done := make(chan struct{})
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        <-sigs
        log.Println("got interruption signal")
        if err := srv.Shutdown(context.TODO()); err != nil {
            log.Printf("server shutdown returned an err: %v\n", err)
        }
        close(done)
    }()

    if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        log.Fatalf("listen and serve returned err: %v", err)
    }
    <-done
    log.Println("final")
}

type UserService struct {
    amqp amqp
    db   db
}

func (s *UserService) RegisterUser(ctx context.Context, name string) error {
    log.Println("start user registration")

    userID, err := s.db.InsertUser(ctx, name)
    if err != nil {
        return fmt.Errorf("db insertion failed: %v", err)
    }

    go s.amqp.PublishUserInserted(ctx, userID)

    return nil
}

type db struct{}

func (d db) InsertUser(ctx context.Context, name string) (int, error) {
    log.Println("user insert")
    return 1, nil
}

type amqp struct{}

func (a amqp) PublishUserInserted(ctx context.Context, id int) {
    log.Println("message publish")
}