package app

import (
	"com.roxorgaming/ja/mvn"
	"sync"
)

func Run(command string, config mvn.pom_config) error {

	ci := mvn.New(config)

	wg := new(sync.WaitGroup)
	wg.Add(1)

	go ci.UpVersionAutomatic(wg)

	wg.Wait()

	return nil
}
