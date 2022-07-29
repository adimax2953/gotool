package nsqtool

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"testing"

	gotool "github.com/adimax2953/go-tool"
	logtool "github.com/adimax2953/log-tool"
	"github.com/nsqio/go-nsq"
)

func Test_SendtoNSQ(t *testing.T) {

	nsqConfig := &NsqConfig{
		Lookups: []string{"192.168.56.1:4161"},
		NSQDs:   []string{"192.168.56.1:4150"},
		NSQD:    "192.168.56.1:4150",
	}

	go InitializeConsumer(nsqConfig, "test", "", NsqunPackTest)
	InitializePublisher(nsqConfig)

	Send("test", []byte("test山豬"))

	// Graceful shutdown -
	ch := make(chan os.Signal, 1)
	signal.Notify(ch,
		// kill -SIGINT XXXX 或 Ctrl+c
		os.Interrupt,
		syscall.SIGINT, // register that too, it should be ok
		// kill -SIGTERM XXXX
		syscall.SIGTERM,
	)
	s := <-ch
	log.Printf("s...%v\n", s)
}

// NsqunPackTest -
func NsqunPackTest(m *nsq.Message) error {
	defer gotool.RecoverPanic()
	logtool.LogDebug("get", string(m.Body))
	return nil
}
