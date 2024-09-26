package shutdown

import (
 "context"
 "log"
 "net/http"
 "os"
 "os/signal"
 "time"
)

func GracefulShutdown(server *http.Server) {
 stop := make(chan os.Signal, 1)
 signal.Notify(stop, os.Interrupt)

 <-stop
 log.Println("Shutting down gracefully...")

 ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
 defer cancel()

 if err := server.Shutdown(ctx); err != nil {
  log.Fatal("Server forced to shutdown:", err)
 }

 log.Println("Server exiting")
}

#