package config

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

// CloudlensAlias manages Cloudlens aliases.
var CloudlensAlias = filepath.Join(CloudlensHome(), "alias.yml")

// Alias tracks shortname to GVR mappings.
type Alias map[string]string

// ShortNames represents a collection of shortnames for aliases.
type ShortNames map[string][]string

// Aliases represents a collection of aliases.
type Aliases struct {
	Alias Alias `yaml:"alias"`
	mx    sync.RWMutex
}

// NewAliases return a new alias.
func NewAliases() *Aliases {
	return &Aliases{
		Alias: make(Alias, 50),
	}
}

// Keys returns all aliases keys.
func (a *Aliases) Keys() []string {
	a.mx.RLock()
	defer a.mx.RUnlock()

	ss := make([]string, 0, len(a.Alias))
	for k := range a.Alias {
		ss = append(ss, k)
	}
	return ss
}

// ShortNames return all shortnames.
func (a *Aliases) ShortNames() ShortNames {
	a.mx.RLock()
	defer a.mx.RUnlock()

	m := make(ShortNames, len(a.Alias))
	for alias, res := range a.Alias {
		if v, ok := m[res]; ok {
			m[res] = append(v, alias)
		} else {
			m[res] = []string{alias}
		}
	}

	return m
}

// Clear remove all aliases.
func (a *Aliases) Clear() {
	a.mx.Lock()
	defer a.mx.Unlock()

	for k := range a.Alias {
		delete(a.Alias, k)
	}
}

// Get retrieves an alias.
func (a *Aliases) Get(k string) (string, bool) {
	a.mx.RLock()
	defer a.mx.RUnlock()

	v, ok := a.Alias[k]
	return v, ok
}

// Define declares a new alias.
func (a *Aliases) Define(resource string, aliases ...string) {
	a.mx.Lock()
	defer a.mx.Unlock()

	for _, alias := range aliases {
		if _, ok := a.Alias[alias]; ok {
			continue
		}
		a.Alias[alias] = resource
	}
}

// Load Cloudlens aliases.
func (a *Aliases) Load() error {
	a.loadDefaultAliases()
	return a.LoadFileAliases(CloudlensAlias)
}

// LoadFileAliases loads alias from a given file.
func (a *Aliases) LoadFileAliases(path string) error {
	f, err := os.ReadFile(path)
	if err == nil {
		var aa Aliases
		if err := yaml.Unmarshal(f, &aa); err != nil {
			return err
		}

		a.mx.Lock()
		defer a.mx.Unlock()
		for k, v := range aa.Alias {
			a.Alias[k] = v
		}
	}

	return nil
}

func (a *Aliases) declare(key string, aliases ...string) {
	a.Alias[key] = key
	for _, alias := range aliases {
		a.Alias[alias] = key
	}
}

func (a *Aliases) loadDefaultAliases() {
	a.mx.Lock()
	defer a.mx.Unlock()

	a.declare("ec2", "Ec2", "EC2")
	a.declare("s3", "S3")
	a.declare("sg", "SG")
	a.declare("iam:u", "IAM:U")
	a.declare("ebs", "EBS")

	a.declare("help", "h", "?")
	a.declare("quit", "q", "q!", "Q")
	a.declare("aliases", "alias", "a")
}

// Save alias to disk.
func (a *Aliases) Save() error {
	log.Debug().Msg("[Config] Saving Aliases...")
	return a.SaveAliases(CloudlensAlias)
}

// SaveAliases saves aliases to a given file.
func (a *Aliases) SaveAliases(path string) error {
	EnsurePath(path, DefaultDirMod)
	cfg, err := yaml.Marshal(a)
	if err != nil {
		return err
	}
	return os.WriteFile(path, cfg, 0644)
}
