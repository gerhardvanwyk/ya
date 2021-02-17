package gocd

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type GoCD interface {

}

type gocd struct {
	Config BuildConfig
}

func New(file string) (GoCD, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil{
		return nil, err
	}

	//1. Load variables - build.yml
	config := BuildConfig{}
	err = yaml.Unmarshal(data, &BuildConfig{})

	return &gocd{
		Config: config,
	}, nil
}

//3. Set the sonar profile

//4. if svm tagging is enabled and integration test are disabled
// then mvn package

//5. if svn taggins is enabled then mvn verify

//6. default mvn deploy

//7. create a directory to for artifacts??

//8. run mvn

//9. Do sonar checks

//10. Do a NexusIQ registration. --> Mandatory

//11. Push to artifactory
