package install

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/kumahq/kuma/app/kumactl/pkg/install/data"
	"github.com/kumahq/kuma/app/kumactl/pkg/install/k8s"
	"github.com/kumahq/kuma/app/kumactl/pkg/install/k8s/kongingress"
)

type kicTemplateArgs struct {
	Namespace string
}

func newInstallKIC() *cobra.Command {
	args := struct {
		Namespace string
	}{
		Namespace: "kong-ingress",
	}
	cmd := &cobra.Command{
		Use:   "kong-ingress",
		Short: "Install Kong Ingress Controller in Kubernetes cluster",
		Long:  `Install Kong Ingress Controller in Kubernetes cluster in a 'kong-ingress' namespace`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			templateArgs := kicTemplateArgs{
				Namespace: args.Namespace,
			}

			templateFiles, err := data.ReadFiles(kongingress.Templates)
			if err != nil {
				return errors.Wrap(err, "Failed to read template files")
			}

			renderedFiles, err := renderFiles(templateFiles, templateArgs, simpleTemplateRenderer)
			if err != nil {
				return errors.Wrap(err, "Failed to render template files")
			}

			sortedResources, err := k8s.SortResourcesByKind(renderedFiles)
			if err != nil {
				return errors.Wrap(err, "Failed to sort resources by kind")
			}

			singleFile := data.JoinYAML(sortedResources)

			if _, err := cmd.OutOrStdout().Write(singleFile.Data); err != nil {
				return errors.Wrap(err, "Failed to output rendered resources")
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&args.Namespace, "namespace", args.Namespace, "namespace to install Kong Ingress to")
	return cmd
}
