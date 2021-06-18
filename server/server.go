package server

import (
	"fmt"
	"os"
)

func Init() {
	PORT := os.Getenv("PORT")
	fmt.Println("Start router ::" + PORT)

	r := NewRouter()
	r.Run(":" + PORT)
}
