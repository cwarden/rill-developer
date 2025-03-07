package rillv1beta

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/go-yaml/yaml"
	runtimev1 "github.com/rilldata/rill/proto/gen/rill/runtime/v1"
	"github.com/rilldata/rill/runtime/drivers"
)

const Version = "rill-beta"

type Codec struct {
	Repo       drivers.RepoStore
	InstanceID string
}

func New(repo drivers.RepoStore, instanceID string) *Codec {
	return &Codec{Repo: repo, InstanceID: instanceID}
}

func (c *Codec) IsInit(ctx context.Context) bool {
	_, err := c.Repo.Get(ctx, c.InstanceID, "rill.yaml")
	return err == nil
}

func (c *Codec) InitEmpty(ctx context.Context, name, rillVersion string) error {
	err := c.Repo.Put(ctx, c.InstanceID, "rill.yaml", strings.NewReader(fmt.Sprintf("compiler: %s\nrill_version: %s\n\nname: %s\n", Version, rillVersion, name)))
	if err != nil {
		return err
	}

	err = c.Repo.Put(ctx, c.InstanceID, ".gitignore", strings.NewReader("*.db\n*.db.wal\ndata/\n"))
	if err != nil {
		return err
	}

	err = c.Repo.Put(ctx, c.InstanceID, "sources/.gitkeep", strings.NewReader(""))
	if err != nil {
		return err
	}

	err = c.Repo.Put(ctx, c.InstanceID, "models/.gitkeep", strings.NewReader(""))
	if err != nil {
		return err
	}

	err = c.Repo.Put(ctx, c.InstanceID, "dashboards/.gitkeep", strings.NewReader(""))
	if err != nil {
		return err
	}

	return nil
}

func (c *Codec) PutSource(ctx context.Context, repo drivers.RepoStore, instanceID string, source *runtimev1.Source, force bool) (string, error) {
	props := source.Properties.AsMap()

	out := Source{
		Type: source.Connector,
	}

	if val, ok := props["uri"].(string); ok {
		out.URI = val
	}

	if val, ok := props["path"].(string); ok {
		out.Path = val
	}

	if val, ok := props["aws.region"].(string); ok {
		out.Region = val
	}

	if val, ok := props["csv.delimiter"].(string); ok {
		out.CSVDelimiter = val
	}

	blob, err := yaml.Marshal(out)
	if err != nil {
		return "", err
	}

	p := path.Join("sources", source.Name+".yaml")

	// TODO: Use create and createOnly when they're added to repo.Put
	if _, err := os.Stat(path.Join(repo.DSN(), p)); err == nil {
		if !force {
			return "", os.ErrExist
		}
	}

	err = repo.Put(ctx, c.InstanceID, p, bytes.NewReader(blob))
	if err != nil {
		return "", err
	}

	return p, nil
}

func (c *Codec) DeleteSource(ctx context.Context, name string) (string, error) {
	p := path.Join("sources", name+".yaml")
	err := c.Repo.Delete(ctx, c.InstanceID, p)
	if err != nil {
		return "", err
	}
	return p, nil
}
