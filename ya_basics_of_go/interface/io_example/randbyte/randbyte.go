package randbyte

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"time"
)

type generator struct {
	rnd rand.Source // Генератор случайных чисел. Вообще rand.Rand уже реализует интерфейс io.Reader, но для примера мы реализуем его самостоятельно.
}

func New(seed int64) io.Reader {
	return &generator{
		rnd: rand.NewSource(seed),
	}
}

func (g *generator) Read(bytes []byte) (n int, err error) {
	resultBytes := make([][]byte, 8)

	for i := range bytes {
		resultBytes[i] = make([]byte, 8)
		randInt := g.rnd.Int63()
		binary.LittleEndian.PutUint64(resultBytes[i], uint64(randInt))
	}

	return len(resultBytes), nil
}

func Execute() {
	generator := New(time.Now().UnixNano())

	buf := make([]byte, 8)

	for range 5 {
		n, _ := generator.Read(buf)
		fmt.Println(n)
	}
}
