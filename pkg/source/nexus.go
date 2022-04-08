package source

import (
	"fmt"
	"net/url"
	"path"
	"strings"
	"path/filepath"
	"regexp"

	// getter "github.com/hashicorp/go-getter"
)

type Nexus struct {
	ModulePath string
	ModuleName  string
	ModuleVersion string
}

func parseNexusURL(s string) (*Nexus, error) {
	u, err := url.Parse(s)
	if err != nil {
		return nil, fmt.Errorf("parse nexus url: %w", err)
	}
	var out Nexus
	out.ModulePath = u.Path
// ADD MODULE REPO
	out.ModuleName = strings.SplitN(strings.TrimSuffix(path.Base(u.Path), filepath.Ext(path.Base(u.Path))), ".", 2)[0]
	reg, err := regexp.Compile("^[^0-9]+")
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	// out.ModuleVersion = strings.SplitN(strings.TrimSuffix(path.Base(u.Path), filepath.Ext(path.Base(u.Path))), ".", 2)[1]
	out.ModuleVersion = reg.ReplaceAllString(strings.TrimSuffix(path.Base(u.Path), filepath.Ext(path.Base(u.Path))), "")
	// if refValue := u.Query().Get("ref"); refValue != "" {
	// 	out.RefValue = &refValue
	// 	query := u.Query()
	// 	query.Del("ref")
	// 	u.RawQuery = query.Encode()
	// }
	// out.Remote = u.String()
	// if dir, subDir := getter.SourceDirSubdir(out.Remote); subDir != "" {
	// 	out.Remote = dir
	// 	out.RemotePath = &subDir
	// }
	return &out, nil
}