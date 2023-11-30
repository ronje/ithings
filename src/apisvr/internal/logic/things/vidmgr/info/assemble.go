package info

import (
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/apisvr/internal/logic"
	"github.com/i-Things/things/src/apisvr/internal/types"
	"github.com/i-Things/things/src/vidsvr/pb/vid"
)

func VidmgrInfoToApi(v *vid.VidmgrInfo) *types.VidmgrInfo {
	return &types.VidmgrInfo{
		CreatedTime: v.CreatedTime, //创建时间 只读
		InfoCommon: types.InfoCommon{
			VidmgrID:     v.VidmgrID,                 //服务id 只读
			VidmgrName:   v.VidmgrName,               //服务名称
			VidmgrIpV4:   v.VidmgrIpV4,               //服务IP
			VidmgrPort:   v.VidmgrPort,               //服务端口
			VidmgrType:   v.VidmgrType,               //服务类型:1:zlmediakit,2:srs,3:monibuca
			VidmgrSecret: v.VidmgrSecret,             //服务秘钥
			VidmgrStatus: v.VidmgrStatus,             //服务状态: 1：离线 2：在线  3：未激活
			Desc:         utils.ToNullString(v.Desc), //描述
			Tags:         logic.ToTagsType(v.Tags),
		},
	}
}

