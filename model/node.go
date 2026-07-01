package model

// Node is an asset-tree node.
type Node struct {
	ID            string `json:"id"`
	Key           string `json:"key"`
	Value         string `json:"value"`
	FullValue     string `json:"full_value"`
	OrgID         string `json:"org_id"`
	AssetsAmount  int    `json:"assets_amount"`
	Parent        string `json:"parent"`
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

// NodeTreeItem is a node in the children tree response.
type NodeTreeItem struct {
	ID        string `json:"id"`
	Key       string `json:"key"`
	Value     string `json:"value"`
	OrgID     string `json:"org_id"`
	Name      string `json:"name"`
	FullValue string `json:"full_value"`
	OrgName   string `json:"org_name"`
}

// NodeTreeItemPage is the paginated list envelope for NodeTreeItems.
type NodeTreeItemPage = Page[NodeTreeItem]
