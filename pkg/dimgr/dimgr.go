package dimgr

type Provider struct {
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

func (m *Module) AddProvider(name string, provider *Provider) {
    m.providers[name] = provider
}

func (m Module) GetProvider(name string) *Provider {
    return m.providers[name]
}
