package dimgr

import (
    "errors"
)

type Provider struct {
    Create   func() interface{}
    Name     string
    Instance interface{}
}

type Module struct {
    providers map[string]*Provider
}

func NewModule() *Module {
    return &Module{
        providers: make(map[string]*Provider),
    }
}

func (module *Module) Add(providers ...Provider) {
    for _, provider := range providers {
        module.providers[provider.Name] = &provider
    }
}

func (module Module) Get(name string) (interface{}, error) {
    provider := module.providers[name]
    if provider == nil {
        return nil, errors.New("No existing provider: " + name)
    }

    // Creates the instance the first time
    if provider.Instance == nil {
        provider.Instance = provider.Create()
    }
    return provider.Instance, nil
}

