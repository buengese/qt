// +build asteroid,!minimal

// this is a very stupid asteroid specific workaround to provide these structures as stubs to svg, sql, quick and multimedia
// classes bases on this are blacklisted in these modules but still generate stubs depending on this stuff being defined ugh.
// need to figure out a better way to prvent generation of these stubs and remove the includes/imports in the relevant modules
package widgets

import (
	"unsafe"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/internal"
)

type QWidget struct {
	core.QObject
	gui.QPaintDevice
}

type QWidget_ITF interface {
	core.QObject_ITF
	gui.QPaintDevice_ITF
	QWidget_PTR() *QWidget
}

func (ptr *QWidget) QWidget_PTR() *QWidget {
	return ptr
}

func (ptr *QWidget) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QWidget) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
		ptr.QPaintDevice_PTR().SetPointer(p)
	}
}

func PointerFromQWidget(ptr QWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func NewQWidgetFromPointer(ptr unsafe.Pointer) (n *QWidget) {
	n = new(QWidget)
	n.SetPointer(ptr)
	return
}

type QGraphicsObject struct {
	core.QObject
	QGraphicsItem
}

type QGraphicsObject_ITF interface {
	core.QObject_ITF
	QGraphicsItem_ITF
	QGraphicsObject_PTR() *QGraphicsObject
}

func (ptr *QGraphicsObject) QGraphicsObject_PTR() *QGraphicsObject {
	return ptr
}

func (ptr *QGraphicsObject) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QGraphicsObject) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
		ptr.QGraphicsItem_PTR().SetPointer(p)
	}
}

func PointerFromQGraphicsObject(ptr QGraphicsObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsObject_PTR().Pointer()
	}
	return nil
}

type QGraphicsItem struct {
	internal.Internal
}

type QGraphicsItem_ITF interface {
	QGraphicsItem_PTR() *QGraphicsItem
}

func (ptr *QGraphicsItem) QGraphicsItem_PTR() *QGraphicsItem {
	return ptr
}

func (ptr *QGraphicsItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QGraphicsItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQGraphicsItem(ptr QGraphicsItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsItem_PTR().Pointer()
	}
	return nil
}

type QItemDelegate struct {
	QAbstractItemDelegate
}

type QItemDelegate_ITF interface {
	QAbstractItemDelegate_ITF
	QItemDelegate_PTR() *QItemDelegate
}

func (ptr *QItemDelegate) QItemDelegate_PTR() *QItemDelegate {
	return ptr
}

func (ptr *QItemDelegate) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractItemDelegate_PTR().Pointer()
	}
	return nil
}

func (ptr *QItemDelegate) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractItemDelegate_PTR().SetPointer(p)
	}
}

func PointerFromQItemDelegate(ptr QItemDelegate_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QItemDelegate_PTR().Pointer()
	}
	return nil
}

type QAbstractItemDelegate struct {
	core.QObject
}

type QAbstractItemDelegate_ITF interface {
	core.QObject_ITF
	QAbstractItemDelegate_PTR() *QAbstractItemDelegate
}

func (ptr *QAbstractItemDelegate) QAbstractItemDelegate_PTR() *QAbstractItemDelegate {
	return ptr
}

func (ptr *QAbstractItemDelegate) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QAbstractItemDelegate) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQAbstractItemDelegate(ptr QAbstractItemDelegate_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractItemDelegate_PTR().Pointer()
	}
	return nil
}
