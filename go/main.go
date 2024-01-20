package main

import (
	"fmt"
	"hash/fnv"
	"log"
	"math/rand"
	"strconv"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	// generate random number in array
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, 1000000)

	// store time of different data structures based on the size of the array
	key_size := []int{}
	default_dict_time := []int64{}
	chained_dict_time := []int64{}
	linear_dict_time := []int64{}

	// print the map
	var k, v int

	for t := 1; t < 100000; t *= 2 {
		key_size = append(key_size, t)
		for i := range arr {
			arr[i] = rand.Intn(t)
		}

		// record start time
		startTime := time.Now()

		// store it in map
		dict := make(map[int]int)
		for _, i := range arr {
			dict[i] = 1
		}

		for key, value := range dict {
			k = key
			v = value
		}
		// fmt.Println(k, v)

		// record end time
		endTime := time.Now()

		// print the time taken
		// fmt.Println("Time taken for map:", endTime.Sub(startTime))
		default_dict_time = append(default_dict_time, endTime.Sub(startTime).Nanoseconds())

		// record start time
		startTime = time.Now()

		// store it in ChainedDictionary
		chainingDict := NewChainedDictionary()
		for _, i := range arr {
			chainingDict.add(i, 1)
		}

		// print the ChainedDictionary
		for _, i := range arr {
			k = i
			v = chainingDict.get(i).(int)
		}
		// fmt.Println(k, v)

		// record end time
		endTime = time.Now()

		// print the time taken
		// fmt.Println("Time taken for ChainedDictionary:", endTime.Sub(startTime))
		chained_dict_time = append(chained_dict_time, endTime.Sub(startTime).Nanoseconds())

		// store it in LinearProbingDictionary

		// record start time
		startTime = time.Now()

		linearDict := NewLinearProbingDictionary()
		for _, i := range arr {
			linearDict.add(i, 1)
		}

		// print data
		for _, i := range arr {
			k = i
			v = linearDict.get(i).(int)
		}
		// fmt.Println(k, v)

		// record end time
		endTime = time.Now()

		// print the time taken
		// fmt.Println("Time taken for LinearProbingDictionary:", endTime.Sub(startTime))
		linear_dict_time = append(linear_dict_time, endTime.Sub(startTime).Nanoseconds())
	}

	// fmt.Println("Time taken for map:", default_dict_time)
	// fmt.Println("Time taken for ChainedDictionary:", chained_dict_time)
	// fmt.Println("Time taken for LinearProbingDictionary:", linear_dict_time)
	// fmt.Println("Key size:", key_size)

	fmt.Println(k, v)

	plotTimeTakenGraph(key_size, default_dict_time, chained_dict_time, linear_dict_time, "time_taken_line_plot.png")

}

// ***************************************************ChainedDictionary*****************************************************************

// ChainedDictionary struct
type ChainedDictionary struct {
	arr           [][][2]interface{}
	loadFactor    float64
	usedKeys      int
	totalKeys     int
	resizingCount int
	resizingSum   int
}

// NewChainedDictionary creates a new ChainedDictionary instance
func NewChainedDictionary() *ChainedDictionary {
	return &ChainedDictionary{
		arr:        make([][][2]interface{}, 10000),
		loadFactor: 0,
		usedKeys:   0,
		totalKeys:  10000,
	}
}

func (cd *ChainedDictionary) calculateLoadFactor() {
	cd.loadFactor = float64(cd.usedKeys) / float64(cd.totalKeys)
}

func (cd *ChainedDictionary) shouldResize() bool {
	return cd.loadFactor >= 0.5
}

func (cd *ChainedDictionary) resizeArray() {
	cd.totalKeys *= 2
	cd.calculateLoadFactor()
	newArr := make([][][2]interface{}, cd.totalKeys)
	for i := range cd.arr {
		if cd.arr[i] != nil {
			newArr[cd.getHashKey(cd.arr[i][0][0])] = cd.arr[i]
			cd.resizingSum++
		}
	}
	cd.arr = newArr
	cd.resizingCount++
}

