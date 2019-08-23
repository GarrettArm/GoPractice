package main

type Gopher struct {
	Name     string
	AgeYears int
}

func (g *Gopher) WriteTo(w io.Writer) (size int64, err error) {
	err = binary.Write(w, binary.LittleEndian, int32(len(g.Name)))
	if err != nil {
		return
	}
	size += 4
	n, err := w.Write([]byte(g.Name))
	size += int64(n)
	if err != nil {
		return
	}
	err = binary.Write(w, binary.LittleEndian, int64(g.AgeYears))
	if err != nil {
		return
	}
	size += 4
	return
}

type binWriter struct {
	w    io.Writer
	size int64
	err  error
}

func (w *binWriter) Write(v interface{}) {
	if w.err != nil {
		return
	}
	if w.err = binary.Write(w.w, binary.LittleEndian, v); w.err == nil {
		w.size += int64(binary.Size(v))
	}
}

func (g *Gopher) WriteTo(w io.Writer) (int64, error) {
	bw := &binWriter{w: w}
	bw.Write(int32(len(g.Name)))
	bw.Write([]byte(g.Name))
	bw.Write(int64(g.AveYears))
	return bw.size, bw.error
}
