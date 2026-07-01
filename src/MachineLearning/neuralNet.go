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

type NeuralNet struct{
	trainingInput []mat.Matrix;
	trainingOutput []mat.Matrix;

	nnLayers []NNLayer;

	int outputDimR, outputDimC;

	lossFuction func(output mat.Matrix, target mat.Matrix) float64;
}

func NewNeuralNet() NeuralNet{
	nnn := new(NeuralNet);
	nnn.trainingInput = make([]mat.Matrix, 0)
	nnn.trainingOutput = make([]mat.Matrix, 0)
	nnLayers = make([]NNLayer, 0)

	outputDimR = -1;
	outputDimC = -1;
	
	return *nnn
}

func AddTrainingData(nnn *NeuralNet, input []mat.Matrix, output []mat.Matrix){
	if len(input) != len(output){
		panic("input and output must have the same length")
	}

	for i := 0; i < len(input); i++ {
		nnn.trainingInput = append(nnn.trainingInput, input[i])
		nnn.trainingOutput = append(nnn.trainingOutput, output[i])
	}
	for i := 0; i < len(nnn.trainingInput); ++i{
		if (nnn.trainingInput[i] != nnn.trainingInput[i-1]){
			panic("All input matrixies must have same length!")
		}
		if (nnn.trainingOutput[i] != nnn.trainingOutput[i-1]){
			panic("All output matrixies must have same length!")
		}
	}
	if (outputDimR == -1 || outputDimC == -1) && len(nnn.trainingInput) > 0{
		(outputDimR, outputDimC) = nnn.trainingInput[0].Dims();
	}
}

func addLayer(nnn *NeuralNet, newLayer NNLayer){
	(newOutputDimR, newOutputDimC) = (newOutputDimR, NNLayer.weights.Dims()[1])

	if NNLayer.biases.Dims() != (newOutputDimR, newOutputDimC) {
		fmt.Println("Biases not compatible with Output of Multiplying previous output with weight matrix")
		fmt.Printf("Previous Output * Weight matrix has dimensions (%d Rows, %d Cols) \n", newOutputDimR, newOutputDimC)
		fmt.Printf("Bias matrix has dimensions (%d Rows, %d Cols). Cannot be added!", NNLayer.biases.Dims()[0], NNLayer.biases.Dims() [1] )
		return
	}


} 

func PrintInfo(nnn *NeuralNet){

}