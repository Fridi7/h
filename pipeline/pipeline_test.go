package pipeline

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPipeline(t *testing.T) {
	numbers := []int{1, 2, 30, 4, 5}

	// Пример стейджа, который умножает значения
	stage := func(in In) Out {
		out := make(Bi)
		go func() {
			defer close(out)
			for val := range in {
				// Эмулируем работу стейджа
				time.Sleep(time.Millisecond * 100)
				out <- val.(int) * 2
			}
		}()
		return out
	}

	done := make(Bi)
	in := generator(done, numbers)
	defer close(done)

	now := time.Now()

	result := ExecutePipeline(in, done, stage, stage, stage, stage)

	numbersOut := make([]int, 0, len(numbers))
	for res := range result {
		require.IsType(t, 1, res)
		v := res.(int)
		numbersOut = append(numbersOut, v)
	}

	require.Equal(t, []int{16, 32, 480, 64, 80}, numbersOut)
	require.True(t, time.Since(now) < time.Second)
	fmt.Println(time.Since(now)) //

}
