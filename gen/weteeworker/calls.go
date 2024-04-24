package weteeworker

import (
	types1 "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	types "github.com/wetee-dao/go-sdk/gen/types"
)

// Worker cluster register
// 集群注册
func MakeClusterRegisterCall(name0 []byte, ip1 []types.Ip, port2 uint32, level3 byte) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeteeWorker: true,
		AsWeteeWorkerField0: &types.WeteeWorkerPalletCall{
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
func MakeClusterProofUploadCall(id0 uint64, proof1 []byte) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeteeWorker: true,
		AsWeteeWorkerField0: &types.WeteeWorkerPalletCall{
			IsClusterProofUpload:       true,
			AsClusterProofUploadId0:    id0,
			AsClusterProofUploadProof1: proof1,
		},
	}
}

// Worker cluster mortgage
// 质押硬件
func MakeClusterMortgageCall(id0 uint64, cpu1 uint32, mem2 uint32, disk3 uint32, gpu4 uint32, deposit5 types1.UCompact) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeteeWorker: true,
		AsWeteeWorkerField0: &types.WeteeWorkerPalletCall{
			IsClusterMortgage:         true,
			AsClusterMortgageId0:      id0,
			AsClusterMortgageCpu1:     cpu1,
			AsClusterMortgageMem2:     mem2,
			AsClusterMortgageDisk3:    disk3,
			AsClusterMortgageGpu4:     gpu4,
			AsClusterMortgageDeposit5: deposit5,
		},
	}
}

// Worker cluster unmortgage
// 解抵押
func MakeClusterUnmortgageCall(id0 uint64, blockNum1 uint64) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeteeWorker: true,
		AsWeteeWorkerField0: &types.WeteeWorkerPalletCall{
			IsClusterUnmortgage:          true,
			AsClusterUnmortgageId0:       id0,
			AsClusterUnmortgageBlockNum1: blockNum1,
		},
	}
}

// Work proof of work data upload
// 提交工作证明
func MakeWorkProofUploadCall(workId0 types.WorkId, proof1 types.OptionTProofOfWork, report2 types.OptionTByteSlice) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeteeWorker: true,
		AsWeteeWorkerField0: &types.WeteeWorkerPalletCall{
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
		IsWeteeWorker: true,
		AsWeteeWorkerField0: &types.WeteeWorkerPalletCall{
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
		IsWeteeWorker: true,
		AsWeteeWorkerField0: &types.WeteeWorkerPalletCall{
			IsClusterStop:    true,
			AsClusterStopId0: id0,
		},
	}
}

// Worker cluster report
// 投诉集群
func MakeClusterReportCall(clusterId0 uint64, workId1 types.WorkId, reason2 []byte) types.RuntimeCall {
	return types.RuntimeCall{
		IsWeteeWorker: true,
		AsWeteeWorkerField0: &types.WeteeWorkerPalletCall{
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
		IsWeteeWorker: true,
		AsWeteeWorkerField0: &types.WeteeWorkerPalletCall{
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
		IsWeteeWorker: true,
		AsWeteeWorkerField0: &types.WeteeWorkerPalletCall{
			IsWorkStop:        true,
			AsWorkStopWorkId0: workId0,
		},
	}
}
