package clusters

import "strings"

type Cluster struct {
	Name     string
	Labels   map[string]bool
	Resouces SupportResources
}

var allClusters = make(map[string]*Cluster)

func init() {
	ParseConfig()
	for _, config := range Configs {
		labels := make(map[string]bool)
		for _, l := range config.Labels {
			labels[strings.ToLower(l)] = true
		}
		name := strings.ToLower(config.ClusterName)
		allClusters[name] = &Cluster{
			Name:     name,
			Labels:   labels,
			Resouces: InitResources(),
		}
	}
}

func GetClustersByLabel(labels []string) (matchedClusters []*Cluster) {
	for n, c := range allClusters {
		if matchClusterByLabel(c, labels) {
			matchedClusters = append(matchedClusters, allClusters[n])
		}
	}
	return matchedClusters
}

func matchClusterByLabel(cluster *Cluster, labels []string) bool {
	for _, l := range labels {
		if !cluster.Labels[strings.ToLower(l)] {
			return false
		}
	}
	return true
}

func GetClustersByName(names string) (cluster *Cluster, ok bool) {
	cluster, ok = allClusters[strings.ToLower(names)]
	return
}