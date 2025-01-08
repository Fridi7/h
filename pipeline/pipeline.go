package pipeline

import "fmt"

type (
	In    = <-chan interface{}
	Out   = In
	Bi    = chan interface{}
	Stage func(in In) (out Out)
)

// generator отправляет данные в канал
func generator(doneCh Bi, numbers []int) Bi {
	outputCh := make(Bi)

	go func() {
		defer close(outputCh)

		for _, num := range numbers {
			select {
			case <-doneCh:
				return
			case outputCh <- num:
			}
		}
	}()

	return outputCh
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	wrapStage := func(in In, done In, stage Stage) Out {
		out := make(chan interface{})
		go func() {
			defer close(out)
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered from panic in stage:", r)
				}
			}()
			stageOut := stage(in)
			for {
				select {
				case <-done:
					return
				case val, ok := <-stageOut:
					if !ok {
						return // Если выходной стейдж канал закрыт, выходим
					}
					select {
					case out <- val: // Отправляем данные дальше
					case <-done:
						return
					}
				}
			}
		}()
		return out
	}

	// Последовательно оборачиваем входной канал через все стейджи
	out := in
	for _, stage := range stages {
		out = wrapStage(out, done, stage)
	}
	return out
}