func ToVidmgrConfigRpc(pi *types.ServerConfig) *vid.VidmgrConfig {
	dpi := &vid.VidmgrConfig{
		GeneralMediaServerId:           pi.GeneralMediaServerId,
		ApiDebug:                       pi.ApiDebug,
		ApiDefaultSnap:                 pi.ApiDefaultSnap,
		ApiSecret:                      pi.ApiSecret,
		ApiSnapRoot:                    pi.ApiSnapRoot,
		ClusterOriginUrl:               pi.ClusterOriginUrl,
		ClusterRetryCount:              pi.ClusterRetryCount,
		ClusterTimeoutSec:              pi.ClusterTimeoutSec,
		FfmpegBin:                      pi.FfmpegBin,
		FfmpegCmd:                      pi.FfmpegCmd,
		FfmpegLog:                      pi.FfmpegLog,
		FfmpegRestartSec:               pi.FfmpegRestartSec,
		FfmpegSnap:                     pi.FfmpegSnap,
		GeneralCheckNvidiaDev:          pi.GeneralCheckNvidiaDev,
		GeneralEnableVhost:             pi.GeneralEnableVhost,
		GeneralEnableFfmpegLog:         pi.GeneralEnableFfmpegLog,
		GeneralFlowThreshold:           pi.GeneralFlowThreshold,
		GeneralMaxStreamWaitMS:         pi.GeneralMaxStreamWaitMS,
		GeneralMergeWriteMS:            pi.GeneralMergeWriteMS,
		GeneralResetWhenRePlay:         pi.GeneralResetWhenRePlay,
		GeneralStreamNoneReaderDelayMS: pi.GeneralStreamNoneReaderDelayMS,
		GeneralUnreadyFrameCache:       pi.GeneralUnreadyFrameCache,
		GeneralWaitAddTrackMs:          pi.GeneralWaitAddTrackMs,
		GeneralWaitTrackReadyMs:        pi.GeneralWaitTrackReadyMs,
		HlsBroadcastRecordTs:           pi.HlsBroadcastRecordTs,
		HlsDeleteDelaySec:              pi.HlsDeleteDelaySec,
		HlsFileBufSize:                 pi.HlsFileBufSize,
		HlsSegDur:                      pi.HlsSegDur,
		HlsSegKeep:                     pi.HlsSegKeep,
		HlsSegNum:                      pi.HlsSegNum,
		HlsSegRetain:                   pi.HlsSegRetain,
		HookAliveInterval:              pi.HookAliveInterval,
		HookEnable:                     pi.HookEnable,
		HookOnFlowReport:               pi.HookOnFlowReport,
		HookOnHttpAccess:               pi.HookOnHttpAccess,
		HookOnPlay:                     pi.HookOnPlay,
		HookOnPublish:                  pi.HookOnPublish,
		HookOnRecordMp4:                pi.HookOnRecordMp4,
		HookOnRecordTs:                 pi.HookOnRecordTs,
		HookOnRtpServerTimeout:         pi.HookOnRtpServerTimeout,
		HookOnRtspAuth:                 pi.HookOnRtspAuth,
		HookOnRtspRealm:                pi.HookOnRtspRealm,
		HookOnSendRtpStopped:           pi.HookOnSendRtpStopped,
		HookOnServerExited:             pi.HookOnServerExited,
		HookOnServerKeepalive:          pi.HookOnServerKeepalive,
		HookOnServerStarted:            pi.HookOnServerStarted,
		HookOnShellLogin:               pi.HookOnShellLogin,
		HookOnStreamChanged:            pi.HookOnStreamChanged,
		HookOnStreamNoneReader:         pi.HookOnStreamNoneReader,
		HookOnStreamNotFound:           pi.HookOnStreamNotFound,
		HookRetry:                      pi.HookRetry,
		HookRetryDelay:                 pi.HookRetryDelay,
		HookStreamChangedSchemas:       pi.HookStreamChangedSchemas,
		HookTimeoutSec:                 pi.HookTimeoutSec,
		HttpAllowCrossDomains:          pi.HttpAllowCrossDomains,
		HttpAllowIpRange:               pi.HttpAllowIpRange,
		HttpCharSet:                    pi.HttpCharSet,
		HttpDirMenu:                    pi.HttpDirMenu,
		HttpForbidCacheSuffix:          pi.HttpForbidCacheSuffix,
		HttpForwardedIpHeader:          pi.HttpForwardedIpHeader,
		HttpKeepAliveSecond:            pi.HttpKeepAliveSecond,
		HttpMaxReqSize:                 pi.HttpMaxReqSize,
		HttpNotFound:                   pi.HttpNotFound,
		HttpPort:                       pi.HttpPort,
		HttpRootPath:                   pi.HttpRootPath,
		HttpSendBufSize:                pi.HttpSendBufSize,
		HttpSslport:                    pi.HttpSslport,
		HttpVirtualPath:                pi.HttpVirtualPath,
		MulticastAddrMax:               pi.MulticastAddrMax,
		MulticastAddrMin:               pi.MulticastAddrMin,
		MulticastUdpTTL:                pi.MulticastUdpTTL,
		ProtocolAddMuteAudio:           pi.ProtocolAddMuteAudio,
		ProtocolAutoClose:              pi.ProtocolAutoClose,
		ProtocolContinuePushMs:         pi.ProtocolContinuePushMs,
		ProtocolEnableAudio:            pi.ProtocolEnableAudio,
		ProtocolEnableFmp4:             pi.ProtocolEnableFmp4,
		ProtocolEnableHls:              pi.ProtocolEnableHls,
		ProtocolEnableHlsFmp4:          pi.ProtocolEnableHlsFmp4,
		ProtocolEnableMp4:              pi.ProtocolEnableMp4,
		ProtocolEnableRtmp:             pi.ProtocolEnableRtmp,
		ProtocolEnableRtsp:             pi.ProtocolEnableRtsp,
		ProtocolEnableTs:               pi.ProtocolEnableTs,
		ProtocolFmp4Demand:             pi.ProtocolFmp4Demand,
		ProtocolHlsDemand:              pi.ProtocolHlsDemand,
		ProtocolHlsSavePath:            pi.ProtocolHlsSavePath,
		ProtocolModifyStamp:            pi.ProtocolModifyStamp,
		ProtocolMp4AsPlayer:            pi.ProtocolMp4AsPlayer,
		ProtocolMp4MaxSecond:           pi.ProtocolMp4MaxSecond,
		ProtocolMp4SavePath:            pi.ProtocolMp4SavePath,
		ProtocolRtmpDemand:             pi.ProtocolRtmpDemand,
		ProtocolRtspDemand:             pi.ProtocolRtspDemand,
		ProtocolTsDemand:               pi.ProtocolTsDemand,
		RecordAppName:                  pi.RecordAppName,
		RecordFastStart:                pi.RecordFastStart,
		RecordFileBufSize:              pi.RecordFileBufSize,
		RecordFileRepeat:               pi.RecordFileRepeat,
		RecordSampleMS:                 pi.RecordSampleMS,
		RtcExternIP:                    pi.RtcExternIP,
		RtcPort:                        pi.RtcPort,
		RtcPreferredCodecA:             pi.RtcPreferredCodecA,
		RtcPreferredCodecV:             pi.RtcPreferredCodecV,
		RtcRembBitRate:                 pi.RtcRembBitRate,
		RtcTcpPort:                     pi.RtcTcpPort,
		RtcTimeoutSec:                  pi.RtcTimeoutSec,
		RtmpHandshakeSecond:            pi.RtmpHandshakeSecond,
		RtmpKeepAliveSecond:            pi.RtmpKeepAliveSecond,
		RtmpPort:                       pi.RtmpPort,
		RtmpSslport:                    pi.RtmpSslport,
		RtpAudioMtuSize:                pi.RtpAudioMtuSize,
		RtpH264StapA:                   pi.RtpH264StapA,
		RtpLowLatency:                  pi.RtpLowLatency,
		RtpRtpMaxSize:                  pi.RtpRtpMaxSize,
		RtpVideoMtuSize:                pi.RtpVideoMtuSize,
		RtpProxyDumpDir:                pi.RtpProxyDumpDir,
		RtpProxyGopCache:               pi.RtpProxyGopCache,
		RtpProxyH264Pt:                 pi.RtpProxyH264Pt,
		RtpProxyH265Pt:                 pi.RtpProxyH265Pt,
		RtpProxyOpusPt:                 pi.RtpProxyOpusPt,
		RtpProxyPort:                   pi.RtpProxyPort,
		RtpProxyPortRange:              pi.RtpProxyPortRange,
		RtpProxyPsPt:                   pi.RtpProxyPsPt,
		RtpProxyTimeoutSec:             pi.RtpProxyTimeoutSec,
		RtspAuthBasic:                  pi.RtspAuthBasic,
		RtspDirectProxy:                pi.RtspDirectProxy,
		RtspHandshakeSecond:            pi.RtspHandshakeSecond,
		RtspKeepAliveSecond:            pi.RtspKeepAliveSecond,
		RtspLowLatency:                 pi.RtspLowLatency,
		RtspPort:                       pi.RtspPort,
		RtspRtpTransportType:           pi.RtspRtpTransportType,
		RtspSslport:                    pi.RtspSslport,
		ShellMaxReqSize:                pi.ShellMaxReqSize,
		ShellPort:                      pi.ShellPort,
		SrtLatencyMul:                  pi.SrtLatencyMul,
		SrtPktBufSize:                  pi.SrtPktBufSize,
		SrtPort:                        pi.SrtPort,
		SrtTimeoutSec:                  pi.SrtTimeoutSec,
	}
	return dpi
}

