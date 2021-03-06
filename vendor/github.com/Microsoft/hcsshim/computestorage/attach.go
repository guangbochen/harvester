package computestorage

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Microsoft/hcsshim/internal/oc"
	"go.opencensus.io/trace"
)

// AttachLayerStorageFilter sets up the layer storage filter on a writable
// container layer.
//
// `layerPath` is a path to a directory the writable layer is mounted. If the
// path does not end in a `\` the platform will append it automatically.
//
// `layerData` is the parent read-only layer data.
func AttachLayerStorageFilter(ctx context.Context, layerPath string, layerData LayerData) (err error) {
	title := "hcsshim.AttachLayerStorageFilter"
	ctx, span := trace.StartSpan(ctx, title)
	defer span.End()
	defer func() { oc.SetSpanStatus(span, err) }()
	span.AddAttributes(
		trace.StringAttribute("layerPath", layerPath),
	)

	bytes, err := json.Marshal(layerData)
	if err != nil {
		return err
	}

	err = hcsAttachLayerStorageFilter(layerPath, string(bytes))
	if err != nil {
		return fmt.Errorf("failed to attach layer storage filter: %s", err)
	}
	return nil
}
