package gotorch_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	torch "github.com/wangkuiyi/gotorch"
)

// >>> t = torch.tensor([[-0.5, -1.], [1., 0.5]])
// >>> s = torch.tensor([[-0.5, -1.], [1., 0.5]])
// >>> t+s
// tensor([[-1., -2.],
//         [ 2.,  1.]])
func TestAdd(t *testing.T) {
	r := torch.NewTensor([][]float32{{-0.5, -1}, {1, 0.5}})
	s := torch.NewTensor([][]float32{{-0.5, -1}, {1, 0.5}})
	q := r.Add(s, 1)
	g := "-1 -2\n 2  1\n[ CPUFloatType{2,2} ]"
	assert.Equal(t, g, q.String())
}

func TestAddI(t *testing.T) {
	x := torch.RandN([]int64{2, 3}, false)
	y := torch.RandN([]int64{2, 3}, false)
	z := torch.Add(x, y, 1)
	x.AddI(y, 1)
	assert.True(t, torch.Equal(x, z))
}

// >>> torch.eq(torch.tensor([[1, 2], [3, 4]]), torch.tensor([[1, 1], [4, 4]]))
// tensor([[ True, False],
//         [False, True]])
func TestEq(t *testing.T) {
	a := torch.NewTensor([][]int16{{1, 2}, {3, 4}})
	b := torch.NewTensor([][]int16{{1, 3}, {2, 4}})
	c := torch.Eq(a, b)
	g := " 1  0\n 0  1\n[ CPUBoolType{2,2} ]"
	assert.Equal(t, g, c.String())
}

func TestTensorEq(t *testing.T) {
	a := torch.NewTensor([][]int16{{1, 2}, {3, 4}})
	b := torch.NewTensor([][]int16{{1, 3}, {2, 4}})
	c := a.Eq(b)
	g := " 1  0\n 0  1\n[ CPUBoolType{2,2} ]"
	assert.Equal(t, g, c.String())
}

func TestEqual(t *testing.T) {
	a := torch.NewTensor([]int64{1, 2})
	b := torch.NewTensor([]int64{1, 2})
	assert.True(t, torch.Equal(a, b))
}

// >>> s = torch.tensor([1,2])
// >>> t = torch.tensor([[1,2],[3,4]])
// >>> s.expand_as(t)
// tensor([[1, 2],
//         [1, 2]])
func TestExpandAs(t *testing.T) {
	a := torch.NewTensor([]int8{'a', 'b'})
	b := torch.NewTensor([][]int8{{1, 2}, {3, 4}})
	c := torch.ExpandAs(a, b)
	g := " 97  98\n 97  98\n[ CPUCharType{2,2} ]"
	assert.Equal(t, g, c.String())
}

func TestTensorExpandAs(t *testing.T) {
	a := torch.NewTensor([]int8{'a', 'b'})
	b := torch.NewTensor([][]int8{{1, 2}, {3, 4}})
	c := a.ExpandAs(b)
	g := " 97  98\n 97  98\n[ CPUCharType{2,2} ]"
	assert.Equal(t, g, c.String())
}

// >>> torch.flatten(torch.tensor([[-0.5, -1.], [1., 0.5]]), 0, 1)
// tensor([-0.5000, -1.0000,  1.0000,  0.5000])
func TestFlatten(t *testing.T) {
	r := torch.Flatten(torch.NewTensor([][]float32{{-0.5, -1}, {1, 0.5}}),
		0, 1)
	g := "-0.5000\n-1.0000\n 1.0000\n 0.5000\n[ CPUFloatType{4} ]"
	assert.Equal(t, g, r.String())
}

func TestIndexSelect(t *testing.T) {
	x := torch.RandN([]int64{3, 4}, false)
	indices := torch.NewTensor([]int64{0, 2})
	y := torch.IndexSelect(x, 0, indices)
	assert.Equal(t, int64(2), y.Shape()[0])
	assert.Equal(t, int64(4), y.Shape()[1])
}

func TestItem(t *testing.T) {
	x := torch.NewTensor([]float32{1})
	y := x.Item()
	assert.Equal(t, float32(1), y)
}

// >>> torch.nn.functional.leaky_relu(torch.tensor([[-0.5, -1.], [1., 0.5]]))
// tensor([[-0.0050, -0.0100],
//         [ 1.0000,  0.5000]])
func TestLeakyRelu(t *testing.T) {
	r := torch.LeakyRelu(torch.NewTensor([][]float32{{-0.5, -1}, {1, 0.5}}),
		0.01)
	g := "-0.0050 -0.0100\n 1.0000  0.5000\n[ CPUFloatType{2,2} ]"
	assert.Equal(t, g, r.String())
}

