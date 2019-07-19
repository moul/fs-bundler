package main // import "moul.io/fs-bundler"

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"

	zglob "github.com/mattn/go-zglob"
	"github.com/pkg/errors"
	cli "gopkg.in/urfave/cli.v2"
	yaml "gopkg.in/yaml.v3"
)

type Dump struct {
	Files []File `json:"files"`
}

type File struct {
	Path    string `json:"path"`
	Name    string `json:"name"`
	Content []byte `json:"content" yaml:"content,flow"`
	// Type string `json:"type"` // mime
}

func main() {
	app := &cli.App{
		Name: "fs-bundler",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "format", Aliases: []string{"f"}, Value: "json", Usage: `output format ("json", "yaml")`},
			&cli.BoolFlag{Name: "indent", Aliases: []string{"i"}, Usage: `use indented output (only for "json" format)`},
			// compress
		},
		Action: func(c *cli.Context) error {
			switch c.String("format") {
			case "json", "yaml":
			default:
				return fmt.Errorf("unsupported output format %q", c.String("format"))
			}
			targets := c.Args().Slice()
			if len(targets) == 0 {
				targets = append(targets, ".")
			}
			allMap := map[string]bool{}
			for _, target := range targets {
				matches, err := zglob.Glob(target)
				if err != nil {
					return errors.Wrapf(err, "failed to list files from %q", target)
				}
				for _, match := range matches {
					allMap[match] = true
				}
			}
			all := []string{}
			for fullPath := range allMap {
				all = append(all, fullPath)
			}
			sort.Strings(all)
			dump := Dump{}
			for _, fullPath := range all {
				content, err := ioutil.ReadFile(fullPath)
				if err != nil {
					log.Printf("failed to read content from %q: %v", fullPath, err)
					continue
				}
				dump.Files = append(dump.Files, File{
					Path:    fullPath,
					Name:    path.Base(fullPath),
					Content: content,
				})
			}
			switch c.String("format") {
			case "json":
				if c.Bool("indent") {
					out, err := json.MarshalIndent(dump, "", "  ")
					if err != nil {
						return err
					}
					fmt.Println(string(out))
				} else {
					out, err := json.Marshal(dump)
					if err != nil {
						return err
					}
					fmt.Println(string(out))
				}
			case "yaml":
				out, err := yaml.Marshal(dump)
				if err != nil {
					return err
				}
				fmt.Println(string(out))
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Printf("error: %v", err)
	}
}