// func ToVidmgrConfigRpc(pi *vid.VidmgrConfig) *types.ServerConfig {
func ToVidmgrConfigApi(pi *vid.VidmgrConfig) *types.ServerConfig {
	dpi := &types.ServerConfig{
		GeneralMediaServerId:           pi.GeneralMediaServerId,
		ApiDebug:                       pi.ApiDebug,
		ApiDefaultSnap:                 pi.ApiDefaultSnap,
		ApiSecret:                      pi.ApiSecret,
		ApiSnapRoot:                    pi.ApiSnapRoot,
		ClusterOriginUrl:               pi.ClusterOriginUrl,
		ClusterRetryCount:              pi.ClusterRetryCount,
		ClusterTimeoutSec:              pi.ClusterTimeoutSec,
		FfmpegBin:                      pi.FfmpegBin,
		FfmpegCmd:                      pi.FfmpegCmd,
		FfmpegLog:                      pi.FfmpegLog,
		FfmpegRestartSec:               pi.FfmpegRestartSec,
		FfmpegSnap:                     pi.FfmpegSnap,
		GeneralCheckNvidiaDev:          pi.GeneralCheckNvidiaDev,
		GeneralEnableVhost:             pi.GeneralEnableVhost,
		GeneralEnableFfmpegLog:         pi.GeneralEnableFfmpegLog,
		GeneralFlowThreshold:           pi.GeneralFlowThreshold,
		GeneralMaxStreamWaitMS:         pi.GeneralMaxStreamWaitMS,
		GeneralMergeWriteMS:            pi.GeneralMergeWriteMS,
		GeneralResetWhenRePlay:         pi.GeneralResetWhenRePlay,
		GeneralStreamNoneReaderDelayMS: pi.GeneralStreamNoneReaderDelayMS,
		GeneralUnreadyFrameCache:       pi.GeneralUnreadyFrameCache,
		GeneralWaitAddTrackMs:          pi.GeneralWaitAddTrackMs,
		GeneralWaitTrackReadyMs:        pi.GeneralWaitTrackReadyMs,
		HlsBroadcastRecordTs:           pi.HlsBroadcastRecordTs,
		HlsDeleteDelaySec:              pi.HlsDeleteDelaySec,
		HlsFileBufSize:                 pi.HlsFileBufSize,
		HlsSegDur:                      pi.HlsSegDur,
		HlsSegKeep:                     pi.HlsSegKeep,
		HlsSegNum:                      pi.HlsSegNum,
		HlsSegRetain:                   pi.HlsSegRetain,
		HookAliveInterval:              pi.HookAliveInterval,
		HookEnable:                     pi.HookEnable,
		HookOnFlowReport:               pi.HookOnFlowReport,
		HookOnHttpAccess:               pi.HookOnHttpAccess,
		HookOnPlay:                     pi.HookOnPlay,
		HookOnPublish:                  pi.HookOnPublish,
		HookOnRecordMp4:                pi.HookOnRecordMp4,
		HookOnRecordTs:                 pi.HookOnRecordTs,
		HookOnRtpServerTimeout:         pi.HookOnRtpServerTimeout,
		HookOnRtspAuth:                 pi.HookOnRtspAuth,
		HookOnRtspRealm:                pi.HookOnRtspRealm,
		HookOnSendRtpStopped:           pi.HookOnSendRtpStopped,
		HookOnServerExited:             pi.HookOnServerExited,
		HookOnServerKeepalive:          pi.HookOnServerKeepalive,
		HookOnServerStarted:            pi.HookOnServerStarted,
		HookOnShellLogin:               pi.HookOnShellLogin,
		HookOnStreamChanged:            pi.HookOnStreamChanged,
		HookOnStreamNoneReader:         pi.HookOnStreamNoneReader,
		HookOnStreamNotFound:           pi.HookOnStreamNotFound,
		HookRetry:                      pi.HookRetry,
		HookRetryDelay:                 pi.HookRetryDelay,
		HookStreamChangedSchemas:       pi.HookStreamChangedSchemas,
		HookTimeoutSec:                 pi.HookTimeoutSec,
		HttpAllowCrossDomains:          pi.HttpAllowCrossDomains,
		HttpAllowIpRange:               pi.HttpAllowIpRange,
		HttpCharSet:                    pi.HttpCharSet,
		HttpDirMenu:                    pi.HttpDirMenu,
		HttpForbidCacheSuffix:          pi.HttpForbidCacheSuffix,
		HttpForwardedIpHeader:          pi.HttpForwardedIpHeader,
		HttpKeepAliveSecond:            pi.HttpKeepAliveSecond,
		HttpMaxReqSize:                 pi.HttpMaxReqSize,
		HttpNotFound:                   pi.HttpNotFound,
		HttpPort:                       pi.HttpPort,
		HttpRootPath:                   pi.HttpRootPath,
		HttpSendBufSize:                pi.HttpSendBufSize,
		HttpSslport:                    pi.HttpSslport,
		HttpVirtualPath:                pi.HttpVirtualPath,
		MulticastAddrMax:               pi.MulticastAddrMax,
		MulticastAddrMin:               pi.MulticastAddrMin,
		MulticastUdpTTL:                pi.MulticastUdpTTL,
		ProtocolAddMuteAudio:           pi.ProtocolAddMuteAudio,
		ProtocolAutoClose:              pi.ProtocolAutoClose,
		ProtocolContinuePushMs:         pi.ProtocolContinuePushMs,
		ProtocolEnableAudio:            pi.ProtocolEnableAudio,
		ProtocolEnableFmp4:             pi.ProtocolEnableFmp4,
		ProtocolEnableHls:              pi.ProtocolEnableHls,
		ProtocolEnableHlsFmp4:          pi.ProtocolEnableHlsFmp4,
		ProtocolEnableMp4:              pi.ProtocolEnableMp4,
		ProtocolEnableRtmp:             pi.ProtocolEnableRtmp,
		ProtocolEnableRtsp:             pi.ProtocolEnableRtsp,
		ProtocolEnableTs:               pi.ProtocolEnableTs,
		ProtocolFmp4Demand:             pi.ProtocolFmp4Demand,
		ProtocolHlsDemand:              pi.ProtocolHlsDemand,
		ProtocolHlsSavePath:            pi.ProtocolHlsSavePath,
		ProtocolModifyStamp:            pi.ProtocolModifyStamp,
		ProtocolMp4AsPlayer:            pi.ProtocolMp4AsPlayer,
		ProtocolMp4MaxSecond:           pi.ProtocolMp4MaxSecond,
		ProtocolMp4SavePath:            pi.ProtocolMp4SavePath,
		ProtocolRtmpDemand:             pi.ProtocolRtmpDemand,
		ProtocolRtspDemand:             pi.ProtocolRtspDemand,
		ProtocolTsDemand:               pi.ProtocolTsDemand,
		RecordAppName:                  pi.RecordAppName,
		RecordFastStart:                pi.RecordFastStart,
		RecordFileBufSize:              pi.RecordFileBufSize,
		RecordFileRepeat:               pi.RecordFileRepeat,
		RecordSampleMS:                 pi.RecordSampleMS,
		RtcExternIP:                    pi.RtcExternIP,
		RtcPort:                        pi.RtcPort,
		RtcPreferredCodecA:             pi.RtcPreferredCodecA,
		RtcPreferredCodecV:             pi.RtcPreferredCodecV,
		RtcRembBitRate:                 pi.RtcRembBitRate,
		RtcTcpPort:                     pi.RtcTcpPort,
		RtcTimeoutSec:                  pi.RtcTimeoutSec,
		RtmpHandshakeSecond:            pi.RtmpHandshakeSecond,
		RtmpKeepAliveSecond:            pi.RtmpKeepAliveSecond,
		RtmpPort:                       pi.RtmpPort,
		RtmpSslport:                    pi.RtmpSslport,
		RtpAudioMtuSize:                pi.RtpAudioMtuSize,
		RtpH264StapA:                   pi.RtpH264StapA,
		RtpLowLatency:                  pi.RtpLowLatency,
		RtpRtpMaxSize:                  pi.RtpRtpMaxSize,
		RtpVideoMtuSize:                pi.RtpVideoMtuSize,
		RtpProxyDumpDir:                pi.RtpProxyDumpDir,
		RtpProxyGopCache:               pi.RtpProxyGopCache,
		RtpProxyH264Pt:                 pi.RtpProxyH264Pt,
		RtpProxyH265Pt:                 pi.RtpProxyH265Pt,
		RtpProxyOpusPt:                 pi.RtpProxyOpusPt,
		RtpProxyPort:                   pi.RtpProxyPort,
		RtpProxyPortRange:              pi.RtpProxyPortRange,
		RtpProxyPsPt:                   pi.RtpProxyPsPt,
		RtpProxyTimeoutSec:             pi.RtpProxyTimeoutSec,
		RtspAuthBasic:                  pi.RtspAuthBasic,
		RtspDirectProxy:                pi.RtspDirectProxy,
		RtspHandshakeSecond:            pi.RtspHandshakeSecond,
		RtspKeepAliveSecond:            pi.RtspKeepAliveSecond,
		RtspLowLatency:                 pi.RtspLowLatency,
		RtspPort:                       pi.RtspPort,
		RtspRtpTransportType:           pi.RtspRtpTransportType,
		RtspSslport:                    pi.RtspSslport,
		ShellMaxReqSize:                pi.ShellMaxReqSize,
		ShellPort:                      pi.ShellPort,
		SrtLatencyMul:                  pi.SrtLatencyMul,
		SrtPktBufSize:                  pi.SrtPktBufSize,
		SrtPort:                        pi.SrtPort,
		SrtTimeoutSec:                  pi.SrtTimeoutSec,
	}
	return dpi
}
