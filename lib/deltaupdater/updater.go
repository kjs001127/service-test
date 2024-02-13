package deltaupdater

import (
	"context"
)

type DeltaUpdater[T any, ID comparable] struct {
	IDOf func(T) ID

	DoUpdate func(ctx context.Context, oldbie T, newbie T) error
	DoInsert func(ctx context.Context, newbie T) error
	DoDelete func(ctx context.Context, oldbie T) error
}

func (u *DeltaUpdater[T, ID]) Update(ctx context.Context, oldbies []T, newbies []T) error {
	if err := u.deltaUpsert(ctx, oldbies, newbies); err != nil {
		return err
	}
	if err := u.deltaDelete(ctx, oldbies, newbies); err != nil {
		return err
	}
	return nil
}

func (u *DeltaUpdater[T, ID]) deltaUpsert(ctx context.Context, oldbies []T, newbies []T) error {
	oldbieMap := make(map[ID]T)
	for _, oldbie := range oldbies {
		oldbieMap[u.IDOf(oldbie)] = oldbie
	}

	for _, newbie := range newbies {
		if oldbie, exists := oldbieMap[u.IDOf(newbie)]; exists {
			if err := u.DoUpdate(ctx, oldbie, newbie); err != nil {
				return err
			}
		} else {
			if err := u.DoInsert(ctx, newbie); err != nil {
				return err
			}
		}
	}
	return nil
}

func (u *DeltaUpdater[T, ID]) deltaDelete(ctx context.Context, oldbies []T, newbies []T) error {
	newbieMap := make(map[ID]T)
	for _, newbie := range newbies {
		newbieMap[u.IDOf(newbie)] = newbie
	}

	for _, oldbie := range oldbies {
		if _, exists := newbieMap[u.IDOf(oldbie)]; !exists {
			if err := u.DoDelete(ctx, oldbie); err != nil {
				return err
			}
		}
	}
	return nil
}
