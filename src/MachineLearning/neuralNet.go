package mCraftnBlockmLearning

/* I hate Machine Learning. I don't like nueral networks.
 * "Nueral network": That is what I would use if my only job was to maximize resume keywords.
 * Perhaps I can console myself that this is only a convineince struct:
 * This has the name "NueralNet" becaus I couldn't think of a better name.
 * (I'll still call this a Neural Net on my resume the minute this is done.)
 */

 import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

type NNLayer struct{
	weights mat.Matrix
	biases mat.Matrix
}

type PoolLayer struct{
	// Has to support max pool & average pool.
}

func newFullLayer(rows int, cols int){
	var returnNNLayer NNLayer = new(NNLayer);
	returnNNLayer.weights = mat.NewDense(rows, cols);
	returnNNLayer.biases = mat.NewDense(rows, cols);

	return returnNNLayer;
}


type NeuralNet struct{
	trainingInput []mat.Matrix;
	trainingOutput []mat.Matrix;

	nnLayers []NNLayer;

	outputDimR, outputDimC int;

	lossFuction func(output mat.Matrix, target mat.Matrix) float64;
}

func NewNeuralNet() NeuralNet{
	nnn := new(NeuralNet);
	nnn.trainingInput = make([]mat.Matrix, 0)
	nnn.trainingOutput = make([]mat.Matrix, 0)
	nnn.nnLayers = make([]NNLayer, 0)

	nnn.outputDimR = -1;
	nnn.outputDimC = -1;
	
	return *nnn
}

func AddTrainingData(nn *NeuralNet, input []mat.Matrix, output []mat.Matrix){
	for i := 0; i < len(input); i++ {
		nn.trainingInput = append(nn.trainingInput, input[i])
		nn.trainingOutput = append(nn.trainingOutput, output[i])
	}
	for i := 0; i < len(nn.trainingInput); i++{
		if (nn.trainingInput[i] != nn.trainingInput[i-1]){
			panic("All input matrixies must have same length!")
		}
		if (nn.trainingOutput[i] != nn.trainingOutput[i-1]){
			panic("All output matrixies must have same length!")
		}
	}
	if (nn.outputDimR == -1 || nn.outputDimC == -1) && len(nn.trainingInput) > 0{
		nn.outputDimR, nn.outputDimC = nn.trainingInput[0].Dims();
	}
}

func addLayer(nn *NeuralNet, newLayer NNLayer){
	weightMatR, weightMatC := newLayer.weights.Dims()
	biasMatR, biasMatC := newLayer.weights.Dims()

	if (nn.outputDimC != weightMatR){
		fmt.Println("Cannot Multiply new layer with previous output")
		fmt.Printf("Previous output matrix had %d columns", nn.outputDimC)
		fmt.Printf("While new layer matrix has %d rows", weightMatR)
		return
	}

	newOutputDimR, newOutputDimC := nn.outputDimR, weightMatC

	if !(biasMatR == newOutputDimR && biasMatC== newOutputDimC) {
		fmt.Println("Biases not compatible with Output of Multiplying previous output with weight matrix")
		fmt.Printf("Previous Output * Weight matrix has dimensions (%d Rows, %d Cols) \n", newOutputDimR, newOutputDimC)
		fmt.Printf("Bias matrix has dimensions (%d Rows, %d Cols). Cannot be added!", biasMatR, biasMatC )
		return
	}


} 

func PrintInfo(nnn *NeuralNet) {
	if len(nnn.trainingInput) > 0 {
		r, c := nnn.trainingInput[0].Dims()
		fmt.Printf("Input dimensions: %d rows, %d cols\n", r, c)
	} else {
		fmt.Println("Input dimensions: (no training input)")
	}

	fmt.Printf("Layers: %d\n", len(nnn.nnLayers))
	for i, layer := range nnn.nnLayers {
		wr, wc := 0, 0
		br, bc := 0, 0
		if layer.weights != nil {
			wr, wc = layer.weights.Dims()
		}
		if layer.biases != nil {
			br, bc = layer.biases.Dims()
		}
		fmt.Printf("  Layer %d: weights (%d rows, %d cols), biases (%d rows, %d cols)\n", i, wr, wc, br, bc)
	}

	if len(nnn.trainingOutput) > 0 {
		r, c := nnn.trainingOutput[0].Dims()
		fmt.Printf("Output dimensions: %d rows, %d cols\n", r, c)
	} else if nnn.outputDimR >= 0 && nnn.outputDimC >= 0 {
		fmt.Printf("Output dimensions: %d rows, %d cols\n", nnn.outputDimR, nnn.outputDimC)
	} else {
		fmt.Println("Output dimensions: (unset)")
	}
}