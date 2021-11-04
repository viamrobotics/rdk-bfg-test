package robotimpl

import (
	"context"

	"github.com/go-errors/errors"
	"go.uber.org/multierr"

	"go.viam.com/core/config"
	"go.viam.com/core/metadata/service"
	"go.viam.com/core/resource"
)

// Reconfigure will safely reconfigure a robot based on the given config. It will make
// a best effort to remove no longer in use parts, but if it fails to do so, they could
// possibly leak resources.
func (r *localRobot) Reconfigure(ctx context.Context, newConfig *config.Config) error {
	diff, err := config.DiffConfigs(r.config, newConfig)
	if err != nil {
		return err
	}
	if diff.Equal {
		return nil
	}
	r.logger.Debugf("reconfiguring with %s", diff)
	draft := newDraftRobot(r, diff)
	err = draft.ProcessAndCommit(ctx)
	if err != nil {
		return err
	}

	// if metadata exists, update it
	if svc := service.ContextService(ctx); svc != nil {
		if err := r.UpdateMetadata(svc); err != nil {
			return err
		}
	}

	// update web service
	webSvc, ok := r.ServiceByName(WebSvcName)
	if ok {
		updateable, ok := webSvc.(resource.Updateable)
		if ok {
			if err := updateable.Update(r.parts.resources); err != nil {
				return err
			}
		}
	}

	return nil
}

// A draftRobot is responsible for the workflow of turning in
// a newly proposed robot into a robot ready to be swapped in
// for an existing one. It understands how to rollback and commit
// changes as safe as it possibly can.
type draftRobot struct {
	original *localRobot
	diff     *config.Diff
	parts    *robotParts

	// additions and removals consist of modifications as well since we treat
	// any modification as a removal to commit and an addition to rollback.
	additions     *robotParts
	modifications *robotParts
	removals      *robotParts
}

// newDraftRobot returns a new draft of a robot based on the given
// original robot and the diff describing what the new robot
// should look like.
func newDraftRobot(r *localRobot, diff *config.Diff) *draftRobot {
	return &draftRobot{
		original:      r,
		diff:          diff,
		parts:         r.parts.Clone(),
		additions:     newRobotParts(r.logger),
		modifications: newRobotParts(r.logger),
		removals:      newRobotParts(r.logger),
	}
}

// Rollback rolls back any intermediate changes made.
func (draft *draftRobot) Rollback() error {
	return draft.additions.Close()
}

// ProcessAndCommit processes all changes in an all-or-nothing fashion
// and then finally commits them; otherwise any changes made along the
// way are rolled back.
func (draft *draftRobot) ProcessAndCommit(ctx context.Context) (err error) {
	defer func() {
		if err != nil {
			draft.original.logger.Infow("rolling back draft changes due to error", "error", err)
			if rollbackErr := draft.Rollback(); rollbackErr != nil {
				err = multierr.Combine(err, errors.Errorf("error rolling back draft changes: %w", rollbackErr))
			}
		}
	}()

	if err := draft.Process(ctx); err != nil {
		return errors.Errorf("error processing draft changes: %w", err)
	}

	draft.original.logger.Info("committing draft changes")
	if err := draft.Commit(ctx); err != nil {
		return errors.Errorf("error committing draft changes: %w", err)
	}
	return nil
}

// Commit commits all changes and updates the original
// robot in place.
func (draft *draftRobot) Commit(ctx context.Context) error {
	draft.original.mu.Lock()
	defer draft.original.mu.Unlock()

	addResult, err := draft.parts.MergeAdd(draft.additions)
	if err != nil {
		return err
	}
	modifyResult, err := draft.parts.MergeModify(ctx, draft.modifications, draft.diff)
	if err != nil {
		return err
	}
	draft.parts.MergeRemove(draft.removals)
	draft.original.parts = draft.parts
	draft.original.config = draft.diff.Right

	if err := addResult.Process(draft.removals); err != nil {
		draft.original.logger.Errorw("error processing add result but still committing changes", "error", err)
	}
	if err := modifyResult.Process(draft.removals); err != nil {
		draft.original.logger.Errorw("error processing modify result but still committing changes", "error", err)
	}
	if err := draft.removals.Close(); err != nil {
		draft.original.logger.Errorw("error closing parts removed but still committing changes", "error", err)
	}
	return nil
}

// Process processes all types changes into the draft robot.
func (draft *draftRobot) Process(ctx context.Context) error {
	// We specifically add, modify, and remove parts of the robot
	// in order to provide the best chance of reconfiguration/compatibility.
	// This assumes the addition/modification of parts does not cause
	// any adverse effects before any removals.
	if err := draft.ProcessAddChanges(ctx); err != nil {
		return err
	}
	if err := draft.ProcessModifyChanges(ctx); err != nil {
		return err
	}
	if err := draft.ProcessRemoveChanges(); err != nil {
		return err
	}
	return nil
}

// ProcessAddChanges processes only additive changes.
func (draft *draftRobot) ProcessAddChanges(ctx context.Context) error {
	return draft.additions.processConfig(ctx, draft.diff.Added, draft.original, draft.original.logger)
}

// ProcessModifyChanges processes only modificative changes.
func (draft *draftRobot) ProcessModifyChanges(ctx context.Context) error {
	return draft.modifications.processModifiedConfig(ctx, draft.diff.Modified, draft.original, draft.original.logger)
}

// ProcessRemoveChanges processes only subtractive changes.
func (draft *draftRobot) ProcessRemoveChanges() error {
	filtered, err := draft.parts.FilterFromConfig(draft.diff.Removed, draft.original.logger)
	if err != nil {
		return err
	}
	draft.removals = filtered
	return nil
}
