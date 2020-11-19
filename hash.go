package main

import (
	"bytes"
	"fmt"
	"os"
	"runtime/pprof"
	"runtime/trace"
	"strconv"
	"time"
	"unsafe"
)

func main() {
	calcs := map[string]func(){
		"calcV1":calcV1,
		"calcV2":calcV2,
		"calcV3":calcV3,
		"calcV4":calcV4,
		"calcV5":calcV5,
		"calcV6":calcV6,
		"calcV8":calcV8,
		"calcV9":calcV9,
		"calcV10":calcV10,
		"calcV11":calcV11,
		"calcV12":calcV12,
		"calcV13":calcV13,
		"calcV14":calcV14,
		"calcV15":calcV15,
	}

	key := "calcV"+os.Args[len(os.Args) - 2]
	isRecord := os.Args[len(os.Args) - 1] == "t"

	func(key string){
		t := time.Now()
		run(key,calcs[key],isRecord)
		fmt.Println(key,time.Now().Sub(t))
	}(key)

}

const l = 10 * 1 + 90 * 2 + 900 * 3 + 9000 * 4 + 90000 * 5 + 900000*6 + 11 * 1000000
func calcV15()  {
	arr := make(map[string]int64, 1000000)
	//arr := map[string]int64{}
	b := [l]byte{}
	start := 0
	end := 0
	tArr := map[int64][]byte{}
	ib := make([]byte,0,7)
	// 1000000
	for i := int64(0); i < 1000000; i++ {
		start = end
		value := time.Now().Unix()
		ibSub := ib[:0]
		ibSub = strconv.AppendInt(ibSub,i,10)
		for _, b2 := range ibSub {
			b[end] = b2
			end++
		}
		b[end] = '_'
		end++
		tb,ok := tArr[value]
		if !ok {
			tb = make([]byte,0,10)
			tb = strconv.AppendInt(tb,value,10)
			tArr[value] = tb
		}
		for _, b2 := range tb {
			b[end] = b2
			end++
		}
		bs := b[start:end]
		arr[*(*string)(unsafe.Pointer(&bs))] = value
	}
}

func calcV14()  {
	arr := make(map[string]int64, 1000000)
	bArr := [1000000][37]byte{}
	tArr := map[int64][]byte{}
	ib := make([]byte,0,7)
	for i := int64(0); i < 1000000; i++ {
		value := time.Now().Unix()

		b := &bArr[i]
		index := 0

		ibSub := ib[:0]
		ibSub = strconv.AppendInt(ibSub,i,10)
		for _, b2 := range ibSub {
			b[index] = b2
			index++
		}
		b[index] = '_'
		index++
		tb,ok := tArr[value]
		if !ok {
			tb := make([]byte,0,10)
			tb = strconv.AppendInt(tb,value,10)
			tArr[value] = tb
		}
		for _, b2 := range tb {
			b[index] = b2
			index++
		}
		bs := b[:index]
		arr[*(*string)(unsafe.Pointer(&bs))] = value
	}
}

func calcV13()  {
	arr := make(map[string]int64, 1000000)
	bArr := [1000000][37]byte{}
	tArr := map[int64][]byte{}
	ib := make([]byte,0,7)
	for i := int64(0); i < 1000000; i++ {
		b := bArr[i][:0]
		value := time.Now().Unix()

		ibSub := ib[:0]
		ibSub = strconv.AppendInt(ibSub,i,10)
		b = append(b,ibSub...)
		b = append(b,'_')
		tb,ok := tArr[value]
		if !ok {
			tb := make([]byte,0,10)
			tb = strconv.AppendInt(tb,value,10)
			tArr[value] = tb
		}
		b = append(b,tb...)

		arr[*(*string)(unsafe.Pointer(&b))] = value
	}
}

