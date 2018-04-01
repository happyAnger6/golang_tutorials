package container

import "../attributes"

type Container struct {
	config attributes.Config
	description attributes.Description
	uses []attributes.Uses
	containers []Container
}
