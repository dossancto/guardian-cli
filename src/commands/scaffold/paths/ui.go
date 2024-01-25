package paths

import (
	"fmt"

	"github.com/lu-css/guardian-cli/src/commands/scaffold/structs"
)

func GetUILayer(config structs.FeatureConfiguration, folder string) string {
	pathTemplate := "%s.UI/src/%s/"

	return fmt.Sprintf(pathTemplate, config.SlnName, folder)

}

