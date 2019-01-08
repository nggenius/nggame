package entity

import (
	"fmt"
	"time"
)

type AttrAlter func(e *Entity, attr string, val interface{}, old interface{})

type AttrNotifier struct {
	cb       map[string]AttrAlter
	invoking bool
}

func NewAttrNotifier() *AttrNotifier {
	n := &AttrNotifier{}
	n.cb = make(map[string]AttrAlter)
	return n
}

// 回调所有挂钩函数
func (a *AttrNotifier) Invoke(e *Entity, attr string, val interface{}, old interface{}) {
	a.invoking = true
	for n, f := range a.cb {
		start := time.Now()
		f(e, attr, val, old)
		if ns := time.Now().Sub(start).Nanoseconds(); time.Duration(ns) > time.Millisecond*10 {
			if e.core != nil {
				e.core.Logger().LogWarn(n, "running times ", ns/1000000)
			}
		}
	}
	a.invoking = false
}

// 增加一个新的回调
func (a *AttrNotifier) Add(name string, cb AttrAlter) error {
	if a.invoking {
		panic("can't add when invoking")
	}

	if _, has := a.cb[name]; has {
		return fmt.Errorf("notifier add twice %s", name)
	}

	a.cb[name] = cb
	return nil
}

// 移除一个回调
func (a *AttrNotifier) Remove(name string) error {
	if a.invoking {
		panic("can't remove when invoking")
	}

	if _, has := a.cb[name]; has {
		delete(a.cb, name)
		return nil
	}

	return fmt.Errorf("notifier %s not found", name)
}

type TableAlter func(e *Entity, table string, op, row, col int)

// 表格变动通知
type TableNotifier struct {
	cb       map[string]TableAlter
	invoking bool
}

func NewTableNotifier() *TableNotifier {
	n := &TableNotifier{}
	n.cb = make(map[string]TableAlter)
	return n
}

// 回调所有挂钩函数
func (a *TableNotifier) Invoke(e *Entity, table string, op, row, col int) {
	a.invoking = true
	for n, f := range a.cb {
		start := time.Now()
		f(e, table, op, row, col)
		if ns := time.Now().Sub(start).Nanoseconds(); time.Duration(ns) > time.Millisecond*10 {
			if e.core != nil {
				e.core.Logger().LogWarn(n, "running times ", ns/1000000)
			}
		}
	}
	a.invoking = false
}

// 增加一个新的回调
func (a *TableNotifier) Add(name string, cb TableAlter) error {
	if a.invoking {
		panic("can't add when invoking")
	}

	if _, has := a.cb[name]; has {
		return fmt.Errorf("table notifier add twice %s", name)
	}

	a.cb[name] = cb
	return nil
}

// 移除一个回调
func (a *TableNotifier) Remove(name string) error {
	if a.invoking {
		panic("can't remove when invoking")
	}

	if _, has := a.cb[name]; has {
		delete(a.cb, name)
		return nil
	}

	return fmt.Errorf("table notifier %s not found", name)
}
