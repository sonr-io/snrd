// Code generated from Pkl module `models`. DO NOT EDIT.
package models

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type Models struct {
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Models
func LoadFromPath(ctx context.Context, path string) (ret *Models, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// LoadFromURL loads the pkl module at the given URL and evaluates it into a Models
func LoadFromURL(ctx context.Context, url string) (ret *Models, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.UriSource(url))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Models
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*Models, error) {
	var ret Models
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
