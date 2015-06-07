package main

import (
	//"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
  "sync"
  "time"
)

const SQL_CREATE_SHARE_TABLE = "create table if not exists 'shares' ('id' INTEGER PRIMARY KEY AUTOINCREMENT,'name' VARCHAR(128) NULL,path VARCHAR(4096) NULL)"

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func ShareWorker(StopChan chan bool, WakeEvent *sync.Cond) {
  var lcvrun bool = true
  for lcvrun {
    select {
    case x, ok := <-StopChan:
        if ok {
            fmt.Println("Value was read:", x)
            if x {
              break
            }
        } else {
            fmt.Println("Channel closed!")
            lcvrun = false
            break
        }
    default:
        fmt.Println("No value ready, Waiting on Event.")
        WakeEvent.L.Lock()
        WakeEvent.Wait()
        WakeEvent.L.Unlock()
    }
  }
  return 
}

func processWalk(path string, info os.FileInfo, err error) error {
  var clean_path string
  clean_path = filepath.Clean(path)
	fmt.Println(clean_path, info)
	return nil
}

func main() {
	fmt.Println("Taco")
  var lock sync.Mutex
  var stopchan chan bool = make(chan bool)
  wakeevent := sync.NewCond(&lock)
  go ShareWorker(stopchan,wakeevent)
  stopchan <- false
  fmt.Println("Sleeping in main")
  time.Sleep(3 * time.Second)
  close(stopchan)
  wakeevent.Broadcast()
  time.Sleep(3 * time.Second)
	//db, err := sql.Open("sqlite3", "./foo.db")
	//checkErr(err)

	//res, err := db.Exec(SQL_CREATE_SHARE_TABLE)
	//_ = res
	//checkErr(err)
	//filepath.Walk("/home/scott/", processWalk)

	//db.Close()
	fmt.Println("Taco2")
}
