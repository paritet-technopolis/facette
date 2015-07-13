// Code generated by protoc-gen-gogo.
// source: request_vote_responses.proto
// DO NOT EDIT!

package protobuf

import proto "code.google.com/p/gogoprotobuf/proto"
import math "math"

// discarding unused import gogoproto "code.google.com/p/gogoprotobuf/gogoproto/gogo.pb"

import io4 "io"
import fmt16 "fmt"
import code_google_com_p_gogoprotobuf_proto8 "code.google.com/p/gogoprotobuf/proto"

import fmt17 "fmt"
import strings8 "strings"
import reflect8 "reflect"

import fmt18 "fmt"
import strings9 "strings"
import code_google_com_p_gogoprotobuf_proto9 "code.google.com/p/gogoprotobuf/proto"
import sort4 "sort"
import strconv4 "strconv"
import reflect9 "reflect"

import fmt19 "fmt"
import bytes4 "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type RequestVoteResponse struct {
	Term             *uint64 `protobuf:"varint,1,req" json:"Term,omitempty"`
	VoteGranted      *bool   `protobuf:"varint,2,req" json:"VoteGranted,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RequestVoteResponse) Reset()      { *m = RequestVoteResponse{} }
func (*RequestVoteResponse) ProtoMessage() {}

func (m *RequestVoteResponse) GetTerm() uint64 {
	if m != nil && m.Term != nil {
		return *m.Term
	}
	return 0
}

func (m *RequestVoteResponse) GetVoteGranted() bool {
	if m != nil && m.VoteGranted != nil {
		return *m.VoteGranted
	}
	return false
}

func init() {
}
func (m *RequestVoteResponse) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io4.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt16.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io4.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Term = &v
		case 2:
			if wireType != 0 {
				return fmt16.Errorf("proto: wrong wireType = %d for field VoteGranted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io4.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.VoteGranted = &b
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto8.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io4.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (this *RequestVoteResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings8.Join([]string{`&RequestVoteResponse{`,
		`Term:` + valueToStringRequestVoteResponses(this.Term) + `,`,
		`VoteGranted:` + valueToStringRequestVoteResponses(this.VoteGranted) + `,`,
		`XXX_unrecognized:` + fmt17.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringRequestVoteResponses(v interface{}) string {
	rv := reflect8.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect8.Indirect(rv).Interface()
	return fmt17.Sprintf("*%v", pv)
}
func (m *RequestVoteResponse) Size() (n int) {
	var l int
	_ = l
	if m.Term != nil {
		n += 1 + sovRequestVoteResponses(uint64(*m.Term))
	}
	if m.VoteGranted != nil {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRequestVoteResponses(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRequestVoteResponses(x uint64) (n int) {
	return sovRequestVoteResponses(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func NewPopulatedRequestVoteResponse(r randyRequestVoteResponses, easy bool) *RequestVoteResponse {
	this := &RequestVoteResponse{}
	v1 := uint64(r.Uint32())
	this.Term = &v1
	v2 := bool(r.Intn(2) == 0)
	this.VoteGranted = &v2
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedRequestVoteResponses(r, 3)
	}
	return this
}

type randyRequestVoteResponses interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneRequestVoteResponses(r randyRequestVoteResponses) rune {
	res := rune(r.Uint32() % 1112064)
	if 55296 <= res {
		res += 2047
	}
	return res
}
func randStringRequestVoteResponses(r randyRequestVoteResponses) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneRequestVoteResponses(r)
	}
	return string(tmps)
}
func randUnrecognizedRequestVoteResponses(r randyRequestVoteResponses, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldRequestVoteResponses(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldRequestVoteResponses(data []byte, r randyRequestVoteResponses, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateRequestVoteResponses(data, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		data = encodeVarintPopulateRequestVoteResponses(data, uint64(v4))
	case 1:
		data = encodeVarintPopulateRequestVoteResponses(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateRequestVoteResponses(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateRequestVoteResponses(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateRequestVoteResponses(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateRequestVoteResponses(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *RequestVoteResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RequestVoteResponse) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Term != nil {
		data[i] = 0x8
		i++
		i = encodeVarintRequestVoteResponses(data, i, uint64(*m.Term))
	}
	if m.VoteGranted != nil {
		data[i] = 0x10
		i++
		if *m.VoteGranted {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func encodeFixed64RequestVoteResponses(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32RequestVoteResponses(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintRequestVoteResponses(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (this *RequestVoteResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings9.Join([]string{`&protobuf.RequestVoteResponse{` + `Term:` + valueToGoStringRequestVoteResponses(this.Term, "uint64"), `VoteGranted:` + valueToGoStringRequestVoteResponses(this.VoteGranted, "bool"), `XXX_unrecognized:` + fmt18.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func valueToGoStringRequestVoteResponses(v interface{}, typ string) string {
	rv := reflect9.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect9.Indirect(rv).Interface()
	return fmt18.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringRequestVoteResponses(e map[int32]code_google_com_p_gogoprotobuf_proto9.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort4.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv4.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings9.Join(ss, ",") + "}"
	return s
}
func (this *RequestVoteResponse) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt19.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*RequestVoteResponse)
	if !ok {
		return fmt19.Errorf("that is not of type *RequestVoteResponse")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt19.Errorf("that is type *RequestVoteResponse but is nil && this != nil")
	} else if this == nil {
		return fmt19.Errorf("that is type *RequestVoteResponsebut is not nil && this == nil")
	}
	if this.Term != nil && that1.Term != nil {
		if *this.Term != *that1.Term {
			return fmt19.Errorf("Term this(%v) Not Equal that(%v)", *this.Term, *that1.Term)
		}
	} else if this.Term != nil {
		return fmt19.Errorf("this.Term == nil && that.Term != nil")
	} else if that1.Term != nil {
		return fmt19.Errorf("Term this(%v) Not Equal that(%v)", this.Term, that1.Term)
	}
	if this.VoteGranted != nil && that1.VoteGranted != nil {
		if *this.VoteGranted != *that1.VoteGranted {
			return fmt19.Errorf("VoteGranted this(%v) Not Equal that(%v)", *this.VoteGranted, *that1.VoteGranted)
		}
	} else if this.VoteGranted != nil {
		return fmt19.Errorf("this.VoteGranted == nil && that.VoteGranted != nil")
	} else if that1.VoteGranted != nil {
		return fmt19.Errorf("VoteGranted this(%v) Not Equal that(%v)", this.VoteGranted, that1.VoteGranted)
	}
	if !bytes4.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt19.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *RequestVoteResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*RequestVoteResponse)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Term != nil && that1.Term != nil {
		if *this.Term != *that1.Term {
			return false
		}
	} else if this.Term != nil {
		return false
	} else if that1.Term != nil {
		return false
	}
	if this.VoteGranted != nil && that1.VoteGranted != nil {
		if *this.VoteGranted != *that1.VoteGranted {
			return false
		}
	} else if this.VoteGranted != nil {
		return false
	} else if that1.VoteGranted != nil {
		return false
	}
	if !bytes4.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}