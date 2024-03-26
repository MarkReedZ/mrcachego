package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
  //"bytes"
  //"encoding/binary"


	"github.com/panjf2000/gnet/v2"
)

//var ErrIncompletePacket = errors.New("incomplete packet")
var resp_miss = []byte{0, 0, 0, 0}
var resp_dbg  = []byte{4, 0, 0, 0, 65, 65, 65, 65}

type mrcacheServer struct {
	gnet.BuiltinEventEngine

	eng       gnet.Engine
	addr      string
  cac       map[string][]byte
}
//dupes := make(map[string][]string)

    //dupes[string(hash)] = []string{"a", "b"}
    //hash[len(hash)-1]++
    //dupes[string(hash)] = []string{"b", "c"}k
func (es *mrcacheServer) OnBoot(eng gnet.Engine) gnet.Action {
	es.eng = eng
	log.Printf("mrcache is listening on %s\n",  es.addr)
	return gnet.None
}

//func (s *mrcacheServer) OnClose(c gnet.Conn, err error) (action gnet.Action) {
  //action = gnet.Shutdown
  //return
//}


func (es *mrcacheServer) OnTraffic(c gnet.Conn) gnet.Action {
	buf, _ := c.Next(-1)
  i := 0
  end := len(buf)
  for {
	  //buf, _ := c.Peek(4)
    if end-i < 4 {
      //fmt.Println("Partial cmd")
      return gnet.None
    }
    //fmt.Println("len buf ", len(buf))
    //keylen := int(binary.LittleEndian.Uint16(buf[2:]))
    keylen := int(buf[i] << 8) | int(buf[i+1])
    //fmt.Println("Key len =", keylen)

	  ///buf, _ = c.Peek(keylen)
    ///_, _   = c.Discard(keylen+4)
    i += 4
    if end-i < keylen {
      //fmt.Println("Partial cmd")
      return gnet.None
    }

    if buf[1] == 1  { // GET
      //fmt.Println("DELME GET")
	    c.Write(resp_dbg)
    }
    if buf[1] == 2  { // SET
      fmt.Println("DELME SET")
    }
  }
	//c.Write(buf)
	return gnet.None
}

func main() {
	var port int

	// Example command: go run echo.go --port 9000 --multicore=true
	flag.IntVar(&port, "port", 7000, "--port 7000")
	flag.Parse()

	echo := &mrcacheServer{addr: fmt.Sprintf("tcp://:%d", port)}
	go func ()  {
		log.Println("gnet server exits:", gnet.Run(echo, echo.addr, gnet.WithLockOSThread(true)))
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	gnet.Stop(context.TODO(), echo.addr)
}
