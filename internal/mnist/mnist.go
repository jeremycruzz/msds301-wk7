package mnist

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/petar/GoMNIST"
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/trees"
)

type Mnist struct {
	data    *base.DenseInstances
	specs   []base.AttributeSpec
	results []float64
}

const (
	IMAGE_SIZE  = 28 * 28
	DATA_FOLDER = "./jump-start-mnist-iforest/data"
	OUTPUT_FILE = "./jump-start-mnist-iforest/results/isotreeGoScores.csv"

	TREES       = 100
	MAX_DEPTH   = 16 //log2(60,000)
	SAMPLE_SIZE = 256
)

func NewMnist() *Mnist {
	return &Mnist{
		data:  base.NewDenseInstances(),
		specs: make([]base.AttributeSpec, IMAGE_SIZE+1),
	}
}

func (m *Mnist) WriteCsv() {
	file, err := os.Create(OUTPUT_FILE)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, num := range m.results {
		err = writer.Write([]string{strconv.FormatFloat(num, 'f', 15, 64)})
		if err != nil {
			panic(err)
		}
	}
}

func (m *Mnist) IsolationForest() {
	forest := trees.NewIsolationForest(TREES, MAX_DEPTH, SAMPLE_SIZE)
	forest.Fit(m.data)
	m.results = forest.Predict(m.data)
}

func (m *Mnist) LoadData() {

	//load mnist data
	train, _, err := GoMNIST.Load(DATA_FOLDER)
	if err != nil {
		panic(err)
	}

	//add attributes for each pixel
	for i := 0; i < IMAGE_SIZE+1; i++ {
		attr := strconv.Itoa(i)
		m.specs[i] = m.data.AddAttribute(base.NewFloatAttribute(attr))
	}

	//add category attribute
	// classAttribute := base.NewCategoricalAttribute()
	// classAttribute.SetName("label")
	classAttribute := base.NewFloatAttribute("label")

	m.specs = append(m.specs, m.data.AddAttribute(classAttribute))
	m.data.AddClassAttribute(classAttribute)

	err = m.data.AddClassAttribute(classAttribute) //should return a spec...
	if err != nil {
		panic(err)
	}

	//allocate mem for rows
	m.data.Extend(train.Count())

	for i := 0; i < train.Count(); i++ {
		image, label := train.Get(i)
		//for each pixel byte convert it to a normalized float and set it
		for j, pixel := range image {
			m.data.Set(m.specs[j], i, base.PackFloatToBytes(float64(pixel)/255.0))
		}
		m.data.Set(m.specs[IMAGE_SIZE], i, base.PackFloatToBytes(float64(label))) //package requires this to be float instead of uint
	}
}
