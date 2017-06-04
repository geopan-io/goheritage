package main

import "log"
import "os"

func main() {
	a := App{}
	a.Initialize(os.Getenv("APP_DB_URL"))
	a.Run(":8000")

}
