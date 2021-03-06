package types

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *BlockExtra) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "key":
			err = dc.ReadExtension(&z.KeyHash)
			if err != nil {
				return
			}
		case "abi":
			z.Abi, err = dc.ReadBytes(z.Abi)
			if err != nil {
				return
			}
		case "issuer":
			err = dc.ReadExtension(&z.Issuer)
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *BlockExtra) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "key"
	err = en.Append(0x83, 0xa3, 0x6b, 0x65, 0x79)
	if err != nil {
		return
	}
	err = en.WriteExtension(&z.KeyHash)
	if err != nil {
		return
	}
	// write "abi"
	err = en.Append(0xa3, 0x61, 0x62, 0x69)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.Abi)
	if err != nil {
		return
	}
	// write "issuer"
	err = en.Append(0xa6, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72)
	if err != nil {
		return
	}
	err = en.WriteExtension(&z.Issuer)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *BlockExtra) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "key"
	o = append(o, 0x83, 0xa3, 0x6b, 0x65, 0x79)
	o, err = msgp.AppendExtension(o, &z.KeyHash)
	if err != nil {
		return
	}
	// string "abi"
	o = append(o, 0xa3, 0x61, 0x62, 0x69)
	o = msgp.AppendBytes(o, z.Abi)
	// string "issuer"
	o = append(o, 0xa6, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72)
	o, err = msgp.AppendExtension(o, &z.Issuer)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BlockExtra) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "key":
			bts, err = msgp.ReadExtensionBytes(bts, &z.KeyHash)
			if err != nil {
				return
			}
		case "abi":
			z.Abi, bts, err = msgp.ReadBytesBytes(bts, z.Abi)
			if err != nil {
				return
			}
		case "issuer":
			bts, err = msgp.ReadExtensionBytes(bts, &z.Issuer)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *BlockExtra) Msgsize() (s int) {
	s = 1 + 4 + msgp.ExtensionPrefixSize + z.KeyHash.Len() + 4 + msgp.BytesPrefixSize + len(z.Abi) + 7 + msgp.ExtensionPrefixSize + z.Issuer.Len()
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Enum) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 string
		zb0001, err = dc.ReadString()
		if err != nil {
			return
		}
		(*z) = parseString(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Enum) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteString((Enum).String(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Enum) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendString(o, (Enum).String(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Enum) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 string
		zb0001, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			return
		}
		(*z) = parseString(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Enum) Msgsize() (s int) {
	s = msgp.StringPrefixSize + len((Enum).String(z))
	return
}

// DecodeMsg implements msgp.Decodable
func (sc *SmartContractBlock) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "type":
			{
				var zb0002 string
				zb0002, err = dc.ReadString()
				if err != nil {
					return
				}
				sc.Type = parseString(zb0002)
			}
		case "previous":
			err = dc.ReadExtension(&sc.PreviousHash)
			if err != nil {
				return
			}
		case "representative":
			err = dc.ReadExtension(&sc.Representative)
			if err != nil {
				return
			}
		case "balance":
			err = dc.ReadExtension(&sc.Balance)
			if err != nil {
				return
			}
		case "link":
			err = dc.ReadExtension(&sc.Link)
			if err != nil {
				return
			}
		case "signature":
			err = dc.ReadExtension(&sc.Signature)
			if err != nil {
				return
			}
		case "extra":
			err = dc.ReadExtension(&sc.Extra)
			if err != nil {
				return
			}
		case "work":
			err = dc.ReadExtension(&sc.Work)
			if err != nil {
				return
			}
		case "owner":
			err = dc.ReadExtension(&sc.Owner)
			if err != nil {
				return
			}
		case "issuer":
			err = dc.ReadExtension(&sc.Issuer)
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (sc *SmartContractBlock) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 10
	// write "type"
	err = en.Append(0x8a, 0xa4, 0x74, 0x79, 0x70, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString((Enum).String(sc.Type))
	if err != nil {
		return
	}
	// write "previous"
	err = en.Append(0xa8, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73)
	if err != nil {
		return
	}
	err = en.WriteExtension(&sc.PreviousHash)
	if err != nil {
		return
	}
	// write "representative"
	err = en.Append(0xae, 0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x76, 0x65)
	if err != nil {
		return
	}
	err = en.WriteExtension(&sc.Representative)
	if err != nil {
		return
	}
	// write "balance"
	err = en.Append(0xa7, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65)
	if err != nil {
		return
	}
	err = en.WriteExtension(&sc.Balance)
	if err != nil {
		return
	}
	// write "link"
	err = en.Append(0xa4, 0x6c, 0x69, 0x6e, 0x6b)
	if err != nil {
		return
	}
	err = en.WriteExtension(&sc.Link)
	if err != nil {
		return
	}
	// write "signature"
	err = en.Append(0xa9, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65)
	if err != nil {
		return
	}
	err = en.WriteExtension(&sc.Signature)
	if err != nil {
		return
	}
	// write "extra"
	err = en.Append(0xa5, 0x65, 0x78, 0x74, 0x72, 0x61)
	if err != nil {
		return
	}
	err = en.WriteExtension(&sc.Extra)
	if err != nil {
		return
	}
	// write "work"
	err = en.Append(0xa4, 0x77, 0x6f, 0x72, 0x6b)
	if err != nil {
		return
	}
	err = en.WriteExtension(&sc.Work)
	if err != nil {
		return
	}
	// write "owner"
	err = en.Append(0xa5, 0x6f, 0x77, 0x6e, 0x65, 0x72)
	if err != nil {
		return
	}
	err = en.WriteExtension(&sc.Owner)
	if err != nil {
		return
	}
	// write "issuer"
	err = en.Append(0xa6, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72)
	if err != nil {
		return
	}
	err = en.WriteExtension(&sc.Issuer)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (sc *SmartContractBlock) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, sc.Msgsize())
	// map header, size 10
	// string "type"
	o = append(o, 0x8a, 0xa4, 0x74, 0x79, 0x70, 0x65)
	o = msgp.AppendString(o, (Enum).String(sc.Type))
	// string "previous"
	o = append(o, 0xa8, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73)
	o, err = msgp.AppendExtension(o, &sc.PreviousHash)
	if err != nil {
		return
	}
	// string "representative"
	o = append(o, 0xae, 0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x76, 0x65)
	o, err = msgp.AppendExtension(o, &sc.Representative)
	if err != nil {
		return
	}
	// string "balance"
	o = append(o, 0xa7, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65)
	o, err = msgp.AppendExtension(o, &sc.Balance)
	if err != nil {
		return
	}
	// string "link"
	o = append(o, 0xa4, 0x6c, 0x69, 0x6e, 0x6b)
	o, err = msgp.AppendExtension(o, &sc.Link)
	if err != nil {
		return
	}
	// string "signature"
	o = append(o, 0xa9, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65)
	o, err = msgp.AppendExtension(o, &sc.Signature)
	if err != nil {
		return
	}
	// string "extra"
	o = append(o, 0xa5, 0x65, 0x78, 0x74, 0x72, 0x61)
	o, err = msgp.AppendExtension(o, &sc.Extra)
	if err != nil {
		return
	}
	// string "work"
	o = append(o, 0xa4, 0x77, 0x6f, 0x72, 0x6b)
	o, err = msgp.AppendExtension(o, &sc.Work)
	if err != nil {
		return
	}
	// string "owner"
	o = append(o, 0xa5, 0x6f, 0x77, 0x6e, 0x65, 0x72)
	o, err = msgp.AppendExtension(o, &sc.Owner)
	if err != nil {
		return
	}
	// string "issuer"
	o = append(o, 0xa6, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72)
	o, err = msgp.AppendExtension(o, &sc.Issuer)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (sc *SmartContractBlock) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "type":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				sc.Type = parseString(zb0002)
			}
		case "previous":
			bts, err = msgp.ReadExtensionBytes(bts, &sc.PreviousHash)
			if err != nil {
				return
			}
		case "representative":
			bts, err = msgp.ReadExtensionBytes(bts, &sc.Representative)
			if err != nil {
				return
			}
		case "balance":
			bts, err = msgp.ReadExtensionBytes(bts, &sc.Balance)
			if err != nil {
				return
			}
		case "link":
			bts, err = msgp.ReadExtensionBytes(bts, &sc.Link)
			if err != nil {
				return
			}
		case "signature":
			bts, err = msgp.ReadExtensionBytes(bts, &sc.Signature)
			if err != nil {
				return
			}
		case "extra":
			bts, err = msgp.ReadExtensionBytes(bts, &sc.Extra)
			if err != nil {
				return
			}
		case "work":
			bts, err = msgp.ReadExtensionBytes(bts, &sc.Work)
			if err != nil {
				return
			}
		case "owner":
			bts, err = msgp.ReadExtensionBytes(bts, &sc.Owner)
			if err != nil {
				return
			}
		case "issuer":
			bts, err = msgp.ReadExtensionBytes(bts, &sc.Issuer)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (sc *SmartContractBlock) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len((Enum).String(sc.Type)) + 9 + msgp.ExtensionPrefixSize + sc.PreviousHash.Len() + 15 + msgp.ExtensionPrefixSize + sc.Representative.Len() + 8 + msgp.ExtensionPrefixSize + sc.Balance.Len() + 5 + msgp.ExtensionPrefixSize + sc.Link.Len() + 10 + msgp.ExtensionPrefixSize + sc.Signature.Len() + 6 + msgp.ExtensionPrefixSize + sc.Extra.Len() + 5 + msgp.ExtensionPrefixSize + sc.Work.Len() + 6 + msgp.ExtensionPrefixSize + sc.Owner.Len() + 7 + msgp.ExtensionPrefixSize + sc.Issuer.Len()
	return
}

