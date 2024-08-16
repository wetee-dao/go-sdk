package weteeworker

import (
	types1 "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	types "github.com/wetee-dao/go-sdk/pallet/types"
)

// Worker cluster register
// 集群注册
func MakeClusterRegisterCall(name0 []byte, ip1 []types.Ip, port2 uint32, level3 byte) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeTEEWorker: true,
		AsWeTEEWorkerField0: &types.WeteeWorkerPalletCall{
			IsClusterRegister:       true,
			AsClusterRegisterName0:  name0,
			AsClusterRegisterIp1:    ip1,
			AsClusterRegisterPort2:  port2,
			AsClusterRegisterLevel3: level3,
		},
	}
}

// Worker cluster upload proof of work data
// 提交集群的工作证明
// Worker cluster mortgage
// 质押硬件
func MakeClusterMortgageCall(id0 uint64, cpu1 uint32, mem2 uint32, cvmCpu3 uint32, cvmMem4 uint32, disk5 uint32, gpu6 uint32, deposit7 types1.UCompact) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeTEEWorker: true,
		AsWeTEEWorkerField0: &types.WeteeWorkerPalletCall{
			IsClusterMortgage:         true,
			AsClusterMortgageId0:      id0,
			AsClusterMortgageCpu1:     cpu1,
			AsClusterMortgageMem2:     mem2,
			AsClusterMortgageCvmCpu3:  cvmCpu3,
			AsClusterMortgageCvmMem4:  cvmMem4,
			AsClusterMortgageDisk5:    disk5,
			AsClusterMortgageGpu6:     gpu6,
			AsClusterMortgageDeposit7: deposit7,
		},
	}
}

// Worker cluster unmortgage
// 解抵押
func MakeClusterUnmortgageCall(id0 uint64, blockNum1 uint64) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeTEEWorker: true,
		AsWeTEEWorkerField0: &types.WeteeWorkerPalletCall{
			IsClusterUnmortgage:          true,
			AsClusterUnmortgageId0:       id0,
			AsClusterUnmortgageBlockNum1: blockNum1,
		},
	}
}

// 启动或重启服务
// 启动或重启服务
func MakeWorkStartCall(workId0 types.WorkId, report1 types.OptionTByteSlice, deployKey2 [32]byte) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeTEEWorker: true,
		AsWeTEEWorkerField0: &types.WeteeWorkerPalletCall{
			IsWorkStart:           true,
			AsWorkStartWorkId0:    workId0,
			AsWorkStartReport1:    report1,
			AsWorkStartDeployKey2: deployKey2,
		},
	}
}

// Work proof of work data upload
// 提交工作证明
func MakeWorkProofUploadCall(workId0 types.WorkId, proof1 types.OptionTProofOfWork, report2 types.OptionTByteSlice) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeTEEWorker: true,
		AsWeTEEWorkerField0: &types.WeteeWorkerPalletCall{
			IsWorkProofUpload:        true,
			AsWorkProofUploadWorkId0: workId0,
			AsWorkProofUploadProof1:  proof1,
			AsWorkProofUploadReport2: report2,
		},
	}
}

// Worker cluster withdrawal
// 提现余额
func MakeClusterWithdrawalCall(workId0 types.WorkId, amount1 types1.U128) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeTEEWorker: true,
		AsWeTEEWorkerField0: &types.WeteeWorkerPalletCall{
			IsClusterWithdrawal:        true,
			AsClusterWithdrawalWorkId0: workId0,
			AsClusterWithdrawalAmount1: amount1,
		},
	}
}

// Worker cluster stop
// 停止集群
func MakeClusterStopCall(id0 uint64) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeTEEWorker: true,
		AsWeTEEWorkerField0: &types.WeteeWorkerPalletCall{
			IsClusterStop:    true,
			AsClusterStopId0: id0,
		},
	}
}

// Worker cluster report
// 投诉集群
func MakeClusterReportCall(clusterId0 uint64, workId1 types.WorkId, reason2 []byte) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeTEEWorker: true,
		AsWeTEEWorkerField0: &types.WeteeWorkerPalletCall{
			IsClusterReport:           true,
			AsClusterReportClusterId0: clusterId0,
			AsClusterReportWorkId1:    workId1,
			AsClusterReportReason2:    reason2,
		},
	}
}

// Worker report stop
// 停止投诉
func MakeReportCloseCall(clusterId0 uint64, workId1 types.WorkId) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeTEEWorker: true,
		AsWeTEEWorkerField0: &types.WeteeWorkerPalletCall{
			IsReportClose:           true,
			AsReportCloseClusterId0: clusterId0,
			AsReportCloseWorkId1:    workId1,
		},
	}
}

// Work stop
// 停止应用
func MakeWorkStopCall(workId0 types.WorkId) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeTEEWorker: true,
		AsWeTEEWorkerField0: &types.WeteeWorkerPalletCall{
			IsWorkStop:        true,
			AsWorkStopWorkId0: workId0,
		},
	}
}

// Set boot peers
// 设置引导节点
func MakeSetBootPeersCall(boots0 []types.P2PAddr) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeTEEWorker: true,
		AsWeTEEWorkerField0: &types.WeteeWorkerPalletCall{
			IsSetBootPeers:       true,
			AsSetBootPeersBoots0: boots0,
		},
	}
}
func MakeSetStageCall(stage0 uint32) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeTEEWorker: true,
		AsWeTEEWorkerField0: &types.WeteeWorkerPalletCall{
			IsSetStage:       true,
			AsSetStageStage0: stage0,
		},
	}
}

// 上传共识节点代码
// update consensus node code
func MakeUploadCodeCall(mrenclave0 []byte, mrsigner1 []byte) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeTEEWorker: true,
		AsWeTEEWorkerField0: &types.WeteeWorkerPalletCall{
			IsUploadCode:           true,
			AsUploadCodeMrenclave0: mrenclave0,
			AsUploadCodeMrsigner1:  mrsigner1,
		},
	}
}
