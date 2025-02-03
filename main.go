package main

import (
	"fmt"

	containerimage "github.com/google/go-containerregistry/pkg/name"
)

func main() {
	// this code copied from https://github.com/aquasecurity/trivy-operator/blob/f6d43a84892a0ccc56119a8eedd39ec0cfd9539d/pkg/plugins/trivy/plugin.go#L216
	imageRef := "my-private-repo.company.com/my-app:some-tag@sha256:1420cefd4b20014b3361951c22593de6e9a2476bbbadd1759464eab5bfc0d34f"
	imageDigest := "2bc57c6bcb194869d18676e003dfed47b87d257fce49667557fb8eb1f324d5d6"

	ref, err := containerimage.ParseReference(imageRef)
	if err != nil {
		fmt.Println("failed parse:", err)
		return
	}
	registryServer := ref.Context().RegistryStr()
	artifactRepository := ref.Context().RepositoryStr()
	artifactTag := ""
	artifactDigest := ""

	switch t := ref.(type) {
	case containerimage.Tag:
		artifactTag = t.TagStr()
	case containerimage.Digest:
		artifactDigest = t.DigestStr()
	}
	if artifactDigest == "" {
		artifactDigest = imageDigest
	}
	fmt.Println("artifactTag =", artifactTag)
	fmt.Println("artifactDigest =", artifactDigest)
	fmt.Println("artifactRepository =", artifactRepository)
	fmt.Println("artifactRepository =", registryServer)
}
