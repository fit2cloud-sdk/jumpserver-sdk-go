package model

// Node is an asset-tree node.
type Node struct {
	ID           string `json:"id"`
	Key          string `json:"key"`
	Value        string `json:"value"`
	FullValue    string `json:"full_value"`
	OrgID        string `json:"org_id"`
	AssetsAmount int    `json:"assets_amount"`
	Parent       string `json:"parent"`
}

// NodeRequest is the create/update payload.
type NodeRequest struct {
	ID     string `json:"id,omitempty"`
	Value  string `json:"value"`
	Parent string `json:"parent,omitempty"`
}

// NodePage is the paginated list envelope for Nodes.
type NodePage = Page[Node]

// NodeChildRequest is the payload for creating a child node.
type NodeChildRequest struct {
	Value string `json:"value"`
}

// NodeTreeMeta contains the metadata for a node tree item.
type NodeTreeMeta struct {
	Data NodeTreeMetaData `json:"data"`
	Type string           `json:"type"`
}

// NodeTreeMetaData contains the actual node data inside a tree item.
type NodeTreeMetaData struct {
	ID    string `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

// NodeTreeItem is a node in the children tree response (zTree format).
type NodeTreeItem struct {
	ID       string       `json:"id"`
	Name     string       `json:"name"`
	PId      string       `json:"pId"`
	IsParent bool         `json:"isParent"`
	Open     bool         `json:"open"`
	Title    string       `json:"title"`
	Meta     NodeTreeMeta `json:"meta"`
}
