package main

import (
	"flag"
	"os"

	"github.com/atlassian/smith/pkg/resources"
	"github.com/atlassian/voyager/cmd"
	comp_crd "github.com/atlassian/voyager/pkg/composition/crd"
	form_crd "github.com/atlassian/voyager/pkg/formation/crd"
	ops_crd "github.com/atlassian/voyager/pkg/ops/crd"
	orch_crd "github.com/atlassian/voyager/pkg/orchestration/crd"
	"github.com/pkg/errors"
	apiext_v1b1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
)

func main() {
	cmd.ExitOnError(innerMain())
}

func innerMain() error {
	outputFormat := flag.String("output-format", "yaml", "Print a CRD and exit (specify format: json or yaml)")
	resource := flag.String("resource", "state", "Select which CRD to print (state, route, sd, ld)")
	flag.Parse()

	var crd *apiext_v1b1.CustomResourceDefinition
	switch *resource {
	case "state":
		crd = orch_crd.StateCrd()
	case "route":
		crd = ops_crd.RouteCrd()
	case "sd":
		crd = comp_crd.ServiceDescriptorCrd()
	case "ld":
		crd = form_crd.LocationDescriptorCrd()
	default:
		return errors.Errorf("unsupported CRD %q", *resource)
	}

	return resources.PrintCleanedObject(os.Stdout, *outputFormat, crd)
}
