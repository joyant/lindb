package discovery

import (
	"context"
	"encoding/json"
	"time"

	"github.com/eleme/lindb/models"
	"github.com/eleme/lindb/pkg/logger"
	"github.com/eleme/lindb/pkg/pathutil"
	"github.com/eleme/lindb/pkg/state"

	"go.uber.org/zap"
)

// Registry represents server node register
type Registry interface {
	// Register registers node info, add it to active node list for discovery
	Register(node models.Node) error
	// Deregister deregister node info, remove it from active list
	Deregister(node models.Node) error
	// Close closes registry, releases resources
	Close() error
}

// registry implements registry interface for server node register with prefix
type registry struct {
	prefix string
	ttl    int64
	repo   state.Repository

	ctx    context.Context
	cancel context.CancelFunc

	log *zap.Logger
}

// NewRegistry returns a new registry with prefix and ttl
func NewRegistry(repo state.Repository, prefix string, ttl int64) Registry {
	ctx, cancel := context.WithCancel(context.Background())
	return &registry{
		prefix: prefix,
		ttl:    ttl,
		repo:   repo,
		ctx:    ctx,
		cancel: cancel,
		log:    logger.GetLogger(),
	}
}

// Register registers node info, add it to active node list for discovery
func (r *registry) Register(node models.Node) error {
	nodeBytes, err := json.Marshal(node)
	if err != nil {
		r.log.Error("convert node to byte error when register node info", zap.Error(err))
		return err
	}
	// register node info
	path := pathutil.GetNodePath(r.prefix, node.String())
	// register node if fail retry it
	go r.register(path, nodeBytes)
	return nil
}

// Deregister deregisters node info, remove it from active list
func (r *registry) Deregister(node models.Node) error {
	return r.repo.Delete(r.ctx, pathutil.GetNodePath(r.prefix, node.String()))
}

// Close closes registry, releases resources
func (r *registry) Close() error {
	r.cancel()
	return nil
}

// register registers node info, if fail do retry
func (r *registry) register(path string, node []byte) {
	for {
		// if ctx happen err, exit register loop
		if r.ctx.Err() != nil {
			return
		}
		closed, err := r.repo.Heartbeat(r.ctx, path, node, r.ttl)
		if err != nil {
			r.log.Error("register storage node error", zap.Error(err))
			time.Sleep(500 * time.Millisecond)
			continue
		}

		r.log.Info("register storage node successfully", zap.String("path", path))

		select {
		case <-r.ctx.Done():
			r.log.Warn("context is canceled, exit register loop")
			return
		case <-closed:
			r.log.Warn("the heartbeat channel is closed, retry register")
		}
	}
}
