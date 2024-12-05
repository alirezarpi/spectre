package docker

import (
    "os/exec"
)

type ImageDetails struct {
    Name    string
    Layers  []string
    Metadata map[string]string
}

func InspectImage(image string) (*ImageDetails, error) {
    // Example: Mocked interaction with Docker CLI for inspection
    cmd := exec.Command("docker", "inspect", image)
    _, err := cmd.Output()
    if err != nil {
        return nil, err
    }

    // Parse output to ImageDetails
    details := &ImageDetails{
        Name: image,
        Layers: []string{"layer1", "layer2"},
        Metadata: map[string]string{"mockKey": "mockValue"},
    }
    return details, nil
}

