package project

import (
	"fmt"
	"log"
	"os"

	"github.com/spycode-io/spycli/model"
)

func NewProject(
	base *model.Scaffold,
	plaform string,
	stack string,
	blueprint string,
	useRemoteState bool,
	remoteStateBucket string,
	remoteStateRegion string,
	environments []string,
	regions []string) (project *ProjectScaffold, err error) {

	project = &ProjectScaffold{
		Scaffold:       *base,
		Platform:       plaform,
		Stack:          stack,
		Blueprint:      blueprint,
		ProjectPath:    fmt.Sprintf("%s/%s", base.BasePath, base.SlugName),
		UseRemoteState: useRemoteState,
		RemoteBucket:   remoteStateBucket,
		RemoteRegion:   remoteStateRegion,
	}

	for _, env := range environments {
		project.Environments = append(project.Environments,
			model.Environment{
				Name: env,
				Path: fmt.Sprintf("%s/%s", project.ProjectPath, env),
			},
		)
	}

	for _, reg := range regions {
		project.Regions = append(project.Regions,
			model.Region{Region: reg},
		)
	}

	project.FileSet.WithSet(DefaultFileSet.Set)

	err = project.Init()

	return
}

func (p *ProjectScaffold) Init() (err error) {

	log.Printf("Initializing %s [%s] %s project on path: %s", p.Name, p.SlugName, p.Platform, p.BasePath)

	//Create base folder if necessary
	_, err = os.Stat(p.ProjectPath)
	if os.IsNotExist(err) {
		os.MkdirAll(p.ProjectPath, 0755)
		err = nil
	}

	//Write project level files
	err = p.FileSet.WriteObjToPath(p.Platform, "platform", p.PlatformPath, p)
	if nil != err {
		return
	}

	//Write project level files
	err = p.FileSet.WriteObjToPath(p.Platform, "project", p.ProjectPath, p)
	if nil != err {
		return
	}

	//Create folder structure
	for _, env := range p.Environments {
		for _, reg := range p.Regions {
			regPath := fmt.Sprintf("%s/%s", env.Path, reg.Region)
			err = p.FileSet.WriteObjToPath(p.Platform, "region", regPath, reg)
			if nil != err {
				return
			}
		}
		err = p.FileSet.WriteObjToPath(p.Platform, "environment", env.Path, env)
		if nil != err {
			return
		}
	}

	return
}