func calcV12()  {
	arr := make(map[string]int64, 1000000)
	bArr := [1000000][37]byte{}

	nb := make([]byte,0,7)
	tb := make([]byte,0,10)

	for i := int64(0); i < 1000000; i++ {
		b := &bArr[i]
		index := 0

		value := time.Now().Unix()

		nbSub := nb[:]
		iStr := strconv.AppendInt(nbSub,i,10)
		for si,sl:=0,len(iStr);si < sl;si++ {
			b[index+si] = iStr[si]
		}
		index += len(iStr)
		b[index] = '_'
		index += 1

		tbSub := tb[:]
		t := strconv.AppendInt(tbSub,value,10)
		for si,sl:=0,len(t);si < sl;si++ {
			b[index+si] = t[si]
		}
		index = len(t)

		bs := b[:index]
		arr[*(*string)(unsafe.Pointer(&bs))] = value
	}
}

func calcV11()  {
	arr := make(map[string]int64, 1000000)
	tArr := map[int64]string{}
	bArr := [1000000][37]byte{}

	nb := make([]byte,0,7)

	for i := int64(0); i < 1000000; i++ {
		b := &bArr[i]
		index := 0

		value := time.Now().Unix()

		nbSub := nb[:]
		iStr := strconv.AppendInt(nbSub,i,10)
		for si,sl:=0,len(iStr);si < sl;si++ {
			b[index+si] = iStr[si]
		}
		index += len(iStr)
		b[index] = '_'
		index += 1

		t,ok := tArr[value]
		if !ok {
			t = strconv.FormatInt(value,10)
			tArr[value] = t
		}
		for si,sl:=0,len(t);si < sl;si++ {
			b[index+si] = t[si]
		}
		index = len(t)

		bs := b[:index]
		arr[*(*string)(unsafe.Pointer(&bs))] = value
	}
}

func calcV10()  {
	arr := make(map[string]int64, 1000000)
	tArr := map[int64]string{}
	bArr := [1000000][37]byte{}
	for i := int64(0); i < 1000000; i++ {
		b := &bArr[i]
		index := 0

		value := time.Now().Unix()
		iStr := strconv.FormatInt(i,10)
		for si,sl:=0,len(iStr);si < sl;si++ {
			b[index+si] = iStr[si]
		}
		index += len(iStr)
		b[index] = '_'
		index += 1

		t,ok := tArr[value]
		if !ok {
			t = strconv.FormatInt(value,10)
			tArr[value] = t
		}
		for si,sl:=0,len(t);si < sl;si++ {
			b[index+si] = t[si]
		}
		index = len(t)

		bs := b[:index]
		arr[*(*string)(unsafe.Pointer(&bs))] = value
	}
}
func calcV9()  {
	arr := make(map[string]int64, 1000000)
	tArr := map[int64]string{}
	for i := int64(0); i < 1000000; i++ {
		b := [37]byte{}
		index := 0

		value := time.Now().Unix()
		iStr := strconv.FormatInt(i,10)
		for si,sl:=0,len(iStr);si < sl;si++ {
			b[index+si] = iStr[si]
		}
		index += len(iStr)
		b[index] = '_'
		index += 1

		t,ok := tArr[value]
		if !ok {
			t = strconv.FormatInt(value,10)
			tArr[value] = t
		}
		for si,sl:=0,len(t);si < sl;si++ {
			b[index+si] = t[si]
		}
		index = len(t)

		bs := b[:index]
		arr[*(*string)(unsafe.Pointer(&bs))] = value
	}
}

func calcV8()  {
	arr := make(map[string]int64, 1000000)
	tArr := map[int64]string{}
	for i := int64(0); i < 1000000; i++ {
		b := [37]byte{}
		index := 0

		value := time.Now().Unix()
		iStr := strconv.FormatInt(i,10)
		copy(b[index:],[]byte(iStr))
		index += len(iStr)
		b[index] = '_'
		index += 1

		t,ok := tArr[value]
		if !ok {
			t = strconv.FormatInt(value,10)
			tArr[value] = t
		}
		copy(b[index:],[]byte(t))
		index = len(t)

		bs := b[:index]
		arr[*(*string)(unsafe.Pointer(&bs))] = value
	}
}

