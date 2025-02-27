package testfiles

import _ "embed"

// for health unit tests

//go:embed healthyCephStatus1.json
var HealthyCephStatus1 []byte

//go:embed rebalanceCephStatus1.json
var RebalanceCephStatus1 []byte

//go:embed rebalanceCephStatus2.json
var RebalanceCephStatus2 []byte

//go:embed rebalanceCephStatusFull.json
var RebalanceCephStatusFull []byte

//go:embed rebalanceCephStatusMultinode.json
var RebalanceCephStatusMultinode []byte

//go:embed tooManyPgsPerOsd.json
var TooManyPGSPerOSD []byte

//go:embed autoscalerCephStatus.json
var AutoscalerInProgressCephStatus []byte

//go:embed noreplicasCephStatus.json
var NoReplicasCephStatus []byte

//go:embed recentCrashCephStatus.json
var RecentCrashCephStatus []byte

// lists of pods to use in migrate unit tests

//go:embed "6 blockdevice pods.json"
var SixBlockDevicePods []byte

//go:embed "hostpathpods.json"
var HostpathPods []byte

// lists of deployments to be used in toolbox unit tests

//go:embed "rook-6osd-deployments.json"
var Rook6OSDDeployments []byte

//go:embed "rook-hostpath-deployments.json"
var RookHostpathDeployments []byte

// lists of nodes to be used in unit tests

//go:embed "upgradedNode.json"
var UpgradedNode []byte
