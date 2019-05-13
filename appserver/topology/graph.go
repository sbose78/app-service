package topology

type graph struct {
	nodes  []node
	edges  []edge
	groups []group
}

type node struct {
	id   string
	name string
}

type edge struct {
	source string
	target string
}

type group struct {
	id    string
	name  string
	nodes []node
}

type resource struct {
	name string
	kind string
}

type nodeData struct {
	nodeType  string
	id        string
	resources []resource
}

type nodeID string

type topology map[nodeID]nodeData

type serverMetadata struct {
	Commit string `json:"commit"`
}

type VisualizationResponse struct {
	graph          `json:"graph"`
	topology       `json:"topology"`
	serverMetadata `json:"serverData"`
}

func GetSampleTopology() VisualizationResponse {
	return VisualizationResponse{
		graph:    graph{},
		topology: topology{},
	}
}
