package blueprint

import (
	"embed"
	"log"
	"os"
)

type Blueprint struct {
	BasePath string
	Url      string
	Version  string
}

type BlueprintInterface interface {
	InitBlueprint() error
}

func NewBlueprint(
	assetsData embed.FS,
	basePath string,
	url string,
	version string) (*Blueprint, error) {

	blueprint := &Blueprint{
		Url:     url,
		Version: version,
	}

	return blueprint, blueprint.InitBlueprint()
}

func (b *Blueprint) InitBlueprint() (err error) {

	log.Printf("Initializing blueprint %s %s", b.Url, b.Version)

	//Create base folder if necessary
	_, err = os.Stat(b.BasePath)
	if os.IsNotExist(err) {
		os.MkdirAll(b.BasePath, os.ModePerm)
	}

	return
}
