package gameobject

import (
	"fmt"
)

// SetCap 设置容量
func (b *BaseBehavior) SetCap(cap int) error {
	if b.c == nil {
		b.c = NewContainer(cap)
		return nil
	}

	return b.c.Resize(cap)
}

// Cap 获取容量
func (b *BaseBehavior) Cap() int {
	if b.c == nil {
		return 0
	}
	return b.c.Cap
}

// CanAdd 是否可以增加子对象
func (b *BaseBehavior) CanAdd(pos int, g GameObject) bool {
	return true
}

// AddChildIf 增加一个对象
func (b *BaseBehavior) AddChildIf(pos int, g interface{}) (int, error) {
	if gb, ok := g.(GameObject); ok {
		return b.AddChild(pos, gb)
	}
	return -1, fmt.Errorf("object not implement GameObject")
}

// AddChild 增加一个对象
func (b *BaseBehavior) AddChild(pos int, g GameObject) (int, error) {
	if b.c.Childs >= b.c.Cap {
		return -1, ERR_CHILD_FULL
	}

	if pos >= 0 {
		if pos >= len(b.c.Child) && b.c.Child[pos] != nil {
			return -1, ERR_POS_NOT_EMPTY
		}
	}

	if pos < 0 {
		pos = b.c.freeIndex()
		if pos == -1 {
			return -1, ERR_CHILD_FULL
		}
	}

	if !b.CanAdd(pos, g) {
		return -1, ERR_ADD_FAILED
	}

	if err := b.c.Add(pos, g); err != nil {
		return -1, err
	}

	g.Behavior().SetParent(b.gameobject)
	g.Behavior().SetParentIndex(pos)
	return pos, nil
}

// RemoveChild 移除一个对象
func (b *BaseBehavior) RemoveChild(pos int, g GameObject) error {

	if g.Behavior().Parent() != b.gameobject {
		return fmt.Errorf("parent not equal")
	}

	if g.Behavior().ParentIndex() != pos {
		return fmt.Errorf("container pos not equal")
	}

	// TODO: 事件回调 移除前
	if err := b.c.Remove(pos, g); err != nil {
		return err
	}

	g.Behavior().SetParent(nil)
	g.Behavior().SetParentIndex(-1)
	// TODO: 事件回调 移除后

	return nil
}

// ChildAt 取子对象
func (b *BaseBehavior) ChildAt(pos int) GameObject {
	return b.c.ChildAt(pos)
}

// ChildAtIf 取子对象接口
func (b *BaseBehavior) ChildAtIf(pos int) interface{} {
	g := b.c.ChildAt(pos)
	if g == nil {
		return nil
	}

	return g
}

// ChildCount 子对象数量
func (b *BaseBehavior) ChildCount() int {
	return b.c.Childs
}

// FirstChild 获取第一个对象，返回的索引可以用于迭代
func (b *BaseBehavior) FirstChild() (int, GameObject) {
	return b.c.FirstChild()
}

// NextChild 获取下一个对象，和FirstChild配套使用。index为FirstChild返回的索引
func (b *BaseBehavior) NextChild(index int) (int, interface{}) {
	return b.c.NextChild(index)
}

// FirstChildIf 获取第一个对象，返回interface类型，返回的索引可以用于迭代
func (b *BaseBehavior) FirstChildIf() (int, interface{}) {
	i, g := b.c.FirstChild()
	if g == nil {
		return i, nil
	}

	return i, g
}

// NextChildIf 获取下一个对象，和FirstChildIf配套使用。index为FirstChildIf返回的索引
func (b *BaseBehavior) NextChildIf(index int) (int, interface{}) {
	i, g := b.c.NextChild(index)
	if g == nil {
		return i, nil
	}
	return i, g
}
