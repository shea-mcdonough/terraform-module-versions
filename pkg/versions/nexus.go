package versions

import (
	"fmt"
	// "strings"
	"github.com/sonatype-nexus-community/gonexus/rm"
	"github.com/Masterminds/semver/v3"
	// "path/filepath"
	// "os"
	"sort"
)

func Nexus(modulePath string, moduleName string, moduleVersion string) ([]*semver.Version, error) {
	// nexusRepositoryManager, err := nexusrm.New("https://artifacts.dox.support", os.Getenv("NEXUS_USER"), os.Getenv("NEXUS_PASS"))
	nexusRepositoryManager, err := nexusrm.New("http://localhost:8081", "admin", "bd5406de-2f74-45c1-8d6a-fef78af4abd1")
	if err != nil {
		return nil, fmt.Errorf("nexus repository manager: %w", err)
	}

	// repo := strings.TrimSuffix(strings.Replace(modulePath, "/repository", "", -1), "/"+moduleName+"."+moduleVersion+filepath.Ext(strings.Replace(modulePath, "/repository", "", -1)))
	repo := "maven-releases"
	fmt.Printf("\nRepo: %v\n", repo)
	// assets, err := nexusrm.GetAssets(nexusRepositoryManager, repo)
	// if err != nil {
	// 	fmt.Printf("Error: %v", err)
	// 	return nil, fmt.Errorf("GetAssets Error: %w", err)
	// }

	query := nexusrm.NewSearchQueryBuilder().Repository(repo).Name("artifactID")
	components, err := nexusrm.SearchComponents(nexusRepositoryManager, query)
	var versionsList []string

	out := make([]*semver.Version, 0, len(components))

	for _, component := range components {
		versionsList = append(versionsList, component.Version)
		version, err := semver.NewVersion(component.Version)
		fmt.Printf("check version: %v\n", version)
		if err != nil {
			continue
		}
		out = append(out, version)
	}
	fmt.Printf("\nVersionsList: %v\n", versionsList)
	fmt.Printf("\nOut: %v\n", out)

	sort.Sort(semver.Collection(out))

	return out, nil
}