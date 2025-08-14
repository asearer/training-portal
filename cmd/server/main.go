// File: cmd/server/main.go
// Entry point of the training portal backend application

package main

import (
    "log"
    "training-portal/internal/interface/http/router"
)

func main() {
    log.Println("Starting Training Portal server...")
    router.SetupAndRun()
}
