// Code generated by fastssz. DO NOT EDIT.
// Hash: fdc7d541c3ff5d95bba223e838544d6afd5816ba13b1cd9810d4abb2200096b1
package services

import (
	ssz "github.com/prysmaticlabs/fastssz"
)

// MarshalSSZ ssz marshals the bxCompressedTransaction object
func (b *bxCompressedTransaction) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the bxCompressedTransaction object to a target array
func (b *bxCompressedTransaction) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(5)

	// Field (0) 'IsFullTransaction'
	dst = ssz.MarshalBool(dst, b.IsFullTransaction)

	// Offset (1) 'Transaction'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.Transaction)

	// Field (1) 'Transaction'
	if size := len(b.Transaction); size > 1073741824 {
		err = ssz.ErrBytesLengthFn("--.Transaction", size, 1073741824)
		return
	}
	dst = append(dst, b.Transaction...)

	return
}

// UnmarshalSSZ ssz unmarshals the bxCompressedTransaction object
func (b *bxCompressedTransaction) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 5 {
		return ssz.ErrSize
	}

	tail := buf
	var o1 uint64

	// Field (0) 'IsFullTransaction'
	b.IsFullTransaction = ssz.UnmarshalBool(buf[0:1])

	// Offset (1) 'Transaction'
	if o1 = ssz.ReadOffset(buf[1:5]); o1 > size {
		return ssz.ErrOffset
	}

	if o1 < 5 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (1) 'Transaction'
	{
		buf = tail[o1:]
		if len(buf) > 1073741824 {
			return ssz.ErrBytesLength
		}
		if cap(b.Transaction) == 0 {
			b.Transaction = make([]byte, 0, len(buf))
		}
		b.Transaction = append(b.Transaction, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the bxCompressedTransaction object
func (b *bxCompressedTransaction) SizeSSZ() (size int) {
	size = 5

	// Field (1) 'Transaction'
	size += len(b.Transaction)

	return
}

// HashTreeRoot ssz hashes the bxCompressedTransaction object
func (b *bxCompressedTransaction) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the bxCompressedTransaction object with a hasher
func (b *bxCompressedTransaction) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'IsFullTransaction'
	hh.PutBool(b.IsFullTransaction)

	// Field (1) 'Transaction'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(b.Transaction))
		if byteLen > 1073741824 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.PutBytes(b.Transaction)
		if ssz.EnableVectorizedHTR {
			hh.MerkleizeWithMixinVectorizedHTR(elemIndx, byteLen, (1073741824+31)/32)
		} else {
			hh.MerkleizeWithMixin(elemIndx, byteLen, (1073741824+31)/32)
		}
	}

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

// MarshalSSZ ssz marshals the BxBlockSSZ object
func (b *BxBlockSSZ) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BxBlockSSZ object to a target array
func (b *BxBlockSSZ) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(16)

	// Offset (0) 'Block'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.Block)

	// Offset (1) 'Txs'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(b.Txs); ii++ {
		offset += 4
		offset += b.Txs[ii].SizeSSZ()
	}

	// Field (2) 'Number'
	dst = ssz.MarshalUint64(dst, b.Number)

	// Field (0) 'Block'
	if size := len(b.Block); size > 367832 {
		err = ssz.ErrBytesLengthFn("--.Block", size, 367832)
		return
	}
	dst = append(dst, b.Block...)

	// Field (1) 'Txs'
	if size := len(b.Txs); size > 1073741825 {
		err = ssz.ErrListTooBigFn("--.Txs", size, 1073741825)
		return
	}
	{
		offset = 4 * len(b.Txs)
		for ii := 0; ii < len(b.Txs); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += b.Txs[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(b.Txs); ii++ {
		if dst, err = b.Txs[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BxBlockSSZ object
func (b *BxBlockSSZ) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 16 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o1 uint64

	// Offset (0) 'Block'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 16 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (1) 'Txs'
	if o1 = ssz.ReadOffset(buf[4:8]); o1 > size || o0 > o1 {
		return ssz.ErrOffset
	}

	// Field (2) 'Number'
	b.Number = ssz.UnmarshallUint64(buf[8:16])

	// Field (0) 'Block'
	{
		buf = tail[o0:o1]
		if len(buf) > 367832 {
			return ssz.ErrBytesLength
		}
		if cap(b.Block) == 0 {
			b.Block = make([]byte, 0, len(buf))
		}
		b.Block = append(b.Block, buf...)
	}

	// Field (1) 'Txs'
	{
		buf = tail[o1:]
		num, err := ssz.DecodeDynamicLength(buf, 1073741825)
		if err != nil {
			return err
		}
		b.Txs = make([]*bxCompressedTransaction, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if b.Txs[indx] == nil {
				b.Txs[indx] = new(bxCompressedTransaction)
			}
			if err = b.Txs[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BxBlockSSZ object
func (b *BxBlockSSZ) SizeSSZ() (size int) {
	size = 16

	// Field (0) 'Block'
	size += len(b.Block)

	// Field (1) 'Txs'
	for ii := 0; ii < len(b.Txs); ii++ {
		size += 4
		size += b.Txs[ii].SizeSSZ()
	}

	return
}

// HashTreeRoot ssz hashes the BxBlockSSZ object
func (b *BxBlockSSZ) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BxBlockSSZ object with a hasher
func (b *BxBlockSSZ) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Block'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(b.Block))
		if byteLen > 367832 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.PutBytes(b.Block)
		if ssz.EnableVectorizedHTR {
			hh.MerkleizeWithMixinVectorizedHTR(elemIndx, byteLen, (367832+31)/32)
		} else {
			hh.MerkleizeWithMixin(elemIndx, byteLen, (367832+31)/32)
		}
	}

	// Field (1) 'Txs'
	{
		subIndx := hh.Index()
		num := uint64(len(b.Txs))
		if num > 1073741825 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range b.Txs {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		if ssz.EnableVectorizedHTR {
			hh.MerkleizeWithMixinVectorizedHTR(subIndx, num, 1073741825)
		} else {
			hh.MerkleizeWithMixin(subIndx, num, 1073741825)
		}
	}

	// Field (2) 'Number'
	hh.PutUint64(b.Number)

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}