func (cd *ChainedDictionary) getHashKey(key interface{}) int {
	h := fnv.New32a()
	_, _ = h.Write([]byte(fmt.Sprintf("%v", key)))
	return int(h.Sum32()) % cd.totalKeys
}

func (cd *ChainedDictionary) add(key, value interface{}) {
	if cd.shouldResize() {
		cd.resizeArray()
	}

	hashKey := cd.getHashKey(key)
	if cd.arr[hashKey] != nil {
		for i := range cd.arr[hashKey] {
			if cd.arr[hashKey][i][0] == key {
				cd.arr[hashKey][i][1] = value
				return
			}
		}
		cd.arr[hashKey] = append(cd.arr[hashKey], [2]interface{}{key, value})
	} else {
		cd.arr[hashKey] = [][2]interface{}{{key, value}}
		cd.usedKeys++
	}
	cd.calculateLoadFactor()
}

func (cd *ChainedDictionary) get(key interface{}) interface{} {
	hashKey := cd.getHashKey(key)
	if cd.arr[hashKey] != nil {
		for _, kv := range cd.arr[hashKey] {
			if kv[0] == key {
				return kv[1]
			}
		}
	}
	return nil
}

func (cd *ChainedDictionary) remove(key interface{}) bool {
	hashKey := cd.getHashKey(key)
	if cd.arr[hashKey] != nil {
		for i := range cd.arr[hashKey] {
			if cd.arr[hashKey][i][0] == key {
				cd.arr[hashKey] = append(cd.arr[hashKey][:i], cd.arr[hashKey][i+1:]...)
				return true
			}
		}
	}
	return false
}

func (cd *ChainedDictionary) getMetadata() map[string]interface{} {
	return map[string]interface{}{
		"totalKeys":     cd.totalKeys,
		"usedKeys":      cd.usedKeys,
		"loadFactor":    cd.loadFactor,
		"resizingCount": cd.resizingCount,
		"resizingSum":   cd.resizingSum,
	}
}

// ***************************************************LinearProbingDictionary*****************************************************************

// LinearProbingDictionary struct
type LinearProbingDictionary struct {
	arr           [][2]interface{}
	loadFactor    float64
	usedKeys      int
	totalKeys     int
	resizingCount int
	resizingSum   int
}

// NewLinearProbingDictionary creates a new LinearProbingDictionary instance
func NewLinearProbingDictionary() *LinearProbingDictionary {
	return &LinearProbingDictionary{
		arr:        make([][2]interface{}, 2),
		loadFactor: 0,
		usedKeys:   0,
		totalKeys:  2,
	}
}

func (lpd *LinearProbingDictionary) calculateLoadFactor() {
	lpd.loadFactor = float64(lpd.usedKeys) / float64(lpd.totalKeys)
}

func (lpd *LinearProbingDictionary) shouldResize() bool {
	return lpd.loadFactor >= 0.5
}

func (lpd *LinearProbingDictionary) resizeArray() {
	lpd.totalKeys *= 2
	lpd.calculateLoadFactor()
	newArr := make([][2]interface{}, lpd.totalKeys)
	for i := range lpd.arr {
		if lpd.arr[i][0] != nil {
			hashKey := lpd.getHashKey(lpd.arr[i][0])
			newArr[hashKey] = lpd.arr[i]
			lpd.resizingSum++
		}
	}
	lpd.arr = newArr
	lpd.resizingCount++
}

func (lpd *LinearProbingDictionary) getHashKey(key interface{}) int {
	h := fnv.New32a()
	_, _ = h.Write([]byte(fmt.Sprintf("%v", key)))
	return int(h.Sum32()) % lpd.totalKeys
}

func (lpd *LinearProbingDictionary) linearProbe(index int) int {
	return (index + 1) % lpd.totalKeys
}

