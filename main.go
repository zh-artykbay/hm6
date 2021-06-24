package main

import (
	"errors"
	"fmt"
	"sync"
)

func Execute(tasks []func() error, E int) error {

	wg := sync.WaitGroup{}
	cnt := 0

	for _ , task := range tasks {
		wg.Add(1)
		v := task()

		go func() {
			defer wg.Done()
			if v != nil  {
				cnt++
			}
		}()

		if cnt >= E {
			break
		}
	}
	wg.Wait()

	return nil
}

func main()  {

	t := []func() error{func() error{
		return errors.New("dsfds")
	}, func() error{
		return errors.New("a")
	},func() error{
		fmt.Println("sfaafa")
		return errors.New("b")
	},func() error{
		return errors.New("c")
	},func() error{
		return errors.New("d")
	},func() error{
		fmt.Println("okoooo")
		return errors.New("e")
	} }

	Execute(t, 5)

}
