package session

import (
	"github.com/openlibrecommunity/olcrtc/internal/transport"
	"github.com/openlibrecommunity/olcrtc/internal/transport/seichannel"
	"github.com/openlibrecommunity/olcrtc/internal/transport/videochannel"
	"github.com/openlibrecommunity/olcrtc/internal/transport/vp8channel"
)

// buildTransportOptions packs per-transport tuning fields from cfg into the
// typed Options value the chosen transport expects. Transports without
// tunable options (datachannel) return nil.
func buildTransportOptions(cfg Config) transport.Options {
	switch cfg.Transport {
	case transportVideo:
		return videochannel.Options{
			Width:      cfg.VideoWidth,
			Height:     cfg.VideoHeight,
			FPS:        cfg.VideoFPS,
			Bitrate:    cfg.VideoBitrate,
			HW:         cfg.VideoHW,
			QRSize:     cfg.VideoQRSize,
			QRRecovery: cfg.VideoQRRecovery,
			Codec:      cfg.VideoCodec,
			TileModule: cfg.VideoTileModule,
			TileRS:     cfg.VideoTileRS,
		}
	case transportVP8:
		return vp8channel.Options{
			FPS:       cfg.VP8FPS,
			BatchSize: cfg.VP8BatchSize,
		}
	case transportSEI:
		return seichannel.Options{
			FPS:          cfg.SEIFPS,
			BatchSize:    cfg.SEIBatchSize,
			FragmentSize: cfg.SEIFragmentSize,
			AckTimeoutMS: cfg.SEIAckTimeoutMS,
		}
	default:
		return nil
	}
}