func (lpd *LinearProbingDictionary) add(key, value interface{}) {
	if lpd.shouldResize() {
		lpd.resizeArray()
	}

	hashKey := lpd.getHashKey(key)
	for lpd.arr[hashKey][0] != nil {
		if lpd.arr[hashKey][0] == key {
			lpd.arr[hashKey][1] = value
			return
		}
		hashKey = lpd.linearProbe(hashKey)
	}
	lpd.arr[hashKey] = [2]interface{}{key, value}
	lpd.usedKeys++
	lpd.calculateLoadFactor()
}

func (lpd *LinearProbingDictionary) get(key interface{}) interface{} {
	hashKey := lpd.getHashKey(key)
	for lpd.arr[hashKey][0] != nil {
		if lpd.arr[hashKey][0] == key {
			return lpd.arr[hashKey][1]
		}
		hashKey = lpd.linearProbe(hashKey)
	}
	return nil
}

func (lpd *LinearProbingDictionary) getMetadata() map[string]interface{} {
	return map[string]interface{}{
		"totalKeys":     lpd.totalKeys,
		"usedKeys":      lpd.usedKeys,
		"loadFactor":    lpd.loadFactor,
		"resizingCount": lpd.resizingCount,
		"resizingSum":   lpd.resizingSum,
	}
}

// ***************************************************Plotting*****************************************************************
func plotTimeTakenGraph(key_size []int, default_dict_time, chained_dict_time, linear_dict_time []int64, filename string) {
	// Provided data
	labels := make([]string, len(key_size))
	for i, num := range key_size {
		labels[i] = strconv.Itoa(num)
	}
	nanoToMilli := 1000000.0
	timesMap := int64ArrayToFloat(default_dict_time, nanoToMilli)
	timesChainedDict := int64ArrayToFloat(chained_dict_time, nanoToMilli)
	timesLinearProbe := int64ArrayToFloat(linear_dict_time, nanoToMilli)

	fmt.Println(labels)
	fmt.Println(timesMap)
	fmt.Println(timesChainedDict)
	fmt.Println(timesLinearProbe)

	// Create a new plot
	p := plot.New()

	// Create line plots
	lineMap, err := plotter.NewLine(createXYs(labels, timesMap))
	if err != nil {
		log.Fatal(err)
	}
	lineMap.LineStyle.Width = vg.Points(2)
	lineMap.Color = plotutil.Color(0)

	lineChainedDict, err := plotter.NewLine(createXYs(labels, timesChainedDict))
	if err != nil {
		log.Fatal(err)
	}
	lineChainedDict.LineStyle.Width = vg.Points(2)
	lineChainedDict.Color = plotutil.Color(1)

	lineLinearProbe, err := plotter.NewLine(createXYs(labels, timesLinearProbe))
	if err != nil {
		log.Fatal(err)
	}
	lineLinearProbe.LineStyle.Width = vg.Points(2)
	lineLinearProbe.Color = plotutil.Color(2)

	// Add lines to the plot
	p.Add(lineMap, lineChainedDict, lineLinearProbe)

	// Set axis labels and title
	p.X.Label.Text = "Key Size"
	p.Y.Label.Text = "Time (ms)"
	p.Title.Text = "Time taken for different data structures"

	// Add legend
	p.Legend.Add("Map", lineMap)
	p.Legend.Add("ChainedDictionary", lineChainedDict)
	p.Legend.Add("LinearProbingDictionary", lineLinearProbe)

	// Save the plot to a file
	if err := p.Save(6*vg.Inch, 4*vg.Inch, filename); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Line plot saved to %s\n", filename)
}

// Helper function to create XYs for the given labels and values
func createXYs(labels []string, values []float64) plotter.XYs {
	pts := make(plotter.XYs, len(labels))
	for i, val := range labels {
		pts[i].X, _ = strconv.ParseFloat(val, 64)
		pts[i].Y = values[i]
	}
	return pts
}

func int64ArrayToFloat(arr []int64, divisor float64) []float64 {
	floatArr := make([]float64, len(arr))
	for i, val := range arr {
		floatArr[i] = float64(val) / divisor
	}
	return floatArr
}
