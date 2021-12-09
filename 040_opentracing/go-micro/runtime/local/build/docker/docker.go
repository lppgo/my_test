// Package docker builds docker images
package docker

import (
	"archive/tar"
	"bytes"
	"io"
	"os"
	"path/filepath"

	docker "github.com/fsouza/go-dockerclient"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/runtime/local/build"
)

type Builder struct {
	Options build.Options
	Client  *docker.Client
}

func (d *Builder) Build(s *build.Source) (*build.Package, error) {
	image := filepath.Join(s.Repository.Path, s.Repository.Name)

	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	dockerFile := "Dockerfile"

	// open docker file
	f, err := os.Open(filepath.Join(s.Repository.Path, s.Repository.Name, dockerFile))
	if err != nil {
		return nil, err
	}
	// read docker file
	by, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	tarHeader := &tar.Header{
		Name: dockerFile,
		Size: int64(len(by)),
	}
	err = tw.WriteHeader(tarHeader)
	if err != nil {
		return nil, err
	}
	_, err = tw.Write(by)
	if err != nil {
		return nil, err
	}
	tr := bytes.NewReader(buf.Bytes())

	err = d.Client.BuildImage(docker.BuildImageOptions{
		Name:           image,
		Dockerfile:     dockerFile,
		InputStream:    tr,
		OutputStream:   io.Discard,
		RmTmpContainer: true,
		SuppressOutput: true,
	})
	if err != nil {
		return nil, err
	}
	return &build.Package{
		Name:   image,
		Path:   image,
		Type:   "docker",
		Source: s,
	}, nil
}

func (d *Builder) Clean(b *build.Package) error {
	image := filepath.Join(b.Path, b.Name)
	return d.Client.RemoveImage(image)
}

func NewBuilder(opts ...build.Option) build.Builder {
	options := build.Options{}
	for _, o := range opts {
		o(&options)
	}
	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		logger.Fatal(err)
	}
	return &Builder{
		Options: options,
		Client:  client,
	}
}
