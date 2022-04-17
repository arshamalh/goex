package rot13

import "io"

type Rot13Reader struct {
	R io.Reader
}

func (rot *Rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot.R.Read(p)
	for i := 0; i < len(p); i++ {
		if p[i] >= 'A' && p[i] < 'Z' {
			p[i] = 'A' + ((p[i] - 'A' + 13) % 26)
		} else if p[i] >= 'a' && p[i] < 'z' {
			p[i] = 'a' + ((p[i] - 'a' + 13) % 26)
		}
	}
	return
}

// Rot13 exercise main
// s := strings.NewReader("Lbh penpxrq gur pbqr!")
// r := rot13.Rot13Reader{R: s}
// io.Copy(os.Stdout, &r)
