package main

import (
  "log"
)

func init(){
  log.SetPrefix("INFO:")
  log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
  log.Println("############################")
  log.Println("Initializing System")
  log.Println("############################")
}


func main(){
  log.Println("Finally Started") 

}

