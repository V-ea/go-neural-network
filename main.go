package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
)

//every layer nodes.output * weight = next layer
type neuralLayer struct {
	nodes [] float32
	weight [][] float32
	derive [] float32

	nextLayer *neuralLayer
}

type neuralNetwork struct {
	layers [] *neuralLayer
}

/**
检查可用之前要初始化no
 */
func (nl *neuralLayer) checkValid() bool{
	//nodes不允许为空至少要被初始化成一个数组
	if nl.nodes==nil{
		return false
	}
	nodeSize := len(nl.nodes)
	//nodes 必须是一个有长度的数组
	if nodeSize==0{
		return false
	}
	//最后一个节点可以是权重和下一层为空
	if 0==len(nl.weight)  && nl.nextLayer ==nil{
		return true
	}
	//正常节点
	if nodeSize !=0 && len(nl.weight)== nodeSize && len(nl.weight[0])== len(nl.nextLayer.nodes) {
		return true
	}
	return false
}

func (nl *neuralLayer) forward() bool{
	if nl.nextLayer == nil || nl.nextLayer.nodes == nil{
		return false
	}
	for i := 0; i<len(nl.nextLayer.nodes);i++  {
		nl.nextLayer.nodes[i] = 0.0
		for  j := 0; j < len(nl.nodes);j++{
			nl.nextLayer.nodes[i]+=nl.weight[j][i] * nl.nodes[j]
		}
	}
	return true
}

func (nn * neuralNetwork) value() []float32{
	//last layer is value of the network
	return nn.layers[len(nn.layers)-1].nodes
}

func (nl * neuralLayer) value() []float32 {
	return nl.nodes
}

func (nl * neuralLayer) weights() [][] float32  {
	return nl.weight
}

/**
given layers initialize the weights
 */
func (nn * neuralNetwork) addLayer(nl * neuralLayer, currSize int)  {
	nl.nodes = []float32{}
	for j:=0;j<currSize;j++{
		nl.nodes = append(nl.nodes, 0)
	}
	nn.layers = append(nn.layers, nl)
	//fmt.Println("1")
	//fmt.Println("2")
	if len(nn.layers)==1{
		return
	}
	size := len(nn.layers)
	//fmt.Println(size)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for g:=0;g<len(nn.layers[size-2].nodes);g++{
		array := make([]float32, currSize)
		nn.layers[size-2].weight = append(nn.layers[size-2].weight, array)
		for k:=0;k<currSize;k++{
			//fmt.Println(k, g)

			nn.layers[size-2].weight[g][k] = r.Float32() * 30
		}
	}
	nn.layers[size-2].nextLayer= nn.layers[size-1]
}
func (nn * neuralNetwork) forward(input [] float32) float32{
	if len(input)!=len(nn.layers[0].nodes){
		//exception
		panic("shape not equals.")
	}
	nn.layers[0].nodes = input
	result := true
	for _,nl :=range nn.layers {
		result = result && nl.forward()
	}

	return 0.0
}

func (nn * neuralNetwork) backward(output [] float32, learnRate float32)  float32{
	size := len(nn.layers)
	derive := output
	for i:=size-1;i>0;i--{
		oldDerive := nn.layers[i].backward(derive)
		fmt.Println(derive, oldDerive)
		//calc lower layer's derives
		var deriveLow []float32
		for indexLow,value :=range nn.layers[i-1].nodes{
			deriveCur := float32(0.0)
			for index,d :=range oldDerive{
				deriveCur = deriveCur + d * nn.layers[i-1].weight[indexLow][index]
				//change lower weights with learning rate
				nn.layers[i-1].weight[indexLow][index] = nn.layers[i-1].weight[indexLow][index] - d * (0 - value) * learnRate
			}
			deriveLow =  append(deriveLow, deriveCur)
		}
		derive= deriveLow


	}
	sumLoss:=float32(0.0)
	for i,_ := range nn.layers[len(nn.layers)-1].nodes{
		sumLoss = sumLoss  + (nn.layers[len(nn.layers)-1].nodes[i] - output[i]) *  (nn.layers[len(nn.layers)-1].nodes[i] - output[i])
	}
	return sumLoss
}

func (nl * neuralLayer) backward(output [] float32) [] float32 {
	size := len(nl.nodes)
	if size!=len(output){
		panic("backward output's size not equals to nodes")
	}
	var newDerive []float32
	for i,v :=range output{
		newDerive = append(newDerive, v - nl.nodes[i])
	}
	nl.derive = newDerive
	return newDerive
}

func (nl * neuralLayer) printLayer(){
	fmt.Println("nodes")
	fmt.Println(nl.nodes)
	fmt.Println("weights")
	fmt.Println(nl.weights())
}
func (nn * neuralNetwork) print()  {
	for i:=0; i<len(nn.layers);i++{
		fmt.Println("=========layer"+ strconv.Itoa(i) + "==========")
		nn.layers[i].printLayer()
	}
}
func (nn * neuralNetwork) train(input[] float32, output[] float32, learningRate float32) float32{
	nn.forward(input)
	nn.print()
	fmt.Println("backward")
	result :=  nn.backward(output, learningRate)
	nn.print()
	return result

}
func main(){
	defer func() {
		if err:=recover();err!=nil{
			fmt.Println("ERROR:", err) // 这里的err其实就是panic传入的内容
		}
	}()
	var arr []float32
	nl1 := new(neuralLayer)
	nl2 := new(neuralLayer)
	//nl3 := new(neuralLayer)
	nn := new(neuralNetwork)
	nn.addLayer(nl1, 3)
	nn.addLayer(nl2,1)
	for i:=0;i<20000;i++ {
		fmt.Println("Round" , i)
		arr = append(arr,nn.train([]float32{1, 2, 2}, []float32{14}, 0.01))
		arr = append(arr,nn.train([]float32{3, 4, 4}, []float32{30}, 0.01))
		arr = append(arr,nn.train([]float32{2, 3, 7}, []float32{33}, 0.01))
	}
	fmt.Println(arr)

}