// DecodeMsg implements msgp.Decodable
func (z *StateBlock) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "type":
			{
				var zb0002 string
				zb0002, err = dc.ReadString()
				if err != nil {
					return
				}
				z.Type = parseString(zb0002)
			}
		case "addresses":
			err = dc.ReadExtension(&z.Address)
			if err != nil {
				return
			}
		case "previous":
			err = dc.ReadExtension(&z.PreviousHash)
			if err != nil {
				return
			}
		case "representative":
			err = dc.ReadExtension(&z.Representative)
			if err != nil {
				return
			}
		case "balance":
			err = dc.ReadExtension(&z.Balance)
			if err != nil {
				return
			}
		case "link":
			err = dc.ReadExtension(&z.Link)
			if err != nil {
				return
			}
		case "signature":
			err = dc.ReadExtension(&z.Signature)
			if err != nil {
				return
			}
		case "token":
			err = dc.ReadExtension(&z.Token)
			if err != nil {
				return
			}
		case "extra":
			err = dc.ReadExtension(&z.Extra)
			if err != nil {
				return
			}
		case "work":
			err = dc.ReadExtension(&z.Work)
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *StateBlock) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 10
	// write "type"
	err = en.Append(0x8a, 0xa4, 0x74, 0x79, 0x70, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString((Enum).String(z.Type))
	if err != nil {
		return
	}
	// write "addresses"
	err = en.Append(0xa9, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73)
	if err != nil {
		return
	}
	err = en.WriteExtension(&z.Address)
	if err != nil {
		return
	}
	// write "previous"
	err = en.Append(0xa8, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73)
	if err != nil {
		return
	}
	err = en.WriteExtension(&z.PreviousHash)
	if err != nil {
		return
	}
	// write "representative"
	err = en.Append(0xae, 0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x76, 0x65)
	if err != nil {
		return
	}
	err = en.WriteExtension(&z.Representative)
	if err != nil {
		return
	}
	// write "balance"
	err = en.Append(0xa7, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65)
	if err != nil {
		return
	}
	err = en.WriteExtension(&z.Balance)
	if err != nil {
		return
	}
	// write "link"
	err = en.Append(0xa4, 0x6c, 0x69, 0x6e, 0x6b)
	if err != nil {
		return
	}
	err = en.WriteExtension(&z.Link)
	if err != nil {
		return
	}
	// write "signature"
	err = en.Append(0xa9, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65)
	if err != nil {
		return
	}
	err = en.WriteExtension(&z.Signature)
	if err != nil {
		return
	}
	// write "token"
	err = en.Append(0xa5, 0x74, 0x6f, 0x6b, 0x65, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteExtension(&z.Token)
	if err != nil {
		return
	}
	// write "extra"
	err = en.Append(0xa5, 0x65, 0x78, 0x74, 0x72, 0x61)
	if err != nil {
		return
	}
	err = en.WriteExtension(&z.Extra)
	if err != nil {
		return
	}
	// write "work"
	err = en.Append(0xa4, 0x77, 0x6f, 0x72, 0x6b)
	if err != nil {
		return
	}
	err = en.WriteExtension(&z.Work)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *StateBlock) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 10
	// string "type"
	o = append(o, 0x8a, 0xa4, 0x74, 0x79, 0x70, 0x65)
	o = msgp.AppendString(o, (Enum).String(z.Type))
	// string "addresses"
	o = append(o, 0xa9, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73)
	o, err = msgp.AppendExtension(o, &z.Address)
	if err != nil {
		return
	}
	// string "previous"
	o = append(o, 0xa8, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73)
	o, err = msgp.AppendExtension(o, &z.PreviousHash)
	if err != nil {
		return
	}
	// string "representative"
	o = append(o, 0xae, 0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x76, 0x65)
	o, err = msgp.AppendExtension(o, &z.Representative)
	if err != nil {
		return
	}
	// string "balance"
	o = append(o, 0xa7, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65)
	o, err = msgp.AppendExtension(o, &z.Balance)
	if err != nil {
		return
	}
	// string "link"
	o = append(o, 0xa4, 0x6c, 0x69, 0x6e, 0x6b)
	o, err = msgp.AppendExtension(o, &z.Link)
	if err != nil {
		return
	}
	// string "signature"
	o = append(o, 0xa9, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65)
	o, err = msgp.AppendExtension(o, &z.Signature)
	if err != nil {
		return
	}
	// string "token"
	o = append(o, 0xa5, 0x74, 0x6f, 0x6b, 0x65, 0x6e)
	o, err = msgp.AppendExtension(o, &z.Token)
	if err != nil {
		return
	}
	// string "extra"
	o = append(o, 0xa5, 0x65, 0x78, 0x74, 0x72, 0x61)
	o, err = msgp.AppendExtension(o, &z.Extra)
	if err != nil {
		return
	}
	// string "work"
	o = append(o, 0xa4, 0x77, 0x6f, 0x72, 0x6b)
	o, err = msgp.AppendExtension(o, &z.Work)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *StateBlock) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "type":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Type = parseString(zb0002)
			}
		case "addresses":
			bts, err = msgp.ReadExtensionBytes(bts, &z.Address)
			if err != nil {
				return
			}
		case "previous":
			bts, err = msgp.ReadExtensionBytes(bts, &z.PreviousHash)
			if err != nil {
				return
			}
		case "representative":
			bts, err = msgp.ReadExtensionBytes(bts, &z.Representative)
			if err != nil {
				return
			}
		case "balance":
			bts, err = msgp.ReadExtensionBytes(bts, &z.Balance)
			if err != nil {
				return
			}
		case "link":
			bts, err = msgp.ReadExtensionBytes(bts, &z.Link)
			if err != nil {
				return
			}
		case "signature":
			bts, err = msgp.ReadExtensionBytes(bts, &z.Signature)
			if err != nil {
				return
			}
		case "token":
			bts, err = msgp.ReadExtensionBytes(bts, &z.Token)
			if err != nil {
				return
			}
		case "extra":
			bts, err = msgp.ReadExtensionBytes(bts, &z.Extra)
			if err != nil {
				return
			}
		case "work":
			bts, err = msgp.ReadExtensionBytes(bts, &z.Work)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *StateBlock) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len((Enum).String(z.Type)) + 10 + msgp.ExtensionPrefixSize + z.Address.Len() + 9 + msgp.ExtensionPrefixSize + z.PreviousHash.Len() + 15 + msgp.ExtensionPrefixSize + z.Representative.Len() + 8 + msgp.ExtensionPrefixSize + z.Balance.Len() + 5 + msgp.ExtensionPrefixSize + z.Link.Len() + 10 + msgp.ExtensionPrefixSize + z.Signature.Len() + 6 + msgp.ExtensionPrefixSize + z.Token.Len() + 6 + msgp.ExtensionPrefixSize + z.Extra.Len() + 5 + msgp.ExtensionPrefixSize + z.Work.Len()
	return
}