func calcV7()  {
	arr := make(map[string]int64, 1000000)
	tArr := map[int64]string{}
	for i := int64(0); i < 1000000; i++ {
		buf := bytes.Buffer{}
		buf.Grow(37)
		buf.Reset()
		value := time.Now().Unix()
		buf.WriteString(strconv.FormatInt(i,10))
		buf.WriteString("_")
		t,ok := tArr[value]
		if !ok {
			t = strconv.FormatInt(value,10)
			tArr[value] = t
		}
		buf.WriteString(t)
		bs := buf.Bytes()
		arr[*(*string)(unsafe.Pointer(&bs))] = value
	}
}

func calcV6()  {
	arr := make(map[string]int64, 1000000)

	for i := int64(0); i < 1000000; i++ {
		buf := bytes.Buffer{}
		buf.Grow(37)
		buf.Reset()
		value := time.Now().Unix()
		buf.WriteString(strconv.FormatInt(i,10))
		buf.WriteString("_")
		buf.WriteString(strconv.FormatInt(value,10))
		bs := buf.Bytes()
		arr[*(*string)(unsafe.Pointer(&bs))] = value
	}

}
func calcV5()  {
	arr := make(map[string]int64, 1000000)
	buf := bytes.Buffer{}
	buf.Grow(37)
	for i := int64(0); i < 1000000; i++ {
		buf.Reset()
		value := time.Now().Unix()
		buf.WriteString(strconv.FormatInt(i,10))
		buf.WriteString("_")
		buf.WriteString(strconv.FormatInt(value,10))
		arr[buf.String()] = value
	}

}

func calcV1()  {
	arr := map[string]int64{}
	for i := int64(0); i < 1000000; i++ {
		value := time.Now().Unix()
		key := strconv.FormatInt(i,10) + "_" + strconv.FormatInt(value,10)
		arr[key] = value
	}
}

func calcV2()  {
	arr := map[string]int64{}
	for i := int64(0); i < 1000000; i++ {
		buf := bytes.Buffer{}
		value := time.Now().Unix()
		buf.WriteString(strconv.FormatInt(i,10))
		buf.WriteString("_")
		buf.WriteString(strconv.FormatInt(value,10))
		arr[buf.String()] = value
	}
}

func calcV3()  {
	arr := map[string]int64{}
	for i := int64(0); i < 1000000; i++ {
		buf := bytes.Buffer{}
		buf.Grow(37)
		value := time.Now().Unix()
		buf.WriteString(strconv.FormatInt(i,10))
		buf.WriteString("_")
		buf.WriteString(strconv.FormatInt(value,10))
		arr[buf.String()] = value
	}
}

func calcV4()  {
	arr := map[string]int64{}
	buf := bytes.Buffer{}
	buf.Grow(37)
	for i := int64(0); i < 1000000; i++ {
		buf.Reset()
		value := time.Now().Unix()
		buf.WriteString(strconv.FormatInt(i,10))
		buf.WriteString("_")
		buf.WriteString(strconv.FormatInt(value,10))
		arr[buf.String()] = value
	}

}

func run(name string, f func(),isRecord bool) {
	if isRecord {
		traceFile,err := os.Create(name+"_trace.out")
		if err != nil {
			panic(err.Error())
		}
		trace.Start(traceFile)
		defer trace.Stop()

		cpuFile,err := os.Create(name+"_cpu.out")
		if err != nil {
			panic(err.Error())
		}
		pprof.StartCPUProfile(cpuFile)
		defer pprof.StopCPUProfile()

		memFile,err := os.Create(name+"_mem.out")
		if err != nil {
			panic(err.Error())
		}
		defer pprof.WriteHeapProfile(memFile)
	}


	f()
}
