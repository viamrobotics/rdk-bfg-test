package pointcloud

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"image/color"
	"path/filepath"

	"github.com/edaniels/golog"
	"github.com/jblindsay/lidario"
)

func NewFromFile(fn string, logger golog.Logger) (*PointCloud, error) {
	switch filepath.Ext(fn) {
	case ".las":
		return NewFromLASFile(fn, logger)
	default:
		return nil, fmt.Errorf("do not know how to read file %q", fn)
	}
}

const pointValueDataTag = "rc|pv"

func NewFromLASFile(fn string, logger golog.Logger) (*PointCloud, error) {
	lf, err := lidario.NewLasFile(fn, "r")
	if err != nil {
		return nil, err
	}
	defer lf.Close()

	var hasValue bool
	var valueData []byte
	for _, d := range lf.VlrData {
		if d.Description == pointValueDataTag {
			hasValue = true
			valueData = d.BinaryData
			break
		}
	}

	pc := New(logger)
	for i := 0; i < lf.Header.NumberPoints; i++ {
		p, err := lf.LasPoint(i)
		if err != nil {
			return nil, err
		}
		data := p.PointData()

		x, y, z := data.X, data.Y, data.Z
		if x < minExactFloat64Integer || x > maxExactFloat64Integer ||
			y < minExactFloat64Integer || y > maxExactFloat64Integer ||
			z < minExactFloat64Integer || z > maxExactFloat64Integer {
			logger.Warnf("potential floating point lossiness for LAS point",
				"point", data, "range", fmt.Sprintf("[%d,%d]", minExactFloat64Integer, maxExactFloat64Integer))
		}
		pToSet := NewBasicPoint(int(x), int(y), int(z))

		if lf.Header.PointFormatID == 2 && p.RgbData() != nil {
			r := uint8(p.RgbData().Red / 256)
			g := uint8(p.RgbData().Green / 256)
			b := uint8(p.RgbData().Blue / 256)
			pToSet.SetColor(color.NRGBA{r, g, b, 255})
		}

		if hasValue {
			value := int(binary.LittleEndian.Uint64(valueData[i*8 : (i*8)+8]))
			pToSet.SetValue(value)
		}

		if err := pc.Set(pToSet); err != nil {
			return nil, err
		}
	}
	return pc, nil
}

func (pc *PointCloud) WriteToFile(fn string) error {
	lf, err := lidario.NewLasFile(fn, "w")
	if err != nil {
		return err
	}
	var successful bool
	defer func() {
		if !successful {
			if err := lf.Close(); err != nil {
				pc.logger.Debug(err)
			}
		}
	}()

	pointFormatID := 0
	if pc.hasColor {
		pointFormatID = 2
	}
	if err := lf.AddHeader(lidario.LasHeader{
		PointFormatID: byte(pointFormatID),
	}); err != nil {
		return err
	}

	var pVals []int
	if pc.hasValue {
		pVals = make([]int, 0, pc.Size())
	}
	var lastErr error
	pc.Iterate(func(p Point) bool {
		pos := p.Position()
		var lp lidario.LasPointer
		pr0 := &lidario.PointRecord0{
			// floating point losiness validated/warned from set/load
			X:         float64(pos.X),
			Y:         float64(pos.Y),
			Z:         float64(pos.Z),
			Intensity: 0,
			BitField: lidario.PointBitField{
				Value: (1) | (1 << 3) | (0 << 6) | (0 << 7),
			},
			ClassBitField: lidario.ClassificationBitField{
				Value: 0,
			},
			ScanAngle:     0,
			UserData:      0,
			PointSourceID: 1,
		}
		lp = pr0
		if pc.hasColor {
			red, green, blue := 255, 255, 255
			if p.HasColor() {
				r, g, b := p.RGB255()
				red, green, blue = int(r), int(g), int(b)
			}
			lp = &lidario.PointRecord2{
				PointRecord0: pr0,
				RGB: &lidario.RgbData{
					Red:   uint16(red * 256),
					Green: uint16(green * 256),
					Blue:  uint16(blue * 256),
				},
			}
		}
		if pc.hasValue {
			if p.HasValue() {
				pVals = append(pVals, p.Value())
			} else {
				pVals = append(pVals, 0)
			}
		}
		if err := lf.AddLasPoint(lp); err != nil {
			lastErr = err
			return false
		}
		return true
	})
	if pc.hasValue {
		var buf bytes.Buffer
		for _, v := range pVals {
			bytes := make([]byte, 8)
			binary.LittleEndian.PutUint64(bytes, uint64(v))
			buf.Write(bytes)
		}
		if err := lf.AddVLR(lidario.VLR{
			UserID:                  "",
			Description:             pointValueDataTag,
			BinaryData:              buf.Bytes(),
			RecordLengthAfterHeader: buf.Len(),
		}); err != nil {
			return err
		}
	}
	if lastErr != nil {
		return lastErr
	}

	successful = true
	return lf.Close()
}