// >>> torch.nn.functional.log_softmax(torch.tensor([[-0.5, -1.], [1., 0.5]]), dim=1)
// tensor([[-0.4741, -0.9741],
//         [-0.4741, -0.9741]])
func TestLogSoftmax(t *testing.T) {
	r := torch.LogSoftmax(torch.NewTensor([][]float32{{-0.5, -1}, {1, 0.5}}),
		1)
	g := "-0.4741 -0.9741\n-0.4741 -0.9741\n[ CPUFloatType{2,2} ]"
	assert.Equal(t, g, r.String())
}

// >>> torch.mean(torch.tensor([[-0.5, -1.], [1., 0.5]]))
// tensor(0.)
func TestMean(t *testing.T) {
	r := torch.Mean(torch.NewTensor([][]float32{{-0.5, -1}, {1, 0.5}}))
	// BUG: The result should be 0.
	g := "0\n[ CPUFloatType{} ]"
	assert.Equal(t, g, r.String())
}

func TestTensorMean(t *testing.T) {
	x := torch.RandN([]int64{2, 3}, true)
	y := x.Mean()
	z := y.Item()
	assert.NotNil(t, z)
}

// >>> torch.relu(torch.tensor([[-0.5, -1.], [1., 0.5]]))
// tensor([[0.0000, 0.0000],
//         [1.0000, 0.5000]])
func TestRelu(t *testing.T) {
	r := torch.Relu(torch.NewTensor([][]float32{{-0.5, -1}, {1, 0.5}}))
	g := " 0.0000  0.0000\n 1.0000  0.5000\n[ CPUFloatType{2,2} ]"
	assert.Equal(t, g, r.String())
}

// >>> torch.sigmoid(torch.tensor([[-0.5, -1.], [1., 0.5]]))
// tensor([[0.3775, 0.2689],
//         [0.7311, 0.6225]])
func TestSigmoid(t *testing.T) {
	r := torch.Sigmoid(torch.NewTensor([][]float32{{-0.5, -1}, {1, 0.5}}))
	g := " 0.3775  0.2689\n 0.7311  0.6225\n[ CPUFloatType{2,2} ]"
	assert.Equal(t, g, r.String())
}

func TestStack(t *testing.T) {
	t1 := torch.RandN([]int64{2, 3}, false)
	t2 := torch.RandN([]int64{2, 3}, false)
	out := torch.Stack([]torch.Tensor{t1, t2}, 0)
	assert.Equal(t, []int64{2, 2, 3}, out.Shape())
}

func TestSqueeze(t *testing.T) {
	x := torch.RandN([]int64{2, 1, 2, 1, 2}, false)
	y := torch.Squeeze(x)
	assert.NotNil(t, y.T)
	z := torch.Squeeze(x, 1)
	assert.NotNil(t, z.T)
}

func TestSum(t *testing.T) {
	x := torch.NewTensor([]float32{1, 2, 4, 7})
	y := torch.Sum(x)
	z := y.Item()
	assert.Equal(t, float32(14), z)
}

func TestTanh(t *testing.T) {
	a := torch.RandN([]int64{4}, false)
	b := torch.Tanh(a)
	assert.NotNil(t, b.T)
}

// >>> torch.topk(torch.tensor([[-0.5, -1.], [1., 0.5]]), 1, 1, True, True)
// torch.return_types.topk(
// values=tensor([[-0.5000],
//         [ 1.0000]]),
// indices=tensor([[0],
//         [0]]))
func TestTopK(t *testing.T) {
	r, i := torch.TopK(torch.NewTensor([][]float64{{-0.5, -1}, {1, 0.5}}),
		1, 1, true, true)
	gr := "-0.5000\n 1.0000\n[ CPUDoubleType{2,1} ]"
	gi := " 0\n 0\n[ CPULongType{2,1} ]"
	assert.Equal(t, gr, r.String())
	assert.Equal(t, gi, i.String())
}

// >>> torch.transpose(torch.tensor([[-0.5, -1.], [1., 0.5]]), 0, 1)
// tensor([[-0.5000,  1.0000],
//         [-1.0000,  0.5000]])
func TestTranspose(t *testing.T) {
	r := torch.Transpose(torch.NewTensor([][]float32{{-0.5, -1}, {1, 0.5}}),
		0, 1)
	g := "-0.5000  1.0000\n-1.0000  0.5000\n[ CPUFloatType{2,2} ]"
	assert.Equal(t, g, r.String())
}
