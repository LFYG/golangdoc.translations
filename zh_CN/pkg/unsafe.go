// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
	Package unsafe contains operations that step around the type safety of Go programs.

	Packages that import unsafe may be non-portable and are not protected by the
	Go 1 compatibility guidelines.
*/

/*
	unsafe �����й���Go�������Ͱ�ȫ�����в���.
*/
package unsafe

// ArbitraryType is here for the purposes of documentation only and is not actually
// part of the unsafe package.  It represents the type of an arbitrary Go expression.

// ArbitraryType �ڴ˴�ֻ�����ĵ�Ŀ�ģ���ʵ���ϲ����� unsafe ����һ���֡�
// ����������һ��Go���ʽ�����͡�
type ArbitraryType int

// Pointer represents a pointer to an arbitrary type.  There are four special operations
// available for type Pointer that are not available for other types.
//	1) A pointer value of any type can be converted to a Pointer.
//	2) A Pointer can be converted to a pointer value of any type.
//	3) A uintptr can be converted to a Pointer.
//	4) A Pointer can be converted to a uintptr.
// Pointer therefore allows a program to defeat the type system and read and write
// arbitrary memory. It should be used with extreme care.

// Pointer ����һ��ָ���������͵�ָ�롣
// ����������Ĳ�������������ָ������������������͡�
//	1) �������͵�ָ��ֵ����ת��Ϊ Pointer��
//	2) Pointer ����ת��Ϊ�������͵�ָ��ֵ��
//	3) uintptr ����ת��Ϊ Pointer��
//	4) Pointer ����ת��Ϊ uintptr��
// ��� Pointer ��������������ϵͳ����д�����ڴ档��Ӧ�����õ÷ǳ�С�ġ�
type Pointer *ArbitraryType

// Sizeof returns the size in bytes occupied by the value v.  The size is that of the
// "top level" of the value only.  For instance, if v is a slice, it returns the size of
// the slice descriptor, not the size of the memory referenced by the slice.

// Sizeof ���ر�ֵ v ��ռ�õ��ֽڴ�С��
// �ô�Сֻ�����������ֵ�����磬�� v ��һ����Ƭ�����᷵�ظ���Ƭ�������Ĵ�С��
// ���Ǹ���Ƭ���õ��ڴ��С��
func Sizeof(v ArbitraryType) uintptr

// Offsetof returns the offset within the struct of the field represented by v,
// which must be of the form structValue.field.  In other words, it returns the
// number of bytes between the start of the struct and the start of the field.

// Offsetof ������ v ������Ľṹ���ֶε�ƫ�ƣ�������Ϊ structValue.field ����ʽ��
// ���仰˵�������ظýṹ��ʼ������ֶ���ʼ��֮����ֽ�����
func Offsetof(v ArbitraryType) uintptr

// Alignof returns the alignment of the value v.  It is the maximum value m such
// that the address of a variable with the type of v will always be zero mod m.
// If v is of the form structValue.field, it returns the alignment of field f within struct object obj.

// Alignof ���� v ֵ�Ķ��뷽ʽ��
// �䷵��ֵ m ������� v �����͵�ַ�� m ȡģΪ 0 �����ֵ���� v �� structValue.field
// ����ʽ�����᷵���ֶ� f ������Ӧ�ṹ���� obj �еĶ��뷽ʽ��
func Alignof(v ArbitraryType) uintptr
