package goja

import "reflect"

type lazyObject struct {
	val    *Object
	create func(*Object) objectImpl
}

func (o *lazyObject) className() string {
	obj := o.create(o.val)
	o.val.self = obj
	return obj.className()
}

func (o *lazyObject) get(n Value, receiver Value) Value {
	obj := o.create(o.val)
	o.val.self = obj
	return obj.get(n, receiver)
}

func (o *lazyObject) getProp(n Value) Value {
	obj := o.create(o.val)
	o.val.self = obj
	return obj.getProp(n)
}

func (o *lazyObject) getPropStr(name string) Value {
	obj := o.create(o.val)
	o.val.self = obj
	return obj.getPropStr(name)
}

func (o *lazyObject) getStr(name string, receiver Value) Value {
	obj := o.create(o.val)
	o.val.self = obj
	return obj.getStr(name, receiver)
}

func (o *lazyObject) getOwnPropStr(name string) Value {
	obj := o.create(o.val)
	o.val.self = obj
	return obj.getOwnPropStr(name)
}

func (o *lazyObject) getOwnProp(name Value) Value {
	obj := o.create(o.val)
	o.val.self = obj
	return obj.getOwnProp(name)
}

func (o *lazyObject) put(n Value, val Value, throw bool) {
	obj := o.create(o.val)
	o.val.self = obj
	obj.put(n, val, throw)
}

func (o *lazyObject) putStr(name string, val Value, throw bool) {
	obj := o.create(o.val)
	o.val.self = obj
	obj.putStr(name, val, throw)
}

func (o *lazyObject) set(p, v, receiver Value, throw bool) bool {
	obj := o.create(o.val)
	o.val.self = obj
	return obj.set(p, v, receiver, throw)
}

func (o *lazyObject) setStr(p string, v, receiver Value, throw bool) bool {
	obj := o.create(o.val)
	o.val.self = obj
	return obj.setStr(p, v, receiver, throw)
}

func (o *lazyObject) hasProperty(n Value) bool {
	obj := o.create(o.val)
	o.val.self = obj
	return obj.hasProperty(n)
}

func (o *lazyObject) hasPropertyStr(name string) bool {
	obj := o.create(o.val)
	o.val.self = obj
	return obj.hasPropertyStr(name)
}

func (o *lazyObject) hasOwnProperty(n Value) bool {
	obj := o.create(o.val)
	o.val.self = obj
	return obj.hasOwnProperty(n)
}

func (o *lazyObject) hasOwnPropertyStr(name string) bool {
	obj := o.create(o.val)
	o.val.self = obj
	return obj.hasOwnPropertyStr(name)
}

func (o *lazyObject) _putProp(name string, value Value, writable, enumerable, configurable bool) Value {
	obj := o.create(o.val)
	o.val.self = obj
	return obj._putProp(name, value, writable, enumerable, configurable)
}

func (o *lazyObject) defineOwnProperty(name Value, descr PropertyDescriptor, throw bool) bool {
	obj := o.create(o.val)
	o.val.self = obj
	return obj.defineOwnProperty(name, descr, throw)
}

func (o *lazyObject) toPrimitiveNumber() Value {
	obj := o.create(o.val)
	o.val.self = obj
	return obj.toPrimitiveNumber()
}

func (o *lazyObject) toPrimitiveString() Value {
	obj := o.create(o.val)
	o.val.self = obj
	return obj.toPrimitiveString()
}

func (o *lazyObject) toPrimitive() Value {
	obj := o.create(o.val)
	o.val.self = obj
	return obj.toPrimitive()
}

func (o *lazyObject) assertCallable() (call func(FunctionCall) Value, ok bool) {
	obj := o.create(o.val)
	o.val.self = obj
	return obj.assertCallable()
}

func (o *lazyObject) deleteStr(name string, throw bool) bool {
	obj := o.create(o.val)
	o.val.self = obj
	return obj.deleteStr(name, throw)
}

func (o *lazyObject) delete(name Value, throw bool) bool {
	obj := o.create(o.val)
	o.val.self = obj
	return obj.delete(name, throw)
}

func (o *lazyObject) proto() *Object {
	obj := o.create(o.val)
	o.val.self = obj
	return obj.proto()
}

func (o *lazyObject) hasInstance(v Value) bool {
	obj := o.create(o.val)
	o.val.self = obj
	return obj.hasInstance(v)
}

func (o *lazyObject) isExtensible() bool {
	obj := o.create(o.val)
	o.val.self = obj
	return obj.isExtensible()
}

func (o *lazyObject) preventExtensions(throw bool) bool {
	obj := o.create(o.val)
	o.val.self = obj
	return obj.preventExtensions(throw)
}

func (o *lazyObject) enumerate(all, recursive bool) iterNextFunc {
	obj := o.create(o.val)
	o.val.self = obj
	return obj.enumerate(all, recursive)
}

func (o *lazyObject) _enumerate(recursive bool) iterNextFunc {
	obj := o.create(o.val)
	o.val.self = obj
	return obj._enumerate(recursive)
}

func (o *lazyObject) export() interface{} {
	obj := o.create(o.val)
	o.val.self = obj
	return obj.export()
}

func (o *lazyObject) exportType() reflect.Type {
	obj := o.create(o.val)
	o.val.self = obj
	return obj.exportType()
}

func (o *lazyObject) equal(other objectImpl) bool {
	obj := o.create(o.val)
	o.val.self = obj
	return obj.equal(other)
}

func (o *lazyObject) getOwnSymbols() []Value {
	obj := o.create(o.val)
	o.val.self = obj
	return obj.getOwnSymbols()
}

func (o *lazyObject) setProto(proto *Object, throw bool) bool {
	obj := o.create(o.val)
	o.val.self = obj
	return obj.setProto(proto, throw)
}

func (o *lazyObject) sortLen() int64 {
	obj := o.create(o.val)
	o.val.self = obj
	return obj.sortLen()
}

func (o *lazyObject) sortGet(i int64) Value {
	obj := o.create(o.val)
	o.val.self = obj
	return obj.sortGet(i)
}

func (o *lazyObject) swap(i, j int64) {
	obj := o.create(o.val)
	o.val.self = obj
	obj.swap(i, j)
}
