package leaf

import "../attributes"

type Leaf struct {
	Name string
	LeafType LeafType
	Description attributes.Description
	Reference attributes.Reference
}